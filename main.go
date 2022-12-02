package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

const version = "0.0.3"

var (
	number_f = flag.Bool("number", false, "line number flag")
	theme    = flag.String("theme", "monokai", "highlights theme name")
	language = flag.String("language", "", "file language")
)

type withNumberWriter struct {
	writer     io.Writer
	lineNumber uint32
	buffer     *bytes.Buffer
}

func (wnw *withNumberWriter) Write(b []byte) (n int, err error) {
	if !bytes.Contains(b, []byte{'\n'}) {
		wnw.buffer.Write(b)
		return len(b), nil
	}

	var p []byte
	if wnw.buffer.Len() > 0 {
		p = wnw.buffer.Bytes()
		wnw.buffer.Reset()
	}
	p = append(p, b...)

	ln := []byte(fmt.Sprintf("%4d| ", wnw.lineNumber))
	if _, err = wnw.writer.Write(append(ln, p...)); err != nil {
		return len(b), err
	}
	wnw.lineNumber++

	return len(b), nil
}

func newWithNumber(w io.Writer) *withNumberWriter {
	return &withNumberWriter{w, 1, new(bytes.Buffer)}
}

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

	bytes, err := os.ReadFile(flag.Arg(0))
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
	var w io.Writer
	w = os.Stdout
	if *number_f {
		w = newWithNumber(w)
	}

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

	return formatter.Format(w, style, iterator)
}
