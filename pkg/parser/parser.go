package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/kit-systems/dari/pkg/database"
	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"bitbucket.org/kit-systems/dari/pkg/urlhelper"
	"github.com/gocolly/colly"
	"github.com/hashicorp/go-multierror"
	"github.com/volatiletech/null"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"gopkg.in/cheggaaa/pb.v1"
)

// Parser is a basic struct for parser.
type Parser struct {
	logger    *zap.Logger
	db        *database.DB
	limit     uint
	pages     *Pages
	collector *colly.Collector
	helper    *urlhelper.URLHelper
	meta      *Meta
	regs      models.RegistrySlice
	bar       *pb.ProgressBar
	journal   *models.RegistryJournal
}

// New creates new parser instance.
func New(db *database.DB, opts Options) (*Parser, error) {
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
	p.limit = opts.Limit
	p.pages = DefaultPages()
	p.collector = DefaultCollector(opts.Debug)
	p.helper = urlhelper.New()
	p.meta = &Meta{}
	p.regs = make(models.RegistrySlice, 0)
	p.db = db
	p.journal = &models.RegistryJournal{}
	return &p, nil
}

// Run starts parser.
func (p *Parser) Run() error {
	var err error

	defer func() {
		if derr := p.logger.Sync(); derr != nil {
			p.logger.Error("sync", zap.Error(derr))
		}
	}()

	err = p.startProcess(time.Now())
	if err != nil {
		return err
	}

	defer func() {
		if terr := p.finishProcess(time.Now()); terr != nil {
			p.logger.Error("process", zap.Error(terr))
		}
	}()

	err = p.visitRegistry()
	if err != nil {
		return err
	}

	err = p.visitGrid()
	if err != nil {
		return err
	}

	err = p.collectPages()
	if err != nil {
		return err
	}

	err = p.collectRecords()
	if err != nil {
		return err
	}

	err = p.updateDatabase()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) startBar(count int, label string) {
	fmt.Println(label)
	p.bar = pb.StartNew(count)
	p.bar.ShowElapsedTime = true
	p.bar.ShowFinalTime = true
}

func (p *Parser) incrementBar() {
	p.bar.Increment()
	time.Sleep(time.Millisecond)
}

func (p *Parser) finishBar(label string) {
	p.bar.FinishPrint(label)
}

func (p *Parser) addRegistry(r *models.Registry) {
	p.regs = append(p.regs, r)
}

func (p *Parser) onError(r *colly.Response, err error) {
	p.logger.Error(
		"collector on error",
		zap.String("url", p.helper.Get()),
		zap.Int("status code", r.StatusCode),
		zap.Error(err),
	)
}

func (p *Parser) visitRegistry() error {
	err := p.helper.Update(p.pages.Main + p.pages.Registry)
	if err != nil {
		return err
	}
	p.helper.Add(p.helper.Map("ReestrForm[reg_type]", "2"))
	p.collector.OnError(func(r *colly.Response, err error) {
		p.onError(r, err)
	})
	return p.collector.Visit(p.helper.Get())
}

func (p *Parser) visitGrid() error {
	err := p.helper.Update(p.pages.Main + p.pages.Grid)
	if err != nil {
		return err
	}
	p.helper.Add(
		p.helper.Map("rows", "300"),
		p.helper.Map("page", "1"),
		p.helper.Map("sidx", "t.reg_date"),
		p.helper.Map("sord", "desc"),
	)
	p.collector.OnError(func(r *colly.Response, err error) {
		p.onError(r, err)
	})
	p.collector.OnResponse(func(r *colly.Response) {
		xml.Unmarshal(r.Body, p.meta)
	})
	return p.collector.Visit(p.helper.Get())
}

func (p *Parser) collectPages() error {
	pages, err := strconv.Atoi(p.meta.Total)
	if err != nil {
		return err
	}

	lastIndex := pages
	if p.limit == 1 {
		return nil
	}
	if p.limit > 0 && p.limit <= uint(lastIndex) {
		lastIndex = int(p.limit)
	}
	p.startBar(pages, "Collecting meta data...")
	p.incrementBar()
	for i := 2; i <= lastIndex; i++ {
		err := p.visitPage(i)
		p.incrementBar()
		if err != nil {
			return err
		}
	}
	p.finishBar(fmt.Sprintf("Total: %d, Records: %d", pages, len(p.meta.Rows)))

	return nil
}

func (p *Parser) collectRecords() error {
	total := len(p.meta.Rows)
	p.startBar(total, "Collecting records...")
	for _, r := range p.meta.Rows {
		rec, err := p.visitPrint(r.ID)
		p.incrementBar()
		if err != nil {
			rec.RegistryStatusID = uint(dict.ParsingFailed)
			p.journal.FailedCount++
			p.logger.Error(
				"registry",
				zap.String("link", rec.Link),
				zap.String("number", rec.Number),
				zap.Error(err),
			)
		}
	}
	p.finishBar(fmt.Sprintf("Total: %d, Errors: %d", total, p.journal.FailedCount))
	return nil
}

func (p *Parser) visitPage(page int) error {
	p.helper.Set("page", strconv.Itoa(page))
	c := p.collector.Clone()
	c.OnError(func(r *colly.Response, err error) {
		p.onError(r, err)
	})
	c.OnResponse(func(r *colly.Response) {
		var m Meta
		xml.Unmarshal(r.Body, &m)
		p.meta.Rows = append(p.meta.Rows, m.Rows...)
	})
	return c.Visit(p.helper.Get())
}

func (p *Parser) visitPrint(ID string) (*models.Registry, error) {
	var result error
	rowid, err := strconv.ParseUint(ID, 10, 0)
	if err != nil {
		return nil, err
	}
	var r models.Registry
	r.ID = uint(rowid)
	r.RegistryStatusID = uint(dict.Processing)
	if err != nil {
		return nil, err
	}
	c := p.collector.Clone()
	url := p.pages.Main + p.pages.Print + ID
	r.Link = url
	c.OnError(func(r *colly.Response, err error) {
		p.onError(r, err)
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		result = parsePage(&r, e)
	})
	c.OnScraped(func(resp *colly.Response) {
		p.addRegistry(&r)
	})
	if err = c.Visit(url); err != nil {
		return nil, err
	}
	return &r, result
}

func parsePage(r *models.Registry, e *colly.HTMLElement) error {
	var g run.Group
	e.ForEach("table", func(i int, el *colly.HTMLElement) {
		// Основные данные
		if i == dict.MainTableIndex {
			g.Add(func() error {
				return parseMainTable(r, el)
			}, func(err error) {})
		}
		// Производители
		if i == dict.ManufacturesTableIndex {
			g.Add(func() error {
				return parseManufacturesTable(r, el)
			}, func(err error) {})
		}
		// Комплектность
		if i == dict.PartsTableIndex {
			g.Add(func() error {
				return parsePartsTable(r, el)
			}, func(err error) {})
		}
	})
	return g.Run()
}

func parseMainTable(r *models.Registry, el *colly.HTMLElement) error {
	var result error
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		switch j {
		case dict.RegistryName:
			r.Title = ch.ChildText("td")
			break
		case dict.RegistryNumber:
			r.Number = ch.ChildText("td")
			break
		case dict.IssueDate:
			dt, err := time.Parse("02.01.2006", ch.ChildText("td"))
			if err != nil {
				result = multierror.Append(result, err)
				break
			}
			r.IssueDate = dt
			break
		case dict.RegistryDuration:
			d, err := strconv.Atoi(ch.ChildText("td"))
			if err != nil {
				result = multierror.Append(result, err)
				break
			}
			r.Duration = d
			break
		case dict.ExpireDate:
			val := ch.ChildText("td")
			if val != "" {
				dt, err := time.Parse("02.01.2006", ch.ChildText("td"))
				if err != nil {
					result = multierror.Append(result, err)
					break
				}
				r.ExpireDate = null.NewTime(dt, true)
			}
			break
		default:
			break
		}
	})
	return result
}

func parseManufacturesTable(r *models.Registry, el *colly.HTMLElement) error {
	var result error
	var rel = r.R.NewStruct()
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		var mType string
		if j != dict.TableHeaderIndex {
			var m models.RegistryManufacturer
			ch.ForEach("td", func(t int, td *colly.HTMLElement) {
				switch t {
				case dict.ManufacturerName:
					m.Title = td.Text
					break
				case dict.Country:
					m.Country = td.Text
					break
				case dict.ManufacturerType:
					mType = td.Text
					break
				default:
					break
				}
			})
			if mType == dict.ManufacturerRequired {
				rel.RegistryManufacturers = append(rel.RegistryManufacturers, &m)
			}
		}
	})
	r.R = rel
	return result
}

func parsePartsTable(r *models.Registry, el *colly.HTMLElement) error {
	var result error
	var rel = r.R.NewStruct()
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		if j != dict.TableHeaderIndex {
			var b models.RegistryBuild
			ch.ForEach("td", func(t int, td *colly.HTMLElement) {
				switch t {
				case dict.Order:
					o, err := strconv.Atoi(td.Text)
					if err != nil {
						result = multierror.Append(result, err)
						break
					}
					b.Order = o
					break
				case dict.PartName:
					b.Title = td.Text
					break
				case dict.Model:
					b.Model = td.Text
					break
				default:
					break
				}
			})
			rel.RegistryBuilds = append(rel.RegistryBuilds, &b)
		}
	})
	r.R = rel
	return result
}
