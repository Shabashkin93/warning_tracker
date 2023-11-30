package status

import (
	"github.com/Shabashkin93/warning_tracker/internal/domain/status"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/pkg/buildinfo"
)

type service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *service {
	return &service{repos: repos}
}

func (s *service) GetAll() (dto *status.Status, code int, err error) {
	dto = &status.Status{}

	buildinfo := buildinfo.GetBuildInfo()
	dto.Version = buildinfo.Version
	dto.BuildTime = buildinfo.BuildTime
	dto.CommitHash = buildinfo.CommitHash
	dto.StatusDB, err = s.repos.Status.GetStatus()
	return
}
