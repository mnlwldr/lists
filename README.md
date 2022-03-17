# whats-this

I just play around with Generics in Go 1.18

## Usage
```go
text := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
head := lists.Head(text)
tail := lists.Tail(text)

fmt.Println(head)
fmt.Println(tail)
```

## Output
```sh
[The]
[quick brown fox jumps over the lazy dog]
```

[![Go Reference](https://pkg.go.dev/badge/github.com/mnlwldr/lists.svg)](https://pkg.go.dev/github.com/mnlwldr/lists)
