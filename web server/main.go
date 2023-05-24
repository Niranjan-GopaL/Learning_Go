package main

import(
	"fmt"
	"log"
	"net/http"
)



func helloHandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	
	fmt.Fprintf(w,"hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}




func main() {

	// we are initializing and defining at the same time and therefore :=
	// we are telling the server to look at our static folder
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Printf("Server starting at port 8080 \n")

	// We'll assign err the value of nil or whatever err
	// we use log package to LOG THE ERROR 
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}