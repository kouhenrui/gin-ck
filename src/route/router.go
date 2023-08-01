package route

import (
	"gin-ck/src/middleware"
	"github.com/gin-gonic/gin"
)

/**
 * @ClassName router
 * @Description TODO
 * @Author khr
 * @Date 2023/7/28 15:42
 * @Version 1.0
 */
func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	r.Use(middleware.LoggerMiddleWare())           //日志捕捉
	r.Use(middleware.UnifiedResponseMiddleware())  //全局统一返回格式
	r.Use(middleware.GlobalErrorMiddleware())      //错误捕捉
	r.NoMethod(middleware.MethodNotAllowedHandler) //405，方法为找到
	r.Use(middleware.NotFoundHandler)              //404
	InitApi(r)                                     //挂载请求路径
	return r
}
