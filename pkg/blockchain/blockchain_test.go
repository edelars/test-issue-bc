package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain_Flow(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock(222, "qwe")
	err, b := bc.GetBlock(222)

	assert.Len(t, b.Hash, 32)
	assert.Len(t, b.PrevBlockHash, 32)
	assert.NoError(t, err)

	err, _ = bc.GetBlock(333)

	assert.Error(t, err)
}
