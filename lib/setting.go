package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"github.com/Sirupsen/logrus"
)

const DEFAULT_CONFIG_PATH = "./.sakuraio/config.yml"

func getDefaultConfigPath() string {
	homedir, _ := homedir.Dir()
	return filepath.Join(homedir, DEFAULT_CONFIG_PATH)
}

func GetSetting() Settings {
	setting, err := GetUserSetting()
	if err != nil {
		logrus.WithError(err).Warn("Cannot read configure file. Use default config.")
	}

	if setting.BaseURL == "" {
		setting.BaseURL = "https://api.sakura.io/"
	}

	if OverrideSettings.APIToken != "" {
		setting.APIToken = OverrideSettings.APIToken
	}
	if OverrideSettings.APISecret != "" {
		setting.APISecret = OverrideSettings.APISecret
	}
	if OverrideSettings.BaseURL != "" {
		setting.BaseURL = OverrideSettings.BaseURL
	}

	return setting
}

func GetUserSetting() (Settings, error) {
	var setting Settings
	buf, err := ioutil.ReadFile(getDefaultConfigPath())
	if err != nil {
		return setting, err
	}

	err = yaml.Unmarshal(buf, &setting)
	return setting, err
}

func WriteSetting(setting Settings) error {
	buf, err := yaml.Marshal(setting)
	if err != nil {
		return err
	}
	_ = os.MkdirAll(filepath.Dir(getDefaultConfigPath()), 0777)
	err = ioutil.WriteFile(getDefaultConfigPath(), buf, 0600)
	return err
}

type Settings struct {
	APIToken  string
	APISecret string
	BaseURL   string
}

var OverrideSettings Settings = Settings{
	APIToken:  "",
	APISecret: "",
	BaseURL:   "",
}
