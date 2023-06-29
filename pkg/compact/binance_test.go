package compact

import (
	"bytes"
	"testing"
)

func TestBinance(t *testing.T) {
	w := new(bytes.Buffer)

	opts := &BinanceOptions{
		TimeColumn:  0,
		PriceColumn: 1,
		PriceDiff:   0.01,
	}
	c := NewCompactBinance(w, opts)

	data := [][]string{
		{"1000", "0.1234"},
		{"1001", "0.1235"}, // small increase
		{"1003", "0.1210"}, // significant decrease
		{"1003", "0.1220"}, // small
		{"1003", "0.1222"}, // small
		{"1003", "0.1223"}, // 0.1223-0.1210 = 0.0013 > 0.01*0.1223
		{"1004", "0.1224"}, // small
		{"1005", "0.1225"}, // last
	}

	for _, row := range data {
		err := c.Write(row)
		if err != nil {
			t.Fatal(err)
		}
	}
	err := c.Flush()
	if err != nil {
		t.Fatal(err)
	}

	expected := "1000,0.1234\n1003,0.1210\n1003,0.1223\n1005,0.1225\n"
	if w.String() != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, w.String())
	}
}

func TestBinanceTimeOrder(t *testing.T) {
	w := new(bytes.Buffer)

	opts := &BinanceOptions{
		TimeColumn:  1,
		PriceColumn: 2,
		PriceDiff:   0.01,
	}
	c := NewCompactBinance(w, opts)

	_ = c.Write([]string{"test", "1000", "0.1234"})
	_ = c.Write([]string{"test", "1001", "0.1235"})
	err := c.Write([]string{"test", "1000", "0.1236"})
	if err == nil {
		t.Fatal("Wrong time order should return an error")
	}
}
