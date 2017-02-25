package main

import (
	"os"

	"github.com/luckily248/luck_web_server"
)

func main() {
	os.Setenv("PORT", "3000")
	luckweb.Run()
}
