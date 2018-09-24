// Code generated by counterfeiter. DO NOT EDIT.
package enginefakes

import (
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	atc "github.com/concourse/concourse/atc"
	engine "github.com/concourse/concourse/atc/engine"
	exec "github.com/concourse/concourse/atc/exec"
)

type FakeBuildDelegate struct {
	BuildStepDelegateStub        func(atc.PlanID) exec.BuildStepDelegate
	buildStepDelegateMutex       sync.RWMutex
	buildStepDelegateArgsForCall []struct {
		arg1 atc.PlanID
	}
	buildStepDelegateReturns struct {
		result1 exec.BuildStepDelegate
	}
	buildStepDelegateReturnsOnCall map[int]struct {
		result1 exec.BuildStepDelegate
	}
	FinishStub        func(lager.Logger, error, bool)
	finishMutex       sync.RWMutex
	finishArgsForCall []struct {
		arg1 lager.Logger
		arg2 error
		arg3 bool
	}
	GetDelegateStub        func(atc.PlanID) exec.GetDelegate
	getDelegateMutex       sync.RWMutex
	getDelegateArgsForCall []struct {
		arg1 atc.PlanID
	}
	getDelegateReturns struct {
		result1 exec.GetDelegate
	}
	getDelegateReturnsOnCall map[int]struct {
		result1 exec.GetDelegate
	}
	PutDelegateStub        func(atc.PlanID) exec.PutDelegate
	putDelegateMutex       sync.RWMutex
	putDelegateArgsForCall []struct {
		arg1 atc.PlanID
	}
	putDelegateReturns struct {
		result1 exec.PutDelegate
	}
	putDelegateReturnsOnCall map[int]struct {
		result1 exec.PutDelegate
	}
	TaskDelegateStub        func(atc.PlanID) exec.TaskDelegate
	taskDelegateMutex       sync.RWMutex
	taskDelegateArgsForCall []struct {
		arg1 atc.PlanID
	}
	taskDelegateReturns struct {
		result1 exec.TaskDelegate
	}
	taskDelegateReturnsOnCall map[int]struct {
		result1 exec.TaskDelegate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildDelegate) BuildStepDelegate(arg1 atc.PlanID) exec.BuildStepDelegate {
	fake.buildStepDelegateMutex.Lock()
	ret, specificReturn := fake.buildStepDelegateReturnsOnCall[len(fake.buildStepDelegateArgsForCall)]
	fake.buildStepDelegateArgsForCall = append(fake.buildStepDelegateArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	fake.recordInvocation("BuildStepDelegate", []interface{}{arg1})
	fake.buildStepDelegateMutex.Unlock()
	if fake.BuildStepDelegateStub != nil {
		return fake.BuildStepDelegateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildStepDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeBuildDelegate) BuildStepDelegateCallCount() int {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	return len(fake.buildStepDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) BuildStepDelegateArgsForCall(i int) atc.PlanID {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	argsForCall := fake.buildStepDelegateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildDelegate) BuildStepDelegateReturns(result1 exec.BuildStepDelegate) {
	fake.BuildStepDelegateStub = nil
	fake.buildStepDelegateReturns = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) BuildStepDelegateReturnsOnCall(i int, result1 exec.BuildStepDelegate) {
	fake.BuildStepDelegateStub = nil
	if fake.buildStepDelegateReturnsOnCall == nil {
		fake.buildStepDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.BuildStepDelegate
		})
	}
	fake.buildStepDelegateReturnsOnCall[i] = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) Finish(arg1 lager.Logger, arg2 error, arg3 bool) {
	fake.finishMutex.Lock()
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct {
		arg1 lager.Logger
		arg2 error
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("Finish", []interface{}{arg1, arg2, arg3})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		fake.FinishStub(arg1, arg2, arg3)
	}
}

func (fake *FakeBuildDelegate) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}

func (fake *FakeBuildDelegate) FinishArgsForCall(i int) (lager.Logger, error, bool) {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	argsForCall := fake.finishArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBuildDelegate) GetDelegate(arg1 atc.PlanID) exec.GetDelegate {
	fake.getDelegateMutex.Lock()
	ret, specificReturn := fake.getDelegateReturnsOnCall[len(fake.getDelegateArgsForCall)]
	fake.getDelegateArgsForCall = append(fake.getDelegateArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	fake.recordInvocation("GetDelegate", []interface{}{arg1})
	fake.getDelegateMutex.Unlock()
	if fake.GetDelegateStub != nil {
		return fake.GetDelegateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeBuildDelegate) GetDelegateCallCount() int {
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	return len(fake.getDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) GetDelegateArgsForCall(i int) atc.PlanID {
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	argsForCall := fake.getDelegateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildDelegate) GetDelegateReturns(result1 exec.GetDelegate) {
	fake.GetDelegateStub = nil
	fake.getDelegateReturns = struct {
		result1 exec.GetDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) GetDelegateReturnsOnCall(i int, result1 exec.GetDelegate) {
	fake.GetDelegateStub = nil
	if fake.getDelegateReturnsOnCall == nil {
		fake.getDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.GetDelegate
		})
	}
	fake.getDelegateReturnsOnCall[i] = struct {
		result1 exec.GetDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) PutDelegate(arg1 atc.PlanID) exec.PutDelegate {
	fake.putDelegateMutex.Lock()
	ret, specificReturn := fake.putDelegateReturnsOnCall[len(fake.putDelegateArgsForCall)]
	fake.putDelegateArgsForCall = append(fake.putDelegateArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	fake.recordInvocation("PutDelegate", []interface{}{arg1})
	fake.putDelegateMutex.Unlock()
	if fake.PutDelegateStub != nil {
		return fake.PutDelegateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeBuildDelegate) PutDelegateCallCount() int {
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	return len(fake.putDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) PutDelegateArgsForCall(i int) atc.PlanID {
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	argsForCall := fake.putDelegateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildDelegate) PutDelegateReturns(result1 exec.PutDelegate) {
	fake.PutDelegateStub = nil
	fake.putDelegateReturns = struct {
		result1 exec.PutDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) PutDelegateReturnsOnCall(i int, result1 exec.PutDelegate) {
	fake.PutDelegateStub = nil
	if fake.putDelegateReturnsOnCall == nil {
		fake.putDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.PutDelegate
		})
	}
	fake.putDelegateReturnsOnCall[i] = struct {
		result1 exec.PutDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) TaskDelegate(arg1 atc.PlanID) exec.TaskDelegate {
	fake.taskDelegateMutex.Lock()
	ret, specificReturn := fake.taskDelegateReturnsOnCall[len(fake.taskDelegateArgsForCall)]
	fake.taskDelegateArgsForCall = append(fake.taskDelegateArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	fake.recordInvocation("TaskDelegate", []interface{}{arg1})
	fake.taskDelegateMutex.Unlock()
	if fake.TaskDelegateStub != nil {
		return fake.TaskDelegateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.taskDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeBuildDelegate) TaskDelegateCallCount() int {
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	return len(fake.taskDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) TaskDelegateArgsForCall(i int) atc.PlanID {
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	argsForCall := fake.taskDelegateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildDelegate) TaskDelegateReturns(result1 exec.TaskDelegate) {
	fake.TaskDelegateStub = nil
	fake.taskDelegateReturns = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) TaskDelegateReturnsOnCall(i int, result1 exec.TaskDelegate) {
	fake.TaskDelegateStub = nil
	if fake.taskDelegateReturnsOnCall == nil {
		fake.taskDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.TaskDelegate
		})
	}
	fake.taskDelegateReturnsOnCall[i] = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildDelegate) recordInvocation(key string, args []interface{}) {
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

var _ engine.BuildDelegate = new(FakeBuildDelegate)