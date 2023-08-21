package auth

import (
	"gin-ck/src/dto/reqDto"
	"gin-ck/src/service/authService"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @ClassName authController
 * @Description TODO
 * @Author khr
 * @Date 2023/7/31 16:56
 * @Version 1.0
 */

var (
	authServiceImpl = authService.AuthService{}
)

func Login(c *gin.Context) {
	var logins reqDto.Login
	if err := c.ShouldBindJSON(&logins); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//异步线程操作
	resErr := make(chan error)
	resData := make(chan interface{})
	go authServiceImpl.Login(logins, resData, resErr)
	endErr := <-resErr
	endData := <-resData
	if endErr != nil {
		c.Error(endErr)
		return
	}
	c.Set("res", endData)
	return
}
