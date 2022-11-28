package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

const version = "0.0.1"

var (
	theme    = flag.String("theme", "monokai", "highlights theme name")
	language = flag.String("language", "", "file language")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: golang-cat [options...] filepath")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := printText(&bytes); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printText(text *[]byte) (err error) {
	var lexer chroma.Lexer
	if *language != "" {
		lexer = lexers.Get(*language)
	} else {
		lexer = lexers.Analyse(string(*text))
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(*theme)
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	iterator, err := lexer.Tokenise(nil, string(*text))
	if err != nil {
		return err
	}

	return formatter.Format(os.Stdout, style, iterator)
}
