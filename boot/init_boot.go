package boot

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/dao/mongodb"
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/gofuncchan/ginger/dao/redis"
	"github.com/gofuncchan/ginger/logger"
	"github.com/gofuncchan/ginger/mq"
	"github.com/gofuncchan/ginger/oauth2"
	"github.com/gofuncchan/ginger/oss"
	"github.com/gofuncchan/ginger/util/jwt"
	"io/ioutil"
	"os"
)

// 系统启动时的运行各种初始化
func Init() {
	// 从启动参数获取配置文件路径
	confPath := getConfigPath()

	// 初始化配置参数
	config.Init(confPath)

	// 初始化zap日志
	logger.Init()

	// 初始化三方登录管理器
	oauth2.InitOAuth2Manager()

	// 消息中间件初始化
	mq.Init()

	// jwt初始化
	jwt.JwtInit()

	// 如开启mongo记录日志，则mongo连接必须初始化
	// 初始化mongodb连接池
	mongodb.Init()
	// 初始化Mysql连接池
	mysql.Init()
	// 初始化Redis连接池
	redis.Init()

	// 对象存储连接初始化
	oss.Init()

}

func getConfigPath() string {
	// 启动时获取配置文件目录参数,如没有则在当前目录的config文件夹查找
	var configRootPath string
	if len(os.Args) > 1 {
		configRootPath = os.Args[1]
	}
	if configRootPath == "" {
		configRootPath = "./config"
	}

	// 先从环境变量获取环境的信息 (debug|release|test)
	// export GIN_MODE=release
	var currentEnv string
	currentEnv = gin.Mode()
	if currentEnv == "" {
		currentEnv = config.DefaultEnv
	}
	confPath := configRootPath + "/" + currentEnv + "/"

	dirs, err := ioutil.ReadDir(confPath)
	if os.IsNotExist(err) {
		common.EF(errors.New("the args confPath `" + confPath + "` is not exist"))
	}
	if len(dirs) == 0 {
		common.EF(errors.New("not any yaml file in this directory"))
	}

	return confPath
}
