package gerrit

type CommonFields struct {
	Type           string  `json:"type"`
	EventCreatedOn float64 `json:"eventCreatedOn"`
	RepoURL        string  `json:"repoURL,omitempty"`
}

type Submitter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type RefUpdate struct {
	OldRev  string `json:"oldRev"`
	NewRev  string `json:"newRev"`
	RefName string `json:"refName"`
	Project string `json:"project"`
}

// Gerrit event types

type ProjectCreated struct {
	ProjectName string `json:"projectName"`
	HeadName    string `json:"headName"`
	CommonFields
}
type ProjectHeadUpdated struct {
	ProjectName string `json:"projectName"`
	OldHead     string `json:"oldHead"`
	NewHead     string `json:"newHead"`
	CommonFields
}
type RefUpdated struct {
	Submitter Submitter `json:"submitter"`
	RefUpdate RefUpdate `json:"refUpdate"`
	CommonFields
}
