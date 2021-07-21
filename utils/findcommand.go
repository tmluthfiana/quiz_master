package utils

import (
	"errors"
	"path/filepath"
)

var cmdpath string = "./commands/"

// FindCmd - Find a command inside ./commands CWD
func FindCmd(command []string) (fullpath string, err error) {
	for _, v := range command {
		found, err := filepath.Glob(cmdpath + v)
		if err != nil || len(found) == 0 {
			return "", errors.New("Command not found")
		}

		if len(found) != 0 {
			return found[0], nil
		}
	}
	return
}
