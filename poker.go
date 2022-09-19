package main

import (
	"math/rand"
	"time"
)

func shuffle(list []string) []string {
	if len(list) < 2 {
		return list
	}
	re := make([]string, 0, len(list))
	ra := rand.New(rand.NewSource(time.Now().Unix()))
	for _, val := range ra.Perm(len(list)) {
		re = append(re, list[val])
	}
	return re
}
