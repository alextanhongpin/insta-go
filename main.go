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
	// New router
	router := httprouter.New()

	// Serve files from the static directory
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	// Init database
	common.InitDatabase()

	// Init routes (can add feature toggle later)
	router = authsvc.Init(router)
	router = likesvc.Init(router)
	router = photosvc.Init(router)

	Port := common.Config.Port
	fmt.Printf("Listening to port *%s", Port)
	log.Fatal(http.ListenAndServe(Port, router))
}
