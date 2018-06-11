package logeric

// Fields is an string map.
type Fields map[string]interface{}

// Add will add a key to a Fields map and keep track of key order.
func (f Fields) Add(key string, value interface{}, order *[]string) bool {
	_, ok := f[key]
	f[key] = value

	if !ok && order != nil {
		*order = append(*order, key)
	}

	return ok
}

func combineFields(a, b Fields, order []string) (Fields, []string) {
	o := make([]string, len(order))
	copy(o, order)

	f := make(Fields, len(a))
	for k, v := range a {
		f.Add(k, v, nil)
	}

	for k, v := range b {
		f.Add(k, v, &o)
	}

	return f, o
}
