package config

import (
	"os"
)

var (
	APIKEY    string
	NAMESPACE string
	EndPoint  = "https://api.testmail.app/api/json"
)

func InitConfig() {
	APIKEY = os.Getenv("APIKEY")
	NAMESPACE = os.Getenv("NAMESPACE")
}
