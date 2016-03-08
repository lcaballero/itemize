package web

import (
	"bytes"
	eng "html/template"
	"log"
)

type Render struct{}

func NewRender(tmplBase, layoutBase string) *Render {
	return &Render{}
}

func (r *Render) Render(data interface{}, templates ...string) (string, error) {
	tmpl, err := eng.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return buf.String(), nil
}
