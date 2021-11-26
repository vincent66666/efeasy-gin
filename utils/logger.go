package utils

import (
	"efeasy-gin/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"reflect"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func Logger(group string) *zap.Logger {
	ShowLine := global.App.ConfigViper.GetBool("logger." + group + ".show_line")

	// 创建根目录
	createRootDir(group)

	// 设置日志等级
	setLogLevel(group)

	if ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(group), options...)
}

func createRootDir(group string) {
	AppLogger := reflect.ValueOf(global.App.Config.Logger)
	LoggerGroup := AppLogger.FieldByName(group)
	RootDir := LoggerGroup.FieldByName("RootDir").String()
	if ok, _ := PathExists(RootDir); !ok {
		_ = os.Mkdir(RootDir, os.ModePerm)
	}
}

func setLogLevel(group string) {
	AppLogger := reflect.ValueOf(global.App.Config.Logger)
	LoggerGroup := AppLogger.FieldByName(group)
	LogLevel := LoggerGroup.FieldByName("Level").String()
	switch LogLevel {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore(group string) zapcore.Core {
	AppLogger := reflect.ValueOf(global.App.Config.Logger)
	LoggerGroup := AppLogger.FieldByName(group)
	LogFormat := LoggerGroup.FieldByName("Format").String()

	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.App.Env + "." + l.String())
	}

	// 设置编码器
	if LogFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(group), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter(group string) zapcore.WriteSyncer {
	AppLogger := reflect.ValueOf(global.App.Config.Logger)
	LoggerGroup := AppLogger.FieldByName(group)
	RootDir := LoggerGroup.FieldByName("RootDir").String()
	Filename := LoggerGroup.FieldByName("Filename").String()
	MaxSize := LoggerGroup.FieldByName("MaxSize").Int()
	MaxBackups := LoggerGroup.FieldByName("MaxBackups").Int()
	MaxAge := LoggerGroup.FieldByName("MaxAge").Int()
	Compress := LoggerGroup.FieldByName("Compress").Bool()
	file := &lumberjack.Logger{
		Filename:   RootDir + "/" + GetNowFormatTodayTime() + "/" + Filename,
		MaxSize:    int(MaxSize),
		MaxBackups: int(MaxBackups),
		MaxAge:     int(MaxAge),
		Compress:   Compress,
	}

	return zapcore.AddSync(file)
}
