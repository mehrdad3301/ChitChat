package handler 


import ( 

	"fmt"
	"net/http"
	"github.com/mehrdad3301/ChitChat/db"
)

func Home( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	
	//reading threads
	//reading session if exists 
	ok, err := checkSession(r) 
	if err != nil { 
		fmt.Println("Home: ", err)
	}
	//executing templates 
	
	if ok { 
		generateHTML(w, new(interface{}), "layout", "private.navbar", "index")
	} else {	
		generateHTML(w, new(interface{}), "layout", "public.navbar", "index")
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
