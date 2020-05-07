package flatten

// Flatten accepts `in` as a interface{} as the only parameter,
// to finally return a []interface.
func Flatten(in interface{}) []interface{} {
	flatarr := make([]interface{}, 0)
	slicearr := in.([]interface{})

	for _, val := range slicearr {
		switch val.(type) {
		case int:
			flatarr = append(flatarr, val)
			break
		case []interface{}:
			flatarr = append(flatarr, Flatten(val)...)
			break
		}
	}

	return flatarr
}
