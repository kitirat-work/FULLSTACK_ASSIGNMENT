package config

import (
	"os"
	"reflect"
)

type Configuration struct {
	SQLConnStr string `env:"SQL_CONN_STR" default:"myuser:mypassword@tcp(db:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"`
	PORT       string `env:"PORT" default:"8080"`
}

func New() Configuration {
	conf := Configuration{}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		envKey := fieldType.Tag.Get("env")
		envValue, ok := os.LookupEnv(envKey)
		switch ok {
		case true:
			field.SetString(envValue)
		case false:
			field.SetString(fieldType.Tag.Get("default"))
		}
	}
	return conf
}
