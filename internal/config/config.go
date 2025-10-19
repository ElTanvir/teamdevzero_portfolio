package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort              string        `mapstructure:"APP_PORT"`
	Environment          string        `mapstructure:"APP_ENV"`

}

func Load() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AllowEmptyEnv(true)
	v.SetDefault("APP_PORT", "8123")
	v.SetDefault("APP_ENV", "development")
	// Bind environment variables
	bindEnvs(v, Config{})

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
	return &cfg
}
func bindEnvs(v *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		vf := ifv.Field(i)
		tf := ift.Field(i)
		tv := tf.Tag.Get("mapstructure")
		if tv == "" {
			continue
		}
		switch vf.Kind() {
		case reflect.Struct:
			bindEnvs(v, vf.Interface(), append(parts, tv)...)
		default:
			_ = v.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
