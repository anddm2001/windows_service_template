package logger

import (
    "log/slog"
    "os"
)

type Logger = slog.Logger

func New(logFilePath string) (*Logger, error) {
    file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    handler := slog.NewTextHandler(file, nil)
    logger := slog.New(handler)
    return logger, nil
}
