package speedtest

import (
	"flag"
	"time"
)

type Opts struct {
	SpeedInBytes bool
	Quiet        bool
	List         bool
	Server       ServerID
	Interface    string
	Timeout      time.Duration
	Secure       bool
	Help         bool
	Version      bool
	Download     bool
	Upload       bool
	Latency      bool
	Verbose      bool
}

func ParseOpts() *Opts {
	opts := new(Opts)

	flag.BoolVar(&opts.SpeedInBytes, "bytes", false, "Display values in bytes instead of bits")
	flag.BoolVar(&opts.Quiet, "quiet", false, "Suppress all output (overrides -verbose)")
	flag.BoolVar(&opts.List, "list", false, "Display a list of speedtest.net servers sorted by distance")
	flag.BoolVar(&opts.Latency, "latency", false, "Run a latency test only")
	flag.BoolVar(&opts.Download, "download", false, "Run a download test only (also runs latency test)")
	flag.BoolVar(&opts.Upload, "upload", false, "Run an upload test only (also runs latency test)")
	flag.Uint64Var((*uint64)(&opts.Server), "server", 0, "Specify a server ID to test against")
	flag.StringVar(&opts.Interface, "interface", "", "IP address of network interface to bind to")
	flag.DurationVar(&opts.Timeout, "timeout", 10*time.Second, "HTTP timeout duration")
	flag.BoolVar(&opts.Secure, "secure", true,
		"Use HTTPS instead of HTTP when communicating with speedtest.net operated servers")
	flag.BoolVar(&opts.Help, "help", false, "Show usage information and exit")
	flag.BoolVar(&opts.Version, "version", false, "Show the version number and exit")
	flag.BoolVar(&opts.Verbose, "verbose", false, "Show debugging and extraneous information")

	flag.Parse()

	// Quiet overrides Verbose
	if opts.Quiet {
		opts.Verbose = false
	}

	// If neither are set, then both are set
	if !opts.Download && !opts.Upload && !opts.Latency {
		opts.Latency = true
		opts.Download = true
		opts.Upload = true
	}

	return opts
}
