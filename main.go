package main

import (
	"github.com/daily-utils/iLLM-backend/src"
	"github.com/daily-utils/iLLM-backend/src/utils"
)

func main() {
	utils.LoadEnv()
	src.Run()
}