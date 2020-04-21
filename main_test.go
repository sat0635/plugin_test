package main

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReturnFive(t *testing.T){
	result := ReturnFive("hello")
	assert.Equal(t,result,5)
}

func TestReturnFive2(t *testing.T){
	result := ReturnFive("hello")
        assert.Equal(t,result,5)
}
