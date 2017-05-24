// Code generated by counterfeiter. DO NOT EDIT.
package gcfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeContainerFactory struct {
	FindContainersForDeletionStub        func() ([]dbng.CreatingContainer, []dbng.CreatedContainer, []dbng.DestroyingContainer, error)
	findContainersForDeletionMutex       sync.RWMutex
	findContainersForDeletionArgsForCall []struct{}
	findContainersForDeletionReturns     struct {
		result1 []dbng.CreatingContainer
		result2 []dbng.CreatedContainer
		result3 []dbng.DestroyingContainer
		result4 error
	}
	findContainersForDeletionReturnsOnCall map[int]struct {
		result1 []dbng.CreatingContainer
		result2 []dbng.CreatedContainer
		result3 []dbng.DestroyingContainer
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerFactory) FindContainersForDeletion() ([]dbng.CreatingContainer, []dbng.CreatedContainer, []dbng.DestroyingContainer, error) {
	fake.findContainersForDeletionMutex.Lock()
	ret, specificReturn := fake.findContainersForDeletionReturnsOnCall[len(fake.findContainersForDeletionArgsForCall)]
	fake.findContainersForDeletionArgsForCall = append(fake.findContainersForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("FindContainersForDeletion", []interface{}{})
	fake.findContainersForDeletionMutex.Unlock()
	if fake.FindContainersForDeletionStub != nil {
		return fake.FindContainersForDeletionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.findContainersForDeletionReturns.result1, fake.findContainersForDeletionReturns.result2, fake.findContainersForDeletionReturns.result3, fake.findContainersForDeletionReturns.result4
}

func (fake *FakeContainerFactory) FindContainersForDeletionCallCount() int {
	fake.findContainersForDeletionMutex.RLock()
	defer fake.findContainersForDeletionMutex.RUnlock()
	return len(fake.findContainersForDeletionArgsForCall)
}

func (fake *FakeContainerFactory) FindContainersForDeletionReturns(result1 []dbng.CreatingContainer, result2 []dbng.CreatedContainer, result3 []dbng.DestroyingContainer, result4 error) {
	fake.FindContainersForDeletionStub = nil
	fake.findContainersForDeletionReturns = struct {
		result1 []dbng.CreatingContainer
		result2 []dbng.CreatedContainer
		result3 []dbng.DestroyingContainer
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeContainerFactory) FindContainersForDeletionReturnsOnCall(i int, result1 []dbng.CreatingContainer, result2 []dbng.CreatedContainer, result3 []dbng.DestroyingContainer, result4 error) {
	fake.FindContainersForDeletionStub = nil
	if fake.findContainersForDeletionReturnsOnCall == nil {
		fake.findContainersForDeletionReturnsOnCall = make(map[int]struct {
			result1 []dbng.CreatingContainer
			result2 []dbng.CreatedContainer
			result3 []dbng.DestroyingContainer
			result4 error
		})
	}
	fake.findContainersForDeletionReturnsOnCall[i] = struct {
		result1 []dbng.CreatingContainer
		result2 []dbng.CreatedContainer
		result3 []dbng.DestroyingContainer
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeContainerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findContainersForDeletionMutex.RLock()
	defer fake.findContainersForDeletionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerFactory) recordInvocation(key string, args []interface{}) {
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
