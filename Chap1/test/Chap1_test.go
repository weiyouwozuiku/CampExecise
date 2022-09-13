package test

import (
	"testing"

	chap1 "github.com/weiyouwozuiku/CampExecise/Chap1"
)

// 可以使用nc 127.0.0.1 1080实验
func Test_Server(t *testing.T) {
	t.Log("demo")
	chap1.Server()
}
