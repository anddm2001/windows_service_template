package config

import (
    "gopkg.in/ini.v1"
)

type Config struct {
    SomeValue   string
    LogFilePath string
}

func Load(path string) (*Config, error) {
    cfgFile, err := ini.Load(path)
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        SomeValue:   cfgFile.Section("").Key("some_value").String(),
        LogFilePath: cfgFile.Section("").Key("log_file_path").String(),
    }
    return cfg, nil
}
