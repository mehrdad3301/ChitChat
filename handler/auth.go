package handler 


import ( 
	"log"
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
		user, err := db.GetUser(email)
		if err != nil { 
			generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
			log.Println("Login: ", err)
		}

		if !db.ValidUser(user, pass) { 
			generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
			log.Println("Login: ", "wrong password")
		} else { 

			uuid, err := db.CreateSession(user)
			if err != nil { 
				log.Println("Login: ", err)
			}
			cookie := http.Cookie { 
				Name: "session_cookie",
				Value: uuid, 
				HttpOnly : true, 
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		}
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
	
		err := db.CreateUser(name, pass, email) 
		if err != nil { 
			log.Println("SignUp: ", err)
			generateHTML(w, new(interface{}), "login.layout", "public.navbar", "signup")
			return 
		}
		http.Redirect(w, r, "/login", http.StatusFound) 
	}
}

