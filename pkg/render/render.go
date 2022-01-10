package render

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, htmltmpl string) {

	parsed, err := template.ParseFiles("./templates/" + htmltmpl)

	//Log error and exit program if template folder is missing
	if err != nil {
		log.Fatal("Unable to locate template folder :", err)
	}

	//check error passing the template
	err = parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
