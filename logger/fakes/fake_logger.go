// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/EngineerBetter/bosh-utils/logger"
)

type FakeLogger struct {
	DebugStub        func(tag, msg string, args ...interface{})
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	DebugWithDetailsStub        func(tag, msg string, args ...interface{})
	debugWithDetailsMutex       sync.RWMutex
	debugWithDetailsArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	InfoStub        func(tag, msg string, args ...interface{})
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	WarnStub        func(tag, msg string, args ...interface{})
	warnMutex       sync.RWMutex
	warnArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	ErrorStub        func(tag, msg string, args ...interface{})
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	ErrorWithDetailsStub        func(tag, msg string, args ...interface{})
	errorWithDetailsMutex       sync.RWMutex
	errorWithDetailsArgsForCall []struct {
		tag  string
		msg  string
		args []interface{}
	}
	HandlePanicStub        func(tag string)
	handlePanicMutex       sync.RWMutex
	handlePanicArgsForCall []struct {
		tag string
	}
	ToggleForcedDebugStub        func()
	toggleForcedDebugMutex       sync.RWMutex
	toggleForcedDebugArgsForCall []struct{}
	FlushStub                    func() error
	flushMutex                   sync.RWMutex
	flushArgsForCall             []struct{}
	flushReturns                 struct {
		result1 error
	}
	FlushTimeoutStub        func(time.Duration) error
	flushTimeoutMutex       sync.RWMutex
	flushTimeoutArgsForCall []struct {
		arg1 time.Duration
	}
	flushTimeoutReturns struct {
		result1 error
	}
}

func (fake *FakeLogger) Debug(tag string, msg string, args ...interface{}) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		fake.DebugStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugArgsForCall(i int) (string, string, []interface{}) {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return fake.debugArgsForCall[i].tag, fake.debugArgsForCall[i].msg, fake.debugArgsForCall[i].args
}

func (fake *FakeLogger) DebugWithDetails(tag string, msg string, args ...interface{}) {
	fake.debugWithDetailsMutex.Lock()
	fake.debugWithDetailsArgsForCall = append(fake.debugWithDetailsArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.debugWithDetailsMutex.Unlock()
	if fake.DebugWithDetailsStub != nil {
		fake.DebugWithDetailsStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) DebugWithDetailsCallCount() int {
	fake.debugWithDetailsMutex.RLock()
	defer fake.debugWithDetailsMutex.RUnlock()
	return len(fake.debugWithDetailsArgsForCall)
}

func (fake *FakeLogger) DebugWithDetailsArgsForCall(i int) (string, string, []interface{}) {
	fake.debugWithDetailsMutex.RLock()
	defer fake.debugWithDetailsMutex.RUnlock()
	return fake.debugWithDetailsArgsForCall[i].tag, fake.debugWithDetailsArgsForCall[i].msg, fake.debugWithDetailsArgsForCall[i].args
}

func (fake *FakeLogger) Info(tag string, msg string, args ...interface{}) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoArgsForCall(i int) (string, string, []interface{}) {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].tag, fake.infoArgsForCall[i].msg, fake.infoArgsForCall[i].args
}

func (fake *FakeLogger) Warn(tag string, msg string, args ...interface{}) {
	fake.warnMutex.Lock()
	fake.warnArgsForCall = append(fake.warnArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.warnMutex.Unlock()
	if fake.WarnStub != nil {
		fake.WarnStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) WarnCallCount() int {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	return len(fake.warnArgsForCall)
}

func (fake *FakeLogger) WarnArgsForCall(i int) (string, string, []interface{}) {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	return fake.warnArgsForCall[i].tag, fake.warnArgsForCall[i].msg, fake.warnArgsForCall[i].args
}

func (fake *FakeLogger) Error(tag string, msg string, args ...interface{}) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorArgsForCall(i int) (string, string, []interface{}) {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return fake.errorArgsForCall[i].tag, fake.errorArgsForCall[i].msg, fake.errorArgsForCall[i].args
}

func (fake *FakeLogger) ErrorWithDetails(tag string, msg string, args ...interface{}) {
	fake.errorWithDetailsMutex.Lock()
	fake.errorWithDetailsArgsForCall = append(fake.errorWithDetailsArgsForCall, struct {
		tag  string
		msg  string
		args []interface{}
	}{tag, msg, args})
	fake.errorWithDetailsMutex.Unlock()
	if fake.ErrorWithDetailsStub != nil {
		fake.ErrorWithDetailsStub(tag, msg, args...)
	}
}

func (fake *FakeLogger) ErrorWithDetailsCallCount() int {
	fake.errorWithDetailsMutex.RLock()
	defer fake.errorWithDetailsMutex.RUnlock()
	return len(fake.errorWithDetailsArgsForCall)
}

func (fake *FakeLogger) ErrorWithDetailsArgsForCall(i int) (string, string, []interface{}) {
	fake.errorWithDetailsMutex.RLock()
	defer fake.errorWithDetailsMutex.RUnlock()
	return fake.errorWithDetailsArgsForCall[i].tag, fake.errorWithDetailsArgsForCall[i].msg, fake.errorWithDetailsArgsForCall[i].args
}

func (fake *FakeLogger) HandlePanic(tag string) {
	fake.handlePanicMutex.Lock()
	fake.handlePanicArgsForCall = append(fake.handlePanicArgsForCall, struct {
		tag string
	}{tag})
	fake.handlePanicMutex.Unlock()
	if fake.HandlePanicStub != nil {
		fake.HandlePanicStub(tag)
	}
}

func (fake *FakeLogger) HandlePanicCallCount() int {
	fake.handlePanicMutex.RLock()
	defer fake.handlePanicMutex.RUnlock()
	return len(fake.handlePanicArgsForCall)
}

func (fake *FakeLogger) HandlePanicArgsForCall(i int) string {
	fake.handlePanicMutex.RLock()
	defer fake.handlePanicMutex.RUnlock()
	return fake.handlePanicArgsForCall[i].tag
}

func (fake *FakeLogger) ToggleForcedDebug() {
	fake.toggleForcedDebugMutex.Lock()
	fake.toggleForcedDebugArgsForCall = append(fake.toggleForcedDebugArgsForCall, struct{}{})
	fake.toggleForcedDebugMutex.Unlock()
	if fake.ToggleForcedDebugStub != nil {
		fake.ToggleForcedDebugStub()
	}
}

func (fake *FakeLogger) ToggleForcedDebugCallCount() int {
	fake.toggleForcedDebugMutex.RLock()
	defer fake.toggleForcedDebugMutex.RUnlock()
	return len(fake.toggleForcedDebugArgsForCall)
}

func (fake *FakeLogger) Flush() error {
	fake.flushMutex.Lock()
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct{}{})
	fake.flushMutex.Unlock()
	if fake.FlushStub != nil {
		return fake.FlushStub()
	} else {
		return fake.flushReturns.result1
	}
}

func (fake *FakeLogger) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *FakeLogger) FlushReturns(result1 error) {
	fake.FlushStub = nil
	fake.flushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLogger) FlushTimeout(arg1 time.Duration) error {
	fake.flushTimeoutMutex.Lock()
	fake.flushTimeoutArgsForCall = append(fake.flushTimeoutArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.flushTimeoutMutex.Unlock()
	if fake.FlushTimeoutStub != nil {
		return fake.FlushTimeoutStub(arg1)
	} else {
		return fake.flushTimeoutReturns.result1
	}
}

func (fake *FakeLogger) FlushTimeoutCallCount() int {
	fake.flushTimeoutMutex.RLock()
	defer fake.flushTimeoutMutex.RUnlock()
	return len(fake.flushTimeoutArgsForCall)
}

func (fake *FakeLogger) FlushTimeoutArgsForCall(i int) time.Duration {
	fake.flushTimeoutMutex.RLock()
	defer fake.flushTimeoutMutex.RUnlock()
	return fake.flushTimeoutArgsForCall[i].arg1
}

func (fake *FakeLogger) FlushTimeoutReturns(result1 error) {
	fake.FlushTimeoutStub = nil
	fake.flushTimeoutReturns = struct {
		result1 error
	}{result1}
}

var _ logger.Logger = new(FakeLogger)
