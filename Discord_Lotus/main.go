package main

import (
	"fmt"
	"github.com/KaiKaizxc/Golang_repo/Goalng_pro/Discord_Lotus/bot"
	"github.com/KaiKaizxc/Golang_repo/Goalng_pro/Discord_Lotus/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	return
}
