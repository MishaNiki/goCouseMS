package tasks

import (
	"testing"
)

func TestStrToInt(t *testing.T) {

	tests := []struct {
		str    string
		result int
	}{
		{
			str:    "1",
			result: 1,
		},
		{
			str:    "123233",
			result: 123233,
		},
		{
			str:    "5644fvdfvd",
			result: 5644,
		},
		{
			str:    "csd5644fvdfvd",
			result: 5644,
		},
		{
			str:    "c1sd5644fvdfvd",
			result: 1,
		},
		{
			str:    "csd5644fvdfvd1",
			result: 5644,
		},
		{
			str:    "bbbb",
			result: -1,
		},
		{
			str:    "bb0bb",
			result: 0,
		},
	}

	for i, v := range tests {
		n, err := StrToInt(v.str)
		if err != nil {
			if v.result >= 0 {
				t.Error(err)
			}
		} else if n != v.result {
			t.Errorf("the answers did not agree, test: %v, result: %v", i, n)
		}
	}
}
