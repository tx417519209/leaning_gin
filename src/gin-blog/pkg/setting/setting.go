package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	JwtIssuer       string
	PageSize        int
	RuntimeRootPath string
	ImagePrefixUrl  string
	ImageSavePath   string
	ImageMaxSize    int64
	ImageAllowExts  []string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	LogTimeFormat   string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	Host         string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to parse 'conf/app.ini':%s", err.Error())
	}
	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Failed to mapto  AppSetting :%s ", err.Error())
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Failed to mapto  ServerSetting :%s ", err.Error())
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}
