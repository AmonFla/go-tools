package tools

/*
	Busca un elemento en un array
*/
func InStringArray(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
