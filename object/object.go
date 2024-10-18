package object

func GetKeys[K comparable, V any](obj map[K]V) []K {
	ret := make([]K, 0, len(obj))
	for k := range obj {
		ret = append(ret, k)
	}
	return ret
}

func GetValues[K comparable, V any](obj map[K]V) []V {
	ret := make([]V, 0, len(obj))
	for _, v := range obj {
		ret = append(ret, v)
	}
	return ret
}
