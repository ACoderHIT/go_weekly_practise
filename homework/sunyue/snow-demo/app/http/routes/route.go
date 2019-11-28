package routes

/**
 * 配置路由
 */
import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/http/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/http/controllers"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/http/middlewares"
)

//api路由配置
func RegisterRoute(router *gin.Engine) {
	//middleware: 服务错误处理 => 生成请求id => access log
	router.Use(middlewares.ServerRecovery(), middleware.GenRequestId, middleware.GenContextKit, middleware.AccessLog())

	router.NoRoute(controllers.Error404)
	router.GET("/hello", controllers.HandleHello)
	router.POST("/test", controllers.HandleTest)
    router.POST("/test_validator", controllers.HandleTestValidator)
	router.POST("/quene/process", controllers.HandleQueneProcess)
	router.POST("/user/query", controllers.HandleUserQueryProcess)
	
    //api版本
	v1 := router.Group("/v1")
	{
		v1.GET("/banner_list", controllers.GetBannerList)
	}
    
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/week3/abservice/a/get", controllers.HandleAGetDemoProcess)
	router.POST("/week3/abservice/a/post", controllers.HandleAPostDemoProcess)
	router.POST("/week3/abservice/a/post_json", controllers.HandleAPostJsonDemoProcess)
	router.GET("/week3/abservice/b/get", controllers.HandleBGetRequestProcess)
	router.POST("/week3/abservice/b/post", controllers.HandleBPostRequestProcess)
	router.POST("/week3/abservice/b/post_json", controllers.HandleBPostJsonRequestProcess)
}
