package q146

import (
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	inputCmds := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	inputData := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	want := []string{"null", "null", "null", "1", "null", "-1", "null", "-1", "3", "4"}

	lru := &LRUCache{}
	var got []string
	for i := range inputCmds {
		got = append(got, execute(lru, inputCmds[i], inputData[i]))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("\ngot:  %v,\nwant: %v\n", got, want)
		}
	}
}

func execute(lru *LRUCache, cmd string, data []int) string {
	switch cmd {
	case "LRUCache":
		*lru = Constructor(data[0])
		return "null"
	case "put":
		lru.Put(data[0], data[1])
		return "null"
	case "get":
		return strconv.Itoa(lru.Get(data[0]))
	default:
		return "invalid"
	}
}
