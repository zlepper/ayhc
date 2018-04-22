package yahc

// Query params to send a long with requests
// These will be added to the query params given in the url string
type QueryParam map[string][]string

// Adds a value to the given parameter
func (p QueryParam) Add(key, value string) {
	p[key] = append(p[key], value)
}

// Overwrites the given parameter
func (p QueryParam) Set(key, value string) {
	p[key] = []string{value}
}

// Removes the given parameter
func (p QueryParam) Del(key string) {
	delete(p, key)
}
