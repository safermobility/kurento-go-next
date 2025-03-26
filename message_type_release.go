package kurento

type Release struct {
	SessionID string `json:"sessionId,omitempty"`
	ObjectID  string `json:"object"`
}

func BuildRelease(id string) *Release {
	return &Release{
		ObjectID: id,
	}
}

func (req *Release) GetMethod() string {
	return "release"
}

func (req *Release) SetSessionID(id string) {
	req.SessionID = id
}
