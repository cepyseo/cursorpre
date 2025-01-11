package main

import (
	"github.com/cepyseo/cursorpre/auth"
	"github.com/cepyseo/cursorpre/tui"
	"github.com/cepyseo/cursorpre/tui/params"
	"github.com/cepyseo/cursorpre/tui/shortcut"
	"github.com/cepyseo/cursorpre/tui/tool"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productSelected, modelIndexSelected := tui.Run()
	startServer(productSelected, modelIndexSelected)
}

func startServer(productSelected string, modelIndexSelected int) {
	lock, pidFilePath, _ := tool.EnsureSingleInstance("cursor-vip")
	params.Sigs = make(chan os.Signal, 1)
	signal.Notify(params.Sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-params.Sigs
		_ = lock.Unlock()
		_ = os.Remove(pidFilePath)
		auth.UnSetClient(productSelected)
		os.Exit(0)
	}()
	go shortcut.Do()
	auth.Run(productSelected, modelIndexSelected)
}
