package main

type I1 interface {
	M1()
	M2()
}
type I2 interface {
	M1()
	I3
}
type I3 interface {
	M2()
}
type T struct{}

func (T) M1() {}
func (T) M2() {}

func main() {

	var v1 I1 = T{}
	var v2 I2 = v1
	_ = v2

}
