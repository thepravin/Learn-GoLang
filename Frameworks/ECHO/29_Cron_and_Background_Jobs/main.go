package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("* 5 16 * 11 1", func() {
		fmt.Println("Function is running : ", time.Now().Format("15:04:05"))
	})

	c.Start()

	select {}
}
