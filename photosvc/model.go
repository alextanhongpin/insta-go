package photosvc

import (
	"time"
)

type (
	// Photo struct {
	// 	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	// 	Alt       string        `json:"alt"`
	// 	CreatedAt time.Time     `json:"created_at"`
	// 	UpdatedAt time.Time     `json:"updated_at"`
	// 	Src       string        `json:"src"`
	// }
	Photo struct {
		Src     string `json:"src"`
		Caption string `json:"caption"`
	}

	allRequest struct {
		Query string `json:"query"`
	}
	allResponse struct {
		Data []Photo `json:"data"`
	}
	oneRequest struct {
		ID string `json:"id"`
	}
	oneResponse struct {
		Data Photo `json:"data"`
	}
	createRequest struct {
		PhotoID     string    `db="photo_id" json="photo_id"`
		Src         string    `db="src" json="src"`
		Caption     string    `db="caption" json="caption"`
		UserID      string    `db="user_id json="user_id"`
		DateCreated time.Time `db="date_created" json="date_created"`
		DateUpdated time.Time `db="date_updated" json="date_updated"`
	}
	createResponse struct {
		Status string `json:"status"`
	}
)
