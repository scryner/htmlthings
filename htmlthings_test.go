package htmlthings

import (
	"os"
	"testing"
)

const (
	_sampleFile = "sample.html"
)

var things *HtmlThings

func TestInit(t *testing.T) {
	f, err := os.Open(_sampleFile)
	if err != nil {
		t.Errorf("open err: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	things, err = NewHtmlThings(f)
	if err != nil {
		t.Errorf("while initializing: %v", err)
		os.Exit(1)
	}
}

func TestFindAttrValues(t *testing.T) {
	cond := NewCondition(NewElement("select").Attr("name", "device")).Include(NewElement("option"))

	ret := things.FindAttrValues(cond, "value")

	if len(ret) < 1 {
		t.Errorf("no match")
		t.FailNow()
	}
}

func TestFindTexts(t *testing.T) {
	cond := NewCondition(NewElement("select").Attr("name", "device")).Include(NewElement("option"))

	ret := things.FindTexts(cond)

	if len(ret) < 1 {
		t.Errorf("no match")
		t.FailNow()
	}
}
