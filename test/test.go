package main

import (
	"github.com/pkg/errors"
)

func testPkgErr() error {
	return errors.New("test pkg err") //底层Wrap或者使用pkg/errors.New()
}

func callFunc() error {
	return testPkgErr() // 应用层直接return
}

func main() {
	//
}
