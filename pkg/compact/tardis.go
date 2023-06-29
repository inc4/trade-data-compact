package compact

import (
	"encoding/csv"
	"fmt"
	"io"
)

type tardis struct {
	w               *csv.Writer
	opts            *TardisOptions
	headerStart     string
	isHeaderWritten bool
	indices         map[string]int
	prices          map[string]string
}

type TardisOptions struct {
	// The column name that contains symbol
	NameColumn string `yaml:"name_column"`

	// The column name that contains price
	PriceColumn string `yaml:"price_column"`

	// All columns that should be saved
	Columns []string `yaml:"columns"`
}

func NewCompactTardis(w io.Writer, opts *TardisOptions) Writer {
	indices := make(map[string]int)
	for _, col := range opts.Columns {
		indices[col] = -1
	}
	return &tardis{
		w:       csv.NewWriter(w),
		opts:    opts,
		indices: indices,
		prices:  make(map[string]string),
	}
}

func (c *tardis) Write(record []string) error {
	if !c.isHeaderWritten {
		c.headerStart = record[0]
		for i, col := range record {
			if _, ok := c.indices[col]; ok {
				c.indices[col] = i
			}
		}
		for col, i := range c.indices {
			if i == -1 {
				return fmt.Errorf("Column %s not found", col)
			}
		}
	}

	rec := make([]string, 0, len(c.indices))
	if !c.isHeaderWritten {
		c.isHeaderWritten = true
		return c.w.Write(append(rec, c.opts.Columns...))
	}

	if record[0] == c.headerStart {
		return nil
	}
	name := record[c.indices[c.opts.NameColumn]]
	price := record[c.indices[c.opts.PriceColumn]]
	if oldPrice, ok := c.prices[name]; ok && oldPrice == price {
		return nil
	}
	c.prices[name] = price

	for _, col := range c.opts.Columns {
		rec = append(rec, record[c.indices[col]])
	}
	return c.w.Write(rec)
}

func (c *tardis) Flush() error {
	c.w.Flush()
	return c.w.Error()
}
