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

	AccountClientHostPost   = "account.client.hostport"
	AccountClientConnNum    = "account.client.connnum"
	PaymentClientHostPost   = "payment.client.hostport"
	PaymentClientConnNum    = "payment.client.connnum"
	WarehouseClientHostPost = "warehouse.client.hostport"
	WarehouseClientConnNum  = "warehouse.client.connnum"
)

var (
	Config = NewConfig()
)

func NewConfig() (v *viper.Viper) {
	v = viper.New()
	v.SetDefault(configPathKey, defaultConfigPath)
	defaultClientConfig(v)
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

func defaultClientConfig(v *viper.Viper) {
	v.SetDefault(AccountClientHostPost, []string{"[::1]:8810"})
	v.SetDefault(AccountClientConnNum, 1)
	v.SetDefault(PaymentClientHostPost, []string{"[::1]:8812"})
	v.SetDefault(PaymentClientConnNum, 1)
	v.SetDefault(WarehouseClientHostPost, []string{"[::1]:8811"})
	v.SetDefault(WarehouseClientConnNum, 1)
}
