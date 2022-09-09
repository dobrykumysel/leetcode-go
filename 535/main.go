package main

import "fmt"

func main() {
	fmt.Println("start")
	url := "https://leetcode.com/problems/design-tinyurl"
	//result := "https://leetcode.com/problems/design-tinyurl"
	fmt.Println(url)
	obj := Constructor()
	fmt.Println(obj)

	// url := obj.encode(longUrl)
	// ans := obj.decode(url)

}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	return ""
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return ""
}
