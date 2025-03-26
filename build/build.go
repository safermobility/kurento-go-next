package main

import (
	"bytes"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const DOCLINELENGTH = 79

var re = regexp.MustCompile(`(.+)\[\]`)
var mapRe = regexp.MustCompile(`(.+)\<\>`)

var CPXTYPES = make(map[string]string, 0)
var CLSTYPES = make(map[string]string, 0)

type Return struct {
	Doc  string
	Type string
}

type constructor struct {
	Name   string
	Doc    string
	Params []map[string]interface{}
}

type method struct {
	constructor
	Return map[string]interface{}
}

type class struct {
	Name        string
	Extends     string
	Doc         string
	Abstract    bool
	Properties  []map[string]interface{}
	Events      []string
	Constructor constructor
	Methods     []method

	Package string
}

type core struct {
	RemoteClasses []class
	ComplexTypes  []ComplexType
}

type ComplexType struct {
	TypeFormat string
	Doc        string
	Values     []string
	Name       string
	Properties []map[string]interface{}

	Package string
}

const (
	CORE     = "build/kms-core-valid-json/core.kmd.json"
	ELEMENTS = "build/kms-elements-valid-json/"
)

// template func that change MediaXXX to IMediaXXX to
// be sure to work with interface.
func tplCheckElement(thisPkg, p string) string {
	if len(p) > 5 && p[:5] == "Media" {
		if p[len(p)-4:] != "Type" {
			return "IMedia" + p[5:]
		}
	}

	return maybePrefixType(thisPkg, p)
}

func maybePrefixType(thisPkg, p string) string {
	original := p
	useSlice := ""
	if p[0] == '[' {
		useSlice = "[]"
		p = p[2:]
	}
	useStar := ""
	if p[0] == '*' {
		useStar = "*"
		p = p[1:]
	}
	if pkg, ok := CLSTYPES[p]; ok && pkg != thisPkg {
		return fmt.Sprintf("%s%s%s.%s", useSlice, useStar, pkg, p)
	}
	if pkg, ok := CPXTYPES[p]; ok && pkg != thisPkg {
		return fmt.Sprintf("%s%s%s.%s", useSlice, useStar, pkg, p)
	}
	return original
}

func isComplexType(t string) bool {
	if _, ok := CPXTYPES[t]; ok {
		return true
	}
	return false
}

var funcMap = template.FuncMap{
	"title":        strings.Title,
	"uppercase":    strings.ToUpper,
	"checkElement": tplCheckElement,
	"prefixType":   maybePrefixType,
	"paramValue": func(pkg string, p map[string]interface{}) string {
		name := p["name"].(string)
		t := p["type"].(string)
		t = tplCheckElement(pkg, t)

		ctype := isComplexType(t)
		switch t {
		case "float64", "int", "int64":
			return fmt.Sprintf("\"%s\" = %s", name, name)
		case "string", "bool":
			return fmt.Sprintf("\"%s\" = %s", name, name)
		default:
			// If param is not complexType, we have Id from String() method
			if !ctype && t[0] == 'I' { /* TODO: fix isInterface */
				return fmt.Sprintf("\"%s\" = fmt.Sprintf(\"%%s\", %s)", name, name)
			}
		}
		// Default is to set value to param
		return fmt.Sprintf("\"%s\" = %s", name, name)
	},
}

func formatDoc(doc string) string {

	doc = strings.Replace(doc, ":rom:cls:", "", -1)
	doc = strings.Replace(doc, ":term:", "", -1)
	doc = strings.Replace(doc, "``", `"`, -1)

	lines := strings.Split(doc, "\n")
	for i, line := range lines {
		lines[i] = "// " + strings.TrimSpace(line)
	}
	ret := strings.Join(lines, "\n")
	return ret
}

func formatTypes(p map[string]interface{}) map[string]interface{} {
	p["doc"] = formatDoc(p["doc"].(string))
	if p["type"] == "String[]" {
		p["type"] = "[]string"
	}

	if p["type"] == "String<>" {
		p["type"] = "map[string]interface{}"
	}

	if p["type"] == "String" {
		p["type"] = "string"
	}

	if p["type"] == "float" {
		p["type"] = "float64"
	}
	if p["type"] == "boolean" {
		p["type"] = "bool"
	}

	if p["type"] == "double" {
		p["type"] = "float64"
	}

	if re.MatchString(p["type"].(string)) {
		found := re.FindAllStringSubmatch(p["type"].(string), -1)
		p["type"] = "[]" + found[0][1]
	}

	if mapRe.MatchString(p["type"].(string)) {
		found := mapRe.FindAllStringSubmatch(p["type"].(string), -1)
		p["type"] = "map[string]" + found[0][1]
	}

	if p["defaultValue"] == "" || p["defaultValue"] == nil {
		switch p["type"] {
		case "string":
			p["defaultValue"] = `""`
		case "bool":
			p["defaultValue"] = "false"
		case "int", "int64", "float64":
			p["defaultValue"] = "0"
		}
	}

	return p
}

func getModel(path string) core {

	i := core{}
	data, _ := os.ReadFile(path)
	err := json.Unmarshal(data, &i)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to parse %s: %w", path, err))
	}
	return i
}

func getInterfaces(templates *template.Template) error {
	paths, _ := filepath.Glob(ELEMENTS + "elements.*.kmd.json")
	for _, p := range paths {
		r := getModel(p).RemoteClasses
		classes, err := parse(r, "elements", templates)
		if err != nil {
			return err
		}
		for name, class := range classes {
			CLSTYPES[name] = "elements"
			err = writePackageFile("elements", name, class)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parse(c []class, pkg string, templates *template.Template) (map[string][]byte, error) {
	ret := make(map[string][]byte, len(c))
	for _, cl := range c {
		log.Println("Generating ", cl.Name)

		// rewrite types
		for j, p := range cl.Properties {
			p = formatTypes(p)
			switch p["type"] {
			case "string", "float64", "int", "int64", "bool", "[]string":
			default:
				if typeString, ok := p["type"].(string); ok {
					if typeString[:2] == "[]" {
						t := typeString[2:]
						if _, ok := CPXTYPES[t]; ok {
							p["type"] = fmt.Sprintf("[]*%s", t)
						} else {
							p["type"] = "[]I" + t
						}
					} else {
						if _, ok := CPXTYPES[typeString]; ok {
							p["type"] = fmt.Sprintf("*%s", typeString)
						} else {
							p["type"] = "I" + typeString
						}
					}
				}
			}
			cl.Properties[j] = p
		}

		for j, m := range cl.Methods {
			for i, p := range m.Params {
				p := formatTypes(p)
				m.Params[i] = p
			}
			m.Doc = formatDoc(m.Doc)

			if m.Return["type"] != nil {
				m.Return = formatTypes(m.Return)
				m.Return["doc"] = formatDoc(m.Return["doc"].(string))
			}

			cl.Methods[j] = m

		}
		for j, p := range cl.Constructor.Params {
			p := formatTypes(p)
			cl.Constructor.Params[j] = p
		}

		cl.Doc = formatDoc(cl.Doc)
		cl.Package = pkg

		buff := bytes.NewBufferString("")
		err := templates.ExecuteTemplate(buff, "struct.tmpl", cl)
		if err != nil {
			return ret, err
		}
		ret[cl.Name] = buff.Bytes()
	}
	return ret, nil
}

func parseComplexTypes(templates *template.Template) error {
	paths, _ := filepath.Glob(ELEMENTS + "elements.*.kmd.json")
	paths = append([]string{CORE}, paths...)
	for _, path := range paths {
		ctypes := getModel(path).ComplexTypes
		for _, ctype := range ctypes {
			pkg := "elements"
			if strings.HasSuffix(path, "core.kmd.json") {
				pkg = "core"
			}

			// Add in list
			CPXTYPES[ctype.Name] = pkg

			ctype.Doc = formatDoc(ctype.Doc)

			for i, p := range ctype.Properties {
				ctype.Properties[i] = formatTypes(p)
			}

			ctype.Package = pkg
			buff := bytes.NewBufferString("")
			err := templates.ExecuteTemplate(buff, "type.tmpl", ctype)
			if err != nil {
				return err
			}
			err = writePackageFile(pkg, ctype.Name, buff.Bytes())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writePackageFile(pkg, name string, content []byte) error {
	filename := fmt.Sprintf("%s/%s.go", pkg, name)
	log.Printf("Writing %s", filename)
	return os.WriteFile(filename, content, os.ModePerm)
}

//go:embed *.tmpl
var thisDir embed.FS

func main() {
	templates, err := (&template.Template{}).Funcs(funcMap).ParseFS(thisDir, "*.tmpl")
	if err != nil {
		log.Fatalf("unable to load templates: %s", err)
	}

	// Base interface
	buff := bytes.NewBufferString("")
	err = templates.ExecuteTemplate(buff, "imediaobject.tmpl", nil)
	if err != nil {
		log.Fatalf("unable to make base IMediaObject: %s", err)
	}
	err = writePackageFile("core", "base", buff.Bytes())
	if err != nil {
		log.Fatalf("unable to write base IMediaObject: %s", err)
	}

	// Perpare complexTypes to get the list
	err = parseComplexTypes(templates)
	if err != nil {
		log.Fatalf("error building complex types: %s", err)
	}

	// create base
	c := getModel(CORE).RemoteClasses
	coreclasses, err := parse(c, "core", templates)
	if err != nil {
		log.Fatalf("error building core: %s", err)
	}
	for name, class := range coreclasses {
		CLSTYPES[name] = "core"
		err = writePackageFile("core", name, []byte(class))
		if err != nil {
			log.Fatalf("error building core: %s", err)
		}
	}

	// make same for each interfaces
	err = getInterfaces(templates)
	if err != nil {
		log.Fatalf("error building interfaces: %s", err)
	}

	// copy LICENSE and NOTICE
	// NOTE: This code is only OK because these are small files. For large files,
	// this would be a very bad idea.
	license, err := os.ReadFile("build/kurento/LICENSE")
	if err != nil {
		log.Fatalf("error reading LICENSE file: %s\n", err)
	}
	err = os.WriteFile("kurento.LICENSE", license, 0644)
	if err != nil {
		log.Fatalf("error writing LICENSE file: %s\n", err)
	}
	notice, err := os.ReadFile("build/kurento/NOTICE")
	if err != nil {
		log.Fatalf("error reading NOTICE file: %s\n", err)
	}
	err = os.WriteFile("kurento.NOTICE", notice, 0644)
	if err != nil {
		log.Fatalf("error writing NOTICE file: %s\n", err)
	}
}
