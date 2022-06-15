package configs

type Config struct {
	Fiber    Fiber
	Database MySQL
	Redis    Redis
}

type Fiber struct {
	Host              string
	Port              string
	ServerReadTimeOut string
}

type MySQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DBNumber string
}
