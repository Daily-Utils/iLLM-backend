package main

import (
	"github.com/daily-utils/iLLM-backend/src"
	"github.com/daily-utils/iLLM-backend/src/utils"
)

func main() {
	utils.LoadEnv()
    ctx, cancel := utils.GetContext()
    defer cancel()

    src.Run(ctx)
}