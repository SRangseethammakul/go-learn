package pointers

import "fmt"

// pointer ไว้เก็บ memory address ของตัวแปรอื่น
//

func Learn() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x) // get memory address

	y := x
	fmt.Println(y)
	fmt.Println(&y)

	println("-------")
	// var p *int = &x เขียนแบบเต็ม
	p := &x // make poiner p is int ชี้ไปหาตัว x
	fmt.Println(p)
	fmt.Println(*p) // * อ่านค่า x ผ่าน pointer (dereference)

	println("-------")
	*p = 20
	fmt.Println(p)
	fmt.Println(*p)

	a := student{name: "act"}
	fmt.Println(a)
	setName(&a)
	fmt.Println(a)
}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Tarn"
}
