package httpHandlers

import (
	"html/template"
	"net/http"

	"github.com/QO-Development/cashFlowApp/models/errorLog"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//Home is the http handler for the "/" path
func Home(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		errorLog.Write("Unable to execute template main.go 25: ", err)
		http.Error(w, err.Error(), 500)
	}

}
