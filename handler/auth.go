package handler 


import ( 

	"net/http"
)

func Login( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	if r.Method == "GET" { 
	generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
	} else { 

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

	}
}

