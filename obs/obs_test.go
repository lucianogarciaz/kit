package obs_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/cqs"
	"github.com/lucianogarciaz/kit/obs"
)

func TestCommandHandlerObsMiddleware(t *testing.T) {
	require := require.New(t)

	mockCommand := &CommandMock{
		CommandNameFunc: func() string {
			return "mock_command"
		},
	}

	t.Run(`Given a successful command handler`, func(t *testing.T) {
		mockObs := &ObserverMock{
			LogFunc: func(obs.LogLevel, string, ...obs.PayloadEntry) error {
				return nil
			},
		}

		middleware := obs.CommandHandlerObsMiddleware[cqs.Command](mockObs)

		commandHandlerMock := &CommandHandlerMock[cqs.Command]{
			HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
				return nil, nil
			},
		}

		wrappedHandler := middleware(commandHandlerMock)

		_, err := wrappedHandler.Handle(context.Background(), mockCommand)
		require.NoError(err)
		require.Len(commandHandlerMock.HandleCalls(), 1)
		require.Len(mockObs.LogCalls(), 1)
	})

	t.Run(`Given an error on command handler`, func(t *testing.T) {
		mockObs := &ObserverMock{
			LogFunc: func(obs.LogLevel, string, ...obs.PayloadEntry) error {
				return nil
			},
		}

		middleware := obs.CommandHandlerObsMiddleware[cqs.Command](mockObs)

		expectedErr := errors.New("mock error")
		commandHandlerMock := &CommandHandlerMock[cqs.Command]{
			HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
				return nil, expectedErr
			},
		}

		wrappedHandlerWithError := middleware(commandHandlerMock)

		_, err := wrappedHandlerWithError.Handle(context.Background(), mockCommand)
		require.ErrorIs(err, expectedErr)
		require.Len(mockObs.LogCalls(), 2)
	})
}

func TestQueryHandlerObsMiddleware(t *testing.T) {
	require := require.New(t)

	mockObs := &ObserverMock{
		LogFunc: func(obs.LogLevel, string, ...obs.PayloadEntry) error {
			return nil
		},
	}

	middleware := obs.QueryHandlerObsMiddleware[cqs.Query, cqs.QueryResult](mockObs)

	mockQuery := &QueryMock{
		QueryNameFunc: func() string {
			return "mock_query"
		},
	}

	t.Run(`Given a successful query handler`, func(t *testing.T) {
		queryHandlerMock := &QueryHandlerMock[cqs.Query, cqs.QueryResult]{
			HandleFunc: func(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
				return nil, nil
			},
		}

		wrappedHandler := middleware(queryHandlerMock)

		_, err := wrappedHandler.Handle(context.Background(), mockQuery)
		require.NoError(err)
		require.Len(queryHandlerMock.HandleCalls(), 1)
	})
	t.Run(`Given an error on query handler`, func(t *testing.T) {
		expectedErr := errors.New("mock error")
		queryHandlerMock := &QueryHandlerMock[cqs.Query, cqs.QueryResult]{
			HandleFunc: func(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
				return nil, expectedErr
			},
		}

		wrappedHandlerWithError := middleware(queryHandlerMock)

		_, err := wrappedHandlerWithError.Handle(context.Background(), mockQuery)
		require.ErrorIs(err, expectedErr)
	})
}
