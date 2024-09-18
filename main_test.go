package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkPlusOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOperator()
	}
}
func BenchmarkAppendOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendOperator()
	}
}
func BenchmarkSprintfOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintf()
	}
}
func BenchmarkJoinOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join()
	}
}
func BenchmarkByteBufferOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteBuffer()
	}
}
func BenchmarkStringBuilderOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder()
	}
}

func BenchmarkPlusOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < b.N; i++ {
		s1 = s1 + s2
	}
	_ = s1
}
func BenchmarkAppendOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < b.N; i++ {
		s1 += s2
	}
	_ = s1
}
func BenchmarkSprintfOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < b.N; i++ {
		s1 = fmt.Sprintf("%s%s", s1, s2)
	}
	_ = s1
}
func BenchmarkJoinOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < b.N; i++ {
		s1 = strings.Join([]string{s1, s2}, "")
	}
	_ = s1
}
func BenchmarkJoinOperatorThousandStrings(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	arr := make([]string, b.N+1)
	arr[0] = s1
	for i := 0; i < b.N; i++ {
		arr = append(arr, s2)
	}
	_ = strings.Join(arr, "")
}
func BenchmarkByteBufferOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	var str bytes.Buffer
	str.WriteString(s1)
	for i := 0; i < b.N; i++ {
		str.WriteString(s2)
	}

	_ = str.String()
}
func BenchmarkStringBuilderOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	for i := 0; i < b.N; i++ {
		str.WriteString(s2)
	}

	_ = str.String()
}
