package authsvc

import (
	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/middleware"
)

// Init returns a router than handles the endpoint for users
func Init(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{
		DB: &Service{common.GetDatabaseContext()},
	}

	// For all the API endpoints, add a middleware to validate the access token
	router.GET("/login", endpoint.LoginView)                                   // Render the login view
	router.POST("/login", endpoint.Login)                                      // Handles user login
	router.GET("/register", endpoint.RegisterView)                             // Render the register view
	router.POST("/register", endpoint.Register)                                // Handles user registration
	router.GET("/profile", middleware.Protect(endpoint.Profile))               // Render the current user profile
	router.GET("/users", endpoint.UsersView)                                   // Render the user view
	router.GET("/users/:id", endpoint.UserView)                                // Render the user view
	router.POST("/users/me", middleware.Protect(endpoint.UpdateUser))          // Update the user
	router.POST("/users/userphotos", middleware.Protect(endpoint.UploadPhoto)) // Update the user
	router.POST("/logout", endpoint.Logout)                                    // Clears the cookie and logs the user out
	return router
}
