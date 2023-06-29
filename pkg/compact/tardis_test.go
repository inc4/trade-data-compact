package compact

import (
	"bytes"
	"testing"
)

func TestTardis(t *testing.T) {
	w := new(bytes.Buffer)

	opts := &TardisOptions{
		NameColumn:  "symbol",
		PriceColumn: "mark_price",
		Columns:     []string{"symbol", "type", "timestamp", "strike_price", "expiration", "mark_price", "mark_iv"},
	}
	c := NewCompactTardis(w, opts)

	data := [][]string{
		{"exchange", "symbol", "timestamp", "local_timestamp", "type", "strike_price", "expiration", "open_interest", "last_price", "bid_price", "bid_amount", "bid_iv", "ask_price", "ask_amount", "ask_iv", "mark_price", "mark_iv", "underlying_index", "underlying_price", "delta", "gamma", "vega", "theta", "rho"},
		{"deribit", "BTC-8JAN21-40000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
		{"deribit", "BTC-8JAN21-41000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
		{"deribit", "BTC-8JAN21-40000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
		{"deribit", "BTC-8JAN21-40000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391782", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
		{"deribit", "BTC-8JAN21-41000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
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

	expected := "symbol,type,timestamp,strike_price,expiration,mark_price,mark_iv\nBTC-8JAN21-40000-C,call,1609459199744000,40000,1610092800000000,0.00391783,132.06\nBTC-8JAN21-41000-C,call,1609459199744000,40000,1610092800000000,0.00391783,132.06\nBTC-8JAN21-40000-C,call,1609459199744000,40000,1610092800000000,0.00391782,132.06\n"
	if w.String() != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, w.String())
	}
}

func TestTardisWrongHeader(t *testing.T) {
	w := new(bytes.Buffer)

	opts := &TardisOptions{
		NameColumn:  "symbol",
		PriceColumn: "mark_price",
		Columns:     []string{"symbol", "type", "timestamp", "strike_price", "expiration", "mark_price", "mark_iv"},
	}
	c := NewCompactTardis(w, opts)

	err := c.Write([]string{"exchange", "nosymbol", "timestamp", "local_timestamp", "type", "strike_price", "expiration", "open_interest", "last_price", "bid_price", "bid_amount", "bid_iv", "ask_price", "ask_amount", "ask_iv", "mark_price", "mark_iv", "underlying_index", "underlying_price", "delta", "gamma", "vega", "theta", "rho"})
	if err == nil {
		t.Fatal("Wrong header should return an error")
	}
}

func TestTardisMultiHeader(t *testing.T) {
	w := new(bytes.Buffer)

	opts := &TardisOptions{
		NameColumn:  "symbol",
		PriceColumn: "mark_price",
		Columns:     []string{"symbol", "type", "timestamp", "strike_price", "expiration", "mark_price", "mark_iv"},
	}
	c := NewCompactTardis(w, opts)

	data := [][]string{
		{"exchange", "symbol", "timestamp", "local_timestamp", "type", "strike_price", "expiration", "open_interest", "last_price", "bid_price", "bid_amount", "bid_iv", "ask_price", "ask_amount", "ask_iv", "mark_price", "mark_iv", "underlying_index", "underlying_price", "delta", "gamma", "vega", "theta", "rho"},
		{"deribit", "BTC-8JAN21-40000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
		{"exchange", "symbol", "timestamp", "local_timestamp", "type", "strike_price", "expiration", "open_interest", "last_price", "bid_price", "bid_amount", "bid_iv", "ask_price", "ask_amount", "ask_iv", "mark_price", "mark_iv", "underlying_index", "underlying_price", "delta", "gamma", "vega", "theta", "rho"},
		{"deribit", "BTC-8JAN21-41000-C", "1609459199744000", "1609459200173039", "call", "40000", "1610092800000000", "4313.2", "0.004", "0.0035", "30", "129.22", "0.0045", "30.4", "135.68", "0.00391783", "132.06", "SYN.BTC-8JAN21", "29052.23", "0.05318", "0.00002", "4.4603", "-40.1604", "0.2875"},
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

	expected := "symbol,type,timestamp,strike_price,expiration,mark_price,mark_iv\nBTC-8JAN21-40000-C,call,1609459199744000,40000,1610092800000000,0.00391783,132.06\nBTC-8JAN21-41000-C,call,1609459199744000,40000,1610092800000000,0.00391783,132.06\n"
	if w.String() != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, w.String())
	}
}
