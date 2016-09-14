package utils

import "os"

func IsPathExist(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}

	return false
}

func IsPathNotExist(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return true
	}

	return false
}
