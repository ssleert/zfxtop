package conf

import (
	"fmt"
	"time"
	"errors"
	"os"
	"github.com/ssleert/ginip"
	"github.com/ssleert/sfolib"
)

var ErrConfFileNotExist = errors.New("conf file not exists")

var (
	// time between updates
	UpdateTime time.Duration
)

func Parse(ConfFile string) error {
	var configFiles [2]string
	userConfig, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configFiles[0] = fmt.Sprintf("%s/ftop/conf.ini", userConfig)
	configFiles[1] = "/etc/ftop/conf.ini"

	var confDir string
	if sfolib.Exists(ConfFile) {
		confDir = ConfFile
	} else {
		for _, e := range configFiles {
			if sfolib.Exists(e) {
				confDir = e
			}
		}
	}
	if confDir == "" {
		return ErrConfFileNotExist
	}

	ini, err := ginip.Load(confDir)
	if err != nil {
		return err
	}

	conf, err := ini.GetValueInt("", "update")
	if err != nil {
		return err
	}
	UpdateTime = (time.Duration(conf) * time.Millisecond)

	return nil
}

