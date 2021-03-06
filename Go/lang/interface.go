package main

import (
	"fmt"
	"strings"
)

type spi interface {
	pf(int) string
}

type sci interface {
	cf(int) string
}

type sp struct {
	p1 string
}

type sc struct {
	c1 string
}

func (sc) cf(a int) string {
	return "cf"
}

func (sp) pf(a int) string {
	return "pf"
}

func (sp) cf(a int) string {
	return "cf"
}

func switchFunc(v interface{}) {
	switch v.(type) {
	case spi:
		fmt.Println("spi")
	case sci:
		fmt.Println("sci")
	case map[string][]string:
		fmt.Println("map[string][]string")
	default:
		fmt.Println("default")
	}
}

func testInterface() {
	// sp := sc{}
	m := make(map[string][]string)
	switchFunc(m)
	// fmt.Println(strconv.Quote("发发了饭\t"))
	fmt.Println(strings.TrimPrefix("func :xxxaagaga", "func"))
}
