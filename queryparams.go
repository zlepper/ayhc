package yahc

type QueryParam map[string][]string

func (p QueryParam) Add(key, value string) {
	p[key] = append(p[key], value)
}

func (p QueryParam) Set(key, value string) {
	p[key] = []string{value}
}

func (p QueryParam) Del(key string) {
	delete(p, key)
}
