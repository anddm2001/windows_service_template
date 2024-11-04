package main

// Добавьте в начало файла
import (
    "fmt"
    "os"

    "golang.org/x/sys/windows/svc/eventlog"
    "golang.org/x/sys/windows/svc/mgr"
)

// В функции installService добавьте регистрацию источника событий
func installService(name, desc string) error {
    exepath, err := os.Executable()
    if err != nil {
        return err
    }
    m, err := mgr.Connect()
    if err != nil {
        return err
    }
    defer m.Disconnect()
    s, err := m.OpenService(name)
    if err == nil {
        s.Close()
        return fmt.Errorf("сервис %s уже установлен", name)
    }
    s, err = m.CreateService(name, exepath, mgr.Config{DisplayName: desc})
    if err != nil {
        return err
    }
    defer s.Close()

    // Регистрация источника событий
    err = eventlog.InstallAsEventCreate(name, eventlog.Info|eventlog.Warning|eventlog.Error)
    if err != nil {
        return fmt.Errorf("не удалось зарегистрировать источник событий: %v", err)
    }

    return nil
}

// В функции removeService добавьте удаление источника событий
func removeService(name string) error {
    m, err := mgr.Connect()
    if err != nil {
        return err
    }
    defer m.Disconnect()
    s, err := m.OpenService(name)
    if err != nil {
        return fmt.Errorf("сервис %s не найден", name)
    }
    defer s.Close()
    err = s.Delete()
    if err != nil {
        return err
    }

    // Удаление источника событий
    err = eventlog.Remove(name)
    if err != nil {
        return fmt.Errorf("не удалось удалить источник событий: %v", err)
    }

    return nil
}

