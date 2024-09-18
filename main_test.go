package main

import (
	"fmt"
	"strings"
	"testing"
)

// func BenchmarkPlusOperator(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		plusOperator()
// 	}
// }
// func BenchmarkAppendOperator(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		appendOperator()
// 	}
// }
// func BenchmarkSprintfOperator(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sprintf()
// 	}
// }
// func BenchmarkStringsBuilderOperator(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		stringBuilder()
// 	}
// }

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
func BenchmarkStringsBuilderOperatorThousandTimes(b *testing.B) {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	for i := 0; i < b.N; i++ {
		str.WriteString(s2)
	}

	_ = str.String()
}
