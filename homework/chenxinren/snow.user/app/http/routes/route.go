package routes

/**
 * 配置路由
 */
import (
	"snow.user/app/http/controllers"
	"snow.user/app/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/http/middleware"
    "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//api路由配置
func RegisterRoute(router *gin.Engine) {
	//middleware: 服务错误处理 => 生成请求id => access log
	router.Use(middlewares.ServerRecovery(), middleware.GenRequestId, middleware.GenContextKit, middleware.AccessLog())

	router.NoRoute(controllers.Error404)
	router.Any("/hello", controllers.HandleHello)
	router.POST("/test", controllers.HandleTest)
    router.POST("/test_validator", controllers.HandleTestValidator)
	
    //api版本
	v1 := router.Group("/v1")
	{
		v1.GET("/banner_list", controllers.GetBannerList)
	}
    
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/job/entry", controllers.EntryJob)

	router.POST("/v2/insert", controllers.Insert)
	router.POST("/v2/multiCreate", controllers.MulInsert)
	router.POST("/v2/updateNameByMobile", controllers.UpdateNameByMobile)
	router.POST("/v2/updateMulNameById", controllers.UpdateMulNameById)




	router.POST("/v2/getUserInfo", controllers.GetUserInfo)
}
