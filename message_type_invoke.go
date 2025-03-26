package kurento

type InvokeParams interface {
	OperationName() string
}

type Invoke struct {
	SessionID string       `json:"sessionId,omitempty"`
	ObjectID  string       `json:"object"`
	Operation string       `json:"operation"`
	Params    InvokeParams `json:"operationParams"`
}

func BuildInvoke(id string, params InvokeParams) *Invoke {
	return &Invoke{
		ObjectID:  id,
		Operation: params.OperationName(),
		Params:    params,
	}
}

func (req *Invoke) GetMethod() string {
	return "invoke"
}

func (req *Invoke) SetSessionID(id string) {
	req.SessionID = id
}
