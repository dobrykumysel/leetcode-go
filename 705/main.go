package main

import "fmt"

type MyHashSet struct {
	Arr []int
}

func Constructor() MyHashSet {
	return MyHashSet{make([]int, 0)}
}

func (this *MyHashSet) Add(key int) {
	k := false
	for _, i := range this.Arr {
		if i == key {
			k = true
		}
	}
	if !k {
		this.Arr = append(this.Arr, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	num := -1
	for i, value := range this.Arr {
		if value == key {
			num = i
		}
	}
	//append([]int{1,2}, []int{3,4}...)
	if num >= 0 {
		var k []int
		k = append(this.Arr[:num], this.Arr[num+1:]...)
		this.Arr = k
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, i := range this.Arr {
		if i == key {
			return true
		}
	}
	return false
}

//["MyHashSet","add","remove","add","remove","remove","add","add","add","add","remove"]
//[[],[9],[19],[14],[19],[9],[0],[3],[4],[0],[9]]
func main() {
	obj := Constructor()
	obj.Add(9)
	fmt.Println(obj)
	obj.Remove(19)
	//obj.Add(14)
	//obj.Remove(19)
	//obj.Remove(9)
	//obj.Add(0)
	//obj.Add(3)
	//obj.Add(4)
	//obj.Add(0)
	//obj.Remove(9)
	fmt.Println(obj)
	//contain := obj.Contains(1)
	//fmt.Printf("1: %v\n", contain)
	//contain3 := obj.Contains(3)
	//fmt.Printf("3: %v\n", contain3)
	//obj.Add(2)
	//fmt.Println(obj)
	//obj.Remove(2)
	//fmt.Println(obj)
	//contain4 := obj.Contains(2)
	//fmt.Printf("2: %v\n", contain4)
	//fmt.Println(obj)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
