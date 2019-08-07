package decorator

import "fmt"

type Component interface {
	ReturnInt() int
	ReturnString() string
}

type ConcreteComponent struct {
	a interface{}
	b interface{}
}

func (c *ConcreteComponent) ReturnInt() int {
	v, ok := c.a.(int)
	if !ok {
		panic(fmt.Sprintf("error happend while type error: %v", c.a))
	}
	return v
}
func (c *ConcreteComponent) ReturnString() string {
	v, ok := c.b.(string)
	if !ok {
		panic(fmt.Sprintf("error happend while type error: %v", c.b))
	}
	return v
}

type Validator struct {
	Component
}

func WrapValidator(c Component) Component {
	return &Validator{
		Component: c,
	}
}

func (w *Validator) ReturnInt() int {
	if w.Component.ReturnInt() >= 40 {
		return w.Component.ReturnInt() + 1
	}
	return w.Component.ReturnInt()
}

func (w *Validator) ReturnString() string {
	return w.Component.ReturnString() + "test"
}
