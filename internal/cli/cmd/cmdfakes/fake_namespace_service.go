// Copyright © 2021 - 2023 SUSE LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/epinio/epinio/internal/cli/cmd"
	"github.com/epinio/epinio/internal/cli/usercmd"
)

type FakeNamespaceService struct {
	CreateNamespaceStub        func(string) error
	createNamespaceMutex       sync.RWMutex
	createNamespaceArgsForCall []struct {
		arg1 string
	}
	createNamespaceReturns struct {
		result1 error
	}
	createNamespaceReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteNamespaceStub        func([]string, bool, bool) error
	deleteNamespaceMutex       sync.RWMutex
	deleteNamespaceArgsForCall []struct {
		arg1 []string
		arg2 bool
		arg3 bool
	}
	deleteNamespaceReturns struct {
		result1 error
	}
	deleteNamespaceReturnsOnCall map[int]struct {
		result1 error
	}
	GetAPIStub        func() usercmd.APIClient
	getAPIMutex       sync.RWMutex
	getAPIArgsForCall []struct {
	}
	getAPIReturns struct {
		result1 usercmd.APIClient
	}
	getAPIReturnsOnCall map[int]struct {
		result1 usercmd.APIClient
	}
	NamespacesStub        func() error
	namespacesMutex       sync.RWMutex
	namespacesArgsForCall []struct {
	}
	namespacesReturns struct {
		result1 error
	}
	namespacesReturnsOnCall map[int]struct {
		result1 error
	}
	NamespacesMatchingStub        func(string) []string
	namespacesMatchingMutex       sync.RWMutex
	namespacesMatchingArgsForCall []struct {
		arg1 string
	}
	namespacesMatchingReturns struct {
		result1 []string
	}
	namespacesMatchingReturnsOnCall map[int]struct {
		result1 []string
	}
	ShowNamespaceStub        func(string) error
	showNamespaceMutex       sync.RWMutex
	showNamespaceArgsForCall []struct {
		arg1 string
	}
	showNamespaceReturns struct {
		result1 error
	}
	showNamespaceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNamespaceService) CreateNamespace(arg1 string) error {
	fake.createNamespaceMutex.Lock()
	ret, specificReturn := fake.createNamespaceReturnsOnCall[len(fake.createNamespaceArgsForCall)]
	fake.createNamespaceArgsForCall = append(fake.createNamespaceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateNamespaceStub
	fakeReturns := fake.createNamespaceReturns
	fake.recordInvocation("CreateNamespace", []interface{}{arg1})
	fake.createNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) CreateNamespaceCallCount() int {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	return len(fake.createNamespaceArgsForCall)
}

func (fake *FakeNamespaceService) CreateNamespaceCalls(stub func(string) error) {
	fake.createNamespaceMutex.Lock()
	defer fake.createNamespaceMutex.Unlock()
	fake.CreateNamespaceStub = stub
}

func (fake *FakeNamespaceService) CreateNamespaceArgsForCall(i int) string {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	argsForCall := fake.createNamespaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNamespaceService) CreateNamespaceReturns(result1 error) {
	fake.createNamespaceMutex.Lock()
	defer fake.createNamespaceMutex.Unlock()
	fake.CreateNamespaceStub = nil
	fake.createNamespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) CreateNamespaceReturnsOnCall(i int, result1 error) {
	fake.createNamespaceMutex.Lock()
	defer fake.createNamespaceMutex.Unlock()
	fake.CreateNamespaceStub = nil
	if fake.createNamespaceReturnsOnCall == nil {
		fake.createNamespaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createNamespaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) DeleteNamespace(arg1 []string, arg2 bool, arg3 bool) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteNamespaceMutex.Lock()
	ret, specificReturn := fake.deleteNamespaceReturnsOnCall[len(fake.deleteNamespaceArgsForCall)]
	fake.deleteNamespaceArgsForCall = append(fake.deleteNamespaceArgsForCall, struct {
		arg1 []string
		arg2 bool
		arg3 bool
	}{arg1Copy, arg2, arg3})
	stub := fake.DeleteNamespaceStub
	fakeReturns := fake.deleteNamespaceReturns
	fake.recordInvocation("DeleteNamespace", []interface{}{arg1Copy, arg2, arg3})
	fake.deleteNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) DeleteNamespaceCallCount() int {
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	return len(fake.deleteNamespaceArgsForCall)
}

func (fake *FakeNamespaceService) DeleteNamespaceCalls(stub func([]string, bool, bool) error) {
	fake.deleteNamespaceMutex.Lock()
	defer fake.deleteNamespaceMutex.Unlock()
	fake.DeleteNamespaceStub = stub
}

func (fake *FakeNamespaceService) DeleteNamespaceArgsForCall(i int) ([]string, bool, bool) {
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	argsForCall := fake.deleteNamespaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNamespaceService) DeleteNamespaceReturns(result1 error) {
	fake.deleteNamespaceMutex.Lock()
	defer fake.deleteNamespaceMutex.Unlock()
	fake.DeleteNamespaceStub = nil
	fake.deleteNamespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) DeleteNamespaceReturnsOnCall(i int, result1 error) {
	fake.deleteNamespaceMutex.Lock()
	defer fake.deleteNamespaceMutex.Unlock()
	fake.DeleteNamespaceStub = nil
	if fake.deleteNamespaceReturnsOnCall == nil {
		fake.deleteNamespaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteNamespaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) GetAPI() usercmd.APIClient {
	fake.getAPIMutex.Lock()
	ret, specificReturn := fake.getAPIReturnsOnCall[len(fake.getAPIArgsForCall)]
	fake.getAPIArgsForCall = append(fake.getAPIArgsForCall, struct {
	}{})
	stub := fake.GetAPIStub
	fakeReturns := fake.getAPIReturns
	fake.recordInvocation("GetAPI", []interface{}{})
	fake.getAPIMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) GetAPICallCount() int {
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	return len(fake.getAPIArgsForCall)
}

func (fake *FakeNamespaceService) GetAPICalls(stub func() usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = stub
}

func (fake *FakeNamespaceService) GetAPIReturns(result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	fake.getAPIReturns = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeNamespaceService) GetAPIReturnsOnCall(i int, result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	if fake.getAPIReturnsOnCall == nil {
		fake.getAPIReturnsOnCall = make(map[int]struct {
			result1 usercmd.APIClient
		})
	}
	fake.getAPIReturnsOnCall[i] = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeNamespaceService) Namespaces() error {
	fake.namespacesMutex.Lock()
	ret, specificReturn := fake.namespacesReturnsOnCall[len(fake.namespacesArgsForCall)]
	fake.namespacesArgsForCall = append(fake.namespacesArgsForCall, struct {
	}{})
	stub := fake.NamespacesStub
	fakeReturns := fake.namespacesReturns
	fake.recordInvocation("Namespaces", []interface{}{})
	fake.namespacesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) NamespacesCallCount() int {
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	return len(fake.namespacesArgsForCall)
}

func (fake *FakeNamespaceService) NamespacesCalls(stub func() error) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = stub
}

func (fake *FakeNamespaceService) NamespacesReturns(result1 error) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = nil
	fake.namespacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) NamespacesReturnsOnCall(i int, result1 error) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = nil
	if fake.namespacesReturnsOnCall == nil {
		fake.namespacesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.namespacesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) NamespacesMatching(arg1 string) []string {
	fake.namespacesMatchingMutex.Lock()
	ret, specificReturn := fake.namespacesMatchingReturnsOnCall[len(fake.namespacesMatchingArgsForCall)]
	fake.namespacesMatchingArgsForCall = append(fake.namespacesMatchingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.NamespacesMatchingStub
	fakeReturns := fake.namespacesMatchingReturns
	fake.recordInvocation("NamespacesMatching", []interface{}{arg1})
	fake.namespacesMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) NamespacesMatchingCallCount() int {
	fake.namespacesMatchingMutex.RLock()
	defer fake.namespacesMatchingMutex.RUnlock()
	return len(fake.namespacesMatchingArgsForCall)
}

func (fake *FakeNamespaceService) NamespacesMatchingCalls(stub func(string) []string) {
	fake.namespacesMatchingMutex.Lock()
	defer fake.namespacesMatchingMutex.Unlock()
	fake.NamespacesMatchingStub = stub
}

func (fake *FakeNamespaceService) NamespacesMatchingArgsForCall(i int) string {
	fake.namespacesMatchingMutex.RLock()
	defer fake.namespacesMatchingMutex.RUnlock()
	argsForCall := fake.namespacesMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNamespaceService) NamespacesMatchingReturns(result1 []string) {
	fake.namespacesMatchingMutex.Lock()
	defer fake.namespacesMatchingMutex.Unlock()
	fake.NamespacesMatchingStub = nil
	fake.namespacesMatchingReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeNamespaceService) NamespacesMatchingReturnsOnCall(i int, result1 []string) {
	fake.namespacesMatchingMutex.Lock()
	defer fake.namespacesMatchingMutex.Unlock()
	fake.NamespacesMatchingStub = nil
	if fake.namespacesMatchingReturnsOnCall == nil {
		fake.namespacesMatchingReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.namespacesMatchingReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeNamespaceService) ShowNamespace(arg1 string) error {
	fake.showNamespaceMutex.Lock()
	ret, specificReturn := fake.showNamespaceReturnsOnCall[len(fake.showNamespaceArgsForCall)]
	fake.showNamespaceArgsForCall = append(fake.showNamespaceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ShowNamespaceStub
	fakeReturns := fake.showNamespaceReturns
	fake.recordInvocation("ShowNamespace", []interface{}{arg1})
	fake.showNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNamespaceService) ShowNamespaceCallCount() int {
	fake.showNamespaceMutex.RLock()
	defer fake.showNamespaceMutex.RUnlock()
	return len(fake.showNamespaceArgsForCall)
}

func (fake *FakeNamespaceService) ShowNamespaceCalls(stub func(string) error) {
	fake.showNamespaceMutex.Lock()
	defer fake.showNamespaceMutex.Unlock()
	fake.ShowNamespaceStub = stub
}

func (fake *FakeNamespaceService) ShowNamespaceArgsForCall(i int) string {
	fake.showNamespaceMutex.RLock()
	defer fake.showNamespaceMutex.RUnlock()
	argsForCall := fake.showNamespaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNamespaceService) ShowNamespaceReturns(result1 error) {
	fake.showNamespaceMutex.Lock()
	defer fake.showNamespaceMutex.Unlock()
	fake.ShowNamespaceStub = nil
	fake.showNamespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) ShowNamespaceReturnsOnCall(i int, result1 error) {
	fake.showNamespaceMutex.Lock()
	defer fake.showNamespaceMutex.Unlock()
	fake.ShowNamespaceStub = nil
	if fake.showNamespaceReturnsOnCall == nil {
		fake.showNamespaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.showNamespaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespaceService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	fake.namespacesMatchingMutex.RLock()
	defer fake.namespacesMatchingMutex.RUnlock()
	fake.showNamespaceMutex.RLock()
	defer fake.showNamespaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNamespaceService) recordInvocation(key string, args []interface{}) {
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

var _ cmd.NamespaceService = new(FakeNamespaceService)
