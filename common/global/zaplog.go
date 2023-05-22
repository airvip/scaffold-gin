package global

import (
	"scaffold-gin/common/config"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZAPLOGGER = InitLogger()

func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	level := getCoreLevel()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",  //结构化（json）输出：时间的key（info，warn，error等）
		LevelKey:      "level", //结构化（json）输出：日志级别的key（info，warn，error等）
		NameKey:       "logger",
		CallerKey:     "linenum", //结构化（json）输出：打印日志的文件对应的Key
		MessageKey:    "msg",     //结构化（json）输出：msg 的 key
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 将日志级别转换成大写（info，warn，error等）
		// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器  或者使用 zapcore.ShortCallerEncoder 短文件路径编码输出（test/main.go:14 ）
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig), // 打印到控制台
		zapcore.NewJSONEncoder(encoderConfig), // 打印到文件
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		writeSyncer,
		level,
	)

	caller := zap.AddCaller()                                               // 开启开发模式，堆栈跟踪
	development := zap.Development()                                        // 开启文件及行号  日志示例  2020-06-12T18:51:11.457+0800	DEBUG	task/main.go:9	this is debug message
	filed := zap.Fields(zap.String("serviceName", config.Conf.Server.Name)) // 设置初始化字段,如：添加一个服务器名称
	return zap.New(core, caller, development, filed)

	// loger, _ := zap.NewProduction(caller, development, filed)
	// return loger

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := lumberjack.Logger{
		Filename:   config.Conf.Logger.Path,       // 日志文件路径，默认 os.TempDir()
		MaxSize:    config.Conf.Logger.MaxSize,    // 每个日志文件保存1M，默认 100M
		MaxBackups: config.Conf.Logger.MaxBackups, // 保留30个备份，默认不限
		MaxAge:     config.Conf.Logger.MaxAge,     // 保留7天，默认不限
		Compress:   config.Conf.Logger.Compress,   // 是否压缩，默认不压缩
	}
	return zapcore.AddSync(&lumberJackLogger)
}

// 设置日志级别
// debug 可以打印出 info debug warn
// info  级别可以打印 warn info
// warn  只能打印 warn
// debug->info->warn->error
func getCoreLevel() zapcore.Level {
	var level zapcore.Level
	switch config.Conf.Logger.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	return level
}
