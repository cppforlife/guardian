// Code generated by counterfeiter. DO NOT EDIT.
package imagepluginfakes

import (
	"os/exec"
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/imageplugin"
	"code.cloudfoundry.org/lager"
)

type FakeCommandCreator struct {
	CreateCommandStub        func(log lager.Logger, handle string, spec gardener.RootfsSpec) (*exec.Cmd, error)
	createCommandMutex       sync.RWMutex
	createCommandArgsForCall []struct {
		log    lager.Logger
		handle string
		spec   gardener.RootfsSpec
	}
	createCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	createCommandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	DestroyCommandStub        func(log lager.Logger, handle string) *exec.Cmd
	destroyCommandMutex       sync.RWMutex
	destroyCommandArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyCommandReturns struct {
		result1 *exec.Cmd
	}
	destroyCommandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
	}
	MetricsCommandStub        func(log lager.Logger, handle string) *exec.Cmd
	metricsCommandMutex       sync.RWMutex
	metricsCommandArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	metricsCommandReturns struct {
		result1 *exec.Cmd
	}
	metricsCommandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandCreator) CreateCommand(log lager.Logger, handle string, spec gardener.RootfsSpec) (*exec.Cmd, error) {
	fake.createCommandMutex.Lock()
	ret, specificReturn := fake.createCommandReturnsOnCall[len(fake.createCommandArgsForCall)]
	fake.createCommandArgsForCall = append(fake.createCommandArgsForCall, struct {
		log    lager.Logger
		handle string
		spec   gardener.RootfsSpec
	}{log, handle, spec})
	fake.recordInvocation("CreateCommand", []interface{}{log, handle, spec})
	fake.createCommandMutex.Unlock()
	if fake.CreateCommandStub != nil {
		return fake.CreateCommandStub(log, handle, spec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCommandReturns.result1, fake.createCommandReturns.result2
}

func (fake *FakeCommandCreator) CreateCommandCallCount() int {
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	return len(fake.createCommandArgsForCall)
}

func (fake *FakeCommandCreator) CreateCommandArgsForCall(i int) (lager.Logger, string, gardener.RootfsSpec) {
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	return fake.createCommandArgsForCall[i].log, fake.createCommandArgsForCall[i].handle, fake.createCommandArgsForCall[i].spec
}

func (fake *FakeCommandCreator) CreateCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.CreateCommandStub = nil
	fake.createCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandCreator) CreateCommandReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.CreateCommandStub = nil
	if fake.createCommandReturnsOnCall == nil {
		fake.createCommandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.createCommandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandCreator) DestroyCommand(log lager.Logger, handle string) *exec.Cmd {
	fake.destroyCommandMutex.Lock()
	ret, specificReturn := fake.destroyCommandReturnsOnCall[len(fake.destroyCommandArgsForCall)]
	fake.destroyCommandArgsForCall = append(fake.destroyCommandArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("DestroyCommand", []interface{}{log, handle})
	fake.destroyCommandMutex.Unlock()
	if fake.DestroyCommandStub != nil {
		return fake.DestroyCommandStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyCommandReturns.result1
}

func (fake *FakeCommandCreator) DestroyCommandCallCount() int {
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	return len(fake.destroyCommandArgsForCall)
}

func (fake *FakeCommandCreator) DestroyCommandArgsForCall(i int) (lager.Logger, string) {
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	return fake.destroyCommandArgsForCall[i].log, fake.destroyCommandArgsForCall[i].handle
}

func (fake *FakeCommandCreator) DestroyCommandReturns(result1 *exec.Cmd) {
	fake.DestroyCommandStub = nil
	fake.destroyCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeCommandCreator) DestroyCommandReturnsOnCall(i int, result1 *exec.Cmd) {
	fake.DestroyCommandStub = nil
	if fake.destroyCommandReturnsOnCall == nil {
		fake.destroyCommandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
		})
	}
	fake.destroyCommandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeCommandCreator) MetricsCommand(log lager.Logger, handle string) *exec.Cmd {
	fake.metricsCommandMutex.Lock()
	ret, specificReturn := fake.metricsCommandReturnsOnCall[len(fake.metricsCommandArgsForCall)]
	fake.metricsCommandArgsForCall = append(fake.metricsCommandArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("MetricsCommand", []interface{}{log, handle})
	fake.metricsCommandMutex.Unlock()
	if fake.MetricsCommandStub != nil {
		return fake.MetricsCommandStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.metricsCommandReturns.result1
}

func (fake *FakeCommandCreator) MetricsCommandCallCount() int {
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	return len(fake.metricsCommandArgsForCall)
}

func (fake *FakeCommandCreator) MetricsCommandArgsForCall(i int) (lager.Logger, string) {
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	return fake.metricsCommandArgsForCall[i].log, fake.metricsCommandArgsForCall[i].handle
}

func (fake *FakeCommandCreator) MetricsCommandReturns(result1 *exec.Cmd) {
	fake.MetricsCommandStub = nil
	fake.metricsCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeCommandCreator) MetricsCommandReturnsOnCall(i int, result1 *exec.Cmd) {
	fake.MetricsCommandStub = nil
	if fake.metricsCommandReturnsOnCall == nil {
		fake.metricsCommandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
		})
	}
	fake.metricsCommandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeCommandCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	fake.destroyCommandMutex.RLock()
	defer fake.destroyCommandMutex.RUnlock()
	fake.metricsCommandMutex.RLock()
	defer fake.metricsCommandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommandCreator) recordInvocation(key string, args []interface{}) {
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

var _ imageplugin.CommandCreator = new(FakeCommandCreator)
