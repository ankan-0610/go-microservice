package main

import(
	"log"
	"net/http"
	"fmt"
)

const webPort = "80"

type Config struct {}

func main(){
	app := Config{}

	log.Printf("Starting broker service on Port %s\n",webPort)

	//define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	//Start the server
	err:= srv.ListenAndServe()
	if err!=nil{
		log.Panic(err)
	}
}