package lpdice

import (
	mrand "math/rand"
	"sync"
)

type dSource struct {
	mu   sync.Mutex
	drnd *mrand.Rand
}

const (
	defaultSeed int64 = 1
)

func newDeterministicSource() *dSource {
	ms := dSource{}
	ms.drnd = mrand.New(mrand.NewSource(defaultSeed))
	return &ms
}

func (m *dSource) Seed(s int64) {
	m.mu.Lock()
	m.drnd.Seed(s)
	m.mu.Unlock()
}

func (m *dSource) Int63() int64 {
	m.mu.Lock()
	v := m.drnd.Int63()
	m.mu.Unlock()
	return v
}

func (m *dSource) err() error {
	return nil
}

func (m *dSource) assert() {}
