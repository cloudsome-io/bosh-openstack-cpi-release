// Code generated by counterfeiter. DO NOT EDIT.
package computefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/compute"
	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/properties"
)

type FakeAvailabilityZoneProvider struct {
	GetAvailabilityZonesStub        func(properties.CreateVM) []string
	getAvailabilityZonesMutex       sync.RWMutex
	getAvailabilityZonesArgsForCall []struct {
		arg1 properties.CreateVM
	}
	getAvailabilityZonesReturns struct {
		result1 []string
	}
	getAvailabilityZonesReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZones(arg1 properties.CreateVM) []string {
	fake.getAvailabilityZonesMutex.Lock()
	ret, specificReturn := fake.getAvailabilityZonesReturnsOnCall[len(fake.getAvailabilityZonesArgsForCall)]
	fake.getAvailabilityZonesArgsForCall = append(fake.getAvailabilityZonesArgsForCall, struct {
		arg1 properties.CreateVM
	}{arg1})
	stub := fake.GetAvailabilityZonesStub
	fakeReturns := fake.getAvailabilityZonesReturns
	fake.recordInvocation("GetAvailabilityZones", []interface{}{arg1})
	fake.getAvailabilityZonesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZonesCallCount() int {
	fake.getAvailabilityZonesMutex.RLock()
	defer fake.getAvailabilityZonesMutex.RUnlock()
	return len(fake.getAvailabilityZonesArgsForCall)
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZonesCalls(stub func(properties.CreateVM) []string) {
	fake.getAvailabilityZonesMutex.Lock()
	defer fake.getAvailabilityZonesMutex.Unlock()
	fake.GetAvailabilityZonesStub = stub
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZonesArgsForCall(i int) properties.CreateVM {
	fake.getAvailabilityZonesMutex.RLock()
	defer fake.getAvailabilityZonesMutex.RUnlock()
	argsForCall := fake.getAvailabilityZonesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZonesReturns(result1 []string) {
	fake.getAvailabilityZonesMutex.Lock()
	defer fake.getAvailabilityZonesMutex.Unlock()
	fake.GetAvailabilityZonesStub = nil
	fake.getAvailabilityZonesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeAvailabilityZoneProvider) GetAvailabilityZonesReturnsOnCall(i int, result1 []string) {
	fake.getAvailabilityZonesMutex.Lock()
	defer fake.getAvailabilityZonesMutex.Unlock()
	fake.GetAvailabilityZonesStub = nil
	if fake.getAvailabilityZonesReturnsOnCall == nil {
		fake.getAvailabilityZonesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getAvailabilityZonesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeAvailabilityZoneProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAvailabilityZonesMutex.RLock()
	defer fake.getAvailabilityZonesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAvailabilityZoneProvider) recordInvocation(key string, args []interface{}) {
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

var _ compute.AvailabilityZoneProvider = new(FakeAvailabilityZoneProvider)
