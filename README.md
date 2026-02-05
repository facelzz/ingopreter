# InGOpreter - Go interpreter

Simple Go language interpreter.

> [!NOTE]
> This project just an opportunity to use and learn Go language.

The goal is to run the following code without compilation:
```go
package main

import "fmt"

func main() {
	result := 5 + 3 * 2 / (1 + 1)
	fmt.Println(result)
}

```

```go
package main

import "fmt"

func main() {
	array := []string{"Hello", "World", "!"}

	for _, value := range array {
		fmt.Print(value)
	}
	fmt.Print("\n")
}
```
