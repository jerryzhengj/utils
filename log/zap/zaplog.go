package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

const callerSkip  = 1
type zapLog struct{
	*zap.SugaredLogger
}

var log zapLog

var logLevel = zap.NewAtomicLevel()

func init(){
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(defaultEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)
	log = zapLog{
		zap.New(core, zap.AddCaller(),zap.AddCallerSkip(callerSkip)).Sugar(),
	}
}

func SetLevel(level string){
	sLevel := strings.ToLower(level)
	l := zap.InfoLevel
	switch sLevel {
	case "debug":
		l = zap.DebugLevel
	case "info":
		l = zap.InfoLevel
	case "warn":
		l = zap.WarnLevel
	case "error":
		l = zap.ErrorLevel
	case "dpanic":
		l = zap.DPanicLevel
	case "panic":
		l = zap.PanicLevel
	case "fatal":
		l = zap.FatalLevel
	default:
		l = zap.InfoLevel
	}

	logLevel.SetLevel(l)
}

func Debug(args ...interface{}) {
	log.Debug(args ...)
}

func Debugf(template string,args ...interface{}) {
	log.Debugf(template,args ...)
}

func Info(args ...interface{}) {
	log.Info(args ...)
}

func Infof(template string,args ...interface{}) {
	log.Infof(template,args ...)
}

func Warn(args ...interface{}) {
	log.Warn(args ...)
}

func Warnf(template string,args ...interface{}) {
	log.Warnf(template,args ...)
}

func Error(args ...interface{}) {
	log.Error(args ...)
}

func Errorf(template string,args ...interface{}) {
	log.Errorf(template,args ...)
}


func DPanic(args ...interface{}) {
	log.DPanic(args ...)
}

func DPanicf(template string,args ...interface{}) {
	log.DPanicf(template,args ...)
}


func Panic(args ...interface{}) {
	log.Panic(args ...)
}

func Panicf(template string,args ...interface{}) {
	log.DPanicf(template,args ...)
}


func Fatal(args ...interface{}) {
	log.Fatal(args ...)
}

func Fatalf(template string,args ...interface{}) {
	log.DPanicf(template,args ...)
}

func defaultTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func defaultEncoderConfig() zapcore.EncoderConfig {
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
		EncodeTime:     defaultTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func InitLogger(logpath string, loglevel string) *zap.SugaredLogger {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留300个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = defaultTimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(defaultEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		level,
	)
	logger := zap.New(core, zap.AddCaller()).Sugar()
	//logger.Sugar()

	//logger := zap.New(core)
	logger.Info("DefaultLogger init success")


	return logger
}


