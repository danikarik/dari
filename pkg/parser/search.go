package parser

import "go.uber.org/zap"

func (p *Parser) beginSearch() error {
	if !p.autoSearch {
		return nil
	}
	p.startBar(1, "Searching products...")
	prods, err := p.db.FindProducts()
	p.incrementBar()
	if err != nil {
		return err
	}
	p.finishBar("Done.")
	count := len(prods)
	p.startBar(count, "Searching matches...")
	for _, prod := range prods {
		err = p.db.FindRegistry(prod)
		if err != nil {
			p.logger.Error("find registry", zap.Error(err))
		}
		p.incrementBar()
	}
	p.finishBar("Done.")
	p.startBar(1, "Creating notifications...")
	err = p.db.CreateNotifications(prods)
	p.incrementBar()
	if err != nil {
		return err
	}
	p.finishBar("Done.")
	return nil
}
