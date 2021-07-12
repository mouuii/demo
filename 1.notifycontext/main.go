package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func everLoop(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			//receive done singal
			fmt.Println("ctx  cancel")
			break LOOP
		default:
			fmt.Println("time.Sleep")
			time.Sleep(time.Second * 10)
		}
	}
}

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		<-sig
// 		// get terminal signal
// 		fmt.Println("get sig , prepare to shutdown")
// 		cancel()
// 	}()
// 	everLoop(ctx)
// 	fmt.Println("graceful shutdown success")
// }

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	everLoop(ctx)
	fmt.Println("graceful shutdown")
}
