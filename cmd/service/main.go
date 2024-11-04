package main

import (
    "log"
    "myservice/internal/app"
    "myservice/internal/config"
    "myservice/internal/logger"

    "golang.org/x/sys/windows/svc"
)

func main() {
    isInteractive, err := svc.IsAnInteractiveSession()
    if err != nil {
        log.Fatalf("Не удалось определить тип сессии: %v", err)
    }

    // Загружаем конфигурацию
    cfg, err := config.Load("config.ini")
    if err != nil {
        log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
    }

    // Инициализируем логгер с использованием пути из конфигурации
    logr, err := logger.New(cfg.LogFilePath)
    if err != nil {
        log.Fatalf("Не удалось инициализировать логирование: %v", err)
    }

    if !isInteractive {
        // Работаем как сервис
        service := &app.Service{
            Config: cfg,
            Logger: logr,
        }

        err = svc.Run("MyGoService", service)
        if err != nil {
            log.Fatalf("Сервис не смог запуститься: %v", err)
        }
        return
    }

    // Работаем в интерактивном режиме (консоль)
    app.Run(cfg, logr)
}
