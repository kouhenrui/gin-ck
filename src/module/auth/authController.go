package auth

import (
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
func Login(c *gin.Context) {
	var logins login
	if err := c.ShouldBindJSON(&logins); err != nil {
		//errs, ok := err.(validator.ValidationErrors)
		//if !ok {
		//	c.Error(err)
		//	return
		//}
		c.AbortWithError(http.StatusBadRequest, err)
		//c.Error(err)
		return
	}
	c.Set("response", logins.Name)
	return
	//c.AbortWithError(500, errors.New("我不知道哪里错了"))
	//c.Error(errors.New("我不知道哪里错了"))
	//return
}

type login struct {
	Name string `json:"name,omitempty" binding:"required" message:"名称不为空"`
}
