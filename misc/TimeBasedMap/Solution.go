package main

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	mtime  map[string][]int
	mvalue map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{map[string][]int{}, map[string][]string{}}
}

func (tmap *TimeMap) Set(key string, value string, timestamp int) {
	tmap.mtime[key] = append(tmap.mtime[key], timestamp)
	tmap.mvalue[key] = append(tmap.mvalue[key], value)
}

func (tmap *TimeMap) Get(key string, timestamp int) string {
	i := sort.Search(len(tmap.mtime[key]), func(i int) bool { return tmap.mtime[key][i] > timestamp })
	if i > 0 {
		return tmap.mvalue[key][i-1]
	}
	return ""
}

func main() {
	tmap := Constructor()
	tmap.Set("foo", "bar", 1)
	fmt.Println(tmap.Get("foo", 1))
	fmt.Println(tmap.Get("foo", 3))
	tmap.Set("foo", "bar2", 4)
	fmt.Println(tmap.Get("foo", 4))
	fmt.Println(tmap.Get("foo",5))
}
