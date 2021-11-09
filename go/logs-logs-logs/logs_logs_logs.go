package logs

import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {

	reference := "!🔍☀"
	reference2 := '!'

	for _, char := range reference {
		fmt.Println("this is reference", char)
	}
	fmt.Printf("reference2 %d %c %U\n", reference2, reference2, reference2)

	myMapApplication := map[rune]string{
		'!': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	myMapApplication2 := map[string]string{
		"!": "recommendation",
		"🔍": "search",
		"☀": "weather",
	}
	test := "!"
	myRune := []rune(log)
	v, _ := fmt.Printf("%s\n", test)
	fmt.Println(v)

	fmt.Println(myRune)
	fmt.Println(myMapApplication)
	fmt.Println(myMapApplication2)

	for index, char := range log {
		fmt.Printf("%d, %c, %U, %d\n", index, char, char, char)
		fmt.Println("THIS is char", char)

		//convert into rune and then find it and print it
		value, exist := myMapApplication[char]
		fmt.Println(value, exist)

	}

	return ""
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return ""
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return false
}
