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

/*
 & ถ้าวางไว้หน้าตัวแปร เป็น operator ไว้อ่าน memory address
 * ถ้าวางไว้หน้าตัวแปร เป็นการดึงค่าจาก memory address นั้นๆๆ เรียกว่า Dereference

 *  ถ้าวางไว้หน้าชนิดข้อมูล (Type) เช่น *student, *int => เป็นการเก็บค่า memory address ถ้าว่างเป็น nil
*/
