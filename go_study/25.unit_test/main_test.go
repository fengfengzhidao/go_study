package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("测试前")
}

func teardown() {
	fmt.Println("测试后")
}

func TestAdd2(t *testing.T) {
	fmt.Println("测试中...")
	t.Errorf("哈哈哈")
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
