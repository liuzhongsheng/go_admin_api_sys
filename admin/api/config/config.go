package config
import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	// Cfg .
	Cfg *ini.File
)

// 加载配置文件
func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("loade config err")
	}
}