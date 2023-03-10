package cqs

import "context"

// QueryResult is a generic query result type.
type QueryResult any

// Query is the interface to identify the DTO for a given query by name.
type Query interface {
	QueryName() string
}

// QueryName is string to identify a given query when it has not input parameters.
type QueryName string

// QueryName implements Query interface.
func (qn QueryName) QueryName() string {
	return string(qn)
}

// QueryHandler is the interface for handling queries.
type QueryHandler[Q Query, R QueryResult] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

type queryHandlerFunc[Q Query, R QueryResult] func(ctx context.Context, query Q) (R, error)

func (f queryHandlerFunc[Q, R]) Handle(ctx context.Context, query Q) (R, error) {
	return f(ctx, query)
}

// QueryHandlerMiddleware is a type for decorating QueryHandlers.
type QueryHandlerMiddleware[Q Query, R QueryResult] func(h QueryHandler[Q, R]) QueryHandler[Q, R]

// QueryHandlerMultiMiddleware applies a sequence of middlewares to a given query handler.
func QueryHandlerMultiMiddleware[Q Query, R QueryResult](middlewares ...QueryHandlerMiddleware[Q, R]) QueryHandlerMiddleware[Q, R] {
	return func(h QueryHandler[Q, R]) QueryHandler[Q, R] {
		handler := h

		for _, m := range middlewares {
			handler = m(handler)
		}

		return queryHandlerFunc[Q, R](handler.Handle)
	}
}
