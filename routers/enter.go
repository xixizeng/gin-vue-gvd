package routers

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()

	//创建了一个以api开头的分组
	apiGroup := router.Group("api")

	//apiGroup.GET("xxx")
	//等价于
	//router.GET("/api/xxx")

	//将api分组赋值给了RouterGroup
	routerGroup := RouterGroup{apiGroup}
	routerGroup.UserRouter()
	return router
}
