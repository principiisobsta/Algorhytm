package lis_test

import (
	"maciejfirus/lis"
	"reflect"
	"testing"
)

func TestLIS(t *testing.T) {
	test := []int{1}
	expected := []int{1}
	checkResult(t, test, expected)

	test = []int{1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16}
	expected = []int{1, 3, 7, 10, 12, 16}
	checkResult(t, test, expected)

	test = []int{1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16, 20, 7, 19, -10, -9, -8, -11, -5, -4, -3, 1, -4, -6, 23}
	expected = []int{-10, -9, -8, -5, -4, -3, 1, 23}
	checkResult(t, test, expected)

	test = []int{}
	expected = []int{}
	checkResult(t, test, expected)

	test = []int{1, 2, 3}
	expected = []int{1, 2, 3}
	checkResult(t, test, expected)

	test = []int{9, 8, 7}
	expected = []int{7}
	checkResult(t, test, expected)

	test = []int{9, 8}
	expected = []int{8}
	checkResult(t, test, expected)
}

func checkResult(t *testing.T, test []int, expected []int) {
	if !reflect.DeepEqual(lis.Lis(test), expected) {
		t.Fail()
		t.Error("result:", lis.Lis(test), " expected: ", expected)
	}
}
