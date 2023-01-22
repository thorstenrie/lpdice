package lpdice

type Source interface {
	assert()
	err() error
	Int63() int64
	Seed(s int64)
}
