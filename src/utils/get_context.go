package utils

import (
    "context"
    "time"
)

func GetContext() (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
    return ctx, cancel
}