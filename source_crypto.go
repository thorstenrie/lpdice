package lpdice

import (
	"crypto/rand"
	"encoding/binary"
)

type csource struct{}

func (c csource) Seed(s int64) {}

func (c csource) Int63() int64 {
	var v uint64
	e := binary.Read(rand.Reader, binary.BigEndian, &v)
	if e != nil {
		return 0
	}
	return int64(v & ^uint64(1<<63))
}

func (c csource) available() bool {
	b := make([]byte, 1)
	if _, e := rand.Read(b); e != nil {
		return false
	}
	return true
}
