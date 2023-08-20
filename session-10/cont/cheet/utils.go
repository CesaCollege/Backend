package cheet

import (
	"context"
	"fmt"
	"time"
)

func SleepAndTalk(ctx context.Context, duration time.Duration, message string) {
	select {
	case <-time.After(duration):
		fmt.Println(message)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
