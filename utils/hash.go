package utils

import (
	"errors"
	"math"
)

type HashTable struct {
	Table map[int]*List
	Size int
	Capacity int
}

type item struct {
	key string
	value interface{}
}

func New(cap int) *HashTable {
	table := make(map[int]*List, cap)
	return &HashTable{Table:table, Size:0, Capacity:cap}
}

func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.position(key)
	item, err := ht.find(index, key)

	if item == nil{
		return "", errors.New("Not Found")
	}

	return item.value, err
}

func (ht *HashTable) Put(key string, value interface{}){
	index := ht.position(key)

	if ht.Table[index] == nil{
		ht.Table[index] = NewList()
	}

	item := &item{key:key, value:value}

	a, err := ht.find(index, key)
	if err != nil{
		// The key doesn't exist in HashTable
		ht.Table[index].Append(item)
		ht.Size++
	}else{
		a.value = value
	}
}

func (ht *HashTable) Del(key string) error {
	index := ht.position(key)
	l := ht.Table[index]
	var val *item

	l.Each(func(node Node) {
		if node.Value.(*item).key == key{
			val = node.Value.(*item)
		}
	})

	if val == nil{
		return nil
	}

	ht.Size--
	return l.Remove(val)
}

func (ht *HashTable) ForEach(f func(*item))  {
	for k := range ht.Table{
		if ht.Table[k] != nil{
			ht.Table[k].Each(func(node Node) {
				f(node.Value.(*item))
			})
		}
	}
}

func (ht *HashTable) position(s string) int {
	return hashCode(s) % ht.Capacity
}

func (ht *HashTable) find(i int, key string) (*item, error) {
	l := ht.Table[i]
	var val *item

	l.Each(func(node Node) {
		if node.Value.(*item).key == key{
			val = node.Value.(*item)
		}
	})

	if val == nil{
		return nil, errors.New("Not Found")
	}
	return val, nil
}

func hashCode(s string) int{
	hash := int32(0)
	for i := 0; i < len(s); i++{
		hash = int32(hash << 5 - hash) + int32(s[i])
		hash &= hash
	}

	return int(math.Abs(float64(hash)))
}