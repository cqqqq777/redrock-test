package boot

import (
	"github.com/gin-gonic/gin"
	g "redrock-test/global"
	"redrock-test/middleware"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.Cors)
	r.Use(middleware.GinRecovery(g.Logger, true))
	v1 := r.Group("/api/v1")
	public := v1.Group("")
	{
		public.POST("/registration")
	}
	private := v1.Group("")
	private.Use(middleware.JWTAuth)
	{

	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
