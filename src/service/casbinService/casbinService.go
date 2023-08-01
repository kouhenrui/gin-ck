package casbinService

import (
	"gin-ck/src/dto/reqDto"
	"gin-ck/src/global"
	util "gin-ck/src/utils"
)

var err error

/**
 * @ClassName casbinService
 * @Description TODO
 * @Author khr
 * @Date 2023/5/5 16:16
 * @Version 1.0
 */
var casbinDb = global.CasbinDb

/*
 * @MethodName CasbinPolicyAdd
 * @Description
 * @Author khr
 * @Date 2023/5/5 16:38
 */

func CasbinPolicyAdd(casbinAdd reqDto.CasbinAdd, resErr chan error, resData chan string) {
	_, err := casbinDb.AddPolicy(casbinAdd)
	if err != nil {
		resErr <- err
	}
	resData <- util.PERMISSION_ADD_SUCCESS
}

/*
 * @MethodName CasbinGroupDel
 * @Description
 * @Author khr
 * @Date 2023/5/5 16:38
 */

func CasbinGroupDel() {
	//casbinDb.DeleteRole()
}

func CheckRuleName(id uint) (error, string) {

	return nil, ""
}
