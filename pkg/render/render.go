package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//a dictionary of functions to be used in the template
var functions = template.FuncMap{}

//Render renders golang templates
func Render(w http.ResponseWriter, tmpl string) {
	tc, err := MakeTmplCache()
	if err != nil {
		log.Fatal(err)
	}
	//check if the template exists
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	//store the template in the buffer
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	//we don't need the num of bytes, just check if there is no error
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

//MakeTmplCache creates cached templates
func MakeTmplCache() (map[string]*template.Template, error) {
	//map of templates: [name] pointer to a fully parsed template
	tempCache := map[string]*template.Template{}
	//look for any files in the templates that match .page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tempCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}
		//look for any files in the templates that match .layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tempCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}
		tempCache[name] = ts
	}
	return tempCache, nil
}
