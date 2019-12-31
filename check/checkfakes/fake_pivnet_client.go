// Code generated by counterfeiter. DO NOT EDIT.
package checkfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v3"
)

type FakePivnetClient struct {
	GetReleaseStub        func(string, string) (pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	getReleaseReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	ReleaseDependenciesStub        func(string, int) ([]pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		arg1 string
		arg2 int
	}
	releaseDependenciesReturns struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	releaseDependenciesReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	ReleaseTypesStub        func() ([]pivnet.ReleaseType, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct {
	}
	releaseTypesReturns struct {
		result1 []pivnet.ReleaseType
		result2 error
	}
	releaseTypesReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseType
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	releasesForProductSlugReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) GetRelease(arg1 string, arg2 string) (pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	ret, specificReturn := fake.getReleaseReturnsOnCall[len(fake.getReleaseArgsForCall)]
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetRelease", []interface{}{arg1, arg2})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *FakePivnetClient) GetReleaseCalls(stub func(string, string) (pivnet.Release, error)) {
	fake.getReleaseMutex.Lock()
	defer fake.getReleaseMutex.Unlock()
	fake.GetReleaseStub = stub
}

func (fake *FakePivnetClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	argsForCall := fake.getReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) GetReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.getReleaseMutex.Lock()
	defer fake.getReleaseMutex.Unlock()
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetReleaseReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.getReleaseMutex.Lock()
	defer fake.getReleaseMutex.Unlock()
	fake.GetReleaseStub = nil
	if fake.getReleaseReturnsOnCall == nil {
		fake.getReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.getReleaseReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependencies(arg1 string, arg2 int) ([]pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	ret, specificReturn := fake.releaseDependenciesReturnsOnCall[len(fake.releaseDependenciesArgsForCall)]
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ReleaseDependencies", []interface{}{arg1, arg2})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseDependenciesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseDependenciesCalls(stub func(string, int) ([]pivnet.ReleaseDependency, error)) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = stub
}

func (fake *FakePivnetClient) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	argsForCall := fake.releaseDependenciesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ReleaseDependenciesReturns(result1 []pivnet.ReleaseDependency, result2 error) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependenciesReturnsOnCall(i int, result1 []pivnet.ReleaseDependency, result2 error) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = nil
	if fake.releaseDependenciesReturnsOnCall == nil {
		fake.releaseDependenciesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseDependency
			result2 error
		})
	}
	fake.releaseDependenciesReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseTypes() ([]pivnet.ReleaseType, error) {
	fake.releaseTypesMutex.Lock()
	ret, specificReturn := fake.releaseTypesReturnsOnCall[len(fake.releaseTypesArgsForCall)]
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseTypesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseTypesCalls(stub func() ([]pivnet.ReleaseType, error)) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = stub
}

func (fake *FakePivnetClient) ReleaseTypesReturns(result1 []pivnet.ReleaseType, result2 error) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseTypesReturnsOnCall(i int, result1 []pivnet.ReleaseType, result2 error) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = nil
	if fake.releaseTypesReturnsOnCall == nil {
		fake.releaseTypesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseType
			result2 error
		})
	}
	fake.releaseTypesReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleasesForProductSlug(arg1 string) ([]pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	ret, specificReturn := fake.releasesForProductSlugReturnsOnCall[len(fake.releasesForProductSlugArgsForCall)]
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releasesForProductSlugReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *FakePivnetClient) ReleasesForProductSlugCalls(stub func(string) ([]pivnet.Release, error)) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = stub
}

func (fake *FakePivnetClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	argsForCall := fake.releasesForProductSlugArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturns(result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	if fake.releasesForProductSlugReturnsOnCall == nil {
		fake.releasesForProductSlugReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.releasesForProductSlugReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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
