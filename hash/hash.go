package hash

import (
	hashtable "github.com/timtadh/data-structures/hashtable"
	"fmt"
	types "github.com/timtadh/data-structures/types"
)

func test()  {
	hash := hashtable.NewHashTable(10)
	
	hash.Put(types.String(1),"a")
	v , err :=hash.Get(types.String(1))
	if err != nil{
		fmt.Println(err.Error())
	}else{
		if v != nil{
			fmt.Println(v.(string))
		}
	}
	v = hash.Has(types.String(3))
	fmt.Println(v)
}