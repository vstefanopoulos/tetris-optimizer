package tetrisopt

import (
	"errors"
	"strings"
)

func handleArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("Not enough arguments")
	}
	if !strings.HasSuffix(args[1], ".txt") {
		return "", errors.New("invalid filename")
	}
	return "test_files/" + args[1], nil
}
