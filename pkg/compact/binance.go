package compact

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
)

type binance struct {
	w              *csv.Writer
	opts           *BinanceOptions
	lastTime       uint64
	lastSavedPrice *float64
	record         []string
}

type BinanceOptions struct {
	TimeColumn  uint    `yaml:"time_column"`
	PriceColumn uint    `yaml:"price_column"`
	PriceDiff   float64 `yaml:"price_diff"`
}

func NewCompactBinance(w io.Writer, opts *BinanceOptions) Writer {
	return &binance{
		w:    csv.NewWriter(w),
		opts: opts,
	}
}

func (c *binance) Write(record []string) error {
	time := record[c.opts.TimeColumn]

	price := record[c.opts.PriceColumn]
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return err
	}

	// Save the record to send it later in Flush()
	c.record = []string{time, price}
	lastTime, err := strconv.ParseUint(time, 10, 64)
	if err != nil {
		return err
	}
	if lastTime < c.lastTime {
		return fmt.Errorf("Wrong time order")
	}
	c.lastTime = lastTime

	// if lastSavedPrice is nil, it means that it's the first record
	if c.lastSavedPrice != nil {
		// Calculate the difference between the current price
		// and the last saved price
		diff := math.Abs(priceFloat-*c.lastSavedPrice) / priceFloat
		if diff <= c.opts.PriceDiff {
			return nil
		}
	}

	c.lastSavedPrice = &priceFloat
	err = c.w.Write(c.record)
	c.record = nil
	return err
}

func (c *binance) Flush() error {
	if c.record != nil {
		if err := c.w.Write(c.record); err != nil {
			return err
		}
	}
	c.w.Flush()
	return c.w.Error()
}
