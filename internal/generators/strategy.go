package generators

import "errors"

func Strategy(generateStrategy string, urls, values, parameters []string, chunkSize int) ([]string, error) {
	var generateUrls []string
	switch generateStrategy {
	case "normal":
		normal := NewNormal(urls, values, parameters)
		generateUrls = append(generateUrls, normal.normalMode(chunkSize)...)
	case "ignore":
		ignore := NewIgnore(urls, values, parameters)
		generateUrls = append(generateUrls, ignore.ignoreMode(chunkSize)...)
	case "combine":
		combine := NewCombine(urls, values, parameters)
		generateUrls = append(generateUrls, combine.combineMode("suffix", chunkSize)...)
	case "all":
		normal := NewNormal(urls, values, parameters)
		generateUrls = append(generateUrls, normal.normalMode(chunkSize)...)
		ignore := NewIgnore(urls, values, parameters)
		generateUrls = append(generateUrls, ignore.ignoreMode(chunkSize)...)
		combine := NewCombine(urls, values, parameters)
		generateUrls = append(generateUrls, combine.combineMode("suffix", chunkSize)...)
	default:
		return nil, errors.New("invalid generate strategy provided")
	}
	return generateUrls, nil
}
