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

    log.Info("Сервис начинает выполнение метода Execute")

    // Сообщаем системе, что сервис запускается
    s <- svc.Status{State: svc.StartPending}
    log.Info("Отправлен статус StartPending")

    // Инициализация контекста для остановки сервиса
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Запуск основного цикла сервиса в отдельной горутине
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Error("Паника в сервисе: %v", r)
            }
        }()

        log.Info("Основной цикл сервиса запущен")
        for {
            select {
            case <-ctx.Done():
                log.Info("Сервис останавливается...")
                return
            default:
                log.Info("Сервис работает... some_value=%s", cfg.SomeValue)
                time.Sleep(10 * time.Second)
            }
        }
    }()

    // Сообщаем системе, что сервис запущен
    s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}
    log.Info("Отправлен статус Running")

    // Обработка запросов от SCM
    for c := range r {
        switch c.Cmd {
        case svc.Interrogate:
            s <- c.CurrentStatus
        case svc.Stop, svc.Shutdown:
            log.Info("Получена команда остановки")
            s <- svc.Status{State: svc.StopPending}
            cancel()
            return false, 0
        default:
            log.Warning("Получена неизвестная команда: %v", c.Cmd)
        }
    }

    return false, 0
}

func Run(cfg *config.Config, log *logger.Logger) {
    log.Info("Сервис запущен в интерактивном режиме")
    log.Info("Параметр конфигурации: some_value=%s", cfg.SomeValue)

    // Пример работы сервиса
    for {
        log.Info("Сервис работает... some_value=%s", cfg.SomeValue)
        time.Sleep(10 * time.Second)
    }
}
