package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"

	"github.com/alextanhongpin/instago/authsvc"
	"github.com/alextanhongpin/instago/likesvc"
	"github.com/alextanhongpin/instago/photosvc"
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
	router = authsvc.Init(router)
	router = likesvc.Init(router)
	router = photosvc.Init(router)

	fmt.Printf("Listening to port *%s", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, router))
}
