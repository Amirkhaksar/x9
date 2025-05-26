package options

import (
	"github.com/projectdiscovery/goflags"
	"time"
)

type X9EngineOptions struct {
	URL               string
	URLList           string
	Parameters        string
	Chunk             int
	Value             goflags.StringSlice
	ValueFile         string
	GenerateStrategy  string
	ValueStrategy     string
	Output            string
	Silent            bool
	Threads           int
	Mode              string
	Help              bool
	UseFallParams     bool
	Delay             int
	CrawlMode         bool
	MaxDepth          int
	CrawlDuration     time.Duration
	Headless          bool
	CustomHeaders     goflags.StringSlice
	MaxLength         int
	MinLength         int
	RequestHttpMethod string
	ProxyUrl          string
}
