package cqs_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/cqs"
)

func TestQueryMultiMiddleware(t *testing.T) {
	require := require.New(t)

	var mwExecutionOrder []int

	queryHandlerMw := func(count int) cqs.QueryHandlerMiddleware[cqs.Query, cqs.QueryResult] {
		return func(h cqs.QueryHandler[cqs.Query, cqs.QueryResult]) cqs.QueryHandler[cqs.Query, cqs.QueryResult] {
			return &QueryHandlerMock[cqs.Query, cqs.QueryResult]{
				HandleFunc: func(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
					_, _ = h.Handle(ctx, query)

					mwExecutionOrder = append(mwExecutionOrder, count)

					return nil, nil
				},
			}
		}
	}

	multiMw := cqs.QueryHandlerMultiMiddleware(
		queryHandlerMw(1),
		queryHandlerMw(2),
		queryHandlerMw(3),
	)

	handlerExecutionCount := 0
	qh := &QueryHandlerMock[cqs.Query, cqs.QueryResult]{
		HandleFunc: func(context.Context, cqs.Query) (cqs.QueryResult, error) {
			handlerExecutionCount++
			return nil, nil
		},
	}

	_, err := multiMw(qh).Handle(context.Background(), nil)
	require.NoError(err)
	require.Equal([]int{1, 2, 3}, mwExecutionOrder)
	require.Equal(1, handlerExecutionCount)
}
