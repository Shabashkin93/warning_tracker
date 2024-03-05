package warning

import "bytes"

type WarningCreate struct {
	Id        string
	Branch    string
	Commit    string
	Count     int
	CreatedBy string
	CreatedAt string
	BuildLog  *bytes.Buffer
}

func NewWarningCreate() (out WarningCreate) {
	out.BuildLog = bytes.NewBuffer(nil)
	return
}

type WarningResponse struct {
	ID              string
	CountNewWarning int
	WarningList     []string
}
