package test

import (
	chap1 "github.com/weiyouwozuiku/CampExecise/Chap1"
	"testing"
)

// 可以使用nc 127.0.0.1 1088实验
func Test_Server(t *testing.T) {
	t.Log("demo")
	chap1.Server()
}
