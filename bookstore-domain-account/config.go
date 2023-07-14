package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

const (
	configPathKey     = "config.path"
	defaultConfigPath = "config.toml"

	DbHostKey     = "db.host"
	DbPortKey     = "db.port"
	DbUserKey     = "db.user"
	DbPasswordKey = "db.password"
	DbDatabaseKey = "db.database"
)

var (
	Config = NewConfig()
)

func NewConfig() (v *viper.Viper) {
	v = viper.New()
	v.SetDefault(configPathKey, defaultConfigPath)
	defaultDbConfig(v)
	v.AutomaticEnv()
	// 通过ENV获取
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.SetConfigFile(v.GetString(configPathKey))
	err := v.ReadInConfig()
	if err != nil {
		klog.CtxErrorf(context.Background(), err.Error())
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		klog.CtxErrorf(context.Background(), "Config file changed: %s", in.Name)
	})
	return
}

func defaultDbConfig(v *viper.Viper) {
	v.SetDefault(DbHostKey, "127.0.0.1")
	v.SetDefault(DbPortKey, 3306)
	v.SetDefault(DbUserKey, "root")
	v.SetDefault(DbDatabaseKey, "bookstore")
}
