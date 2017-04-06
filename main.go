package main

import (
	"fmt"
	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/photosvc"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	// Setup static directory
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Load the config
	conf := common.GetConfig()
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	common.InitDatabase()

	// Just return the router to make the syntax nicer
	router = photosvc.Init(router)

	fmt.Printf("Listening to port *%s", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, router))

}
