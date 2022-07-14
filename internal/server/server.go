package server

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"net"
	"test-issue-bc/internal/storage"
	"test-issue-bc/pkg/blockchain"
	"time"
)

type server struct {
	l      net.Listener
	logger zerolog.Logger
	bc     *blockchain.Blockchain
	st     *storage.Storage
}

func NewServer(logger zerolog.Logger, bc *blockchain.Blockchain, st *storage.Storage) *server {
	return &server{logger: logger, bc: bc, st: st}
}

func (r *server) Run(ctx context.Context, addr string) (err error) {

	r.l, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := r.l.Accept()

		if ctx.Err() != err {
			return ctx.Err()
		}
		if err != nil {
			r.logger.Err(err).Msg("error accepting")
		}
		go r.handleRequest(conn)
	}
}

func (r *server) Shutdown() error {
	return r.l.Close()
}

func (r *server) handleRequest(conn net.Conn) {

	var msg []byte
	var err error

	defer conn.Close()

	if err, msg = readMsg(conn); err != nil {
		r.logger.Err(err).Msg("error read")
		return
	}

	r.logger.Debug().Msgf("Receive msg: %s", msg)

	//1. should receive 'get'
	if string(msg) != "get" {
		r.logger.Debug().Msg("Receive msg but not get")
		return
	}

	//2. send current unix time and prev hash
	curUnixTime := time.Now().Unix()
	var b *blockchain.Block
	err, b = r.bc.GetBlock(curUnixTime)
	if err != nil {
		r.logger.Debug().Msgf("block %d not found, abort", curUnixTime)
	}

	if _, err := conn.Write(b.PrevBlockHash); err != nil {
		r.logger.Err(err).Msg("error write")
		return
	}

	//3. should receive a hash and verify it
	if err, msg = readMsg(conn); err != nil {
		r.logger.Err(err).Msg("error read")
		return
	}

	if !equal(msg, b.Hash) {
		r.logger.Err(err).Msg("error check pow")
		return
	}

	//4. send our secret data
	data := r.st.GetRandom()
	if _, err := conn.Write([]byte(data)); err != nil {
		r.logger.Err(err).Msg("error write")
		return
	}

	r.logger.Debug().Msgf("Secret data send: %s", data)
}

func readMsg(conn net.Conn) (error, []byte) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)

	if err != nil {
		return errors.Wrap(err, "Error reading"), nil
	}

	if reqLen == 0 {
		return errors.Wrap(err, "Error zero message"), nil
	}

	res := make([]byte, reqLen)
	copy(res, buf)

	return nil, res
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
