package url_utils

import (
	"bufio"
	"os"
	"strings"
)

// GetPayloads generates or retrieves payloads for fuzzing.
func GetPayloads(valueFile string, values []string) ([]string, error) {
	var payloads []string

	// If a valueFile is provided, read it
	if valueFile != "" {
		file, err := os.Open(valueFile)
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
			payloads = append(payloads, strings.TrimSpace(scanner.Text()))
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	// Include values passed directly
	payloads = append(payloads, values...)

	return payloads, nil
}
