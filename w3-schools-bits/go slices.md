
## Go Slices
- Slices are like arrays, but they are more powerful and flexible.

- Like arrays, slices store multiple values of the same type in one variable, Unlike arrays, the length of a slice can grow or shrink whenever you want.
- In Go, there are a few ways to create a slice:
Using []datatype{values}
Making a slice from an array
Using the make() function
You can create a slice with []datatype{values}
You can create a slice from an array
You can create a slice using the make() function
Access Elements of a Slice
You can get a specific element in a slice by using its index number.
In Go, indexes start at 0. That means [0] is the first element, [1] is the second, and so on.
Append Elements to a Slice
You can add elements to the end of a slice using the append() function.