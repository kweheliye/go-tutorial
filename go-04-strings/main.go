package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := "the quick brown fox"
	c := a[:3]
	d := a[:3] + "slow" + a[len(a)-2:]

	fmt.Printf("a type = %T, value = %s, address = %p\n", a, a, unsafe.StringData(a))
	fmt.Printf("c type = %T, value = %s, address = %p\n", c, c, unsafe.StringData(c))
	fmt.Printf("d type = %T, value = %s, address = %p\n", d, d, &d)

	var m map[string]int

	fmt.Printf("m type = %T, value = %v\n", m, m)

	v, ok := m["the"]

	fmt.Printf("ok : %v\n", ok)

	fmt.Printf("value type = %T, value = %v\n", v, v)

}
