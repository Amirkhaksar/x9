package generators

import (
	"fmt"
	"net/url"
	"os"
)

type Combine struct {
	Urls       []string
	Payloads   []string
	Parameters []string
}

func NewCombine(urls, payloads, parameters []string) *Combine {
	return &Combine{
		Urls:       urls,
		Payloads:   payloads,
		Parameters: parameters,
	}
}

func (c *Combine) updateURLParameters(baseURL, payload string, chunkSize int) []string {
	var res []url.Values
	var result []string
	start := 0
	end := len(c.Parameters)
	step := chunkSize

	for i := start; i < end; i += step {
		parsedURL, err := url.Parse(baseURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
			return nil
		}

		queryParams := parsedURL.Query()
		x := i
		for _, param := range c.Parameters[x:min(x+step, end)] {
			queryParams.Set(param, payload)
		}
		res = append(res, queryParams)
	}

	for _, r := range res {
		parsedURL, err := url.Parse(baseURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
			return nil
		}

		parsedURL.RawQuery = r.Encode()
		updatedURL := parsedURL.String()
		result = append(result, updatedURL)
	}

	return result
}

func (c *Combine) replaceSuffix(baseURL, param, payload string, valueStrategy string, chunkSize int) []string {
	var result []string
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
		return nil
	}

	queryParams := parsedURL.Query()
	for _, values := range queryParams {
		for i := range values {
			if values[i] == param {
				if valueStrategy == "suffix" {
					values[i] = param + payload
				} else {
					values[i] = payload
				}
			}
		}
	}

	newQueryString := queryParams.Encode()
	newURL := parsedURL.String()
	parsedURL.RawQuery = newQueryString

	if len(c.Parameters) > 0 {
		for _, payload := range c.Payloads {
			result = append(result, c.updateURLParameters(newURL, payload, chunkSize)...)
			break
		}
	}

	return result
}

func (c *Combine) combineMode(valueStrategy string, chunkSize int) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", r)
		}
	}()

	var result []string
	for _, valUrl := range c.Urls {
		parsedURL, err := url.Parse(valUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
			continue
		}

		queryParams := parsedURL.Query()
		for _, values := range queryParams {
			for _, payload := range c.Payloads {
				result = append(result, c.replaceSuffix(valUrl, values[0], payload, valueStrategy, chunkSize)...)
			}
		}
	}

	return result
}
