package kurento

type Create struct {
	SessionID  string              `json:"sessionId,omitempty"`
	Type       string              `json:"type"`
	Params     IMediaObjectBuilder `json:"constructorParams"`
	Properties any                 `json:"properties"`
}

func BuildCreate(b IMediaObjectBuilder) *Create {
	return &Create{
		Type:   b.GetTypeName(),
		Params: b,
	}
}

func (req *Create) GetMethod() string {
	return "create"
}

func (req *Create) SetSessionID(id string) {
	req.SessionID = id
}
