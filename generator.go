package agen

import (
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	codeLen = 6
)

var (
	epoch = time.Date(2021, 2, 6, 0, 0, 0, 0, time.UTC)
	base = uint64(len(alphabet))
)

type Generator interface {
	UniqueNumber() uint64
	Codes(size int) []string
}

func NewGenerator() Generator {
	return &generator{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// impl

type generator struct {
	rnd *rand.Rand
}

func (g *generator) Codes(size int) []string {
	codes := make([]string, size)

	bs := make([]byte, 6)
	for i := 0; i < codeLen; i++ {
		bs[i] = 'A'
	}

	for i := 0; i < size; i++ {
		n := g.UniqueNumber()

		j := 0
		for n > 0 && j < codeLen {
			bs[j] = alphabet[n % base]

			n /= base
			j += 1
		}

		codes[i] = string(bs)
	}

	return codes
}

func (g *generator) UniqueNumber() uint64 {
	sinceEpoch := uint32(time.Since(epoch).Nanoseconds() % math.MaxUint32)

	timeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(timeBytes, uint32(sinceEpoch))

	randBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(randBytes, g.rnd.Uint32())

	bs := make([]byte, 8)
	for i := 0; i < 4; i++ {
		k := i * 2
		bs[k] = timeBytes[i]
		bs[k+1] = randBytes[i]
	}

	return binary.LittleEndian.Uint64(bs)
}
