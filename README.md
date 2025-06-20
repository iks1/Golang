Learning Golang

Following the Resources 
- https://one2n.io/go-bootcamp
- https://quii.gitbook.io/learn-go-with-tests/

Book PDF Link -  https://github.com/heavykenny/book/blob/master/Go/The.Go.Programming.Language.pdf

### Notes
Go only allows one package inside a single foler (with single or multiple files along with `***_test.go` files)

To initialize a new module
```
go mod init github.com/yourusername/myapp
```
This will create `go.mod` module like
```
module github.com/yourusername/myapp

go 1.21
```

```
myapp/                            ← Your Go project root (Go module root)
├── go.mod                        ← Declares the module name & dependencies
├── go.sum                        ← Dependency checksums for security
├── main.go                       ← Entry point of the application
├── utils/                        ← Public helper functions (own package)
│   └── utils.go                  ← Imports internal/auth (same module)
└── internal/                     ← Not importable from outside the module
    └── auth/
        └── auth.go               ← Internal authentication logic

+-------------------------------+
|      go.mod contents:        |
|-------------------------------|
| module github.com/raj/myapp  |
| go 1.21                       |
| require (...)                 |
+-------------------------------+

     ↓ imports (within same module)

main.go   ─────▶ internal/auth/auth.go ─────▶ "github.com/yourusername/myapp/internal/auth" (importing statement local packages)       
           └──▶ utils/utils.go               
utils.go ─────▶ internal/auth/auth.go        

     ↓ imports (external modules)

utils.go ─────▶ github.com/google/uuid       

     ↓ resolved automatically (through {go run}, no need of {go get} but can if need a specific version)

go.mod updated       ✅
go.sum generated     ✅

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


### Note for ch7
Found a [youtube video](https://www.youtube.com/watch?v=SX1gT5A9H-U) for interface, might help to get started 

### Note for ch8

### Notes on testing

In Go, helper functions used inside test functions (like func TestXxx(t *testing.T)) must be defined before they are used, due to Go's strict top-to-bottom scoping rules. Also, function declarations using func name(...) are only allowed at the package level - inside functions, use anonymous functions assigned to variables (e.g., name := func(...) { ... }).