package submodule

type BinaryArithematic struct {
	First  int
	Second int
}

func (b *BinaryArithematic) Add() int {
	return b.First + b.Second
}
