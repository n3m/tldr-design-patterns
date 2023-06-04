# TLDR; Design Patterns - Creational Patterns :: Factory Method

## Description

The Factory Method pattern is a creational design pattern that provides an interface for creating objects but allows subclasses to decide which class to instantiate. It allows for the creation of objects without specifying the exact class of the object that will be created at runtime.

## Applications

- Use the factory method when you don't know beforehand all the types or implementations of a category in your code.
- Use the Factory Method when you want to extend the usability of a single module, before it is developed.
- Creating objects with complex initialization logic: If the creation of an object requires multiple steps or dependencies, the Factory Method pattern can encapsulate the creation logic in the creator, making it easier to manage.
- Dynamic object creation: The Factory Method pattern allows you to choose different implementations of a product at runtime, based on certain conditions or configurations.
- Hiding the concrete classes from the client: By using the Factory Method pattern, the client code only needs to work with the abstract creator and product interfaces, abstracting away the concrete implementations.

## Examples

[\[Golang\] Transport Example](golang/application.go)

## References

- [Factory Method](https://www.refactoring.guru/design-patterns/factory-method)
