package btree

import (
	utils "fabioseixas/dsa/utils"
	"slices"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		input    []utils.NilInt
		expected []int
	}{
		{[]utils.NilInt{
			{Value: 3, Null: false},
			{Value: 2, Null: false},
			{Value: 4, Null: false},
			{Value: 1, Null: false},
			{Value: 0, Null: true},
			{Value: 0, Null: true},
			{Value: 5, Null: false},
		}, []int{3, 2, 4, 1, 5}},
	}

	for _, test := range tests {
		btree := BuildBinaryTree(test.input)
		values := btree.GetList()
		if slices.Compare(values, test.expected) != 0 {
			t.Errorf("BinaryTree(%v) = %v; want %v", test.input, values, test.expected)
		}
	}
}
