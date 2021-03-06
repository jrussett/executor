// Code generated by counterfeiter. DO NOT EDIT.
package containerstorefakes

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/executor/depot/containerstore"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
)

type FakeCredentialHandler struct {
	CreateDirStub        func(logger lager.Logger, container executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error)
	createDirMutex       sync.RWMutex
	createDirArgsForCall []struct {
		logger    lager.Logger
		container executor.Container
	}
	createDirReturns struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	createDirReturnsOnCall map[int]struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	RemoveDirStub        func(logger lager.Logger, container executor.Container) error
	removeDirMutex       sync.RWMutex
	removeDirArgsForCall []struct {
		logger    lager.Logger
		container executor.Container
	}
	removeDirReturns struct {
		result1 error
	}
	removeDirReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(credentials containerstore.Credential, container executor.Container) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		credentials containerstore.Credential
		container   executor.Container
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func(invalidCredentials containerstore.Credential, container executor.Container) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		invalidCredentials containerstore.Credential
		container          executor.Container
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredentialHandler) CreateDir(logger lager.Logger, container executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error) {
	fake.createDirMutex.Lock()
	ret, specificReturn := fake.createDirReturnsOnCall[len(fake.createDirArgsForCall)]
	fake.createDirArgsForCall = append(fake.createDirArgsForCall, struct {
		logger    lager.Logger
		container executor.Container
	}{logger, container})
	fake.recordInvocation("CreateDir", []interface{}{logger, container})
	fake.createDirMutex.Unlock()
	if fake.CreateDirStub != nil {
		return fake.CreateDirStub(logger, container)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createDirReturns.result1, fake.createDirReturns.result2, fake.createDirReturns.result3
}

func (fake *FakeCredentialHandler) CreateDirCallCount() int {
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	return len(fake.createDirArgsForCall)
}

func (fake *FakeCredentialHandler) CreateDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	return fake.createDirArgsForCall[i].logger, fake.createDirArgsForCall[i].container
}

func (fake *FakeCredentialHandler) CreateDirReturns(result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.CreateDirStub = nil
	fake.createDirReturns = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredentialHandler) CreateDirReturnsOnCall(i int, result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.CreateDirStub = nil
	if fake.createDirReturnsOnCall == nil {
		fake.createDirReturnsOnCall = make(map[int]struct {
			result1 []garden.BindMount
			result2 []executor.EnvironmentVariable
			result3 error
		})
	}
	fake.createDirReturnsOnCall[i] = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredentialHandler) RemoveDir(logger lager.Logger, container executor.Container) error {
	fake.removeDirMutex.Lock()
	ret, specificReturn := fake.removeDirReturnsOnCall[len(fake.removeDirArgsForCall)]
	fake.removeDirArgsForCall = append(fake.removeDirArgsForCall, struct {
		logger    lager.Logger
		container executor.Container
	}{logger, container})
	fake.recordInvocation("RemoveDir", []interface{}{logger, container})
	fake.removeDirMutex.Unlock()
	if fake.RemoveDirStub != nil {
		return fake.RemoveDirStub(logger, container)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeDirReturns.result1
}

func (fake *FakeCredentialHandler) RemoveDirCallCount() int {
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	return len(fake.removeDirArgsForCall)
}

func (fake *FakeCredentialHandler) RemoveDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	return fake.removeDirArgsForCall[i].logger, fake.removeDirArgsForCall[i].container
}

func (fake *FakeCredentialHandler) RemoveDirReturns(result1 error) {
	fake.RemoveDirStub = nil
	fake.removeDirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) RemoveDirReturnsOnCall(i int, result1 error) {
	fake.RemoveDirStub = nil
	if fake.removeDirReturnsOnCall == nil {
		fake.removeDirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeDirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) Update(credentials containerstore.Credential, container executor.Container) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		credentials containerstore.Credential
		container   executor.Container
	}{credentials, container})
	fake.recordInvocation("Update", []interface{}{credentials, container})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(credentials, container)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeCredentialHandler) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeCredentialHandler) UpdateArgsForCall(i int) (containerstore.Credential, executor.Container) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].credentials, fake.updateArgsForCall[i].container
}

func (fake *FakeCredentialHandler) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) UpdateReturnsOnCall(i int, result1 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) Close(invalidCredentials containerstore.Credential, container executor.Container) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		invalidCredentials containerstore.Credential
		container          executor.Container
	}{invalidCredentials, container})
	fake.recordInvocation("Close", []interface{}{invalidCredentials, container})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub(invalidCredentials, container)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeCredentialHandler) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeCredentialHandler) CloseArgsForCall(i int) (containerstore.Credential, executor.Container) {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.closeArgsForCall[i].invalidCredentials, fake.closeArgsForCall[i].container
}

func (fake *FakeCredentialHandler) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredentialHandler) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.CredentialHandler = new(FakeCredentialHandler)
