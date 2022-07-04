package logger

import (
	"log"
	"os"

	"github.com/blastertwist/flag-dash/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()
	Debug(msg string, args ...zapcore.Field)
	Info(msg string, args ...zapcore.Field)
	Warn(msg string, args ...zapcore.Field)
	Error(msg string, args ...zapcore.Field)
	DPanic(msg string, args ...zapcore.Field)
	Panic(msg string, args ...zapcore.Field)
	Fatal(msg string, args ...zapcore.Field)
}

type apiLogger struct{
	cfg *config.Config
	zapLogger *zap.Logger
}

func NewLogger(cfg *config.Config) *apiLogger{
	return &apiLogger{cfg:cfg}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *apiLogger) getWriters() (zapcore.WriteSyncer, zapcore.WriteSyncer){
	var lumberjackConfig = l.cfg.LumberjackLogger

	if l.cfg.Server.Mode == "Development" {
		path := "./logs/dev"
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
        	log.Fatal(err)
    	}
		lumberjackLogger := &lumberjack.Logger{
			Filename: path + "/dev-log.log",
			MaxSize: int(lumberjackConfig.MaxSize),
			MaxAge: int(lumberjackConfig.MaxAge),
			MaxBackups: int(lumberjackConfig.MaxBackups),
			Compress: lumberjackConfig.Compress,
		}

		return zapcore.AddSync(os.Stderr), zapcore.AddSync(lumberjackLogger)
	} else {
		path := "./logs/prod"
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
        	log.Fatal(err)
    	}
		lumberjackLogger := &lumberjack.Logger{
			Filename: path + "/prod-log.log",
			MaxSize: int(lumberjackConfig.MaxSize),
			MaxAge: int(lumberjackConfig.MaxAge),
			MaxBackups: int(lumberjackConfig.MaxBackups),
			Compress: lumberjackConfig.Compress,
		}
		return zapcore.AddSync(os.Stderr), zapcore.AddSync(lumberjackLogger)
	}

}

func (l *apiLogger) getEncoders() (zapcore.Encoder, zapcore.Encoder){
	if l.cfg.Server.Mode == "Development" {
		return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	} else {
		return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()), zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
}

func (l *apiLogger) InitLogger(){
	loggerLevel := l.getLoggerLevel(l.cfg)
	consoleWriterSyncer, fileWriterSyncer := l.getWriters()
	consoleEncoder, fileEncoder := l.getEncoders()
	core := zapcore.NewTee(
        zapcore.NewCore(fileEncoder, fileWriterSyncer, loggerLevel),
        zapcore.NewCore(consoleEncoder, consoleWriterSyncer, loggerLevel),
    )
	l.zapLogger = zap.New(core)
}

func (l *apiLogger) Debug(msg string, args ...zapcore.Field){
	l.zapLogger.Debug(msg, args...)
}

func (l *apiLogger) DPanic(msg string, args ...zapcore.Field){
	l.zapLogger.DPanic(msg, args...)
}

func (l *apiLogger) Info(msg string, args ...zapcore.Field){
	l.zapLogger.Info(msg, args...)
}

func (l *apiLogger) Warn(msg string, args ...zapcore.Field){
	l.zapLogger.Warn(msg, args...)
}

func (l *apiLogger) Error(msg string, args ...zapcore.Field){
	l.zapLogger.Error(msg, args...)
}

func (l *apiLogger) Panic(msg string, args ...zapcore.Field){
	l.zapLogger.Panic(msg, args...)
}

func (l *apiLogger) Fatal(msg string, args ...zapcore.Field){
	l.zapLogger.Fatal(msg, args...)
}