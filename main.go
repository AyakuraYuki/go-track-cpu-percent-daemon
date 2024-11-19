package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	_ "github.com/shirou/gopsutil/v4/cpu"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		for {
			percent, err := cpu.Percent(time.Second, false)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[x] get cpu usage failed: %v\n", err)
				continue
			}
			for _, v := range percent {
				fmt.Printf("%.2f\n", v)
			}
		}
	}()

	gracefulStop(func() {})
}

func gracefulStop(callback func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	s := <-ch
	_, _ = fmt.Fprintf(os.Stderr, "[<] graceful stop at %s by %v\n", time.Now().String(), s)
	callback()
	os.Exit(0)
}
