package handler 


import ( 
	"log"
	"fmt"
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
		user, err := db.GetUser("email", email)
		if err != nil { 
			generateHTML(w, new(interface{}), "login.layout", "public.navbar", "login")
			log.Println("Login: ", err)
		}

		if !db.CheckPassword(user, pass) { 
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

	err := deleteSession(r)
	if err != nil { 
		log.Println("Logout: ", err) 	
	}

	http.Redirect(w, r, "/", http.StatusFound)
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

func checkSession(r *http.Request) (bool, error) { 
	cookie, err := r.Cookie("session_cookie")
	if err != nil { 
		return false, fmt.Errorf("checkSession: ", err)
	}	
	ok, err := db.ValidSession(cookie.Value)
	return ok, err 
}

func deleteSession(r *http.Request) (error) {

	cookie, err := r.Cookie("session_cookie")
	if err != nil { 
		return fmt.Errorf("deleteSession: ", err)
	}	
	err = db.DeleteSession(cookie.Value)
	return  err 
}
