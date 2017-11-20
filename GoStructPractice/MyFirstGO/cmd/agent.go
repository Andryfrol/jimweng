// 1. create a random drawer systemm with function output a random number
// and coresponding to particular person name
// 2. combine wiht OO in go

package main

import (
	"github.com/stretchr/testify/mock"
	"github.com/vektra/cypress"
)

type MockPlugin struct {
	mock.Mock
}

func (m *MockPlugin) Read() ([]*cypress.Message, error) {
	ret := m.Called()

	r0 := ret.Get(0).([]*cypress.Message)
	r1 := ret.Error(1)

	return r0, r1
}

func main() {

}
