package render

import (
	"fmt"
	"log"
	"net/http"
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

var tc = make(map[string]*template.Template)

func RenderTenplate(w http.ResponseWriter, tmpl string) {
	var tmplObj *template.Template
	var err error

	_, isMap := tc[tmpl]
	if !isMap {
		//create for the first time
		err := createTemplateCache(tmpl)
		if err != nil {
			log.Println("Error = ", err)
		}
	} else {
		log.Println("Using from cache")
	}
	tmplObj = tc[tmpl]
	err = tmplObj.Execute(w, nil)
	if err != nil {
		log.Println("Error in execute = ", err)
	}
}

func createTemplateCache(tmp1 string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmp1),
		"./templates/base.layout.tmpl",
	}
	templ, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[tmp1] = templ
	return nil

}
