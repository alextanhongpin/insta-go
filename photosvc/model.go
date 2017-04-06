package photosvc

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
	// createRequest struct {
	// 	File
	// }
	createResponse struct {
		Status string `json:"status"`
	}
)
