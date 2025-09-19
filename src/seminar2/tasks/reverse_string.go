package tasks

// ReverseString переворачивает строку
func ReverseString(s string) string {
	new_s := ""
	for _, char := range s {
		new_s = string(char) + new_s
	}

	return new_s
}
