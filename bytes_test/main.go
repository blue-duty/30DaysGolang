package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	for i := 0; i < 10; i++ {
		b.WriteByte('!')
	}
	fmt.Println(b.String())

	//l := make([]byte, 10)
	//n, _ := b.Read(l)
	//fmt.Println(n, string(l))
	//_,_ =  b.Read(l)
	//fmt.Println(n, string(l))

	//fmt.Println(b.ReadRune())

	s, _ := b.ReadBytes('!')
	fmt.Println(string(s))
	fmt.Println(b.String())
}
