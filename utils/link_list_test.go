package utils

import (
	"testing"
	"fmt"
)

func TestLinkedList(t *testing.T)  {
	l := NewList()

	l.Prepend(NewNode(1))
	l.Prepend(NewNode(2))
	l.Prepend(NewNode(3))

	zero := *slice(l.Get(0))[0].(*Node).Value.(*Node)
	one := *slice(l.Get(1))[0].(*Node).Value.(*Node)
	two := *slice(l.Get(2))[0].(*Node).Value.(*Node)

	if zero != *NewNode(1) ||
		one != *NewNode(2) ||
		two != *NewNode(3){
		fmt.Println(zero.Value)
		fmt.Println(one.Value)
		fmt.Println(two.Value)
		t.Error()
	}

	//Test Add
	k := NewList()
	k.Append(NewNode(1))

}

func slice(args ...interface{}) []interface{} {
	return args
}