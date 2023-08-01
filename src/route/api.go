package route

import (
	"gin-ck/src/module/auth"
	"github.com/gin-gonic/gin"
)

/**
 * @ClassName api
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:18
 * @Version 1.0
 */

func InitApi(route *gin.Engine) {
	api := route.Group("/api")
	{
		authModule := api.Group("/auth")

		{
			authModule.POST("login", auth.Login)
		}
	}

}
