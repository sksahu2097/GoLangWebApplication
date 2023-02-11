package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Render Template using html template
func RenderTenplateTe(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the template ", err)
		return
	}
}

func RenderTenplate(w http.ResponseWriter, tmpl string) {
	//create the template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)

	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range over each file and add to cache
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil

}
