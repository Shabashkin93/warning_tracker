package warning_test

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	domain_warning "github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	"github.com/Shabashkin93/warning_tracker/internal/project_errors"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/internal/usecase/warning"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
	logger "github.com/Shabashkin93/warning_tracker/pkg/logging/slog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	sanity := bluemonday.UGCPolicy()

	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("Failed get config")
		os.Exit(1)
	}

	out := os.Stdout
	loggerEnt := logger.NewLogger(cfg.LogLevel, out)
	defer out.Close()
	logger := logging.NewLogger(loggerEnt)

	repos := repository.Repository{}
	cache := new(repository.MockCache)
	cache.On("Set", "94aa402e73bf057b0fe86ac42af4e13dac469826a2db3c4cdac0ab36d1b8ec92", "").Return(nil)
	cache.On("Get", "94aa402e73bf057b0fe86ac42af4e13dac469826a2db3c4cdac0ab36d1b8ec92").Return("", project_errors.CacheKeyNotFound)
	cache.On("Delete", "94aa402e73bf057b0fe86ac42af4e13dac469826a2db3c4cdac0ab36d1b8ec92").Return(nil)

	cache.On("Set", "1760b5f92c5e878307016c04f821e23abef8c9843f4412188b85f6dd0ea0fe5b", "").Return(nil)
	cache.On("Get", "1760b5f92c5e878307016c04f821e23abef8c9843f4412188b85f6dd0ea0fe5b").Return("", project_errors.CacheKeyNotFound)
	cache.On("Delete", "1760b5f92c5e878307016c04f821e23abef8c9843f4412188b85f6dd0ea0fe5b").Return(nil)

	cache.On("DeleteAll").Return(nil)
	cache.On("Shutdown").Return(nil)
	repos.Cache = cache

	warn := new(repository.MockWarning)
	createIn := &domain_warning.WarningCreate{Branch: "nodevelop",
		Commit:    "beffe2b9a727c481c8a4896edb1783a054ac084c",
		CreatedBy: "Shabashkin",
		CreatedAt: "2024-02-19T19:36:10.103Z"}

	repos.Warning = warn

	usecase := warning.NewService(ctx, sanity, &repos, logger)

	var buf bytes.Buffer
	inBuf := `/main_cli.c:2668:30: warning: passing argument 2 of 'main_signal_set' from incompatible pointer type [-Wincompatible-pointer-types]
/usr/include/features.h:184:3: warning: #warning "_BSD_SOURCE and _SVID_SOURCE are deprecated, use _DEFAULT_SOURCE" [-Wcpp]`
	buf.WriteString(inBuf)

	createIn.BuildLog = &buf
	warn.On("Create", createIn).Return("2eda5cf6-438b-4d27-9e62-e2b667ab831d", nil)

	outStruct, err := usecase.Create(createIn)
	assert.Equal(t, err, nil, "not equal happy err")
	assert.Equal(t, outStruct.ID, "2eda5cf6-438b-4d27-9e62-e2b667ab831d", "not equal happy ID")
	assert.Equal(t, outStruct.CountNewWarning, 2, "not equal happy Count")
	cache.On("Get", "94aa402e73bf057b0fe86ac42af4e13dac469826a2db3c4cdac0ab36d1b8ec92").Return("", nil)
	cache.On("Get", "1760b5f92c5e878307016c04f821e23abef8c9843f4412188b85f6dd0ea0fe5b").Return("", nil)

	outStruct, err = usecase.Create(createIn)
	assert.Equal(t, err, nil, "not equal happy err")
	assert.Equal(t, outStruct.ID, "2eda5cf6-438b-4d27-9e62-e2b667ab831d", "not equal happy ID")
	assert.Equal(t, outStruct.CountNewWarning, 0, "not equal happy Count")
}
