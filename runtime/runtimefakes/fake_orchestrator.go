// Code generated by counterfeiter. DO NOT EDIT.
package runtimefakes

import (
	"context"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/runtime"
	"github.com/concourse/atc/worker"
)

type FakeOrchestrator struct {
	RunTaskStub        func(context.Context, runtime.TaskExecutionDelegate, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes, runtime.IOConfig, atc.TaskConfig) (chan runtime.TaskResult, []worker.VolumeMount, error)
	runTaskMutex       sync.RWMutex
	runTaskArgsForCall []struct {
		arg1 context.Context
		arg2 runtime.TaskExecutionDelegate
		arg3 db.ContainerOwner
		arg4 db.ContainerMetadata
		arg5 worker.ContainerSpec
		arg6 creds.VersionedResourceTypes
		arg7 runtime.IOConfig
		arg8 atc.TaskConfig
	}
	runTaskReturns struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}
	runTaskReturnsOnCall map[int]struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrchestrator) RunTask(arg1 context.Context, arg2 runtime.TaskExecutionDelegate, arg3 db.ContainerOwner, arg4 db.ContainerMetadata, arg5 worker.ContainerSpec, arg6 creds.VersionedResourceTypes, arg7 runtime.IOConfig, arg8 atc.TaskConfig) (chan runtime.TaskResult, []worker.VolumeMount, error) {
	fake.runTaskMutex.Lock()
	ret, specificReturn := fake.runTaskReturnsOnCall[len(fake.runTaskArgsForCall)]
	fake.runTaskArgsForCall = append(fake.runTaskArgsForCall, struct {
		arg1 context.Context
		arg2 runtime.TaskExecutionDelegate
		arg3 db.ContainerOwner
		arg4 db.ContainerMetadata
		arg5 worker.ContainerSpec
		arg6 creds.VersionedResourceTypes
		arg7 runtime.IOConfig
		arg8 atc.TaskConfig
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("RunTask", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.runTaskMutex.Unlock()
	if fake.RunTaskStub != nil {
		return fake.RunTaskStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.runTaskReturns.result1, fake.runTaskReturns.result2, fake.runTaskReturns.result3
}

func (fake *FakeOrchestrator) RunTaskCallCount() int {
	fake.runTaskMutex.RLock()
	defer fake.runTaskMutex.RUnlock()
	return len(fake.runTaskArgsForCall)
}

func (fake *FakeOrchestrator) RunTaskArgsForCall(i int) (context.Context, runtime.TaskExecutionDelegate, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes, runtime.IOConfig, atc.TaskConfig) {
	fake.runTaskMutex.RLock()
	defer fake.runTaskMutex.RUnlock()
	return fake.runTaskArgsForCall[i].arg1, fake.runTaskArgsForCall[i].arg2, fake.runTaskArgsForCall[i].arg3, fake.runTaskArgsForCall[i].arg4, fake.runTaskArgsForCall[i].arg5, fake.runTaskArgsForCall[i].arg6, fake.runTaskArgsForCall[i].arg7, fake.runTaskArgsForCall[i].arg8
}

func (fake *FakeOrchestrator) RunTaskReturns(result1 chan runtime.TaskResult, result2 []worker.VolumeMount, result3 error) {
	fake.RunTaskStub = nil
	fake.runTaskReturns = struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrchestrator) RunTaskReturnsOnCall(i int, result1 chan runtime.TaskResult, result2 []worker.VolumeMount, result3 error) {
	fake.RunTaskStub = nil
	if fake.runTaskReturnsOnCall == nil {
		fake.runTaskReturnsOnCall = make(map[int]struct {
			result1 chan runtime.TaskResult
			result2 []worker.VolumeMount
			result3 error
		})
	}
	fake.runTaskReturnsOnCall[i] = struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrchestrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runTaskMutex.RLock()
	defer fake.runTaskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrchestrator) recordInvocation(key string, args []interface{}) {
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

var _ runtime.Orchestrator = new(FakeOrchestrator)
