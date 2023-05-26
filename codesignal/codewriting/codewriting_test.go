package codewriting_test

import (
	"testing"

	"bensivo.com/leetcode/codesignal/codewriting"
	"github.com/stretchr/testify/assert"
)

func TestCodeWriting_1(t *testing.T) {
	assert.Equal(t, 3, codewriting.CodeWriting("123"))
}

func TestCodeWriting_2(t *testing.T) {
	assert.Equal(t, 2, codewriting.CodeWriting("12"))
}
func TestCodeWriting_3(t *testing.T) {
	assert.Equal(t, 0, codewriting.CodeWriting("0"))
}
func TestCodeWriting_4(t *testing.T) {
	assert.Equal(t, 1, codewriting.CodeWriting("1"))
}
func TestCodeWriting_5(t *testing.T) {
	assert.Equal(t, 233, codewriting.CodeWriting("2871221111122261"))
}
