package lpdice

func NewD4() (*Die, error) {
	return newDie(4)
}

func NewD6() (*Die, error) {
	return newDie(6)
}

func NewD8() (*Die, error) {
	return newDie(8)
}

func NewD10() (*Die, error) {
	return newDie(10)
}

func NewD12() (*Die, error) {
	return newDie(12)
}

func NewD20() (*Die, error) {
	return newDie(20)
}

func newDie(n int) (*Die, error) {
	return (&Die{sides: n}).init()
}
