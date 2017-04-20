package util

import "testing"
import "fmt"

func Test_JsonpToJSON(t *testing.T) {
	jsonp := "hello({Id:1,Name:'yidane'})"
	result := JsonpToJSON(jsonp)
	if result != "{Id:1,Name:'yidane'}" {
		fmt.Println(result)
		t.Error("error jsonp func")
	}
}

func Benchmark_JsonpToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonp := "hello({Id:1,Name:'yidane'})"
		JsonpToJSON(jsonp)
	}
}

func Test_IsNumber(t *testing.T) {
	a := "-1"
	if IsInt(a) {
		t.Error("-1 is not number")
	}

	b := "-1.1"
	if IsInt(b) {
		t.Error("-1.1 is not number")
	}

	c := "a"
	if IsInt(c) {
		t.Error("a is number")
	}
}

func Test_MakeHash(t *testing.T) {
	hash := MakeHash("123")
	fmt.Println("123", "[", hash, "]")
}

func Benchmark_IsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsInt("123")
	}
}
