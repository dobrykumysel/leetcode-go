package main

import "fmt"

type Skiplist struct {
}

func main() {
	obj := Constructor()
	param_1 := obj.Search(1)
	fmt.Println(param_1)
	fmt.Println(obj)
	//obj.Add(num)
	//param_3 := obj.Erase(num)
}

func Constructor() Skiplist {
	return Skiplist{}
}

func (this *Skiplist) Search(target int) bool {
	return false
}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {
	return false
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
