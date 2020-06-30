package main

import (
	"fmt"
)

func main() {

	//Slice
	stars := []string{"*", "**", "***", "****", "*****",
		"******", "*******", "********", "*********", "**********"}

	//Map
	starsmap := map[int]string{1: "*", 2: "**", 3: "***", 4: "****",
		5: "*****", 6: "******", 7: "*******", 8: "********",
		9: "*********", 10: "**********"}

	fmt.Println(stars)
	fmt.Println(starsmap)

	for i := 1; i <= len(starsmap); i++ {
		fmt.Println(starsmap[i])

	}

	for i := 0; i < len(stars); i++ {
		fmt.Println(stars[i])

	}

	//Another method
	// fmt.Println(len(stars))
	// fmt.Println(strings.Join(stars, "\n"))
}
