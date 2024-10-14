package utils

import (
    "context"
    "time"
)

// GetContext returns a context with a timeout and its cancel function.
func GetContext() (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    return ctx, cancel
}