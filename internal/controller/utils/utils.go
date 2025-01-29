package utils

import (
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func IsNotFound(err error) bool {
	return errors.IsNotFound(err)
}

func IsHTTPIngressPathExist(paths []networkingv1.HTTPIngressPath, path string) bool {
	for _, value := range paths {
		if value.Path == path {
			return true
		}
	}
	return false
}
