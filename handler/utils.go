package handler 


import ( 
	"fmt"
	"log"
	"net/http"
	"html/template"
) 


func generateHTML(
	w http.ResponseWriter, 
	data interface{}, 
	fileName ...string,
) {

	files := wrapFileName("./templates/",  ".html", fileName...)
	templates, err := template.ParseFiles(files...)
	if err != nil { 
		log.Fatal("generateHTML: ", err)
	}
	templates.ExecuteTemplate(w, "layout", data)
}

func wrapFileName(
	prefix string, 
	suffix string, 
	filenames ...string, 
) []string { 

	files := make([]string, len(filenames)) 
	for i, filename := range filenames { 
		files[i] = fmt.Sprintf("%s%s%s", prefix, filename, suffix)
	}
	return files 
}
