package main

import "fmt"

type MyHashMap struct {
	Map map[int]int
}

//["MyHashMap","put","put","get","get","put","get","remove","get"]
//[[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]
func main() {
	//Your MyHashMap object will be instantiated and called as such:
	obj := Constructor()
	fmt.Println(obj)
	obj.Put(1, 1)
	fmt.Println(obj)
	obj.Put(2, 2)
	fmt.Println(obj)
	param1 := obj.Get(1)
	fmt.Println(param1)
	param2 := obj.Get(3)
	fmt.Println(param2)

	obj.Put(3, 2)
	obj.Put(2, 1)
	fmt.Printf("after put %v:\n", obj)

	param := obj.Get(3)
	fmt.Printf("Get 3: %v\n", param)

	param4 := obj.Get(2)
	fmt.Printf("Get 2: %v\n", param4)

	param41 := obj.Get(1)
	fmt.Printf("Get 1: %v\n", param41)

	obj.Remove(2)
	param3 := obj.Get(2)
	fmt.Println(param3)
}
func Constructor() MyHashMap {
	return MyHashMap{make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.Map[key] = value
}

func (this *MyHashMap) Get(key int) int {
	//traverse through the map
	if val, ok := this.Map[key]; ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.Map, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
