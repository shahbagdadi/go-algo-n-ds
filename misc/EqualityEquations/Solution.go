package main

import "fmt"

func equationsPossible(equations []string) bool {
    parent := map[byte]byte{}

	for _, e := range equations {
		v1 , v2 := e[0] , e[3]
		if e[1] == '=' {
			x := find(v1,&parent)
			y := find(v2,&parent)
			if x != y {
				parent[y] = x
			}
		}
	}

	for _, e := range equations {
		v1 , v2 := e[0] , e[3]
		if e[1] == '!' {
			if find(v1,&parent) == find(v2,&parent) {
				return false
			}
		}
	}
	return true
}


func find(c byte, par *map[byte]byte) byte {
	if value, ok := (*par)[c]; ok {
		return find(value, par)
	}

		return c
}

func main() {
	ip := []string{"a==b","b!=a"}
	ans := equationsPossible(ip)
	fmt.Println(ans)


}