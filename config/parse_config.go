package config

import (
	"github.com/gofuncchan/ginger/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Base struct {
	Env        string `yaml:"env"`
	ListenPort int64  `yaml:"listen"`
}

type Log struct {
	LogDir                     string `yaml:"logDir"`
	LogMaxDayCount             int    `yaml:"maxDayCount"`
	LogMongoHookOn             bool   `yaml:"logMongoHookOn"`
	LogMongoCollection         string `yaml:"logMongoCol"`
	LogMongoExpireAfterSeconds int    `yaml:"logMongoExpire"`
	WithRotationTime           int    `yaml:"withRotationTime"`
}

type Mysql struct {
	DbHost                  string `yaml:"dbHost"`
	DbPort                  int64  `yaml:"dbPort"`
	DbUser                  string `yaml:"dbUser"`
	DbPasswd                string `yaml:"dbPasswd"`
	DbName                  string `yaml:"dbName"`
	ConnMaxLifetime         int64  `yaml:"connMaxLifetime"`
	MaxIdleConns            int64  `yaml:"maxIdleConns"`
	MaxOpenConns            int64  `yaml:"maxOpenConns"`
	ChartSet                string `yaml:"charset"`
	AllowCleartextPasswords bool   `yaml:"allowCleartextPasswords"`
	InterpolateParams       bool   `yaml:"interpolateParams"`
	Timeout                 int64  `yaml:"timeout"`
	ReadTimeout             int64  `yaml:"readTimeout"`
	ParseTime               bool   `yaml:"parseTime"`
	PING                    bool   `yaml:"ping"`
}

type Redis struct {
	DbHost      string `yaml:"dbHost"`
	DbPort      int64  `yaml:"dbPort"`
	DbAuth      bool   `yaml:"dbAuth"`
	DbPasswd    string `yaml:"dbPasswd"`
	MaxActive   int64  `yaml:"maxActive"`
	MaxIdle     int64  `yaml:"maxIdle"`
	IdleTimeout int64  `yaml:"idleTimeout"`
}

type Mongodb struct {
	DbHost   string `yaml:"dbHost"`
	DbPort   int64  `yaml:"dbPort"`
	DbUser   string `yaml:"dbUser"`
	DbPasswd string `yaml:"dbPasswd"`
	DbName   string `yaml:"dbName"`
}

var (
	BaseConf  Base
	LogConf   Log
	MysqlConf Mysql
	RedisConf Redis
	MongoConf Mongodb
)

// 动态参数配置项，编译后可携yaml配置文件启动
func Init(confPath string){

	baseConfFile, err := ioutil.ReadFile(confPath + "/base.yaml")
	common.EF(err)
	err = yaml.Unmarshal(baseConfFile, &BaseConf)
	common.EF(err)

	logConfFile, err := ioutil.ReadFile(confPath + "/log.yaml")
	common.EF(err)
	err = yaml.Unmarshal(logConfFile, &LogConf)
	common.EF(err)

	mysqlConfFile, err := ioutil.ReadFile(confPath + "/mysql.yaml")
	common.EF(err)
	err = yaml.Unmarshal(mysqlConfFile, &MysqlConf)
	common.EF(err)

	RedisConfFile, err := ioutil.ReadFile(confPath + "/redis.yaml")
	common.EF(err)
	err = yaml.Unmarshal(RedisConfFile, &RedisConf)
	common.EF(err)

	MongoConfFile, err := ioutil.ReadFile(confPath + "/mongodb.yaml")
	common.EF(err)
	err = yaml.Unmarshal(MongoConfFile, &MongoConf)
	common.EF(err)

}