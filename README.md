# Kit
[![Go Report Card](https://goreportcard.com/badge/github.com/lucianogarciaz/kit)](https://goreportcard.com/report/github.com/lucianogarciaz/kit)

Kit is a set of tools that can be used to enhance your service. 

## Installation
To start using latest version of Kit, just run:
To install kit, follow these steps:

1. Open a terminal window.
2. Navigate to the project directory where you want to install the library
3. Run the following command:
```go
go get github.com/lucianogarciaz/kit
```


## Usage
Here are a few examples of how to use Kit:

### Id
```go
import (
    "github.com/lucianogarciaz/kit
)

func main() {
    id := kit.NewID()
    id.String()
}

```

### Pointers

#### IntPtr
```go
func someOtherFunc(a *int) {
    // does something
}

func main() {
	someOtherFunc(IntPtr(123))
}
```

#### IntValue
```go
func pointerFunc() *int {
    var in = 123
    return &in
}

func someOtherFunc(a int) {
    // does something
}

func main() {
    someOtherFunc(IntValue(pointerFunc()))
}
```

#### Int32Ptr
```go
func someOtherFunc(a *int32) {
	// does something
}

func main() {
	someOtherFunc(Int32Ptr(3123))
}
```

#### Int32Value
```go
func pointerFunc() *int32 {
    var in int32 = 123
    return &in
}

func someOtherFunc(a int32) {
    // does something
}

func main() {
    someOtherFunc(Int32Value(pointerFunc()))
}
```

#### Int64Ptr
```go

func someOtherFunc(a *int64) {
	// does something
}

func main() {
	someOtherFunc(Int64Ptr(3123))
}
```
#### Int64Value
```go
func pointerFunc() *int64 {
	var in int64 = 123
	return &in
}

func someOtherFunc(a int64) {
	// does something
}

func main() {
	someOtherFunc(Int64Value(pointerFunc()))
}
```

#### Float32Ptr
```go

func someOtherFunc(a *float32) {
	// does something
}

func main() {
	someOtherFunc(Float32Ptr(123.2))
}
```

#### Float32Value
```go

func pointerFunc() *float32 {
	var in float32 = 123
	return &in
}

func someOtherFunc(a float32) {
	// does something
}

func main() {
	someOtherFunc(Float32Value(pointerFunc()))
}
```

#### Float64Ptr
```go
func someOtherFunc(a *float64) {
	// does something
}

func main() {
	someOtherFunc(Float64Ptr(123.2))
}
```

#### Float64Value
```go
func pointerFunc() *float64 {
	var in float64 = 123
	return &in
}

func someOtherFunc(a float64) {
	// does something
}

func main() {
	someOtherFunc(Float64Value(pointerFunc()))
}
```

#### BoolValue
```go
func pointerFunc() *bool {
    var in = true
    return &in
}

func someOtherFunc(a bool) {
    // does something
}

func main() {
    someOtherFunc(BoolValue(pointerFunc()))
}
```

#### BoolPtr
```go
func someOtherFunc(a *bool) {
    // does something
}

func main() {
	someOtherFunc(BoolPtr(true))
}
```

#### StringPtr
```go
func someOtherFunc(a *string) {
	// does something
}

func main() {
	someOtherFunc(StringPtr("some-string"))
}
```

### StringValue
```go

func pointerFunc() *string {
	var in = "something"
	return &in
}

func someOtherFunc(a string) {
	// does something
}

func main() {
	someOtherFunc(StringValue(pointerFunc()))
}
```

#### TimePtr
```go
func someOtherFunc(a *time.Time) {
	// does something
}

func main() {
	someOtherFunc(TimePtr(time.Now()))
}
```

#### TimeValue
```go
func pointerFunc() *time.Time {
	var in = time.Now()
	return &in
}

func someOtherFunc(a time.Time) {
	// does something
}

func main() {
	someOtherFunc(TimeValue(pointerFunc()))
}
```

## Contributing
I welcome contributions from the community! To contribute to Kit, follow these steps:

1. Fork the repository.
2. Create an issue where we can debate about what you want to add
3. Create a new branch for your changes: git checkout -b my-feature-branch
4. Please document the usage in the README.md file
5. Make your changes and commit them
6. Push your changes to your fork: git push origin my-feature-branch
7. Submit a pull request.


## Acknowledgements
I would like to thank the following contributors, they are the soul of this library:
* @gonzaloserrano: Gonzalo Serrano 
* @xabi93: Xavi Martinez
* @jmaeso: Joan Maeso
* @mountolive: Leandro Guercio
* @jkmrto: Juan Carlos Martinez
* @maitesin: Oscar Forner Martinez
* @theskyinflames: Jaume Araús
