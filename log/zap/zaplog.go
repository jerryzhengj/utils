package zap

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

const CallerSkip int = 1

type zapLog struct{
	*zap.SugaredLogger
}

var log zapLog
var conf *ZaplogConf

var logLevel = zap.NewAtomicLevel()

func init(){
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(defaultEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)
	log = zapLog{
		zap.New(core, zap.AddCaller(),zap.AddCallerSkip(CallerSkip)).Sugar(),
	}
}

func SetLevel(level string){
	logLevel.SetLevel(parseLevel(level))
}

func parseLevel(level string) zapcore.Level{
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

	return l
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

func InitLogger(logconf *ZaplogConf){
	conf = logconf
	if conf.MaxAge < 1{
		conf.MaxAge = 1
	}

	if conf.MaxBackups < 0{
		conf.MaxBackups = 0
	}

	if conf.MaxSize <= 0{
		conf.MaxSize = 1
	}

	setupLog(logconf)

}

func setupLog(logconf *ZaplogConf) *zap.SugaredLogger {
    var w zapcore.WriteSyncer
	if logconf.DailyBackup {
		hook := timeHook(logconf.Filename,logconf.MaxAge,logconf.MaxBackups)
		w = zapcore.AddSync(hook)
	}else{
		hook := sizeHook(logconf.Filename,logconf.MaxSize,logconf.MaxBackups,logconf.MaxAge)
		w = zapcore.AddSync(&hook)
	}

	var level  = parseLevel(logconf.LogLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = defaultTimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(defaultEncoderConfig()),
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),w), //console and file
		zapcore.NewMultiWriteSyncer(w), //just file
		level,
	)

	log = zapLog{
		zap.New(core, zap.AddCaller(),zap.AddCallerSkip(CallerSkip)).Sugar(),
	}
	log.Info("DefaultLogger init success")

	return log.SugaredLogger
}


//按时间分割日志
func timeHook(filename string, maxAge,maxBackups int) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithRotationCount(maxBackups),
		rotatelogs.WithMaxAge(time.Hour * 24 * time.Duration(maxAge)),
		rotatelogs.WithRotationTime(time.Hour * 24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// 按文件大小分割日志
func sizeHook(logpath string, maxSize,maxBackups,maxAge int) lumberjack.Logger{
	hook :=lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    maxSize,     // megabytes
		MaxBackups: maxBackups,      // 最多保留备份的个数
		MaxAge:     maxAge,       // days
		Compress:   false,    // 是否压缩 disabled by default
	}

	return hook

}



