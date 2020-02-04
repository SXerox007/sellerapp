package constant

//constants
const (
	ENV_LOCAL  = "local"
	LOCAL_PORT = "50051"
)

/**
*
* Getting the environment local
*
**/
func GetLocal() string {
	return ENV_LOCAL
}

/**
*
* Getting the port local
*
**/
func GetPort() string {
	return LOCAL_PORT
}
