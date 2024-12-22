package helpers

// This function checks if any element in the one slice exists in the other one
func Contains(s []string, e []string) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}
	return false
}
