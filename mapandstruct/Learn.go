package mapandstruct

import "fmt"

func Learn() {
	//maps (key, value)

	var m2 = make(map[string]int)
	m2["a"] = 10
	m2["b"] = 20
	fmt.Println(m2["b"])

	m := map[string]int{"token": 10, "access": 20}
	fmt.Println(m)
	fmt.Println(m["access"])
	m["access"] = 500
	fmt.Println(m["access"])

	// loop
	for key, value := range m {
		fmt.Printf("%v %v \n", key, value)
	}

	// delete map
	delete(m, "access")
	for key, value := range m {
		fmt.Printf("loop %v %v \n", key, value)
	}

	// add map
	m["golang"] = 400
	for key, value := range m {
		fmt.Printf("golang %v %v \n", key, value)
	}

	// Struct
	type User struct {
		id       int
		username string
		password string
	}

	john := User{
		id:       1,
		username: "Johny Dep",
		password: "123123",
	}
	fmt.Println(john.username)
	john.password = "1234"
	fmt.Println(john.password)

	users := []User{
		{id: 2, username: "Act", password: "123"},
		{id: 3, username: "Act2", password: "1234"},
	}
	fmt.Println(users[0].password)
}
