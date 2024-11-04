package logger

import (
    "fmt"
    "golang.org/x/sys/windows/svc/eventlog"
)

type Logger struct {
    eventLog *eventlog.Log
}

func New(sourceName string) (*Logger, error) {
    // Открываем системный журнал
    el, err := eventlog.Open(sourceName)
    if err != nil {
        return nil, fmt.Errorf("не удалось открыть системный журнал: %v", err)
    }

    return &Logger{
        eventLog: el,
    }, nil
}

func (l *Logger) Close() error {
    return l.eventLog.Close()
}

// Методы для логирования разных уровней
func (l *Logger) Info(msg string, args ...interface{}) {
    message := fmt.Sprintf(msg, args...)
    l.eventLog.Info(1, message)
}

func (l *Logger) Warning(msg string, args ...interface{}) {
    message := fmt.Sprintf(msg, args...)
    l.eventLog.Warning(1, message)
}

func (l *Logger) Error(msg string, args ...interface{}) {
    message := fmt.Sprintf(msg, args...)
    l.eventLog.Error(1, message)
}
