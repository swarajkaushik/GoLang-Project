package routes

import(
	controllers "github.com/swarajkaushik/GoLang-Project/controllers"
	"github.com/swarajkaushik/Golang-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}