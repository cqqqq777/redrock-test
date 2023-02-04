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
		public.PUT("/password/forget", controller.ForgetPassword)          //忘记密码
		public.GET("/users/:uid/info", controller.GetUserInfo)             //获取用户信息
		public.GET("/books/:bid/comments", controller.BookCommentList)     //获取某本书的书评
		public.GET("/comments/:cid/replies", controller.ReplyList)         //获取书评的回复
		public.GET("/tags", controller.TagList)                            //获取所有标签
		public.GET("/tags/:tid/books", controller.TagBookList)             //获取某一标签下的书籍
		public.GET("/books/:bid", controller.BookDetail)                   //获取某本书的详情
		public.GET("/books", controller.SearchBook)                        //搜索某本书
		public.GET("/books/list", controller.BookList)                     //书籍列表
	}
	private := v1.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.PUT("/password", controller.RevisePassword)                    //修改密码
		private.PUT("/username", controller.ReviseUsername)                    //修改用户名
		private.PUT("/users/:uid/info", controller.UpdateUserInfo)             //修改用户信息
		private.POST("/comments", controller.CommentBook)                      //书评
		private.POST("/comments/reply", controller.ReplyComment)               //回复评论
		private.DELETE("/comments/:cid", controller.DeleteComment)             //删除评论
		private.POST("/comments/:cid/star", controller.StarComment)            //给评论点赞
		private.POST("/books/:bid/collection", controller.CollectBook)         //收藏某本书
		private.DELETE("/books/:bid/collection", controller.CancelCollectBook) //取消收藏
		private.POST("/books/:bid/star", controller.StarBook)                  //点赞某本书
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
