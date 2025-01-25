package utils

import (
	"k8s.io/apimachinery/pkg/api/errors"
)

func IsNotFound(err error) bool {
	return errors.IsNotFound(err)
}
