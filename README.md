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
## CQS
<details>
<summary> Summary </summary>
Command-Query Separation (CQS) pattern is for handling command and queries in a software system. 
In CQS, commands and queries are separated into two distinct types of operations, each with its own interface and handler. 
While commands change the state of the system, queries retrieve data from the system without modifying it.

This pattern provides several benefits, including better code organization, easier testing, and improved scalability. By separating queries from commands, developers can focus on each type of operation separately and optimize their implementations for their specific use cases.

The cqs package provides a flexible way to handle queries by defining interfaces for queries, query handlers, 
and query result types.

Additionally, it provides a middleware function that allows developers to add additional
functionality to the query/command handling pipeline, such as caching or logging, without modifying the underlying query handler.
### Queries

<details> 
    <summary> explain more:</summary>

```go

var _ Query = &HelloQuery{}

// Define the Query type.
type HelloQuery struct {
	Id string
}

func (h HelloQuery) QueryName() string {
	return "hello_query"
}

var _ QueryHandler[HelloQuery, QueryResult] = &HelloQueryHandler{}

// Define the QueryHandler type.
type HelloQueryHandler struct {
	someRepo SomeRepository
}

// Implement the Handle method for the QueryHandler type.
func (h HelloQueryHandler) Handle(ctx context.Context, query HelloQuery) (QueryResult, error) {
	hello, err := h.someRepo.GetById(ctx, query.Id)
	if err != nil {
		return nil, err
	}

	return hello, nil
}

// implementation of a logger middlware
func LoggerMiddleware[Q Query, R QueryResult](log Logger) QueryHandlerMiddleware[Q, R] {
	return func(h QueryHandler[Q, R]) QueryHandler[Q, R] {
		return queryHandlerFunc[Q, R](func(ctx context.Context, query Q) (R, error) {
                        log.Info("you will see this message before the handle is called")
			result, err := h.Handle(ctx, query)
                        log.Info("you will see this message after the handle is called")
			if err != nil {
				log.Error(fmt.Errorf("something went wrong, %w", err))
				return result, err
			}

			log.Info(fmt.Sprintf("query: %s was executed correctly", query.QueryName()))
			return result, err
		})
	}
}

type Logger interface {
	Info(string)
	Error(error)
}

func qhMw[Q Query, R QueryResult](logger Logger) QueryHandlerMiddleware[Q, R] {
	return QueryHandlerMultiMiddleware(
    // Be careful ⚠️ the order of the mid. are important
		LoggerMiddleware[Q, R](logger),
	)
}

func main() {
	handler := HelloQueryHandler{}
	qh := qhMw[HelloQuery, QueryResult](JSONLogger{})(handler)

	result, err := qh.Handle(context.Background(), HelloQuery{Id: "some-id"})
	if err != nil {
		// do something
		return
	}
	// do something else
	_ = result
}

``` 

</details>

Additionally, it provides a middleware function that allows developers to add additional 
functionality to the query handling pipeline, such as caching or logging, without modifying the underlying query handler.

</details>

## Value objects
### Id
<details>

<summary>usage examples</summary>

```go
import (
    "github.com/lucianogarciaz/kit
)

func main() {
    id := kit.NewID()
    id.String()
}

```
</details>

### Pointers
<details>

<summary>usage examples</summary>


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
</details>

### DateTime
<details>
<summary>usage examples</summary>

```go
func main() {
    dt := vo.DateTimeNow()
    
    dt.Format(time.RFC3339Nano)
    
    dt2 := vo.DateTimeNow()
    
    dt.Equal(dt2) //false
    
    dt.IsZero() //false
    
    err := dt.Scan(time.Now()) // err = false
    
    // implements marshalJSON
    bt, err := dt.MarshalJSON() //err = false

    var emptyDt vo.DateTime
    err = emptyDt.UnmarshalJSON(bt) //err = false
    emptyDt.Equal(dt) // true
}
```
</details>


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
