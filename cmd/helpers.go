package cmd

// Contains checks if arguments exist in a []string
func Contains(source []string, input string) bool {
	for _, v := range source {
		if v == input {
			return true
		}
	}
	return false
}
