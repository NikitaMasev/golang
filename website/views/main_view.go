package views

import (
	"fmt"
	"html/template"
	"net/http"
    "reflect"
	"github.com/NikitaMasev/golang/website/models"
)

var tmpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Website</title>
		<style>
   table {
    width: 100%; 
    background: white; 
    color: white; 
    border-spacing: 1px; 
   }
   td, th {
    background: maroon; 
    padding: 3px; 
   }
  </style>
	</head>
	<body>
		<table border="1" width="100%" cellpadding="5">
				<tr>
					<td>FirstName</td>
					<td>Email</td>
					<td>Username</td>
				</tr>
				{{range .}}<tr>{{range rangeStruct .}} <td>{{.}}</td>{{end}}</tr>{{end}}
										
			
			
				
			</table>
		
	</body>
</html>`

var templateFuncs = template.FuncMap{"rangeStruct": rangeStructer}

func rangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}
//CreateHTMLResponseView allow to render output html page...
func CreateHTMLResponseView(w http.ResponseWriter, data []models.WebInfo) {
	tmp := template.New("tmp").Funcs(templateFuncs)
	t, err := tmp.Parse(tmpl)
	if err != nil {
		panic(err)
	}
	templErr := t.ExecuteTemplate(w, "tmp", data)
	if templErr != nil {
		fmt.Fprintf(w, templErr.Error())
		fmt.Fprintf(w, t.DefinedTemplates())
	}
}