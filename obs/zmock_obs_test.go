// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package obs_test

import (
	"github.com/lucianogarciaz/kit/obs"
	"sync"
)

// Ensure, that MarshalerMock does implement obs.Marshaler.
// If this is not the case, regenerate this file with moq.
var _ obs.Marshaler = &MarshalerMock{}

// MarshalerMock is a mock implementation of obs.Marshaler.
//
//	func TestSomethingThatUsesMarshaler(t *testing.T) {
//
//		// make and configure a mocked obs.Marshaler
//		mockedMarshaler := &MarshalerMock{
//			MarshalFunc: func(v interface{}) ([]byte, error) {
//				panic("mock out the Marshal method")
//			},
//		}
//
//		// use mockedMarshaler in code that requires obs.Marshaler
//		// and then make assertions.
//
//	}
type MarshalerMock struct {
	// MarshalFunc mocks the Marshal method.
	MarshalFunc func(v interface{}) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Marshal holds details about calls to the Marshal method.
		Marshal []struct {
			// V is the v argument value.
			V interface{}
		}
	}
	lockMarshal sync.RWMutex
}

// Marshal calls MarshalFunc.
func (mock *MarshalerMock) Marshal(v interface{}) ([]byte, error) {
	if mock.MarshalFunc == nil {
		panic("MarshalerMock.MarshalFunc: method is nil but Marshaler.Marshal was just called")
	}
	callInfo := struct {
		V interface{}
	}{
		V: v,
	}
	mock.lockMarshal.Lock()
	mock.calls.Marshal = append(mock.calls.Marshal, callInfo)
	mock.lockMarshal.Unlock()
	return mock.MarshalFunc(v)
}

// MarshalCalls gets all the calls that were made to Marshal.
// Check the length with:
//
//	len(mockedMarshaler.MarshalCalls())
func (mock *MarshalerMock) MarshalCalls() []struct {
	V interface{}
} {
	var calls []struct {
		V interface{}
	}
	mock.lockMarshal.RLock()
	calls = mock.calls.Marshal
	mock.lockMarshal.RUnlock()
	return calls
}
