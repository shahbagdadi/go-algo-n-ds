package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	fmap := map[string][]string{}
	for _, p := range paths {
		data := strings.Split(p, " ")
		root := data[0]
		for _, v := range data[1:] {
			f := strings.Split(v, "(")
			fname , fcontent := f[0] , f[1][:len(f[1])-1]
			fmap[fcontent] = append(fmap[fcontent], root + "/" + fname)
		}
	}
	ans := [][]string{}
	for _, v := range fmap {
		if len(v)> 1{
			ans = append(ans,v)
		}
	}
	return ans
}


func main (){
	ip1 := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"}
	ans := findDuplicate(ip1)
	fmt.Println(ans)
}