package boot

import (
	"github.com/gin-gonic/gin"
	"redrock-test/controller"
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
		public.POST("/verification", controller.PostVerification)          //发送验证码
		public.POST("/registration", controller.Register)                  //注册
		public.POST("/login/password", controller.Login)                   //登录（密码）
		public.POST("/login/verification", controller.LoginByVerification) //登录（验证码）
	}
	private := v1.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.PUT("/password", controller.RevisePassword) //修改密码
		private.PUT("/username", controller.ReviseUsername) //修改用户名
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
