package url_utils

import (
	"bufio"
	"os"
	"strings"
)

// GetParameters reads parameters from a file and returns them as a slice of strings.
func GetParameters(parametersFile string) ([]string, error) {
	var parameters []string

	file, err := os.Open(parametersFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parameters = append(parameters, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parameters, nil
}
