package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type Request struct {
	Name string
}

type Response struct {
	Greeting string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    var request Request
    
    decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err != nil {
		panic("Error while decoding JSON ...")
	}
	
	postResponse := Response{Greeting: "Hello, " + request.Name + " !"}
	json.NewEncoder(rw).Encode(postResponse)
    
}

func main() {
    mux := httprouter.New()
    
    mux.GET("/hello/:name", hello)
    mux.POST("/hello/", helloPost)
    
    server := http.Server{
            Addr: "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}