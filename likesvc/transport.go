package likesvc

import (
	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"
)

func Init(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{
		DB: &Service{common.GetDatabaseContext()},
	}

	router.POST("/api/likes", endpoint.Like)
	router.DELETE("/api/likes", endpoint.Unlike)
	router.GET("/api/likes/:photoID/count", endpoint.Count)
	return router
}
