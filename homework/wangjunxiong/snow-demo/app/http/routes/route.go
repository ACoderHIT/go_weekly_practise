package routes

/**
 * 配置路由
 */
import (
	"github.com/gin-gonic/gin"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/controllers"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/controllers/ucentercontroller"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/controllers/usercontroller"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/middlewares"
	"github.com/qit-team/snow-core/http/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//api路由配置
func RegisterRoute(router *gin.Engine) {
	//middleware: 服务错误处理 => 生成请求id => access log
	router.Use(middlewares.ServerRecovery(), middleware.GenRequestId, middleware.GenContextKit, middleware.AccessLog())

	router.NoRoute(controllers.Error404)
	router.GET("/hello", controllers.HandleHello)
	router.POST("/test", controllers.HandleTest)
	router.POST("/test_validator", controllers.HandleTestValidator)

	//api版本
	v1 := router.Group("/v1")
	{
		v1.GET("/banner_list", controllers.GetBannerList)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/user/login", usercontroller.HandlePostUserLogin)
	router.POST("/user/update_user_info", usercontroller.HandlUpdateUserInfo)


	ucenterGroup := router.Group("ucenter")
	{
		ucenterGroup.GET("/user_info",ucentercontroller.HandleGetUserInfo)
	}

	userServiceGroup := router.Group("/user")
	{
		userServiceGroup.GET("info",usercontroller.HandleInfoByGet)
		userServiceGroup.POST("info",usercontroller.HandleInfoByGet)
	}
}
