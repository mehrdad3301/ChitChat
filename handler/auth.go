package handler 


import ( 

	"net/http"
)

func Login( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
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

}

func SignIn( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

}
