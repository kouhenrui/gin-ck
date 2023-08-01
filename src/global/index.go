package global

import (
	"context"
	"fmt"
	"gin-ck/src/global/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

/**
 * @ClassName index
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:36
 * @Version 1.0
 */
var (
	v      *viper.Viper
	err    error
	day    = time.Hour * 24
	hour   = time.Hour
	minute = time.Minute
)

var (
	Captcha     string //redis缓存验证码前缀
	Port        string //程序使用端口
	HttpVersion bool   //版本控制

	InterceptPrefix string
	CaptchaExp      time.Duration
	ctx             = context.Background()
	RedisClient     *redis.Client
	MysqlDClient    *gorm.DB
	JWTKEY          = "12"
	LANGUAGE        = "zh"
	IpAccess        = []string{"127.0.0.1"}
	WriteList       = []string{}
	EtcdArry        = []string{"192.168.245.22:9092"}
)
var (
	MysqlConfig    conf.MysqlConf    //连接实例化参数
	RedisConfig    conf.RedisConf    //连接实例化参数
	RabbitMQConfig conf.RabbitmqConf //连接实例化参数
	LogConf        conf.LogCof       //连接实例化参数
	CabinConfig    conf.CabinConf    //连接实例化参数
	ClickConfig    conf.ClickConf    //连接实例化参数
)

func init() {
	log.Println("实例化配置文件")
	// 构建 Viper 实例
	v = viper.New()
	v.SetConfigFile("conf.yaml") // 指定配置文件路径
	v.SetConfigName("conf")      // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")      // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	v.AddConfigPath(".") // 还可以在工作目录中查找配置
	// 查找并读取配置文件
	if err = v.ReadInConfig(); err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig() //开启监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})

	viperLoadConf()

}
func viperLoadConf() {
	//读取单条配置文件
	Port = v.GetString("port")
	//设置http1.0还是2.0
	HttpVersion = v.GetBool("protocol")
	Captcha = v.GetString("captcha")
	InterceptPrefix = v.GetString("InterceptPrefix")
	CaptchaExp = time.Duration(v.GetInt("CaptchaExp")) * minute
	//日志路径及名称设置
	logConfig := v.GetStringMap("log")

	//读取mysql,redis,rabbitmq,casbin
	mysql := v.GetStringMap("mysql") //读取MySQL配置
	redis := v.GetStringMap("redis") //读取redis配置
	mq := v.GetStringMap("rabbitmq") //读取rabbitmq配置
	cn := v.GetStringMap("cabin")    //读取casbin配置
	ck := v.GetStringMap("click")    //读取click house配置
	//map转struct
	mapstructure.Decode(mysql, &MysqlConfig)
	mapstructure.Decode(redis, &RedisConfig)
	mapstructure.Decode(mq, &RabbitMQConfig)
	mapstructure.Decode(logConfig, &LogConf)
	mapstructure.Decode(cn, &CabinConfig)
	mapstructure.Decode(ck, &ClickConfig)

	//mapstructure.Decode(ca, &CaptchaConf)
	//etcd := v.GetStringSlice("etcd")
	//kafka := v.GetStringSlice("kafka")
	//oracle := v.GetStringSlice("oracle")
	//EtcdArry = append(EtcdArry, etcd...)
	//KafkaArry = append(KafkaArry, kafka...)
	log.Println("全局配置文件信息读取无误,开始载入")
	//Dbinit()         //mysql初始化
	//Redisinit() //redis初始化
	Loginit() //日志初始化
	//CasbinInit()     //rbac初始化
	//OracleInit()     //Oracle初始化
	//ClickhouseInit() //click house初始化
	//EtcdInit()
	//Mqinit()
}
