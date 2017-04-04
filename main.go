package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/photosvc"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	// Load the config
	conf := common.GetConfig()
	router := httprouter.New()
	common.InitDatabase()

	photosvc.Init(router)
	router.GET("/", Index)

	fmt.Printf("Listening to port *%s", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, router))
}
