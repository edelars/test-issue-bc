package main_test

import (
	"context"
	"test-issue-bc/internal/client"
	"test-issue-bc/internal/config"
	"test-issue-bc/internal/server"
	"test-issue-bc/internal/storage"
	"test-issue-bc/pkg/blockchain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	srvAddr = "127.0.0.1:12211"
)

func TestFlow(t *testing.T) {

	ctx := context.Background()

	logger := config.InitLog("server", "DEBUG")

	bc := blockchain.NewBlockchain()
	st := storage.NewStorage()
	srv := server.NewServer(logger, bc, st)

	go bc.Run(ctx)
	go srv.Run(ctx, srvAddr)
	defer srv.Shutdown()

	time.Sleep(time.Second * 1)
	cl := client.NewClient()
	err := cl.Run(srvAddr)

	assert.NoError(t, err)
}
