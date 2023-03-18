package global

import (
	"log"
	"micro/pkg/setting"
	"os"
	"time"
)

// 服务器配置
type ServerSettingS struct {
	RunMode      string        `json:"runMode"`
	HttpPort     string        `json:"httpPort"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
}

// 数据库配置
type DatabaseSettingS struct {
	DBType       string `yaml:"dbtype"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DBName       string `yaml:"dbname"`
	Charset      string `yaml:"charset"`
	ParseTime    bool   `yaml:"parse_time"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}
type RedisSettingS struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

func (d *DatabaseSettingS) Set(value string) {
	d.Password = value
}

func (d *RedisSettingS) Set(value string) {
	d.Password = value
}

func (d ServerSettingS) Set(value string) {

}

// 定义全局变量
var (
	ServerSetting   *ServerSettingS
	DatabaseSetting *DatabaseSettingS
	RedisSetting    *RedisSettingS
)

// 读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	// var MySQL_PASSWORD, Redis_PASSWORD string
	MySQL_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	log.Printf("获取环境变量 MySQL_PASSWORD:%s\n", MySQL_PASSWORD)
	if MySQL_PASSWORD != "" {
		DatabaseSetting.Password = MySQL_PASSWORD
	}
	Redis_PASSWORD := os.Getenv("REDIS_PASSWORD")
	log.Printf("获取环境变量 Redis_PASSWORD:%s\n", Redis_PASSWORD)
	if Redis_PASSWORD != "" {
		RedisSetting.Password = Redis_PASSWORD
	}
	log.Println("+++++++initial setting+++++++++++:")
	log.Println(*ServerSetting)
	log.Println(*DatabaseSetting)
	log.Println(*RedisSetting)
	return nil
}
