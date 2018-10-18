package parser

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

func (p *Parser) updateDatabase() error {
	errors := 0
	count := len(p.regs)
	if p.limit == 0 {
		total, err := strconv.Atoi(p.meta.Records)
		if err != nil {
			return err
		}
		if total != count {
			return fmt.Errorf("wrong count: %d out of %d", count, total)
		}
	}
	p.startBar(count, "Inserting / updating records...")
	for _, r := range p.regs {
		err := p.db.AddRegistry(r)
		p.incrementBar()
		if err != nil {
			errors++
			p.logger.Error("add registry", zap.Error(err))
		}
	}
	p.finishBar(fmt.Sprintf("Total: %d, Errors: %d", count, errors))
	return nil
}
