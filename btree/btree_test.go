package btree

import (
	utils "fabioseixas/dsa/utils"
	"slices"
	"testing"
)

func TestBuild(t *testing.T) {
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
		btree := Build(test.input)
		got := btree.GetList()
		if slices.Compare(got, test.expected) != 0 {
			t.Errorf("BinaryTree(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		input    int
		expected *Node
	}{
		{input: 4, expected: &Node{4, nil, nil}},
		{input: 9, expected: &Node{9, nil, nil}},
	}

	btree := &BinaryTree{
		value: &Node{
			5,
			&Node{
				3,
				&Node{2, nil, nil},
				&Node{4, nil, nil},
			},
			&Node{
				7,
				&Node{6, nil, nil},
				&Node{9, nil, nil},
			},
		},
	}

	for _, test := range tests {

		if got := btree.Find(test.input); *got != *test.expected {
			t.Errorf("Find(%v) = %v; want %v", test.input, got, test.expected)
		}
	}

	tests = []struct {
		input    int
		expected *Node
	}{
		{input: 10, expected: nil},
	}

	for _, test := range tests {

		if got := btree.Find(test.input); got != test.expected {
			t.Errorf("Find(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestFindWithLevel(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 4, expected: 2},
		{input: 9, expected: 2},
		{input: 3, expected: 1},
		{input: 7, expected: 1},
		{input: 5, expected: 0},
	}

	btree := &BinaryTree{
		value: &Node{
			5,
			&Node{
				3,
				&Node{2, nil, nil},
				&Node{4, nil, nil},
			},
			&Node{
				7,
				&Node{6, nil, nil},
				&Node{9, nil, nil},
			},
		},
	}

	for _, test := range tests {
		gotNode, gotLevel := btree.FindWithLevel(test.input)
		if gotLevel != test.expected {
			t.Errorf("Find(%v) = %v; want %v", test.input, gotLevel, test.expected)
		}

		if gotNode == nil {
			t.Errorf("Find(%v) = nil; want != nil", test.input)
		}
	}

	testNotFound := []struct {
		input    int
		expected struct {
			node  *Node
			level int
		}
	}{
		{input: 10, expected: struct {
			node  *Node
			level int
		}{node: nil, level: 0}},
	}

	for _, test := range testNotFound {

		if gotNode, gotLevel := btree.FindWithLevel(test.input); gotLevel != test.expected.level || gotNode != nil {
			t.Errorf("Find(%v) = %v, %v; want %v, %v", test.input, gotNode, gotLevel, test.expected.node, test.expected.level)
		}
	}
}
