package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/inc4/trade-data-compact/pkg/compact"
	"gopkg.in/yaml.v3"
)

type Config interface {
}

type config struct {
	Type string `yaml:"type"`
}

func findConfig(path string) string {
	for d := path; d != "" && filepath.Dir(d) != d; d = filepath.Dir(d) {
		file := filepath.Join(d, "compact.yaml")
		if _, err := os.Stat(file); err == nil {
			return file
		}
		if d == "/" {
			break
		}
	}
	return ""
}

func readConfig(confFilename string) (Config, error) {
	var conf config

	data, err := os.ReadFile(confFilename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	switch conf.Type {
	case "binance":
		var binanceConfig compact.BinanceOptions
		err := yaml.Unmarshal(data, &binanceConfig)
		if err != nil {
			return nil, err
		}
		return &binanceConfig, nil
	case "tardis":
		var tardisConfig compact.TardisOptions
		err := yaml.Unmarshal(data, &tardisConfig)
		if err != nil {
			return nil, err
		}
		return &tardisConfig, nil
	default:
		return nil, fmt.Errorf("Unknown type '%s'", conf.Type)
	}
}

func main() {
	help := flag.Bool("h", false, "Show help message")
	flag.Parse()
	if *help {
		fmt.Printf("Usage: %s [config-search-path]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	path := "."
	args := flag.Args()
	if len(args) > 0 {
		path = args[0]
	}

	confFile := findConfig(path)
	if confFile == "" {
		log.Fatal("compact.yaml file not found")
	}
	conf, err := readConfig(confFile)
	if err != nil {
		log.Fatal(err)
	}

	var w compact.Writer
	switch confType := conf.(type) {
	case *compact.BinanceOptions:
		w = compact.NewCompactBinance(
			os.Stdout,
			confType,
		)
	case *compact.TardisOptions:
		w = compact.NewCompactTardis(
			os.Stdout,
			confType,
		)
	default:
		log.Fatal("Unknown type")
	}

	r := csv.NewReader(os.Stdin)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if err = w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}
