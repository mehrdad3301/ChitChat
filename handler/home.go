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
	
	ok, err := checkSession(r) 
	if err != nil { 
		fmt.Println("Home: ", err)
	}
	//reading threads
	threads, err := db.GetThreads() 
	if err != nil { 
		fmt.Println("Home: ", err)
	}
	//reading session if exists 
	//executing templates 
	
	if ok { 
		generateHTML(w, threads, "layout", "private.navbar", "index")
	} else {	
		generateHTML(w, threads, "layout", "public.navbar", "index")
	}

}

