package main

import (
	"crypto/md5"
	. "fmt"
	"math/rand"
	"time"
)

var table [10]byte

func MD5(s string) string {
	return Sprintf("%x", md5.Sum([]byte(s)))
}

func encrypt(s string) string {
	res := make([]byte, len(s))
	for i := range s {
		res[i] = table[(s[i]+2)%10]
	}
	return Sprintf("%s", res)
}

func randomGenerateTable() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Perm(26)[:10]
	for i := range tmp {
		table[i] = alphabet[tmp[i]]
	}
	return Sprintf("%s", table)
}
