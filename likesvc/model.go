package likesvc

import "time"

// Like is the model created to store the user/photo pair
type Like struct {
	UserID      string    `json:"user_id,omitempty"`
	PhotoID     string    `json:"photo_id,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
}
