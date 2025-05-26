package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/projectdiscovery/goflags"

	"github.com/amirkhaksar/x9/internal/options"
	"github.com/amirkhaksar/x9/internal/x9engine"
)

// Version information - will be injected during build by goreleaser
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printBanner() {
	banner := `
             ___
     __  __/ _ \
     \ \/ / (_) |
      >  < \__, |
     /_/\_\  /_/

 `
	fmt.Printf("%sVersion: %s (%s) - %s\n", banner, version, commit, date)
}

var (
	myOptions = &options.X9EngineOptions{}
)

// ParseArgs parses the command-line arguments
func ParseArgs() *options.X9EngineOptions {
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Generate URLs for parameter fuzzing.")
	createGroup(flagSet, "input", "Input",
		flagSet.StringVarP(&myOptions.URL, "url", "u", "", "Single URL to edit."),
		flagSet.StringVarP(&myOptions.URLList, "list", "l", "", "Path to a list of URLs to edit."),
		flagSet.StringVarP(&myOptions.Parameters, "parameters", "p", "", "Parameter wordlist to fuzz."),
		flagSet.IntVarP(&myOptions.Chunk, "chunk", "c", 15, "Chunk size for parameter fuzzing."),
		flagSet.VarP(&myOptions.Value, "value", "v", "Values for parameters to fuzz, comma-separated."),
		flagSet.StringVarP(&myOptions.ValueFile, "value_file", "vf", "", "List of values for parameters to fuzz."),
		flagSet.StringVarP(&myOptions.GenerateStrategy, "generate_strategy", "gs", "all", "Generate strategy: normal, ignore, combine, or all."),
		flagSet.StringVarP(&myOptions.ValueStrategy, "value_strategy", "vs", "replace", "Value strategy: replace or suffix."),
		flagSet.StringVarP(&myOptions.Output, "output", "o", "", "Output results file."),
		flagSet.BoolVarP(&myOptions.Silent, "silent", "s", false, "Enable silent mode."),
		flagSet.IntVarP(&myOptions.Threads, "threads", "t", 1, "Number of threads for processing."),
		flagSet.BoolVarP(&myOptions.UseFallParams, "use_fallparams", "ufp", false, "Use fallparams or not."),
	)
	createGroup(flagSet, "fallparams", "fallparams options",
		flagSet.IntVarP(&myOptions.Delay, "delay", "rd", 0, "Request delay between each request in seconds"),
		flagSet.BoolVarP(&myOptions.CrawlMode, "crawl", "cm", false, "Crawl pages to extract their parameters"),
		flagSet.IntVarP(&myOptions.MaxDepth, "depth", "d", 2, "maximum depth to crawl"),
		flagSet.DurationVarP(&myOptions.CrawlDuration, "crawl-duration", "ct", 0, "maximum duration to crawl the target"),
		flagSet.BoolVarP(&myOptions.Headless, "headless", "hl", false, "Discover parameters with headless browser"),
		flagSet.VarP(&myOptions.CustomHeaders, "header", "H", "Header `\"Name: Value\"`, separated by colon. Multiple -H flags are accepted."),
		flagSet.IntVarP(&myOptions.MaxLength, "max-length", "xl", 30, "Maximum length of words"),
		flagSet.IntVarP(&myOptions.MinLength, "min-length", "nl", 0, "Minimum length of words"),
		flagSet.StringVarP(&myOptions.RequestHttpMethod, "method", "X", "GET", "HTTP method to use"),
		flagSet.StringVarP(&myOptions.ProxyUrl, "proxy", "x", "", "Proxy URL (SOCKS5 or HTTP). For example: http://127.0.0.1:8080 or socks5://127.0.0.1:8080"),
	)
	help := flag.Bool("help", false, "Display help message.")
	// Parse flags
	err := flagSet.Parse()
	if err != nil {
		return nil
	}

	// Print help and exit if --help is set
	if *help {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return myOptions
}

func main() {
	// Parse command-line arguments
	arguments := ParseArgs()

	// Print banner if not in silent mode
	if !arguments.Silent {
		printBanner()
	}

	// Check if valueFile or value is provided
	if arguments.ValueFile == "" && len(arguments.Value) == 0 {
		_, err := fmt.Fprintln(os.Stderr, "Error: Either valueFile or value must be provided.")
		if err != nil {
			return
		}
		os.Exit(1)
	}

	// Create X9EngineOptions struct
	generateURLsArgs := options.X9EngineOptions{
		URL:               arguments.URL,
		URLList:           arguments.URLList,
		Parameters:        arguments.Parameters,
		ValueFile:         arguments.ValueFile,
		Chunk:             arguments.Chunk,
		Value:             arguments.Value,
		GenerateStrategy:  arguments.GenerateStrategy,
		Output:            arguments.Output,
		UseFallParams:     arguments.UseFallParams,
		Threads:           arguments.Threads,
		Delay:             arguments.Delay,
		CrawlMode:         arguments.CrawlMode,
		MaxDepth:          arguments.MaxDepth,
		CrawlDuration:     arguments.CrawlDuration,
		Headless:          arguments.Headless,
		CustomHeaders:     arguments.CustomHeaders,
		MaxLength:         arguments.MaxLength,
		MinLength:         arguments.MinLength,
		RequestHttpMethod: arguments.RequestHttpMethod,
		ProxyUrl:          arguments.ProxyUrl,
	}

	// Generate URLs using the parsed arguments
	completeURLs, err := x9engine.GenerateURLs(generateURLsArgs)

	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error generating URLs: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	// Print the generated URLs
	for _, url := range completeURLs {
		fmt.Println(url)
	}
}

func createGroup(flagSet *goflags.FlagSet, groupName, description string, flags ...*goflags.FlagData) {
	flagSet.SetGroup(groupName, description)
	for _, currentFlag := range flags {
		currentFlag.Group(groupName)
	}
}
