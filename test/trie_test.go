package test

import (
	"testing"

	"../core"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	node := core.GetNode()
	core.Insert(node, "there")
	core.Insert(node, "their")
	assert.Equal(t, true, core.Search(node, "their"))
	assert.Equal(t, false, (core.Search(node, "the")))
}
