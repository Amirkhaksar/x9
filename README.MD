# X9 - URL Parameter Fuzzer

<p align="center">
  <pre>
             ___
     __  __/ _ \
     \ \/ / (_) |
      >  < \__, |
     /_/\_\  /_/
  </pre>
</p>

X9 is a powerful command-line tool for generating URLs with parameter variations for web security testing and fuzzing. It offers multiple strategies for generating parameter combinations and integrates with the fallparams library for enhanced parameter discovery.

## Features

- Multiple URL generation strategies:
  - **Normal**: Replace existing parameters with fuzz values
  - **Ignore**: Ignore existing parameters and add new ones
  - **Combine**: Combine existing parameters with new ones
  - **All**: Use all strategies at once
- Support for parameter chunking to control output size
- Multiple input methods: single URL, URL list file, or stdin
- Parameter discovery using the fallparams library
- Customizable threading for performance optimization
- Headless browser support for JavaScript-heavy applications
- Proxy support for anonymity and traffic routing

## Installation

### Prerequisites

- Go 1.22.6 or higher

### Building from Source

```bash
# Clone the repository
git clone https://github.com/amirkhaksar/x9.git
cd x9

# Build the binary
go build -o x9 ./cmd/x9/main.go
```

## Usage

```
x9 -u <URL> -p <PARAMETER_LIST> -v <VALUES> [options]
```

### Basic Examples

```bash
# Generate URLs with a single URL, parameter list, and single value
x9 -u "https://example.com/" -p parameters.txt -v XSS

# Use multiple values
x9 -u "https://example.com/" -p parameters.txt -v "XSS,SQLi,SSRF"

# Use values from a file
x9 -u "https://example.com/" -p parameters.txt -vf values.txt

# Use a list of URLs
x9 -l urls.txt -p parameters.txt -v XSS
```

### Advanced Examples

```bash
# Use the normal strategy only
x9 -u "https://example.com/" -p parameters.txt -v XSS -gs normal

# Use chunking to control output size (add 10 parameters at a time)
x9 -u "https://example.com/" -p parameters.txt -v XSS -c 10

# Enable parameter discovery with fallparams
x9 -u "https://example.com/" -vf values.txt -ufp

# Use headless browser mode
x9 -u "https://example.com/" -vf values.txt -ufp -hl

# With proxy
x9 -u "https://example.com/" -p parameters.txt -v XSS -x http://127.0.0.1:8080
```

## Command Line Options

### Input Options

| Flag | Description |
|------|-------------|
| `-u, --url` | Single URL to edit |
| `-l, --list` | Path to a list of URLs to edit |
| `-p, --parameters` | Parameter wordlist to fuzz |
| `-c, --chunk` | Chunk size for parameter fuzzing (default: 15) |
| `-v, --value` | Values for parameters to fuzz, comma-separated |
| `-vf, --value_file` | List of values for parameters to fuzz |
| `-gs, --generate_strategy` | Generation strategy: normal, ignore, combine, or all (default: all) |
| `-vs, --value_strategy` | Value strategy: replace or suffix (default: replace) |
| `-o, --output` | Output results file |
| `-s, --silent` | Enable silent mode |
| `-t, --threads` | Number of threads for processing (default: 1) |
| `-ufp, --use_fallparams` | Use fallparams for parameter discovery |

### Fallparams Options

| Flag | Description |
|------|-------------|
| `-rd, --delay` | Request delay between each request in seconds |
| `-cm, --crawl` | Crawl pages to extract their parameters |
| `-d, --depth` | Maximum depth to crawl (default: 2) |
| `-ct, --crawl-duration` | Maximum duration to crawl the target |
| `-hl, --headless` | Discover parameters with headless browser |
| `-H, --header` | Header "Name: Value", separated by colon |
| `-xl, --max-length` | Maximum length of words (default: 30) |
| `-nl, --min-length` | Minimum length of words (default: 0) |
| `-X, --method` | HTTP method to use (default: GET) |
| `-x, --proxy` | Proxy URL (SOCKS5 or HTTP) |

## URL Generation Strategies

- **Normal**: Adds parameters to the URL, replacing any existing parameters with the same name.
- **Ignore**: Keeps existing parameters and adds new ones.
- **Combine**: Combines existing parameter values with new ones.
- **All**: Uses all three strategies.

## Integration with fallparams

X9 can use the [fallparams](https://github.com/ImAyrix/fallparams) library to discover parameters in web applications before fuzzing them. This is enabled with the `-ufp` flag and can be combined with crawling and headless browser options for more thorough discovery.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [fallparams](https://github.com/ImAyrix/fallparams) for parameter discovery functionality
- [go-tld](https://github.com/jpillora/go-tld) for URL parsing
- [goflags](https://github.com/projectdiscovery/goflags) for command-line argument parsing