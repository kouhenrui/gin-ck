package authService

import "gin-ck/src/dto/reqDto"

/**
 * @ClassName authService
 * @Description TODO
 * @Author khr
 * @Date 2023/8/1 14:21
 * @Version 1.0
 */
type AuthService struct{}

func (a *AuthService) Login(loginParam reqDto.Login, resData chan<- interface{}, resErr chan<- error) {

}
