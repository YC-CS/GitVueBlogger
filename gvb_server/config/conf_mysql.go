package config

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
	LogLevel string `json:"log_level"` //日志等级：debug (输出全部sql), dev, release
}
