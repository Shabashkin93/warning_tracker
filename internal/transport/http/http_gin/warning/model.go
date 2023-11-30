package warning

import (
	"io"
	"mime/multipart"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	"github.com/pkg/errors"
)

type httpWarningCreate struct {
	Commit       string                `form:"commit"`
	Branch       string                `form:"branch"`
	CreatedBy    string                `form:"created_by"`
	CreatedAt    string                `form:"created_at"`
	BuildLogFile *multipart.FileHeader `form:"build_log"`
}

type httpWarningResponse struct {
	ID              string   `json:"id"`
	CountNewWarning int      `json:"count"`
	WarningList     []string `json:"newWarningList"`
}

func warningCreateHttpToDomain(in *httpWarningCreate, out *warning.WarningCreate) (err error) {
	if in == nil {
		err = errors.Errorf("Invalid input data")
		return
	}

	if out == nil {
		err = errors.Errorf("Internal server error")
		return
	}

	out.Branch = in.Branch
	out.Commit = in.Commit
	out.CreatedAt = in.CreatedAt
	out.CreatedBy = in.CreatedBy

	if in.BuildLogFile == nil {
		err = errors.Errorf("Invalid input data")
		return
	}

	file, err := in.BuildLogFile.Open()
	if err != nil {
		err = errors.Errorf("Failed open input file")
		return
	}
	defer file.Close()

	if _, err = io.Copy(out.BuildLog, file); err != nil {
		err = errors.Errorf("Failed copy file content")
		return
	}

	return
}

func warningCreateDomainToHttp(in *warning.WarningResponse, out *httpWarningResponse) {
	out.ID = in.ID
	out.CountNewWarning = in.CountNewWarning
	out.WarningList = append(out.WarningList, in.WarningList...)
}
