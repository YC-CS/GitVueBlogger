package config

import "strconv"

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Config   string `json:"config"` //高级配置，例如 charset
	Db       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	LogLevel string `json:"log_level"` //日志等级：debug (输出全部sql), dev, release
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.Config
}
