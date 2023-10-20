package aho_corasick

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLoadStrings(t *testing.T) {
	tb := NewTrieBuilder()

	if err := tb.LoadStrings("doesnt-exists.txt"); err == nil {
		t.Errorf("should fail")
	}

	if err := tb.LoadStrings("./test_data/strings.txt"); err != nil {
		t.Error(err)
	}
	tr := tb.Build()

	ibsen, err := ioutil.ReadFile("./test_data/Ibsen.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 697
	ms := tr.Match(ibsen)
	if len(ms) != expected {
		for _, m := range ms {
			fmt.Println(m)
		}
		t.Errorf("expected %d matches, got %d\n", expected, len(ms))
	}
	mss := map[string]int{}
	for _, m := range ms {
		mss[string(m.match)]++
	}
	fmt.Println(mss)
}

func TestLoadPatterns(t *testing.T) {
	var err error
	tb := NewTrieBuilder()

	if err = tb.LoadPatterns("./test_data/patterns.txt"); err != nil {
		t.Error(err)
	}
	tr := tb.Build()

	ibsen, err := ioutil.ReadFile("./test_data/Ibsen.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 697
	ms := tr.Match(ibsen)
	if len(ms) != expected {
		for _, m := range ms {
			fmt.Println(m)
		}
		t.Errorf("expected %d matches, got %d\n", expected, len(ms))
	}
}

func TestStrings(t *testing.T) {
	tb := NewTrieBuilder()

	tb.AddStrings([]string{"he", "her", "hi"})

	tr := tb.Build()
	ms := tr.MatchString("hi!erheherhi")
	mss := map[string]int{}
	for _, m := range ms {
		mss[string(m.match)]++
	}
	fmt.Println(mss)
}
