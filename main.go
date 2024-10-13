package main

import (
	"github.com/daily-utils/iLLM-backend/src"
	"github.com/daily-utils/iLLM-backend/src/utils"
)

func main() {
	utils.LoadEnv()
	ctx := utils.GetContext()
	src.Run(ctx)
}