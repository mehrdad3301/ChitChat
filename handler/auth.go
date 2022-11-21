package handler 


import ( 
	"net/http"
	"github.com/mehrdad3301/ChitChat/db"
)

func Login( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	if r.Method == "GET" { 
	generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
	} else { 
		r.ParseForm() 
		pass := r.FormValue("password")
		email := r.FormValue("email")
		
	}
}

func Logout( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

}

func SignUp( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	if r.Method == "GET" { 
		generateHTML(w, new(interface{}), "login.layout", "public.navbar", "signup")
	} else { 
		r.ParseForm() 
		name := r.FormValue("name") 
		pass := r.FormValue("password")
		email := r.FormValue("email")
	
		db.CreateUser(name, pass, email) 
		http.Redirect(w, r, "/login", http.StatusFound) 
	}
}

