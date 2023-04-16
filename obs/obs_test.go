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

	// Use the provided mock observer to capture logs
	mockObs := &ObserverMock{
		LogFunc: func(obs.LogLevel, string, ...obs.PayloadEntry) error {
			return nil
		},
	}

	// Create the middleware
	middleware := obs.CommandHandlerObsMiddleware[cqs.Command](mockObs)

	// Create a mock command
	mockCommand := &CommandMock{
		CommandNameFunc: func() string {
			return "mock_command"
		},
	}

	// Create a mock command handler
	handlerExecutionCount := 0
	mockHandler := &CommandHandlerMock[cqs.Command]{
		HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
			handlerExecutionCount++
			return nil, nil
		},
	}

	// Wrap the handler with the middleware
	wrappedHandler := middleware(mockHandler)

	// Execute the wrapped handler
	_, err := wrappedHandler.Handle(context.Background(), mockCommand)
	require.NoError(err)
	require.Equal(1, handlerExecutionCount)

	// Test error handling
	errorMsg := "mock error"
	mockHandlerWithError := &CommandHandlerMock[cqs.Command]{
		HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
			return nil, errors.New(errorMsg)
		},
	}

	// Wrap the handler with the middleware
	wrappedHandlerWithError := middleware(mockHandlerWithError)

	// Execute the wrapped handler with error
	_, err = wrappedHandlerWithError.Handle(context.Background(), mockCommand)
	require.Error(err)
	require.Equal(errorMsg, err.Error())
}
