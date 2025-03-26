package kurento

type Request interface {
	GetMethod() string
	SetSessionID(string)
}

type IMediaObjectBuilder interface {
	GetTypeName() string
}
