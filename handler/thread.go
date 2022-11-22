package handler 


import ( 

	"fmt"
	"net/http"
)

func CreateThread( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 

	ok, err := checkSession(r) 
	if err != nil { 
		fmt.Println("CreateThread", err)
	}
	if !ok { 
		http.Redirect(w, r, "/login", http.StatusFound)
		return 
	}

	if r.Method == "GET" { 
		generateHTML(w, new(interface{}), "layout", "private.navbar", "new.thread")
	} else { 

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

