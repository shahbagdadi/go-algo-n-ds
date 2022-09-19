package main

import (
	"reflect"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	ip1 := args{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}}
	op1 := [][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}}
	ip2 := args{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}}
	op2 := [][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt"}}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", ip1, op1},
		{"2", ip2, op2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
