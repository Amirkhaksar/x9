package generators

import (
	"fmt"
	"net/url"
	"os"
)

type Normal struct {
	Urls       []string
	Payloads   []string
	Parameters []string
}

func NewNormal(urls, payloads, parameters []string) *Normal {
	return &Normal{
		Urls:       urls,
		Payloads:   payloads,
		Parameters: parameters,
	}
}

func (n *Normal) replaceParameters(baseURL, payload string, chunkSize int) []string {
	var result []string
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
		return nil
	}

	queryParams := parsedURL.Query()
	queryParamsCopy := url.Values{}
	for k, v := range queryParams {
		queryParamsCopy[k] = v
	}

	if len(n.Parameters) > 1 {
		for param := range queryParamsCopy {
			delete(queryParams, param)
		}

		start := 0
		end := len(n.Parameters)
		step := chunkSize
		for x := start; x < end; x += step {
			for _, param := range n.Parameters[x:min(x+step, end)] {
				queryParams.Set(param, payload)
			}

			newQuery := queryParams.Encode()
			parsedURL.RawQuery = newQuery
			newURL := parsedURL.String()

			result = append(result, newURL)
			queryParams = parsedURL.Query()
		}
	} else {
		for param := range queryParams {
			queryParams.Set(param, payload)
		}

		newQuery := queryParams.Encode()
		parsedURL.RawQuery = newQuery
		newURL := parsedURL.String()

		result = append(result, newURL)
	}
	return result
}

func (n *Normal) normalMode(chunkSize int) []string {
	var result []string
	for _, url := range n.Urls {
		for _, payload := range n.Payloads {
			result = append(result, n.replaceParameters(url, payload, chunkSize)...)
		}
	}
	return result
}
