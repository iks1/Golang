Learning Golang

Following the Resource - https://one2n.io/go-bootcamp

Book PDF Link -  https://github.com/heavykenny/book/blob/master/Go/The.Go.Programming.Language.pdf

### Notes
Go requires all .go files in the same folder to declare the same package.

```
ch6/ 
├── go.mod
├── main.go              // uses package main
├── geometry/
    └── geometry.go      // uses package geometry
```
### OOPs in Go/ Method Writing

[oops in golang vs other languages](ch6/methods.go)

| OOP Concept    | Go Equivalent                                    |
| -------------- | ------------------------------------------------ |
| Class          | `struct` + method receivers                      |
| Inheritance    | Composition via embedded structs                 |
| Interfaces     | Native interfaces (duck typing)                  |
| Constructors   | Factory functions like `NewType()`               |
| Access control | Exported (Capitalized) vs unexported (lowercase) |

