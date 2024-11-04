package main

import (
    "fmt"
    "golang.org/x/sys/windows/svc/mgr"
    "log"
    "os"
)

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
    return nil
}

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
    return nil
}

// func usage(errmsg string) {
//     fmt.Fprintf(os.Stderr,
//         "%s\n\n"+
//             "Использование:\n"+
//             "  %s install   - установить сервис\n"+
//             "  %s remove    - удалить сервис\n"+
//             "", errmsg, filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
//     os.Exit(2)
// }

func init() {
    if len(os.Args) > 1 {
        cmd := os.Args[1]
        serviceName := "MyGoService"
        serviceDesc := "My Go Windows Service"

        switch cmd {
        case "install":
            err := installService(serviceName, serviceDesc)
            if err != nil {
                log.Fatalf("Ошибка установки сервиса: %v", err)
            }
            fmt.Println("Сервис успешно установлен")
            os.Exit(0)
        case "remove":
            err := removeService(serviceName)
            if err != nil {
                log.Fatalf("Ошибка удаления сервиса: %v", err)
            }
            fmt.Println("Сервис успешно удален")
            os.Exit(0)
        }
    }
}
