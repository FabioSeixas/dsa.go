package main

import (
	btree "fabioseixas/dsa/btree"
	utils "fabioseixas/dsa/utils"
)

func main() {
	x := []utils.NilInt{
		{Value: 3, Null: false},
		{Value: 2, Null: false},
		{Value: 4, Null: false},
		{Value: 1, Null: false},
		{Value: 0, Null: true},
		{Value: 0, Null: true},
		{Value: 5, Null: false},
	}

	btree.Build(x)

}
