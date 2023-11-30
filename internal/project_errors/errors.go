package project_errors

import "github.com/pkg/errors"

var CacheKeyNotFound = errors.Errorf("Key not found")
