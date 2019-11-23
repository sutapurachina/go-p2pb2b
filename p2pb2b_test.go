package p2pb2b

import (
	"encoding/json"
	"reflect"
)

// isEqualJSON checks two json strings for equality
func isEqualJSON(s1 string, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	err := json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(o1, o2), nil
}
