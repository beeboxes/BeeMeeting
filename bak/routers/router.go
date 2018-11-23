package routers

import (
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func handler(c *gin.Context) {
	c.String(200, c.Request.URL.Path)
}

func init() {
	//Router.GET("/", handler)

	sessions := Router.Group("/sessions")
	{
		/*
	 new_user_session GET    /users/sign_in(.:format)      devise/sessions#new
		 user_session POST   /users/sign_in(.:format)      devise/sessions#create
 destroy_user_session DELETE /users/sign_out(.:format)     devise/sessions#destroy
*/
		sessions.GET("/new", handler)
		sessions.POST("/create", handler)
		sessions.DELETE("/destroy", handler)
	}

	users := Router.Group("/users")
	//users.Use(helpers.AuthRequired())
	{
		/*
 cancel_user_registration GET    /users/cancel(.:format)                                                                  devise/registrations#cancel
    new_user_registration GET    /users/sign_up(.:format)                                                                 devise/registrations#new
   edit_user_registration GET    /users/edit(.:format)                                                                    devise/registrations#edit
        user_registration PATCH  /users(.:format)                                                                         devise/registrations#update
                          PUT    /users(.:format)                                                                         devise/registrations#update
                          DELETE /users(.:format)                                                                         devise/registrations#destroy
                          POST   /users(.:format)                                                                         devise/registrations#create
		*/
		users.GET("/cancel", handler)
		users.GET("/sign_up", handler)
		users.GET("/edit", handler)
		users.PATCH("/", handler)
		users.PUT("/", handler)
		users.DELETE("/", handler)
		users.POST("/", handler)


		users.GET("/sign_in", handler)
		users.POST("/sign_in", handler)
		users.DELETE("/sign_out", handler)
	}

	passwords := Router.Group("/password")
	{
		/*
        new_user_password GET    /users/password/new(.:format)                                                            devise/passwords#new
       edit_user_password GET    /users/password/edit(.:format)                                                           devise/passwords#edit
            user_password PATCH  /users/password(.:format)                                                                devise/passwords#update
                          PUT    /users/password(.:format)                                                                devise/passwords#update
                          POST   /users/password(.:format)                                                                devise/passwords#create
		*/
		passwords.GET("/new", handler)
		passwords.GET("/edit", handler)
		passwords.PUT("/", handler)
		passwords.PATCH("/", handler)
		passwords.POST("/", handler)
	}

	dashboard := Router.Group("/dashboard")
	{
		dashboard.GET("/new", handler)
	}


}
