package app

import (
    "context"
    "myservice/internal/config"
    "myservice/internal/logger"
    "time"

    "golang.org/x/sys/windows/svc"
)

type Service struct {
    Config *config.Config
    Logger *logger.Logger
}

func (m *Service) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
    log := m.Logger
    cfg := m.Config

    // Сообщаем системе, что сервис запускается
    s <- svc.Status{State: svc.StartPending}

    // Инициализация
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Сообщаем системе, что сервис запущен
    s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

    // Запуск основного цикла работы сервиса в отдельной горутине
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                log.Info("Сервис работает...", "some_value", cfg.SomeValue)
                time.Sleep(10 * time.Second)
            }
        }
    }()

    // Обработка запросов от системы управления сервисами
    for {
        select {
        case c := <-r:
            switch c.Cmd {
            case svc.Interrogate:
                s <- c.CurrentStatus
            case svc.Stop, svc.Shutdown:
                s <- svc.Status{State: svc.StopPending}
                cancel() // Останавливаем основной цикл
                return false, 0
            default:
                log.Warn("Получена неизвестная команда", "cmd", c.Cmd)
            }
        }
    }
}

func Run(cfg *config.Config, log *logger.Logger) {
    log.Info("Сервис запущен в интерактивном режиме")
    log.Info("Параметр конфигурации", "some_value", cfg.SomeValue)

    // Пример работы сервиса
    for {
        log.Info("Сервис работает...", "some_value", cfg.SomeValue)
        time.Sleep(10 * time.Second)
    }
}
