package handler 


import ( 

	"net/http"
)

func Home( 
	w http.ResponseWriter, 
	r *http.Request, 
) { 
	
	//reading threads
	//reading session if exists 
	//executing templates 
	
	
	generateHTML(w, new(interface{}), "layout", "public.navbar", "index")

}

