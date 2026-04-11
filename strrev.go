package piscine // defines the package name (required by piscine)

func StrRev(s string) string { // function takes a string and returns a string

	runes := []rune(s)   // convert string into rune slice (so we can modify characters safely)
	length := len(runes) // get number of characters in the string

	for i := 0; i < length/2; i++ { // loop from start to middle of the slice

		runes[i], runes[length-1-i] = runes[length-1-i], runes[i] // swap front and back characters
	}

	return string(runes) // convert rune slice back to string and return it
}
