package core

import (
	"MCS_Server/global"
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	_, err := os.Stat(global.MCS_Config.Zap.Dir)
	if err != nil || os.IsNotExist(err) {
		_ = os.Mkdir(global.MCS_Config.Zap.Dir, os.ModePerm)
	}
	level = zap.InfoLevel

	logger = zap.New(getEncoderCore())
	if global.MCS_Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
func getEncoderCore() (core zapcore.Core) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.MCS_Config.Zap.Dir, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName("latest_log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
	}
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	return zapcore.NewCore(getEncoder(), writer, level)
}
func getEncoder() zapcore.Encoder {
	if global.MCS_Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}
func getEncoderConfig() (config zapcore.EncoderConfig) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return encoderConfig
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.MCS_Config.Zap.Header + "2006/01/02 - 15:04:05.000"))
}
