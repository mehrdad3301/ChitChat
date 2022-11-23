package handler 


import ( 
	"fmt"
	"net/http"
	"strconv"
	"github.com/mehrdad3301/ChitChat/db"
)

func CreateThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	r.ParseForm() 
	topic := r.FormValue("topic")	
	
	session, err := getCurrSession(r)
	if err != nil { 
		fmt.Println("CreateThread: ", err)
		http.Redirect(w, r, "/", http.StatusFound)
		return 
	}
	
	err = db.CreateThread(topic, session.UserId)
	if err != nil { 
		fmt.Println("CreateThread: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func NewThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

	ok, err := checkSession(r) 
	if err != nil { 
		fmt.Println("CreateThread: ", err)
	}
	if !ok { 
		http.Redirect(w, r, "/login", http.StatusFound)
		return 
	}

	generateHTML(w, new(interface{}), "layout", "private.navbar", "new.thread")
}

func ReadThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	
	if r.Method == "GET" { 
		r.ParseForm()
		id := r.FormValue("id")
		thread, err := db.GetThread(id)
		if err != nil { 
			fmt.Println("ReadThread: ", err) 
			http.Redirect(w, r, "/", http.StatusFound)
			return 
		}
		generateHTML(w, thread, "layout", "private.navbar", "private.thread")
	}

}

func PostThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

	r.ParseForm() 
	body := r.FormValue("body")	
	threadId, _ := strconv.Atoi(r.FormValue("id"))
	
	session, err := getCurrSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = db.CreatePost(threadId, session.UserId, body) 
		if err != nil { 
			fmt.Println("createPost: ", err)	
		}
		url := "/thread/read" + "?id=" +r.FormValue("id")
		http.Redirect(w, r, url, http.StatusFound)
	}
}

