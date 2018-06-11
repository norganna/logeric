package logeric

import (
	"encoding/json"
	"fmt"
)

type Field struct {
	K string
	V interface{}
}

type FieldList []Field

// Fields is an string map.
type Fields map[string]interface{}

// Add will add a key to a Fields map and keep track of key order.
func (f Fields) Add(key string, value interface{}, order *[]string) bool {
	if key == "error" && value == nil {
		return false
	}

	_, ok := f[key]
	f[key] = value

	if !ok && order != nil {
		*order = append(*order, key)
	}

	return ok
}

// Dump takes the fields and renders them.
func (f Fields) Dump() ([]byte, error) {
	out := map[string]json.RawMessage{}

	for k, v := range f {
		var d []byte
		if s, ok := v.(fmt.Stringer); ok {
			v = s.String()
		} else if e, ok := v.(error); ok {
			v = e.Error()
		}

		if j, jErr := json.Marshal(v); jErr == nil && len(j) > 0 {
			d = j
		} else {
			d, _ = json.Marshal(fmt.Sprintf("%#v", v))
		}
		out[k] = d
	}

	return json.Marshal(out)
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

func combineFieldList(a Fields, b FieldList, order []string) (Fields, []string) {
	o := make([]string, len(order))
	copy(o, order)

	f := make(Fields, len(a))
	for k, v := range a {
		f.Add(k, v, nil)
	}

	for _, x := range b {
		f.Add(x.K, x.V, &o)
	}

	return f, o
}
