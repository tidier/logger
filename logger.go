package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	Console = "console"
	File    = "file"
)

var (
	level  = zap.DebugLevel
	target = Console
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Init() {
	w := zapcore.AddSync(&lumberjackv2.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxAge, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
	})

	var writeSyncer zapcore.WriteSyncer
	if target == Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}
	if target == File {
		writeSyncer = zapcore.NewMultiWriteSyncer(w)
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		writeSyncer,
		level,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}
