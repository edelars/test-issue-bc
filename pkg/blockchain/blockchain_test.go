package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain_Flow(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock(222, "qwe")
	err, _ := bc.GetBlock(222)

	assert.NoError(t, err)

	err, _ = bc.GetBlock(333)

	assert.Error(t, err)
}
