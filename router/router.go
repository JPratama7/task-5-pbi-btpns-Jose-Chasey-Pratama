package router

import (
	"btpn/controller"
	"btpn/database/pictures"
	"btpn/database/users"
	"btpn/helpers"
	"btpn/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func NewRouter(app *gin.Engine) {
	db := helpers.CreateGormConnection("root:@tcp(127.0.0.1:3306)/btpn?parseTime=true", 10, 100, 0)
	userRepo := users.New(db)
	PictureRepo := pictures.New(db)
	jwtMiddle := middleware.NewJWT([]byte("secret"), jwt.SigningMethodHS512)

	userControll := controller.NewUsers(userRepo, jwtMiddle)
	pictureControll := controller.NewPicture(PictureRepo, jwtMiddle)

	user := app.Group("/users")
	{
		user.POST("/register", userControll.Insert)
		user.PUT("/:userId", userControll.Update)
		user.DELETE("/:userId", userControll.Delete)
		user.GET("/login", userControll.Login)
	}

	picture := app.Group("/photos")
	{
		picture.POST("/", pictureControll.Insert)
		picture.PUT("/:pictureId", pictureControll.Update)
		picture.DELETE("/:pictureId", pictureControll.Delete)
		picture.GET("/", pictureControll.GetAll)
	}

}
