package cli

import (
	"github.com/nobuenhombre/suikat/pkg/clivar"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

// Config - App Command Line Configuration
type Config struct {
	ConfigFile string `cli:"config[app config yaml file]:string=config.api.yaml"`
	LogFile    string `cli:"log[app log file]:string=/var/log/api.my-domain.com/api.log"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := clivar.Load(cfg)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return cfg, nil
}
