package variables

import (
	"fmt"
	"strconv"
)

var fullName string //global
var email string = "srangsee@gmail.com"
var c, python = true, false // default bool is false
const vat int = 7           // constant

func Learn() {
	fullName = "Suttipong Rangseethammakul"
	// fmt.Println(fullName)
	fmt.Printf("Hello : %s, Email : %s \n", fullName, email)
	fmt.Println(c, python)

	companyName := "suttipong"
	isShow := true
	age := 15
	fmt.Println(companyName, isShow, age)
	fmt.Printf("%T \n", isShow)

	f := float64(64)
	fmt.Printf("%T \n", f)

	s := strconv.Itoa(vat)
	fmt.Println(" Vat is" + s)
}

/*
 Noted : การประกาศตัวแปรแบบ Global ให้ประกาศแบบเต็มร
 ประกาศตัวแปรใน function สามารถประกาศแบบ shothand ได้ ":="
*/
