package blockchain

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"time"
)

const (
	timeConst = 1 * time.Second
)

type Blockchain struct {
	blocks    map[int64]*Block
	lastBlock int64
	lock      sync.Mutex
}

func (bc *Blockchain) AddBlock(idBlock int64, data string) {

	bc.lock.Lock()

	prevBlock := bc.blocks[bc.lastBlock]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks[idBlock] = newBlock
	bc.lastBlock = idBlock

	bc.lock.Unlock()
}

func (bc *Blockchain) GetBlock(idBlock int64) (error, *Block) {

	var ok bool
	var b *Block

	bc.lock.Lock()

	if b, ok = bc.blocks[idBlock]; !ok {
		return errors.New("block not found"), nil
	}

	bc.lock.Unlock()

	return nil, b
}

func (bc *Blockchain) Run(ctx context.Context) error {

	timer := time.NewTimer(timeConst)

	for {

		select {
		case <-timer.C:

			bc.AddBlock(time.Now().Unix(), "sda")
			timer.Reset(timeConst)

		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
}

func NewBlockchain() *Blockchain {

	newBc := &Blockchain{
		blocks: make(map[int64]*Block),
	}

	newBc.blocks[0] = NewGenesisBlock()
	newBc.lastBlock = 0

	return newBc
}
