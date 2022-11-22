package handler 


import ( 
	"fmt"
	"net/http"
	"github.com/mehrdad3301/ChitChat/db"
)

func CreateThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

	//ok, err := checkSession(r) 
	//if err != nil { 
	//	fmt.Println("CreateThread", err)
	//}
	//if !ok { 
	//	http.Redirect(w, r, "/login", http.StatusFound)
	//	return 
	//}

	if r.Method == "GET" { 
		generateHTML(w, new(interface{}), "layout", "private.navbar", "new.thread")
	} else { 
		r.ParseForm() 
		topic := r.FormValue("topic")	
		
		cookie, err := r.Cookie("session_cookie") 
		if err != nil { 
			fmt.Println("CreateThread: ", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return 
		} 

		session, err := db.GetSession(cookie.Value)
	
		err = db.CreateThread(topic, session.UserId)
		if err != nil { 
			fmt.Println("CreateThread: ", err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func ReadThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

}

func PostThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

}

