package pslice

import (
	"encoding/json"
	"testing"
)


func TestForRange(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	for i, e := range slice {
		t.Log(i, e)
	}
}

func TestSliceInterface(t *testing.T) {
	var slice []int
	t.Log(len(slice))
}

// 待分析
func TestSliceJson(t *testing.T) {
	var slice []int
	var sliceZero = make([]int, 0)
	t.Logf("slice:%+v %d sliceZero:%+v %d", slice, cap(slice), sliceZero, cap(sliceZero))
	j, err := json.Marshal(&slice)
	if err != nil {
		t.Error(err)
	}
	t.Log(j)
	t.Log(string(j))
}
