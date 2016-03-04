package web
import (
	"path/filepath"
	"io/ioutil"
	"html/template"
	"log"
	"bytes"
)


type Render struct {
	base string
}

func NewRender(base string) *Render {
	return &Render{
		base: base,
	}
}

func (r *Render) Render(tpl string, data interface{}) (string, error) {
	templateName := filepath.Join(r.base, tpl)
	bits, _ := ioutil.ReadFile(templateName)
	tmpl, err := template.New("list").Parse(string(bits))
	if err != nil {
		log.Println(err)
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})

	if data == nil {
		data = make(map[string]interface{})
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return buf.String(), nil
}
