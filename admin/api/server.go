package main

import (
	"fmt"
	"goAdminApi/config"
	"goAdminApi/routers"
)

func main() {
	port,err := config.Cfg.Section("server").GetKey("HTTP_PORT")
	if err != nil {
		fmt.Println("loade config fail")
	}
	r := routers.InitRouter();
	r.Run(":" + port.String())
}