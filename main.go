package main 

import ( 
	"net/http"
	"github.com/mehrdad3301/ChitChat/handler"
) 


func main() { 

	fileHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fileHandler))
	http.HandleFunc("/", handler.Home) 
	http.HandleFunc("/login", handler.Login) 
	http.HandleFunc("/signup", handler.SignUp) 

	http.ListenAndServe("localhost:6060", nil)
}

