package warning

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	"github.com/Shabashkin93/warning_tracker/internal/project_errors"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"

	"github.com/microcosm-cc/bluemonday"
)

const (
	warningString = ": warning:"
	mainBranch    = "develop"
)

type usecase struct {
	repos  *repository.Repository
	sanity *bluemonday.Policy
	logger logging.Logger
	ctx    context.Context
}

func NewService(ctx context.Context, sanityCfg interface{}, repos *repository.Repository, logger logging.Logger) *usecase {
	sanity := sanityCfg.(*bluemonday.Policy)
	return &usecase{repos: repos, sanity: sanity, logger: logger, ctx: ctx}
}

func (s *usecase) sanitize(in *warning.WarningCreate) {
	in.Branch = s.sanity.Sanitize(in.Branch)
	in.Commit = s.sanity.Sanitize(in.Commit)
	in.CreatedBy = s.sanity.Sanitize(in.CreatedBy)
	in.CreatedAt = s.sanity.Sanitize(in.CreatedAt)
}

func (s *usecase) Create(in *warning.WarningCreate) (result warning.WarningResponse, err error) {
	uniqueHashStrings := make(map[string]string)
	s.sanitize(in)
	scanner := bufio.NewScanner(in.BuildLog)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), warningString) {
			hash := sha256.New()
			hash.Write([]byte(scanner.Text()))
			bs := hash.Sum(nil)
			encryptWarnStr := hex.EncodeToString(bs)
			if _, exist := uniqueHashStrings[encryptWarnStr]; !exist {
				uniqueHashStrings[encryptWarnStr] = scanner.Text()
			}
		}
	}

	if in.Branch == mainBranch {
		err = s.repos.Cache.DeleteAll()
		if err != nil {
			s.logger.Error(s.ctx, err.Error())
		}

		for hashString, warningString := range uniqueHashStrings {
			err = s.repos.Cache.Set(hashString, "")
			if err != nil {
				s.logger.Error(s.ctx, err.Error())
			}
			result.WarningList = append(result.WarningList, warningString)
		}
	} else {
		for hashString, warningString := range uniqueHashStrings {
			_, err = s.repos.Cache.Get(hashString)
			if err == project_errors.CacheKeyNotFound {
				result.WarningList = append(result.WarningList, warningString)
			} else if err != nil {
				s.logger.Error(s.ctx, err.Error())
				return
			}
		}
	}

	in.Count = len(result.WarningList)
	result.CountNewWarning = in.Count
	result.ID, err = s.repos.Warning.Create(in)
	if err != nil {
		s.logger.Error(s.ctx, err.Error())
		return
	}

	return
}
