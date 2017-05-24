// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"os"
	"sync"

	"github.com/concourse/atc/resource"
)

type FakeFetchSource struct {
	LockNameStub        func() (string, error)
	lockNameMutex       sync.RWMutex
	lockNameArgsForCall []struct{}
	lockNameReturns     struct {
		result1 string
		result2 error
	}
	lockNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	FindInitializedStub        func() (resource.VersionedSource, bool, error)
	findInitializedMutex       sync.RWMutex
	findInitializedArgsForCall []struct{}
	findInitializedReturns     struct {
		result1 resource.VersionedSource
		result2 bool
		result3 error
	}
	findInitializedReturnsOnCall map[int]struct {
		result1 resource.VersionedSource
		result2 bool
		result3 error
	}
	InitializeStub        func(signals <-chan os.Signal, ready chan<- struct{}) (resource.VersionedSource, error)
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	initializeReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 resource.VersionedSource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSource) LockName() (string, error) {
	fake.lockNameMutex.Lock()
	ret, specificReturn := fake.lockNameReturnsOnCall[len(fake.lockNameArgsForCall)]
	fake.lockNameArgsForCall = append(fake.lockNameArgsForCall, struct{}{})
	fake.recordInvocation("LockName", []interface{}{})
	fake.lockNameMutex.Unlock()
	if fake.LockNameStub != nil {
		return fake.LockNameStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lockNameReturns.result1, fake.lockNameReturns.result2
}

func (fake *FakeFetchSource) LockNameCallCount() int {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return len(fake.lockNameArgsForCall)
}

func (fake *FakeFetchSource) LockNameReturns(result1 string, result2 error) {
	fake.LockNameStub = nil
	fake.lockNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) LockNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.LockNameStub = nil
	if fake.lockNameReturnsOnCall == nil {
		fake.lockNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lockNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) FindInitialized() (resource.VersionedSource, bool, error) {
	fake.findInitializedMutex.Lock()
	ret, specificReturn := fake.findInitializedReturnsOnCall[len(fake.findInitializedArgsForCall)]
	fake.findInitializedArgsForCall = append(fake.findInitializedArgsForCall, struct{}{})
	fake.recordInvocation("FindInitialized", []interface{}{})
	fake.findInitializedMutex.Unlock()
	if fake.FindInitializedStub != nil {
		return fake.FindInitializedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findInitializedReturns.result1, fake.findInitializedReturns.result2, fake.findInitializedReturns.result3
}

func (fake *FakeFetchSource) FindInitializedCallCount() int {
	fake.findInitializedMutex.RLock()
	defer fake.findInitializedMutex.RUnlock()
	return len(fake.findInitializedArgsForCall)
}

func (fake *FakeFetchSource) FindInitializedReturns(result1 resource.VersionedSource, result2 bool, result3 error) {
	fake.FindInitializedStub = nil
	fake.findInitializedReturns = struct {
		result1 resource.VersionedSource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetchSource) FindInitializedReturnsOnCall(i int, result1 resource.VersionedSource, result2 bool, result3 error) {
	fake.FindInitializedStub = nil
	if fake.findInitializedReturnsOnCall == nil {
		fake.findInitializedReturnsOnCall = make(map[int]struct {
			result1 resource.VersionedSource
			result2 bool
			result3 error
		})
	}
	fake.findInitializedReturnsOnCall[i] = struct {
		result1 resource.VersionedSource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetchSource) Initialize(signals <-chan os.Signal, ready chan<- struct{}) (resource.VersionedSource, error) {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("Initialize", []interface{}{signals, ready})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub(signals, ready)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.initializeReturns.result1, fake.initializeReturns.result2
}

func (fake *FakeFetchSource) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeFetchSource) InitializeArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return fake.initializeArgsForCall[i].signals, fake.initializeArgsForCall[i].ready
}

func (fake *FakeFetchSource) InitializeReturns(result1 resource.VersionedSource, result2 error) {
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) InitializeReturnsOnCall(i int, result1 resource.VersionedSource, result2 error) {
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 resource.VersionedSource
			result2 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeFetchSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	fake.findInitializedMutex.RLock()
	defer fake.findInitializedMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetchSource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.FetchSource = new(FakeFetchSource)
