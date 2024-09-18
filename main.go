package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

}

func plusOperator() string {
	s1 := "hello"
	s2 := "world"

	return s1 + s2
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
func join() string {
	s1 := "hello"
	s2 := "world"

	return strings.Join([]string{s1, s2}, "")
}
func byteBuffer() string {
	s1 := "hello"
	s2 := "world"

	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)

	return buffer.String()
}
func stringBuilder() string {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	str.WriteString(s2)

	return str.String()
}

func copyString() string {
	s1 := "hello"
	s2 := "world"

	str := make([]byte, len(s1)+len(s2))
	copy(str, s1)
	copy(str[len(s1):], s2)
	return string(str)
}

func plusOperatorThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 1000; i++ {
		s1 = s1 + s2
	}
	return s1
}
func appendOperatorThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 1000; i++ {
		s1 += s2
	}
	return s1
}
func sprintfThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 1000; i++ {
		s1 += fmt.Sprintf("%s%s", s1, s2)
	}
	return s1
}
func joinThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	for i := 0; i < 1000; i++ {
		s1 = strings.Join([]string{s1, s2}, "")
	}
	return strings.Join([]string{s1, s2}, "")
}
func joinThousandStrings() string {
	s1 := "hello"
	s2 := "world"

	arr := make([]string, 1001)
	arr[0] = s1
	for i := 0; i < 1000; i++ {
		arr = append(arr, s2)
	}
	return strings.Join(arr, "")
}
func byteBufferThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	var str bytes.Buffer
	str.WriteString(s1)
	for i := 0; i < 1000; i++ {
		str.WriteString(s2)
	}

	return str.String()
}
func stringBuilderThousandTimes() string {
	s1 := "hello"
	s2 := "world"

	var str strings.Builder
	str.WriteString(s1)
	for i := 0; i < 1000; i++ {
		str.WriteString(s2)
	}

	return str.String()
}
