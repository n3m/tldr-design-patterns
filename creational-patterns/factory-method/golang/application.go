package factory

// Product interface
type Product interface {
	Operation()
}

// Concrete Products
type ConcreteProductA struct{}

func (p *ConcreteProductA) Operation() {
	fmt.Println("ConcreteProductA operation")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Operation() {
	fmt.Println("ConcreteProductB operation")
}

// Creator interface
type Creator interface {
	FactoryMethod() Product
	Operation()
}

// Concrete Creators
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

func (c *ConcreteCreatorA) Operation() {
	product := c.FactoryMethod()
	product.Operation()
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func (c *ConcreteCreatorB) Operation() {
	product := c.FactoryMethod()
	product.Operation()
}

// Usage
creatorA := &ConcreteCreatorA{}
creatorA.Operation() // Output: "ConcreteProductA operation"

creatorB := &ConcreteCreatorB{}
creatorB.Operation() // Output: "ConcreteProductB operation"
