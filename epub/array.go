package tools

/*
	Busca un string en un array
*/
func StringInArray(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
