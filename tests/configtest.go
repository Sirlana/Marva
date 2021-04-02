package main

import (
	"fmt"

	"sirlana.com/sirlana/sso/libs"
)

func mainss() {
	cfg, _ := libs.NewConfig()
	fmt.Println(cfg.Database.DBName)
}
