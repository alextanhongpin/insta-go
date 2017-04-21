package likesvc

import "time"

// Like is the model created to store the user/photo pair
type Like struct {
	UserID      string    `json:"user_id"`
	PhotoID     string    `json:"photo_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// type All struct{}
// type One struct{}
// type Create struct {
//  UserID      string    `json:"user_id"`
//  PhotoID     string    `json:"photo_id"`
//  DateCreated time.Time `json:"date_created"`
//  DateUpdated time.Time `json:"date_updated"`
// }
// type Update struct{}
// type Delete struct {
//  ID string `json:"id"`
// }
