package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize  int
	JwtSecret string

	//mysql
	DB_TYPE string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PWD  string
	DB_NAME string

	//redis
	REDIS_HOST string
	REDIS_PORT string

	//log
	LOG_FILE_PATH string
	LOG_FILE_NAME string

	//ceph
	CEPH_HOST        string
	CEPH_USERNAME    string
	CEPH_ACCESSKEY   string
	CEPH_SECRETKEY   string
	CEPH_BUCKET_NAME string
	CEPH_BLOCK_SIZE  string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/appConfig.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadMysql()
	LoadRedis()
	LoadLog()
	LoadCeph()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadMysql() {
	sec, err := Cfg.GetSection("mysql")
	if err != nil {
		log.Fatalf("Fail to get section 'mysql': %v", err)
	}

	DB_TYPE = sec.Key("DB_TYPE").MustString("mysql")
	DB_HOST = sec.Key("DB_HOST").MustString("127.0.0.1")
	DB_PORT = sec.Key("DB_PORT").MustString("3306")
	DB_USER = sec.Key("DB_USER").MustString("root")
	DB_PWD = sec.Key("DB_PWD").MustString("123456")
	DB_NAME = sec.Key("DB_NAME").MustString("test")
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}

	REDIS_HOST = sec.Key("REDIS_HOST").MustString("127.0.0.1")
	REDIS_PORT = sec.Key("REDIS_PORT").MustString("6379")
}

func LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}

	LOG_FILE_PATH = sec.Key("LOG_FILE_PATH").MustString("/logs/")
	LOG_FILE_NAME = sec.Key("LOG_FILE_NAME").MustString("gin.log")
}

func LoadCeph() {
	sec, err := Cfg.GetSection("ceph")
	if err != nil {
		log.Fatalf("Fail to get section 'ceph': %v", err)
	}

	LOG_FILE_PATH = sec.Key("CEPH_HOST").MustString("")
	LOG_FILE_NAME = sec.Key("CEPH_USERNAME").MustString("eb-kunming")
	LOG_FILE_NAME = sec.Key("CEPH_ACCESSKEY").MustString("")
	LOG_FILE_NAME = sec.Key("CEPH_SECRETKEY").MustString("")
	LOG_FILE_NAME = sec.Key("CEPH_BUCKET_NAME").MustString("eb-kunming")
	LOG_FILE_NAME = sec.Key("CEPH_BLOCK_SIZE").MustString("5*1024*1024")
}
