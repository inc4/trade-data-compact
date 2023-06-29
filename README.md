# Market Trade Data Parser

[![Test status](https://github.com/inc4/trade-data-compact/workflows/Checks/badge.svg)](https://github.com/inc4/trade-data-compact/actions?query=workflow%3A%22Checks%22)

The Parser is a command-line utility that condenses price data from CSV files.
It reads CSV data from the standard input, processes it based on a specified
price difference threshold, and outputs data when the price has changed
significantly.

## Configuration

The program takes its configuration from a `compact.yaml` file that should be
located in the current directory or any upper directory. Alternatively, you can
specify a path where the `compact.yaml` will be searched. The search will
commence from the specified directory and extend upwards.

Several types of data are supported. Each type has its own configuration
structure. Examples can be found in [examples/](./examples).

## Usage

To use the program, pipe the CSV data into the program as follows:

```shell
cat data.csv | ./bin/compact > output.csv
```

This will read the CSV data from `data.csv`, process it, and then write the
filtered data to `output.csv`.

Or, specify the path where the config can be found. Let's assume that the
config is located in `data/deribit/compact.yaml`:

```shell
cat data/deribit/OPTIONS/options_chain/2023/04/2023-04-15.csv | ./bin/compact data/deribit/OPTIONS/options_chain/2023/04 > output.csv
```
