package handler 


import ( 

	"fmt"
	"net/http"
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

