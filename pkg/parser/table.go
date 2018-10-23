package parser

import (
	"github.com/hashicorp/go-multierror"
	"context"
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"bitbucket.org/kit-systems/dari/pkg/compare"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
)

func (p *Parser) beforeRegistryUpsert() models.RegistryHook {
	return func(ctx context.Context, exec boil.ContextExecutor, r *models.Registry) error {
		exists, err := models.FindRegistry(ctx, exec, r.ID)
		if err != nil {
			r.RegistryStatusID = uint(dict.Inserted)
			return nil
		}
		r.RegistryStatusID = uint(dict.Updated)
		return p.compareRegistries(exists, r)
	}
}

func (p *Parser) compareRegistries(old *models.Registry, new *models.Registry) error {
	var fss models.RegistryFieldStatSlice
	// Compare titles.
	if !compare.String(old.Title, new.Title) {
		fss = append(fss, &models.RegistryFieldStat{
			RegistryID:        null.NewUint(old.ID, true),
			RegistryJournalID: p.journal.ID,
			Label:             dict.RegistryNameLabel,
			OldValue:          null.StringFrom(old.Title),
			NewValue:          null.StringFrom(new.Title),
			CreatedAt:         time.Now(),
		})
	}
	// Compare numbers.
	if !compare.String(old.Number, new.Number) {
		fss = append(fss, &models.RegistryFieldStat{
			RegistryID:        null.NewUint(old.ID, true),
			RegistryJournalID: p.journal.ID,
			Label:             dict.RegistryNumberLabel,
			OldValue:          null.StringFrom(old.Number),
			NewValue:          null.StringFrom(new.Number),
			CreatedAt:         time.Now(),
		})
	}
	// Compare issue dates.
	if !compare.Date(old.IssueDate, new.IssueDate) {
		fss = append(fss, &models.RegistryFieldStat{
			RegistryID:        null.NewUint(old.ID, true),
			RegistryJournalID: p.journal.ID,
			Label:             dict.RegistryIssueDateLabel,
			OldValue:          null.StringFrom(old.IssueDate.Format("2006-01-02")),
			NewValue:          null.StringFrom(new.IssueDate.Format("2006-01-02")),
			CreatedAt:         time.Now(),
		})
	}
	// Compare durations.
	if !compare.Int(old.Duration, new.Duration) {
		fss = append(fss, &models.RegistryFieldStat{
			RegistryID:        null.NewUint(old.ID, true),
			RegistryJournalID: p.journal.ID,
			Label:             dict.RegistryDurationLabel,
			OldValue:          null.StringFrom(strconv.Itoa(old.Duration)),
			NewValue:          null.StringFrom(strconv.Itoa(new.Duration)),
			CreatedAt:         time.Now(),
		})
	}
	// Compare expire dates.
	if !compare.Date(old.ExpireDate.Time, new.ExpireDate.Time) {
		oldDt := ""
		newDt := ""
		if old.ExpireDate.Ptr() != nil {
			oldDt = old.ExpireDate.Time.Format("2006-01-02")
		}
		if new.ExpireDate.Ptr() != nil {
			newDt = new.ExpireDate.Time.Format("2006-01-02")
		}
		fss = append(fss, &models.RegistryFieldStat{
			RegistryID:        null.NewUint(old.ID, true),
			RegistryJournalID: p.journal.ID,
			Label:             dict.RegistryExpireDateLabel,
			OldValue:          null.StringFrom(oldDt),
			NewValue:          null.StringFrom(newDt),
			CreatedAt:         time.Now(),
		})
	}
	// Collecting changes.
	var result *multierror.Error
	for _, fs := range fss {
		if err := p.db.InsertFieldStat(fs); err != nil {
			result.Errors = append(result.Errors, err)
		}
	}
	// Result.
	return result.ErrorOrNil()
}

func (p *Parser) logFailed(r *models.Registry, err error) error {
	fs := &models.RegistryFieldStat{
		RegistryID:        null.NewUint(r.ID, true),
		RegistryJournalID: p.journal.ID,
		Label:             dict.RegistryParsingLabel,
		OldValue:          null.StringFrom(""),
		NewValue:          null.StringFrom(""),
		ErrorMSG: 		   null.StringFrom(err.Error()),
		CreatedAt:         time.Now(),
	}
	return p.db.InsertFieldStat(fs)
}

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
	err = p.db.AddRegistryHooks(boil.BeforeUpsertHook, p.beforeRegistryUpsert())
	if err != nil {
		return err
	}
	p.startBar(count, "Inserting / updating records...")
	for _, r := range p.regs {
		if r.RegistryStatusID == uint(dict.ParsingFailed) {
			p.journal.FailedCount++
			lerr := p.logFailed(r, fmt.Errorf("parsing failed"))
			if lerr != nil {
				p.logger.Error("field stat", zap.Error(lerr))
			}
			continue
		}
		err := p.db.AddRegistry(r)
		p.incrementBar()
		if err != nil {
			p.journal.FailedCount++
			lerr := p.logFailed(r, err)
			if lerr != nil {
				p.logger.Error("failed records", zap.Error(lerr))
			}
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
	p.startBar(1, "Setting deleted status...")
	err = p.db.MarkAsDeleted(p.regs)
	p.incrementBar()
	if err != nil {
		return err
	}
	p.finishBar("Done.")
	p.startBar(1, "Updating stats...")
	err = p.db.UpdateCounter(p.journal)
	if err != nil {
		return err
	}
	p.incrementBar()
	p.finishBar("Done.")
	return nil
}
