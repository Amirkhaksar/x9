package url_utils

import (
	"fmt"
	"strings"

	"github.com/jpillora/go-tld"
)

// CleanURL cleans and normalizes the given URL by removing unnecessary components.
func CleanURL(rawURL string) (string, error) {
	var fullURL string
	var sanitizedURL = strings.TrimSpace(rawURL)
	// Extract the URL components
	ext, err := tld.Parse(sanitizedURL)
	if err != nil {
		return "", err
	}

	var extURL string
	if ext.Subdomain != "" {
		extURL = fmt.Sprintf("%s.%s.%s", ext.Subdomain, ext.Domain, ext.TLD)
	} else {
		extURL = fmt.Sprintf("%s.%s", ext.Domain, ext.TLD)
	}

	// Ensure the URL starts with http:// or https://
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		fullURL = fmt.Sprintf("https://%s", rawURL)
	} else {
		fullURL = sanitizedURL
	}

	// Check if the URL contains query parameters
	if !strings.Contains(fullURL, "?") && !strings.Contains(fullURL, "%3F") && !strings.Contains(fullURL, "%3f") {
		if strings.HasPrefix(fullURL, "https:") {
			if fullURL == fmt.Sprintf("https://%s", extURL) {
				if !strings.HasSuffix(fullURL, "/") {
					fullURL = fmt.Sprintf("https://%s/", extURL)
				}
			}
		} else {
			if fullURL == fmt.Sprintf("http://%s", extURL) {
				if !strings.HasSuffix(fullURL, "/") {
					fullURL = fmt.Sprintf("http://%s/", extURL)
				}
			}
		}
	}

	return fullURL, nil
}
