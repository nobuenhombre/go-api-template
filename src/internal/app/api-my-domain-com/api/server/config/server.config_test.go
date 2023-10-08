package config

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nobuenhombre/suikat/pkg/fico"
)

type testConfig struct {
	fileName    string
	fileContent string
	config      *Config
	err         error
}

func TestConfigLoad(t *testing.T) {
	test := &testConfig{
		fileName:    "server.config_test_load.yaml",
		fileContent: "",
		config: &Config{
			Hosts: HostsConfig{
				API: HTTPServerConfig{
					Host: "127.0.0.1",
					Port: "9001",
					Store: RedisConfig{
						Host:     "127.0.0.1",
						Port:     "6379",
						Password: "",
						DB:       0,
					},
				},
			},
		},
		err: nil,
	}

	config := new(Config)
	err := config.Load(test.fileName)

	if !(reflect.DeepEqual(config, test.config) && errors.Is(err, test.err)) {
		t.Errorf(
			"config.Load(%#v),\n Expected (config = %#v, err = %#v),\n Actual (config = %#v, err = %#v).\n",
			test.fileName, test.config, test.err, config, err,
		)
	}
}

func TestConfigSave(t *testing.T) {
	test := &testConfig{
		fileName:    "server.config_test_save.yaml",
		fileContent: "hosts:\n    api:\n        host: 127.0.0.1\n        post: \"9001\"\n        store:\n            host: 127.0.0.1\n            port: \"6379\"\n",
		config: &Config{
			Hosts: HostsConfig{
				API: HTTPServerConfig{
					Host: "127.0.0.1",
					Port: "9001",
					Store: RedisConfig{
						Host:     "127.0.0.1",
						Port:     "6379",
						Password: "",
						DB:       0,
					},
				},
			},
		},
		err: nil,
	}

	config := test.config
	err := config.Save(test.fileName)

	txtConfigFile := fico.TxtFile(test.fileName)
	fileContent, errReadFile := txtConfigFile.Read()

	if errReadFile != nil {
		t.Errorf(
			"txtConfigFile.Read error %#v",
			errReadFile,
		)
	}

	if !(reflect.DeepEqual(fileContent, test.fileContent) && errors.Is(err, test.err)) {
		t.Errorf(
			"config.Save(%#v),\n Expected (fileContent = %#v, err = %#v),\n Actual (fileContent = %#v, err = %#v).\n",
			test.fileName, test.fileContent, test.err, fileContent, err,
		)
	}
}
