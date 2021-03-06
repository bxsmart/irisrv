package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

func Initialize(config ...zap.Config) *zap.Logger {
	var cfg zap.Config
	if len(config) <= 0 {
		var err error
		cfg = zap.NewDevelopmentConfig()
		if logger, err = cfg.Build(); err != nil {
			panic(err)
			return nil
		}
	} else {
		cfg = config[0]

		sink := lumberjack.Logger{
			Filename:   cfg.OutputPaths[0], // 日志文件路径
			MaxSize:    1024,               // megabytes
			MaxBackups: 7,                  // 最多保留3个备份
			MaxAge:     30,                 //days
			Compress:   true,               // 是否压缩 disabled by default
		}

		errSink := lumberjack.Logger{
			Filename:   cfg.ErrorOutputPaths[0], // 日志文件路径
			MaxSize:    1024,                    // megabytes
			MaxBackups: 1,                       // 最多保留3个备份
			MaxAge:     30,                      //days
			Compress:   true,                    // 是否压缩 disabled by default
		}

		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(cfg.EncoderConfig),
			zapcore.AddSync(&sink),
			cfg.Level,
		)

		logger = zap.New(core, zap.ErrorOutput(zapcore.AddSync(&errSink)))
	}

	sugaredLogger = logger.Sugar()
	logger.Info("DefaultLogger INIT success")

	return logger
}

func GetLogger() *zap.Logger {
	return logger
}

func IsInit() bool {
	return nil != logger
}
