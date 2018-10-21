package parser

import (
	"fmt"
	"strconv"

	"bitbucket.org/kit-systems/dari/pkg/dict"
	"go.uber.org/zap"
)

func (p *Parser) updateDatabase() error {
	count := len(p.regs)
	p.journal.TotalCount = uint(count)
	if p.limit == 0 {
		total, err := strconv.Atoi(p.meta.Records)
		if err != nil {
			return err
		}
		if total != count {
			return fmt.Errorf("wrong count: %d out of %d", count, total)
		}
	}
	p.startBar(1, "Setting processing status...")
	err := p.db.MarkAsProcessing(p.regs)
	p.incrementBar()
	if err != nil {
		return err
	}
	p.finishBar("Done.")
	err = p.db.AddRegistryHooks()
	if err != nil {
		return err
	}
	p.startBar(count, "Inserting / updating records...")
	for _, r := range p.regs {
		if r.RegistryStatusID == uint(dict.ParsingFailed) {
			p.journal.FailedCount++
			continue
		}
		err := p.db.AddRegistry(r)
		p.incrementBar()
		if err != nil {
			p.journal.FailedCount++
			p.logger.Error("add registry", zap.Error(err))
		} else {
			if r.RegistryStatusID == uint(dict.Inserted) {
				p.journal.InsertedCount++
			}
			if r.RegistryStatusID == uint(dict.Updated) {
				p.journal.UpdatedCount++
			}
		}
	}
	p.finishBar(fmt.Sprintf("Total: %d, Errors: %d", count, p.journal.FailedCount))
	// p.startBar(1, "Setting deleted status...")
	// err = p.db.MarkAsDeleted(p.regs)
	// p.incrementBar()
	// if err != nil {
	// 	return err
	// }
	// p.finishBar("Done.")
	return nil
}
