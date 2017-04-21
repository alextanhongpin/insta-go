package authsvc

import (
	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/middleware"
)

// Init returns a router than handles the endpoint for users
func Init(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	service := &Service{common.GetDatabaseContext()}

	// For all the API endpoints, add a middleware to validate the access token
	router.GET("/login", endpoint.LoginView(service))                 // Render the login view
	router.POST("/login", endpoint.Login(service))                    // Handles user login
	router.GET("/register", endpoint.RegisterView(service))           // Render the register view
	router.POST("/register", endpoint.Register(service))              // Handles user registration
	router.GET("/me", middleware.Validate(endpoint.Profile(service))) // Render the current user profile
	router.GET("/users/:id", endpoint.UserView(service))              // Render the user view
	router.POST("/logout", endpoint.Logout(service))                  // Clears the cookie and logs the user out
	return router
}
