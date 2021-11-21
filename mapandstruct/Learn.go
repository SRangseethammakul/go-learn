package mapandstruct

import "fmt"

func Learn() {
	//maps (key, value)
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
}
