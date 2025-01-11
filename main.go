package main

import (
	"github.com/cepyseo/cursorpre/auth"
	"github.com/cepyseo/cursorpre/tui"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productIndexSelected := tui.Run()
	startServer(productIndexSelected)
}

func startServer(productIndexSelected string) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		tui.UnSetProxy()
		os.Exit(0)
	}()
	tui.SetProxy("localhost", auth.Port)
	auth.Run(productIndexSelected)
}
