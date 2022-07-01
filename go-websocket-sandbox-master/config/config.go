package config

import (
	"fmt"
	"go/build"
)

func GetPort() string {
	v := getViper()
	return fmt.Sprintf(":%s", v.GetString("server.port"))
}

func GetProjectPath() string {
	v := getViper()
	return fmt.Sprintf("%s/src%s",
		build.Default.GOPATH,
		v.GetString("projectDir"),
	)
}

func GetDataBaseAccess() string {
	v := getViper()
	connection := fmt.Sprintf(`%s:%s@tcp([%s]:3306)/%s?charset=utf8&parseTime=True&loc=`,
		v.GetString("database.userName"),
		v.GetString("database.password"),
		v.GetString("database.host"),
		v.GetString("database.name"),
	)
	connection += `Asia%2FTokyo`
	return connection
}
