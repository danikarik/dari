package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/oklog/run"
	"go.uber.org/zap"

	"bitbucket.org/kit-systems/dari/pkg/urlhelper"
)

// Parser is a basic struct for parser.
type Parser struct {
	logger    *zap.Logger
	pages     *Pages
	collector *colly.Collector
	helper    *urlhelper.URLHelper
	regs      []Registry
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
	p.pages = DefaultPages()
	p.collector = DefaultCollector(opts.Debug)
	p.helper = urlhelper.New()
	p.regs = make([]Registry, 0)
	return &p, nil
}

// Run starts parser.
func (p *Parser) Run() error {
	var err error

	defer func() {
		p.logger.Info(
			"finished",
			zap.Float64("time took", time.Since(time.Now()).Seconds()),
		)
	}()

	defer func() {
		if derr := p.logger.Sync(); derr != nil {
			p.logger.Error("sync", zap.Error(derr))
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

	err = p.collectRegistry()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) addRegistry(r Registry) {
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
		var m Meta
		xml.Unmarshal(r.Body, &m)
		p.pages.Count, _ = strconv.Atoi(m.Total)
		for _, r := range m.Rows {
			p.visitPrint(r.ID)
		}
	})
	return p.collector.Visit(p.helper.Get())
}

func (p *Parser) collectRegistry() error {
	var err error
	// for i := 2; i <= p.pages.Count; i++ {
	for i := 2; i <= 2; i++ {
		err = p.visitPage(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) visitPage(page int) error {
	p.helper.Set("page", strconv.Itoa(page))
	p.collector.OnResponse(func(r *colly.Response) {
		var m Meta
		xml.Unmarshal(r.Body, &m)
		for _, r := range m.Rows {
			err := p.visitPrint(r.ID)
			if err != nil {
				p.logger.Error(
					"print page",
					zap.String("ID", r.ID),
					zap.Error(err),
				)
			}
		}
	})
	return p.collector.Visit(p.helper.Get())
}

func (p *Parser) visitPrint(ID string) error {
	var err error
	var r Registry
	c := p.collector.Clone()
	url := p.pages.Main + p.pages.Print + ID
	r.Link = url
	c.OnHTML("body", func(e *colly.HTMLElement) {
		err = parsePage(&r, e)
	})
	c.OnScraped(func(resp *colly.Response) {
		if err == nil {
			fmt.Println(r.Number)
			p.regs = append(p.regs, r)
		}
	})
	return c.Visit(url)
}

func parsePage(r *Registry, e *colly.HTMLElement) error {
	var g run.Group
	{
		e.ForEach("table", func(i int, el *colly.HTMLElement) {
			// Основные данные
			if i == MainTableIndex {
				g.Add(func() error {
					parseMainTable(r, el)
					return nil
				}, func(error) {
					//
				})
			}
			// Производители
			if i == ManufacturesTableIndex {
				g.Add(func() error {
					parseManufacturesTable(r, el)
					return nil
				}, func(error) {
					//
				})
			}
			// Комплектность
			if i == PartsTableIndex {
				g.Add(func() error {
					parsePartsTable(r, el)
					return nil
				}, func(error) {
					//
				})
			}
		})
	}
	return g.Run()
}

func parseMainTable(r *Registry, el *colly.HTMLElement) {
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		switch j {
		case RegistryName:
			r.Name = ch.ChildText("td")
			break
		case RegistryNumber:
			r.Number = ch.ChildText("td")
			break
		case IssueDate:
			r.IssueDate, _ = time.Parse("02.01.2006", ch.ChildText("td"))
			break
		case RegistryDuration:
			r.Duration, _ = strconv.Atoi(ch.ChildText("td"))
			break
		case ExpireDate:
			r.ExpireDate, _ = time.Parse("02.01.2006", ch.ChildText("td"))
			break
		default:
			break
		}
	})
}

func parseManufacturesTable(r *Registry, el *colly.HTMLElement) {
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		if j != TableHeaderIndex {
			var m Manufacturer
			ch.ForEach("td", func(t int, td *colly.HTMLElement) {
				switch t {
				case ManufacturerName:
					m.Name = td.Text
					break
				case Country:
					m.Country = td.Text
					break
				case ManufacturerType:
					m.Type = td.Text
					break
				default:
					break
				}
			})
			if m.Type == ManufacturerRequired {
				r.Manufacturers = append(r.Manufacturers, m)
			}
		}
	})
}

func parsePartsTable(r *Registry, el *colly.HTMLElement) {
	el.ForEach("tbody tr", func(j int, ch *colly.HTMLElement) {
		if j != TableHeaderIndex {
			var b Build
			ch.ForEach("td", func(t int, td *colly.HTMLElement) {
				switch t {
				case Order:
					b.Order, _ = strconv.Atoi(td.Text)
					break
				case PartName:
					b.Name = td.Text
					break
				case Model:
					b.Model = td.Text
					break
				default:
					break
				}
			})
			r.Builds = append(r.Builds, b)
		}
	})
}
