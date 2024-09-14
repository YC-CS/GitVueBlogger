package config

type Logger struct {
	Level        string `json:"level"`
	Prefix       string `json:"prefix"`
	Director     string `json:"director"`
	ShowLine     bool   `json:"show_line"`      //是否显示行号
	LogInConsole bool   `json:"log_in_console"` //是否显示打印的路径
}
