package html

import (
	"html/template"
	"log"
)

var IndexHTMLTemplate *template.Template
var TailHTMLTemplate *template.Template

func init() {
	var err error
	IndexHTMLTemplate, err = template.New("index").Parse(IndexHTML)
	if err != nil {
		log.Fatal("could not parse Index HTML", err)
		return
	}

	TailHTMLTemplate, err = template.New("index").Parse(TailHTML)
	if err != nil {
		log.Fatal("could not parse Tail HTML", err)
		return
	}
}