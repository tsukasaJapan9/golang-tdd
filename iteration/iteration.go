package iteration

// Repeat input char
func Repeat(char string, count int) string {
	var str string

	for i := 0; i < count; i++ {
		str += char
	}

	return str
}
