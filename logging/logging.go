package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
	"time"
)

func NewLogger(fileName, level string, sessionID zap.Field) *zap.Logger {

	outputPath := "./logs/" + fileName + ".log"

	logLevel := zap.DebugLevel

	if strings.EqualFold(level, "INFO") {
		logLevel = zap.InfoLevel
	} else if strings.EqualFold(level, "WARN") {
		logLevel = zap.WarnLevel
	} else if strings.EqualFold(level, "ERROR") {
		logLevel = zap.ErrorLevel
	}

	logWriter := newLogWriter(outputPath)
	encoder := newEncoderConfig()

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), logWriter, logLevel)
	logger := zap.New(core)
	return logger.With(sessionID)

	//config := zap.Config{
	//	Level:       zap.NewAtomicLevelAt(logLevel),
	//	Development: true,
	//	Sampling: &zap.SamplingConfig{
	//		Initial:    100,
	//		Thereafter: 100,
	//	},
	//	Encoding:      "json",
	//	EncoderConfig: zap.NewProductionEncoderConfig(),
	//	//OutputPaths:      []string{"stderr"},
	//	//ErrorOutputPaths: []string{"stderr"},
	//	OutputPaths:      []string{outputPath},
	//	ErrorOutputPaths: []string{outputPath},
	//}
	//
	//config.EncoderConfig.TimeKey = "TimeUTC"
	//
	//config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	//	// 2019-08-13T04:39:11Z
	//}
	//
	//logger, err := config.Build()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//logger = logger.With(sessionID)
	//return logger, err
}

func newEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "TimeUTC"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	}
	return encoderConfig
}

func newLogWriter(filename string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxAge:     10,
		MaxBackups: 3,
		LocalTime:  false,
		Compress:   false,
	})
}
