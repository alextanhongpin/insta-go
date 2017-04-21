package photosvc

import "time"

const getPhotosByUser string = "/photos/:user_id"

// Photo is the model
type Photo struct {
	Src       string    `json:"src"`        // The source of the image
	Caption   string    `json:"caption"`    // The caption of the image
	ID        int       `json:"id"`         // The id of the image
	Alt       string    `json:"alt"`        // The filename
	CreatedAt time.Time `json:"created_at"` // The date when the image is created
	UpdatedAt time.Time `json:"updated_at"` // The date when the image is updated
}

// allRequest request photos that belongs to a user
type allRequest struct {
	UserID string `json:"user_id"`
}

// allResponse response all photos by users
type allResponse struct {
	Data []Photo `json:"data"`
}
type oneRequest struct {
	ID string `json:"id"`
}
type oneResponse struct {
	Data Photo `json:"data"`
}
type createRequest struct {
	FileName string
}
type createResponse struct {
	Status string `json:"status"`
}
