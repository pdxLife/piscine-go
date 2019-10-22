package piscine

func AlphaCount(str string) int {

	counter := 0

	// for index, letter := range str {

	for _, letter := range str {

		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			counter++

			//fmt.Printf("index is %v     letter is &w \n", index, letter)
		}
	}
	//fmt.Printf("Counter value is %v \n", counter)
	//fmt.Println()
	return counter
}
