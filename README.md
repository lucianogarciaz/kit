# Kit
[![Go Report Card](https://goreportcard.com/badge/github.com/lucianogarciaz/kit)](https://goreportcard.com/report/github.com/lucianogarciaz/kit)
![workflow](https://github.com/lucianogarciaz/kit/actions/workflows/lint.yml/badge.svg)
![workflow](https://github.com/lucianogarciaz/kit/actions/workflows/test.yml/badge.svg)


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

## Observability

Provides a simple interface for logging. It allows the user to define a custom logger with options for log level, message, and payload.

### Logger
The obs package provides a basic implementation of a logger, but you can also create your own custom logger by implementing the Logger interface.


<details>
<summary>Explain more</summary>


```go
type Logger interface {
	Log(level LogLevel, message string, payload ...PayloadEntry) error
}
```
The Log() method takes a log level, a message string, and an optional list of payload entries.
You can define your own implementation of the Log() method to customize how log messages are processed and formatted.

### Creating a Basic Logger
To create a basic logger with default options, you can use the NewBasicLogger() function:

```go
logger := obs.NewBasicLogger()
```
The default logger writes log messages to os.Stdout using the json format.

#### Logging Messages
To log a message, you can use the Log() method of the logger.
The method takes a log level, a message string, and an optional list of payload entries.
The log level can be one of the predefined constants LevelDebug, LevelInfo, LevelWarn, or LevelError.
For example:
```go
logger.Log(obs.LevelInfo, "Hello, world!")
```

#### Customizing the Logger
You can customize the behavior of the logger by passing one or more options to the NewBasicLogger() function.
The available options are:

* *MarshalerOpt*: sets the Marshaler used to encode log messages. The default is jsonMarshaler.
* *WriterOpt*: sets the writer to which log messages are written. The default is os.Stdout.
For example, to create a logger that writes log messages to a file instead of os.Stdout, you can use the following code:

```go
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
	log.Fatal(err)
}

logger := obs.NewBasicLogger(
	obs.WriterOpt(file),
)
logger.Log(obs.LevelInfo, "some log")
```


#### Advanced Usage
The obs package provides a basic implementation of a logger, but you can also create your own custom logger by implementing the Logger interface.

</details>


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

### Command Handlers

<details>

<summary> explain more:</summary>

```go
var _ Command = &HelloCommand{}

// Define the Command type.
type HelloCommand struct {
	Id   vo.ID
	Name string
}

func (h HelloCommand) CommandName() string {
	return "hello_command"
}

var _ CommandHandler[HelloCommand] = &HelloCommandHandler{}

type SomeRepository interface {
	Save(ctx context.Context, id vo.ID, name string) error
}

// Define the CommandHandler type.
type HelloCommandHandler struct {
	someRepo SomeRepository
}

// Implement the Handle method for the CommandHandler type.
func (h HelloCommandHandler) Handle(ctx context.Context, cmd HelloCommand) ([]Event, error) {
	err := h.someRepo.Save(ctx, cmd.Id, cmd.Name)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// implementation of a logger middlware
func LoggerMiddleware[C Command](log Logger) CommandHandlerMiddleware[C] {
	return func(h CommandHandler[C]) CommandHandler[C] {
		return CommandHandlerFunc[C](func(ctx context.Context, cmd C) ([]Event, error) {
			log.Info("you will see this message before the handle is called")
			events, err := h.Handle(ctx, cmd)
			log.Info("you will see this message after the handle is called")
			if err != nil {
				log.Error(fmt.Errorf("something went wrong, %w", err))
				return events, err
			}

			log.Info(fmt.Sprintf("command: %s was executed correctly", cmd.CommandName()))
			return events, err
		})
	}
}

type Logger interface {
	Info(string)
	Error(error)
}

func chMw[C Command](logger Logger) CommandHandlerMiddleware[C] {
	return CommandHandlerMultiMiddleware(
		// Be careful ⚠️ the order of the mid. are important
		OtherMiddlware[C](logger),
		LoggerMiddleware[C](logger),
	)
}

var _ Logger = &JSONLogger{}

type JSONLogger struct{}

func (J JSONLogger) Info(s string) {}
func (J JSONLogger) Error(err error) {}

func main() {
	handler := HelloCommandHandler{}
	ch := chMw[HelloCommand](JSONLogger{})(handler)

	events, err := ch.Handle(context.Background(), HelloCommand{Id: vo.NewID(), Name: "some-name"})
	if err != nil {
		// do something
		return
	}
	// do something else
	_ = events
}
```

</details>

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
I would like to thank also the people who have inspired me:
* @gonzaloserrano: Gonzalo Serrano 
* @xabi93: Xavi Martinez
* @jmaeso: Joan Maeso
* @mountolive: Leandro Guercio
* @jkmrto: Juan Carlos Martinez
* @maitesin: Oscar Forner Martinez
* @theskyinflames: Jaume Araús
* Typeformers and specially subscriptions team 🧡
