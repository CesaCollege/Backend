package main

import (
	"context"
	"fmt"
)

func main() {
	// 1
	//ctx := context.Background()
	//cheet.SleepAndTalk(ctx, 3*time.Second, "hello")

	// 2
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//go func() {
	//	s := bufio.NewScanner(os.Stdin)
	//	s.Scan()
	//	cancel()
	//}()
	//cheet.SleepAndTalk(ctx, 5*time.Second, "hello")

	// 3
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//go func() {
	//	time.Sleep(time.Second)
	//	cancel()
	//}()
	//cheet.SleepAndTalk(ctx, 5*time.Second, "hello")

	// 4
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//time.AfterFunc(time.Second, cancel)
	//cheet.SleepAndTalk(ctx, 5*time.Second, "hello")

	// 5
	//ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()
	//cheet.SleepAndTalk(ctx, 5*time.Second, "hello")

	// 6
	ctx := context.Background()
	requestIDKey := 32
	ctx = context.WithValue(ctx, requestIDKey, 500)
	val, ok := ctx.Value(requestIDKey).(int)
	if ok {
		fmt.Println(val)
	}
}

//func sleepAndTalk(ctx context.Context, duration time.Duration, message string) {
//	select {
//	case <-time.After(duration):
//		fmt.Println(message)
//	case <-ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}
