package generators

import (
	"fmt"
	"net/url"
	"os"
)

type Ignore struct {
	Urls       []string
	Payload    []string
	Parameters []string
}

func NewIgnore(urls, payload, wordlist []string) *Ignore {
	return &Ignore{
		Urls:       urls,
		Payload:    payload,
		Parameters: wordlist,
	}
}

func (i *Ignore) updateURLParameters(baseURL, payload string, chunkSize int) []string {
	var result []string
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
		return nil
	}

	queryParams := parsedURL.Query()

	for _, param := range i.Parameters {
		if _, exists := queryParams[param]; exists {
			delete(queryParams, param)
		}
	}

	start := 0
	end := len(i.Parameters)
	step := chunkSize

	for x := start; x < end; x += step {
		for _, param := range i.Parameters[x:min(x+step, end)] {
			queryParams.Set(param, payload)
		}

		parsedURL.RawQuery = queryParams.Encode()
		updatedURL := parsedURL.String()
		result = append(result, updatedURL)

		queryParams = parsedURL.Query()
	}
	return result
}

func (i *Ignore) ignoreMode(chunkSize int) []string {
	if len(i.Parameters) <= 1 && i.Parameters[0] == "" {
		fmt.Println("Please enter your parameter list as a text file")
		return nil
	}

	var result []string
	for _, valURL := range i.Urls {
		for _, payload := range i.Payload {
			result = append(result, i.updateURLParameters(valURL, payload, chunkSize)...)
		}
	}
	return result
}
