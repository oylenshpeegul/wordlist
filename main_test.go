package main

import (
	"bufio"
	"os"
	"testing"
)

var compressPairTests = []struct {
	scheme string
	s      string
	t      string
	out    string
}{
	{"DAWG", "", "foo", "0foo"},
	{"DAWG", "foo", "foot", "3t"},
	{"DAWG", "foot", "footle", "4le"},
	{"DAWG", "footle", "fubar", "1ubar"},
	{"DAWG", "fubar", "fub", "3"},
	{"DAWG", "fub", "grunt", "0grunt"},

	{"Mike", "", "myxa", "@myxa"},
	{"Mike", "myxa", "myxophyta", "Cophyta"},
	{"Mike", "myxophyta", "myxopod", "Eod"},
	{"Mike", "myxopod", "nab", "@nab"},
	{"Mike", "nab", "nabbed", "Cbed"},
	{"Mike", "nabbed", "nabbing", "Ding"},
	{"Mike", "nabbing", "nabit", "Cit"},
	{"Mike", "nabit", "nabk", "Ck"},
	{"Mike", "nabk", "nabob", "Cob"},
	{"Mike", "nabob", "nacarat", "Bcarat"},
	{"Mike", "nacarat", "nacelle", "Celle"},
}

var decompressPairTests = []struct {
	scheme string
	s      string
	t      string
	out    string
	err    error
}{
	{"DAWG", "", "0foo", "foo", nil},
	{"DAWG", "foo", "3t", "foot", nil},
	{"DAWG", "foot", "4le", "footle", nil},
	{"DAWG", "footle", "1ubar", "fubar", nil},
	{"DAWG", "fubar", "3", "fub", nil},
	{"DAWG", "fub", "0grunt", "grunt", nil},

	{"Mike", "", "@myxa", "myxa", nil},
	{"Mike", "myxa", "Cophyta", "myxophyta", nil},
	{"Mike", "myxophyta", "Eod", "myxopod", nil},
	{"Mike", "myxopod", "@nab", "nab", nil},
	{"Mike", "nab", "Cbed", "nabbed", nil},
	{"Mike", "nabbed", "Ding", "nabbing", nil},
	{"Mike", "nabbing", "Cit", "nabit", nil},
	{"Mike", "nabit", "Ck", "nabk", nil},
	{"Mike", "nabk", "Cob", "nabob", nil},
	{"Mike", "nabob", "Bcarat", "nacarat", nil},
	{"Mike", "nacarat", "Celle", "nacelle", nil},
}

func TestCompressPair(t *testing.T) {

	for _, tt := range compressPairTests {

		w := Wordlist{
			Schemes[tt.scheme],
			bufio.NewScanner(bufio.NewReader(os.Stdin)),
			bufio.NewWriter(os.Stdout),
		}

		p := w.compressPair(tt.s, tt.t)
		if p != tt.out {
			t.Errorf("w.compressPair(%q, %q) => %q, want %q",
				tt.s, tt.t, p, tt.out)
		}
	}

}

func TestDecompressPair(t *testing.T) {

	for _, tt := range decompressPairTests {

		w := Wordlist{
			Schemes[tt.scheme],
			bufio.NewScanner(bufio.NewReader(os.Stdin)),
			bufio.NewWriter(os.Stdout),
		}

		p, err := w.decompressPair(tt.s, tt.t)
		if err != nil {
			t.Errorf("w.decompressPair(%q, %q) => %q, want %q",
				tt.s, tt.t, err, tt.err)
		}
		if p != tt.out {
			t.Errorf("w.decompressPair(%q, %q) => %q, want %q",
				tt.s, tt.t, p, tt.out)
		}
	}

}
