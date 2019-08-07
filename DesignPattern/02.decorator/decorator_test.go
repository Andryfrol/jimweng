package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	// var c Component = &ConcreteComponent{}

	var c Component = &ConcreteComponent{
		a: "test",
		b: "string",
	}

	// c = WarpAddDecorator(c, 10)
	c = WrapValidator(c)
	res := c.ReturnInt()
	res2 := c.ReturnString()
	// fmt.Printf("res %d__%s\n", res, res2)

	assert.Equal(t, 61, res)
	assert.Equal(t, "stringtest", res2)

	// Output:
	// res 80
}
