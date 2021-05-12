package main

import (
	"github.com/kuludi/kuludi-gin-vue/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start failed: ", err.Error())
		os.Exit(-1)
	}

}
