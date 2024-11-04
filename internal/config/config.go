package config

import (
    "gopkg.in/ini.v1"
)

type Config struct {
    SomeValue string
}

func Load(path string) (*Config, error) {
    cfgFile, err := ini.Load(path)
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        SomeValue: cfgFile.Section("").Key("some_value").String(),
    }
    return cfg, nil
}
