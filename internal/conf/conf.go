package conf

import (
	"fmt"
	"github.com/ssleert/ginip"
	"github.com/ssleert/sfolib"
	"github.com/ssleert/sterm"
	"os"
	"time"
)

var (
	// time between updates
	Update     = 110 * time.Millisecond
	Icons      = true
	Borders    = sterm.RoundedBorders
	Colors     = true
	Center     = false
	ColorFaint = 238
	ColorMid   = 245
	ColorLoad  = [6]int{27, 63, 99, 135, 171, 207}
	ColorTempr = [6]int{49, 79, 109, 139, 169, 199}
	ColorList  = [3]int{109, 79, 169}

	updateVar  = "update"
	iconsVar   = "icons"
	bordersVar = "borders"
	colorsVar  = "colors"
	centerVar  = "center"
	faintVar   = "faint"
	midVar     = "mid"
	loadVars   = [6]string{
		"load0",
		"load1",
		"load2",
		"load3",
		"load4",
		"load5",
	}
	temprVars = [6]string{
		"tempr0",
		"tempr1",
		"tempr2",
		"tempr3",
		"tempr4",
		"tempr5",
	}
	listVars = [3]string{
		"list0",
		"list1",
		"list2",
	}
)

func checkColor(c int) bool {
	return !(c <= 255 && c >= 0)
}

func parseTui(ini *ginip.Ini) error {
	sect := "tui"

	var conf any
	conf, err := ini.GetValueInt(sect, updateVar)
	if err != nil {
		return err
	}
	if conf.(int) > 110 {
		Update = time.Duration(conf.(int)) * time.Millisecond
	}

	conf, err = ini.GetValueBool(sect, iconsVar)
	if err != nil {
		return err
	}
	Icons = conf.(bool)

	conf, err = ini.GetValueString(sect, bordersVar)
	if err != nil {
		return err
	}
	switch conf.(string) {
	case "rounded":
		Borders = sterm.RoundedBorders
	case "sharp":
		Borders = sterm.SharpBorders
	case "double":
		Borders = sterm.DoubleBorders
	case "ascii":
		Borders = sterm.AsciiBorders
	case "dot":
		Borders = sterm.DotBorders
	default:
		return &BorderIsIncorrect{conf.(string)}
	}

	conf, err = ini.GetValueBool(sect, colorsVar)
	if err != nil {
		return err
	}
	Colors = conf.(bool)

	conf, _ = ini.GetValueBool(sect, centerVar)
	// now without error)
	// backward compatibility

	Center = conf.(bool)

	return nil
}

func parseColors(ini *ginip.Ini) error {
	sect := "colors"

	var conf int
	conf, err := ini.GetValueInt(sect, faintVar)
	if err != nil {
		return err
	}
	ColorFaint = conf

	conf, err = ini.GetValueInt(sect, midVar)
	if err != nil {
		return err
	}
	ColorMid = conf

	for i, e := range loadVars {
		conf, err := ini.GetValueInt(sect, e)
		if err != nil {
			return err
		}
		if checkColor(conf) {
			return &ColorOutOfRangeError{e}
		}
		ColorLoad[i] = conf
	}

	for i, e := range temprVars {
		conf, err := ini.GetValueInt(sect, e)
		if err != nil {
			return err
		}
		if checkColor(conf) {
			return &ColorOutOfRangeError{e}
		}
		ColorTempr[i] = conf
	}

	for i, e := range listVars {
		conf, err := ini.GetValueInt(sect, e)
		if err != nil {
			return err
		}
		if checkColor(conf) {
			return &ColorOutOfRangeError{e}
		}
		ColorList[i] = conf
	}

	return nil
}

func Parse(ConfFile string) error {
	var configFiles [2]string
	userConfig, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configFiles[0] = fmt.Sprintf("%s/zfxtop/conf.ini", userConfig)
	configFiles[1] = "/etc/zfxtop/conf.ini"

	var confDir string
	if ConfFile == "" {
		for _, e := range configFiles {
			if sfolib.Exists(e) {
				confDir = e
			}
		}
		if confDir == "" {
			return ErrNoConfFiles
		}
	} else if !sfolib.Exists(ConfFile) {
		return ErrConfFileNotExist
	} else {
		confDir = ConfFile
	}

	ini, err := ginip.Load(confDir)
	if err != nil {
		return err
	}

	err = parseColors(&ini)
	if err != nil {
		return err
	}
	err = parseTui(&ini)
	if err != nil {
		return err
	}

	return nil
}
