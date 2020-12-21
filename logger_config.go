package logger

import "go.uber.org/zap/zapcore"

type TLoggerConfig struct {
	Filename   string `json:"filename"`
	MaxSize    int64  `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
}

var config = TLoggerConfig{
	Filename:   "logs/log.log",
	MaxSize:    1024,
	MaxBackups: 10,
	MaxAge:     7,
}

func InitLogger(cfg *TLoggerConfig) {
	if cfg != nil {
		config.Filename = cfg.Filename
		config.MaxAge = cfg.MaxAge
		config.MaxBackups = cfg.MaxBackups
		config.MaxAge = cfg.MaxAge
	}
}

func SetTarget(ta string) {
	target = ta
}

func SetLevel(l zapcore.Level) {
	level = l
}
