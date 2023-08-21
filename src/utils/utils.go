package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-ck/src/dto/reqDto"
	"gin-ck/src/global"
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"reflect"
	"regexp"
)

// json格式化数据
func Marshal(user interface{}) []byte {
	ub, _ := json.Marshal(user)
	return ub
}
func UnMarshal(r []byte, res interface{}) (bool, interface{}) {
	err := json.Unmarshal(r, &res)
	if err != nil {
		return false, REDIS_INFORMATION_ERROR
	}
	return true, res
}

/*
 * @MethodName 参数验证
 * @Description
 * @Author khr
 * @Date 2023/8/21 10:21
 */
func GetValidate(err error, obj any) error {

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return invalid
	}
	//反射获取标签的注释
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		return errs
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return errors.New(msg)
			}
		}
	}
	return err
}

/*
 * @MethodName ExistIn
 * @Description 判断参数是否存在
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func ExistIn(param string, paths []string) bool {
	for _, v := range paths {
		if param == v {
			return true
		}
	}
	return false
}

/*
 * @MethodName FuzzyMatch
 * @Description 正则模糊匹配路径
 * @Author khr
 * @Date 2023/5/9 16:25
 */
func FuzzyMatch(param string, paths []string) bool {
	for _, y := range paths {
		if regexp.MustCompile(y).MatchString(param) {

			//fmt.Print("匹配道路进了")
			return true
		}

	}
	return false
}

var CaptchaStore base64Captcha.Store = &RedisStore{}

/*
 * @MethodName CreateCaptcha
 * @Description 生成图片验证
 * @Author khr
 * @Date 2023/5/8 10:44
 */

func CreateCaptcha() (error, interface{}) {
	var newCaptcha reqDto.Captcha
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          80,                                     //高度
		Width:           240,                                    //宽度
		NoiseCount:      0,                                      //干扰数
		ShowLineOptions: 6,                                      //展示个数
		Length:          6,                                      //长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microfiche.ttc"}, // 字体wqy-microhei.ttc
	}
	driver = driverString.ConvertFonts()
	//生成验证码
	c := base64Captcha.NewCaptcha(driver, CaptchaStore)
	id, content, err := c.Generate()
	newCaptcha.Id = id
	newCaptcha.Content = content
	if err != nil {
		fmt.Println("生成有错:", err)
		return err, nil
	}
	return nil, &newCaptcha

}

/*
 * @MethodName VerifyCaptcha
 * @Description 验证图片验证码
 * @Author khr
 * @Date 2023/5/8 10:45
 */

func VerifyCaptcha(capt reqDto.Captcha) bool {
	// id 验证码id
	// answer 需要校验的内容
	// clear 校验完是否清除
	if CaptchaStore.Verify(capt.Id, capt.Content, true) {

		fmt.Println("验证正确")
		return true
	} else {
		fmt.Println("验证cuowu ")
		return false
	}
	//return nil, store.Verify(capt.Id, capt.Content, true)
}

type RedisStore struct{}

func (s *RedisStore) Set(id string, value string) error {
	key := global.Captcha + id
	return global.SetRedis(key, []byte(value), global.CaptchaExp) //Redis.Set(ctx, key, value, global.CaptchaExp).Err()
}
func (s *RedisStore) Get(id string, clear bool) string {
	key := global.Captcha + id
	val := global.GetRedis(key)
	err := global.DelRedis(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
