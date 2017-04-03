package photosvc

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Photo struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Alt       string        `json:"alt"`
		CreatedAt time.Time     `json:"created_at"`
		UpdatedAt time.Time     `json:"updated_at"`
		Src       string        `json:"src"`
	}

	allRequest struct {
		Query string `json:"query"`
	}
	allResponse struct {
		Data []Photo `json:"data"`
	}
)
