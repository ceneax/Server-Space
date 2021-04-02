package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
	"xcdh.space/space/utils"
)

var (
	Cfg *ini.File

	RunMode string

	ServerHost   string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	LogFilepath  string
	LogFilename  string

	PageSize    int
	IdentityKey string

	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
)

func init() {
	var err error
	var configPath string
	var appPath string
	workPath, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	var filename = "config.ini"
	configPath = filepath.Join(workPath, filename)
	if !utils.FileExists(configPath) {
		if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			panic(err)
		}
		configPath = filepath.Join(appPath, filename)
		if !utils.FileExists(configPath) {
			configPath = "../config.ini"
			if !utils.FileExists(configPath) {
				configPath = "/etc/space/config.ini"
			}
		}
	}
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	ServerHost = sec.Key("SERVER_HOST").MustString("localhost:8080")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	LogFilepath = sec.Key("LOG_FILEPATH").MustString("./")
	LogFilename = sec.Key("LOG_FILENAME").MustString("log.txt")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	IdentityKey = sec.Key("IDENTITY_KEY").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	Type = sec.Key("TYPE").MustString("mysql")
	User = sec.Key("USER").MustString("test")
	Password = sec.Key("PASSWORD").MustString("test")
	Host = sec.Key("HOST").MustString("127.0.0.1")
	Port = sec.Key("PORT").MustString("3306")
	Name = sec.Key("NAME").MustString("space")
}
