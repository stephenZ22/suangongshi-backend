package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Wechat   WechatConfig   `mapstructure:"wechat"`
	// OSS      OSSConfig      `mapstructure:"oss"` --- IGNORE ---
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type WechatConfig struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

var GlobalConfig *Config

func InitConfig() error {
	v := viper.New()

	v.SetConfigName("config")   // 配置文件名 (无扩展名)
	v.SetConfigType("yaml")     // 如果文件没有扩展名，需要指定类型
	v.AddConfigPath("./config") // 查找路径

	// --- 核心功能：支持环境变量 ---
	// 例如：设置环境变量 SG_DATABASE_DSN 会自动覆盖配置文件中的 database.dsn
	v.SetEnvPrefix("SG")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置失败: %s", err)
	}

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		return fmt.Errorf("解析配置失败: %s", err)
	}

	fmt.Printf("配置加载成功，当前环境: %s\n", GlobalConfig.Server.Mode)
	return nil
}
