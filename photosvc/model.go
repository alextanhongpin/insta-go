package photosvc

import (
	"regexp"
	"time"

	"github.com/ventu-io/go-shortid"
)

const getPhotosByUser string = "/photos/:user_id"

// Photo is the model
type Photo struct {
	Src       string    `json:"src"`                   // The source of the image
	Caption   string    `json:"caption"`               // The caption of the image
	ID        string    `json:"id,omitempty"`          // The id of the image
	CreatedAt time.Time `json:"created_at,omitempty" ` // The date when the image is created
	UpdatedAt time.Time `json:"updated_at,omitempty"`  // The date when the image is updated
	UserID    string    `json:"user_id"`               // The id of the owner of the photo
	UserLikes []string  `json:"user_likes"`            // The list of users liking this
	LikeCount int64     `json:"like_count"`
	User
}

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
type Image struct {
	Name string
}

func (i Image) Path(path string) (string, string) {
	id, _ := shortid.Generate()
	reg := regexp.MustCompile(`\.(gif|jpg|jpeg|tiff|png)$`)

	match := reg.FindStringSubmatch(i.Name)
	ext := match[0]

	res := reg.ReplaceAllString(i.Name, "$1W")
	finalPath := path + res + id + ext

	// Return an absolute path and relative path
	return finalPath, "." + finalPath // unique id
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
