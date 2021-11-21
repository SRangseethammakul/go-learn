package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {
	// if
	score := 10
	if score == 10 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Bad")
	}

	// switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Mac os")
	case "windows":
		fmt.Println("Window")
	default:
		fmt.Printf("%s \n", os)
	}

	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	content := string(b)
	fmt.Println(content)

	// For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j := 10; j >= 0; j-- {
		if j == 4 {
			fmt.Println("Boom")
			break
		}
		fmt.Println(j)
		time.Sleep(2 * time.Second)
	}
}
