package environment

import (
	"os"
	constant "sellerapp/base/constants"
)

/**
*
* to get the environment local/or something else
*
**/
func GetEnv() string {
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = constant.GetLocal()
	}
	return env
}

/**
*
*  to get the port if its blank then set to port
*
**/
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = constant.GetPort()
	}
	return port
}
