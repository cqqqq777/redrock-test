package entry

import (
	"github.com/gin-gonic/gin"
	"redrock-test/business/controller"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/enroll", controller.Enroll)
	r.POST("/login", controller.Login)
}
