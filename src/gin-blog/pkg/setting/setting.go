package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg           *ini.File
	RunMode       string
	HTTPPort      int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	PageSize      int
	JwtSecret     string
	JwtIssuer     string
	DBType        string
	DBName        string
	DBUser        string
	DBPwd         string
	DBHost        string
	DBTablePrefix string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to parse 'conf/app.ini':%s", err.Error())
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadDataBase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Failed to get section 'server':%s", err.Error())
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(9000)

	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
func LoadDataBase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	DBType = sec.Key("TYPE").String()
	DBName = sec.Key("NAME").String()
	DBUser = sec.Key("USER").String()
	DBPwd = sec.Key("PASSWORD").String()
	DBHost = sec.Key("HOST").String()
	DBTablePrefix = sec.Key("TABLE_PREFIX").String()
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	JwtIssuer = sec.Key("JWT_ISSUER").MustString("xing")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
