package main

import "fmt"

type Strategy interface {
	Execute() string
}

type ConcreteStrategyA struct {}
func (s ConcreteStrategyA) Execute() string { return "Strategy A" }

func (s ConcreteStrategyB) Execute() string { return "Strategy B" }

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c Context) ExecuteStrategy() string {
	return c.strategy.Execute()
}

func main() {
	context := &Context{}
	context.SetStrategy(ConcreteStrategyA{})
	fmt.Println(context.ExecuteStrategy())
}