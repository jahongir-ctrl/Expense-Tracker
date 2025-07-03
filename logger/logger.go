package logger

import (
	"ExpenceTracker/internal/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

var Log = logrus.New()

func InitLogger() {
	logDir := config.AppConfig.LogParams.LogDirectory
	os.MkdirAll(logDir, os.ModePerm)

	infoFile, _ := os.OpenFile(filepath.Join(logDir, config.AppConfig.LogParams.LogInfo), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errorFile, _ := os.OpenFile(filepath.Join(logDir, config.AppConfig.LogParams.LogError), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	warnFile, _ := os.OpenFile(filepath.Join(logDir, config.AppConfig.LogParams.LogWarn), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	debugFile, _ := os.OpenFile(filepath.Join(logDir, config.AppConfig.LogParams.LogDebug), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Основной формат
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetOutput(io.Discard) // Сначала убираем дефолтный Stdout

	Log.SetLevel(logrus.DebugLevel) // Позволяет логировать все уровни

	Log.AddHook(&writerHook{
		Writer: map[logrus.Level]io.Writer{
			logrus.InfoLevel:  infoFile,
			logrus.WarnLevel:  warnFile,
			logrus.ErrorLevel: errorFile,
			logrus.DebugLevel: debugFile,
		},
		LogLevels: logrus.AllLevels,
	})
}

type writerHook struct {
	Writer    map[logrus.Level]io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	writer, ok := hook.Writer[entry.Level]
	if !ok {
		return nil
	}
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(line))
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}
