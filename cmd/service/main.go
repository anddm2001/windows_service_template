package main

import (
    "log"
    "myservice/internal/app"
    "myservice/internal/config"
    "myservice/internal/logger"

    "golang.org/x/sys/windows/svc"
)

func main() {
    // Проверяем, является ли сессия интерактивной
    isInteractive, err := svc.IsAnInteractiveSession()
    if err != nil {
        log.Fatalf("Не удалось определить тип сессии: %v", err)
    }

    // Имя сервиса
    serviceName := "MyGoService"

    // Загружаем конфигурацию
    cfg, err := config.Load("C:\\ProgramData\\MyService\\config.ini")
    if err != nil {
        log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
    }

    // Инициализируем логгер с использованием имени сервиса
    logr, err := logger.New(serviceName)
    if err != nil {
        log.Fatalf("Не удалось инициализировать логирование: %v", err)
    }
    defer logr.Close()

    if !isInteractive {
        // Работаем как сервис
        service := &app.Service{
            Config: cfg,
            Logger: logr,
        }

        err = svc.Run(serviceName, service)
        if err != nil {
            logr.Error("Сервис не смог запуститься: %v", err)
        }
        return
    }

    // Работаем в интерактивном режиме (консоль)
    app.Run(cfg, logr)
}
