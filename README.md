# ranna-go

`ranna-go` is a Golang client implementation for the code execution sandbox [`ranna`](https://github.com/ranna-go/ranna).

**Please note:** `ranna` already provides a Golang client implementation [here](https://github.com/ranna-go/ranna/tree/master/pkg/client). If this is enough for you, go for it.
With this library I aim at providing a generic client for ranna **plus** the according other microservices ([`snippets`](https://github.com/ranna-go/snippets)).

## Usage

Download the library:

```
go get github.com/lus/ranna-go
```

### Code execution

```go
package main

import "github.com/lus/ranna-go/ranna"

func main() {
    client := ranna.NewClient("https://public.ranna.zekro.de")
    
    // Retrieve all registered language specifications
    specs, err := client.Specs()
    if err != nil {
        panic(err)
    }
    // specs now contains a map of language specifications

    // Execute Go code
    code := `
        package main

        import "fmt"

        func main() {
            fmt.Println("Hello, ranna!")
        }
    `
    request := &ranna.ExecutionRequest{
        Language:    "go",
        Code:        code,
        Arguments:   []string{},
        Environment: map[string]string{},
    }
    result, err := client.Execute(request)
    if err != nil {
        panic(err)
    }
    // result contains stdout, stderr and the execution duration
}
```

### Snippets

```go
package main

import "github.com/lus/ranna-go/snippets"

func main() {
    client := snippets.NewClient("https://snippets.ranna.zekro.de")

    // Create a Go code sippet
    code := `
        package main

        import "fmt"

        func main() {
            fmt.Println("Hello, ranna!")
        }
    `
    snippet := &snippets.Snippet{
        Language: "go",
        Code:     code,
    }
    created, err := client.Create(snippet)
    if err != nil {
        panic(err)
    }
    // created contains the created snippet

    // Retrieve a code snippet
    retrieved, err := client.Snippet("snippet")
    if err != nil {
        panic(err)
    }
    // retrieved contains the retrieved snippet
}
```
