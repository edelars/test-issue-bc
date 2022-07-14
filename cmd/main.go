package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"test-issue-bc/internal/client"
	"test-issue-bc/internal/config"
	"test-issue-bc/internal/server"
	"test-issue-bc/internal/storage"
	"test-issue-bc/pkg/blockchain"
)

const (
	srvAddr = "127.0.0.1:12211"
)

func main() {

	var isServer = true
	flag.BoolVar(&isServer, "server", true, "Server or client")

	switch isServer {

	case true:
		startServer()
	case false:
		startClient()
	}

}
func startServer() {

	var err error
	var wg sync.WaitGroup

	logger := config.InitLog("server", "DEBUG")

	ctx := context.Background()

	errs := make(chan error, 4)
	go waitInterruptSignal(errs)

	bc := blockchain.NewBlockchain()
	st := storage.NewStorage()
	srv := server.NewServer(logger, bc, st)

	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- bc.Run(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- srv.Run(ctx, srvAddr)
	}()

	logger.Info().Msg("started")
	err = <-errs
	logger.Err(err).Msg("trying to shutdown gracefully")

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := srv.Shutdown()
		logger.Err(err).Msg("server stopped")
	}()

}

func startClient() {
	cl := client.NewClient()
	err := cl.Run(srvAddr)
	if err != nil {
		os.Exit(1)
	}
}

func waitInterruptSignal(errs chan<- error) {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errs <- fmt.Errorf("%s", <-c)
	signal.Stop(c)
}
