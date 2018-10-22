package parser

import (
	"time"

	"github.com/volatiletech/null"
)

func (p *Parser) startProcess(t time.Time) error {
	p.journal.StartedAt = t
	p.journal.TotalCount = 0
	p.journal.FailedCount = 0
	p.journal.InsertedCount = 0
	p.journal.UpdatedCount = 0
	p.journal.DeletedCount = 0
	return p.db.StartProcess(p.journal)
}

func (p *Parser) updateCounters() error {
	return p.db.UpdateCounter(p.journal)
}

func (p *Parser) finishProcess(t time.Time) error {
	p.journal.FinishedAt = null.TimeFrom(t)
	return p.db.FinishProcess(p.journal)
}
