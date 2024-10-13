package utils

import (
    "context"
    "time"
)

func GetContext() context.Context {
    ctx, err := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		panic(err)
	}

    return ctx
}