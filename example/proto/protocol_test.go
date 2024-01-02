package proto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGetCommand(t *testing.T) {
	cmd := &CommandGet{
		Key: []byte("Foo"),
	}
	r := bytes.NewReader(cmd.Bytes())
	pcmd, err := ParseCommand(r)
	assert.Nil(t, err)

	assert.Equal(t, cmd, pcmd)
}

func TestParseSetCommand(t *testing.T) {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}
	r := bytes.NewReader(cmd.Bytes())
	pcmd, err := ParseCommand(r)
	assert.Nil(t, err)

	assert.Equal(t, cmd, pcmd)
}

func BenchmarkParseCommand(b *testing.B) {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}

	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(cmd.Bytes())
		_, _ = ParseCommand(r)
	}
}
