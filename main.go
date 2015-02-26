package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type Scheme struct {
	header string
	sub    string
}

var Schemes = map[string]*Scheme{
	"Crack": {
		"#!xdawg",
		"0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"[\\]^_`abcdefghijklmnopqrstuvwxyz{",
	},
	"DAWG": {
		"#!xdawg",
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	},
	"Mike": {
		"",
		"@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{",
	},
}

type Wordlist struct {
	*Scheme
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func (w Wordlist) compress() {
	p := ""
	for w.scanner.Scan() {
		word := w.scanner.Text()
		fmt.Fprintln(w.writer, w.compressPair(p, word))
		p = word
	}
	w.writer.Flush()
}

func (w Wordlist) compressPair(s, t string) string {
	n := 0
	for n < len(s) && n < len(t) && s[n] == t[n] && n < len(w.sub) {
		n++
	}
	return string(w.sub[n]) + t[n:]
}

func (w Wordlist) decompress() error {
	p := ""
	var err error
	for w.scanner.Scan() {
		word := w.scanner.Text()
		if p, err = w.decompressPair(p, word); err != nil {
			return err
		}
		fmt.Fprintln(w.writer, p)
	}
	w.writer.Flush()
	return nil
}

func (w Wordlist) decompressPair(s, t string) (string, error) {
	n := strings.IndexByte(w.sub, t[0])
	if n == -1 {
		msg := fmt.Sprintf(
			"Cannot decompress '%s': '%s' not found in '%s'",
			t,
			string(t[0]),
			w.sub,
		)
		return "", errors.New(msg)
	}
	return s[:n] + t[1:], nil
}

func main() {

	var schemename = flag.String(
		"schemename",
		"Mike",
		"name of scheme to use: Crack, DAWG or [Mike]",
	)
	flag.Parse()

	filename := "-"
	if flag.NArg() == 1 {
		filename = flag.Arg(0)
	}

	var file *os.File
	if filename == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	w := Wordlist{
		Schemes[*schemename],
		bufio.NewScanner(bufio.NewReader(file)),
		bufio.NewWriter(os.Stdout),
	}

	if path.Ext(filename) == ".cpt" {
		if w.header != "" {
			w.scanner.Scan()
			if w.header != w.scanner.Text() {
				log.Fatal(w.scanner.Text(), " != ", w.header)
			}
		}
		if err := w.decompress(); err != nil {
			log.Fatal(err)
		}
	} else {
		if w.header != "" {
			fmt.Fprintln(w.writer, w.header)
		}
		w.compress()
	}

}
