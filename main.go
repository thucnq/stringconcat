package main

import (
	"fmt"
	"strings"
)

func main() {

}

func plusOperator() string {
	s1 := "hello"
	s2 := "world"

	return (s1 + s2)
}
func appendOperator() string {
	s1 := "hello"
	s2 := "world"

	s1 += s2
	return s1
}
func sprintf() string {
	s1 := "hello"
	s2 := "world"

	return fmt.Sprintf("%s%s", s1, s2)
}
func stringBuilder() string {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	str.WriteString(s2)

	return str.String()
}

func plusOperatorThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 100; i++ {
		s1 = s1 + s2
	}
	return s1
}
func appendOperatorThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 100; i++ {
		s1 += s2
	}
	return s1
}
func sprintfThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 100; i++ {
		s1 += fmt.Sprintf("%s%s", s1, s2)
	}
	return s1
}
func stringBuilderThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	for i := 0; i < 100; i++ {
		str.WriteString(s2)
	}

	return str.String()
}
