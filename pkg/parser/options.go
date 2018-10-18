package parser

import (
	"github.com/gocolly/colly"
	collydebug "github.com/gocolly/colly/debug"
	"go.uber.org/zap"
)

// Options is a configuration for parser container.
type Options struct {
	Logger *zap.Logger
	Debug  bool
	Limit  uint
}

// Pages is a configuration for dari.kz urls to be parsed.
type Pages struct {
	Main     string
	Registry string
	Grid     string
	Print    string
}

// DefaultLogger returns basic logger instance.
func DefaultLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// DefaultPages returns dari.kz urls.
func DefaultPages() *Pages {
	return &Pages{
		Main:     "http://www.ndda.kz/register.php/mainpage",
		Registry: "/reestr/lang/ru",
		Grid:     "/getReestr",
		Print:    "/print/",
	}
}

// DefaultCollector returns parser collector.
func DefaultCollector(debug bool) *colly.Collector {
	var c *colly.Collector
	var opts []func(*colly.Collector)
	if debug {
		opts = append(opts, colly.Debugger(&collydebug.LogDebugger{}))
	}
	c = colly.NewCollector(opts...)
	return c
}
