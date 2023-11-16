package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	*logrus.Logger
}

func NewCustomLogger() *CustomLogger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&customFormatter{})
	return &CustomLogger{Logger: logger}
}

func (cl *CustomLogger) LogInfo(emoji string, format string, args ...interface{}) {
	cl.logWithEmoji(logrus.InfoLevel, emoji, format, args...)
}

func (cl *CustomLogger) Info(format string, args ...interface{}) {
	cl.logWithEmoji(logrus.InfoLevel, "💡", format, args...)
}

func (cl *CustomLogger) Warn(format string, args ...interface{}) {
	cl.logWithEmoji(logrus.WarnLevel, "⚠️", format, args...)
}

func (cl *CustomLogger) Debug(format string, args ...interface{}) {
	cl.logWithEmoji(logrus.DebugLevel, "🐛", format, args...)
}

func (cl *CustomLogger) Fatal(format string, args ...interface{}) {
	cl.logWithEmoji(logrus.FatalLevel, "🔥", format, args...)
}

func (cl *CustomLogger) Error(format string, args ...interface{}) {
	cl.logWithEmoji(logrus.ErrorLevel, "🚨", format, args...)
}

func (cl *CustomLogger) logWithEmoji(level logrus.Level, emoji string, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	cl.Logger.WithFields(logrus.Fields{
		"emoji": emoji,
	}).Log(level, message)
}

type customFormatter struct{}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006/01/02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	return []byte(fmt.Sprintf("[%s] %s: %s %s\n", timestamp, level, entry.Data["emoji"], entry.Message)), nil
}
