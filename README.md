`Interface`

1.  Interface is a custom type that specifies a set of one or more method signature.
2.  Interfaces are abstract, it is not possible to instantiate an Interface
3.  Though, it is possible to create a variable whose type is that of Interface, which can be assigned a value of any concrete type which has methods
required by interface.
4. `interface{}` type is interface that species empty set of methods.
5. Every value satisfies `interface{}` type, as it will have its own set of methods which includes empty set of methods. So, `interface{}` type can be used for all values.

`Interface Values`
1. A value of an Inetrface Type, or `interface value` has two components.
	a. dynamic type --> concrete type
	b. dynamic value --> value of that concrete type



`Type Assertions`
1. Type assertion is an operation performed on Intyerfaces.
2. It takes the form of x.(T) where x is expression of interface type and T is a type called "Asserted Type"
3. A type assertion checks whether the dynamic type of its operand matches the asserted type.
