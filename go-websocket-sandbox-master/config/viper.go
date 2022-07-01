package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var vp *viper.Viper

func read() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	vp.AddConfigPath("../")
	vp.AddConfigPath("../../")
	vp.AddConfigPath("../../../")
	vp.AddConfigPath("../../../../")
	vp.AddConfigPath("../../../../../")
	err := vp.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
	return vp
}

func getViper() *viper.Viper {
	if vp == nil {
		newVp := read()
		vp = newVp
	}
	return vp
}
