package parser

import (
	"go.uber.org/zap"
)

// Parser is a basic struct for parser.
type Parser struct {
	logger *zap.Logger
}

// New creates new parser instance.
func New(opts Options) (*Parser, error) {
	var p Parser
	if opts.Logger != nil {
		p.logger = opts.Logger
	} else {
		lg, err := DefaultLogger()
		if err != nil {
			return nil, err
		}
		p.logger = lg
	}
	return &p, nil
}

// Run starts parser.
func (p *Parser) Run() error {
	p.logger.Info("Running...")
	defer p.logger.Sync()
	return nil
}
