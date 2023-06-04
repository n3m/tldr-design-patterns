// Abstract Product
interface Product {
  operation(): void;
}

// Concrete Products
class ConcreteProductA implements Product {
  operation(): void {
    console.log("ConcreteProductA operation");
  }
}

class ConcreteProductB implements Product {
  operation(): void {
    console.log("ConcreteProductB operation");
  }
}

// Abstract Creator
abstract class Creator {
  abstract factoryMethod(): Product;

  operation(): void {
    const product = this.factoryMethod();
    product.operation();
  }
}

// Concrete Creators
class ConcreteCreatorA extends Creator {
  factoryMethod(): Product {
    return new ConcreteProductA();
  }
}

class ConcreteCreatorB extends Creator {
  factoryMethod(): Product {
    return new ConcreteProductB();
  }
}

// Usage
const creatorA: Creator = new ConcreteCreatorA();
creatorA.operation(); // Output: "ConcreteProductA operation"

const creatorB: Creator = new ConcreteCreatorB();
creatorB.operation(); // Output: "ConcreteProductB operation"
