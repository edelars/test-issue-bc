package client

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"test-issue-bc/pkg/blockchain"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Run(addr string) error {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//1. send 'get'
	if _, err := conn.Write([]byte("get")); err != nil {
		fmt.Println(err)
		return err
	}

	//2. receive unixtime with prev block hash
	var recBuf []byte
	if err, recBuf = readMsg(conn); err != nil {
		fmt.Println(err)
		return err
	}

	//3.1 generate hash
	b := blockchain.NewBlock("sda", recBuf)

	//3.2 send hash
	if _, err := conn.Write(b.Hash); err != nil {
		fmt.Println(err)
		return err
	}

	//4. receive secret data
	var data []byte
	if err, data = readMsg(conn); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("got data! " + string(data))

	conn.Close()

	return nil
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
