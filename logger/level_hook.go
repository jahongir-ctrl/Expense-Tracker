package logger

import (
	"ExpenceTracker/internal/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path"
)

type LevelHook struct {
	writers map[logrus.Level]io.Writer
}

func NewLevelHook(params config.LogParams) *LevelHook {
	base := params.LogDirectory

	return &LevelHook{
		writers: map[logrus.Level]io.Writer{
			logrus.InfoLevel: &lumberjack.Logger{
				Filename:   path.Join(base, params.LogInfo),
				MaxSize:    params.MaxSizeMegabytes,
				MaxBackups: params.MaxBackups,
				MaxAge:     params.MaxAge,
				Compress:   params.Compress,
				LocalTime:  params.LocalTime,
			},
			logrus.WarnLevel: &lumberjack.Logger{
				Filename:   path.Join(base, params.LogWarn),
				MaxSize:    params.MaxSizeMegabytes,
				MaxBackups: params.MaxBackups,
				MaxAge:     params.MaxAge,
				Compress:   params.Compress,
				LocalTime:  params.LocalTime,
			},
			logrus.ErrorLevel: &lumberjack.Logger{
				Filename:   path.Join(base, params.LogError),
				MaxSize:    params.MaxSizeMegabytes,
				MaxBackups: params.MaxBackups,
				MaxAge:     params.MaxAge,
				Compress:   params.Compress,
				LocalTime:  params.LocalTime,
			},
			logrus.DebugLevel: &lumberjack.Logger{
				Filename:   path.Join(base, params.LogDebug),
				MaxSize:    params.MaxSizeMegabytes,
				MaxBackups: params.MaxBackups,
				MaxAge:     params.MaxAge,
				Compress:   params.Compress,
				LocalTime:  params.LocalTime,
			},
		},
	}
}

func (hook *LevelHook) Fire(entry *logrus.Entry) error {
	writer, ok := hook.writers[entry.Level]
	if !ok {
		return nil
	}
	msg, err := entry.String()
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(msg))
	return err
}

func (hook *LevelHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.DebugLevel,
	}
}
