package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	header bool
	colon  string
	tab    string
	fields []string
}

func main() {
	header := flag.Bool("h", true, "show header")
	colon := flag.String("c", ":", "colon string")
	tab := flag.String("t", "\t", "tab string")
	only := flag.String("f", "", "fields")
	flag.Parse()

	config := &Config{
		header: *header,
		colon:  *colon,
		tab:    *tab,
		fields: strings.Split(*only, ","),
	}

	switch flag.NArg() {
	case 0:
		fromStdin(config)
	case 1:
		fromFile(config, flag.Arg(0))
	default:
		showHelp()
	}
}

func fromStdin(config *Config) {
	ltsv(config, bufio.NewScanner(os.Stdin))
}

func fromFile(config *Config, path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ltsv(config, bufio.NewScanner(file))
}

func ltsv(config *Config, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		fmt.Println(line)
	}
}

func showHelp() {
	fmt.Printf(`Usage: %v [options] [file] fields,...

options:
    -c: colon string, default is ':'
    -t: tab string, default is '\t'
`, os.Args[0])
}
