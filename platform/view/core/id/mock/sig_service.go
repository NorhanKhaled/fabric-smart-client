// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/core/id"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type SigService struct {
	RegisterSignerStub        func(identity view.Identity, signer api.Signer, verifier api.Verifier) error
	registerSignerMutex       sync.RWMutex
	registerSignerArgsForCall []struct {
		identity view.Identity
		signer   api.Signer
		verifier api.Verifier
	}
	registerSignerReturns struct {
		result1 error
	}
	registerSignerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SigService) RegisterSigner(identity view.Identity, signer api.Signer, verifier api.Verifier) error {
	fake.registerSignerMutex.Lock()
	ret, specificReturn := fake.registerSignerReturnsOnCall[len(fake.registerSignerArgsForCall)]
	fake.registerSignerArgsForCall = append(fake.registerSignerArgsForCall, struct {
		identity view.Identity
		signer   api.Signer
		verifier api.Verifier
	}{identity, signer, verifier})
	fake.recordInvocation("RegisterSigner", []interface{}{identity, signer, verifier})
	fake.registerSignerMutex.Unlock()
	if fake.RegisterSignerStub != nil {
		return fake.RegisterSignerStub(identity, signer, verifier)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.registerSignerReturns.result1
}

func (fake *SigService) RegisterSignerCallCount() int {
	fake.registerSignerMutex.RLock()
	defer fake.registerSignerMutex.RUnlock()
	return len(fake.registerSignerArgsForCall)
}

func (fake *SigService) RegisterSignerArgsForCall(i int) (view.Identity, api.Signer, api.Verifier) {
	fake.registerSignerMutex.RLock()
	defer fake.registerSignerMutex.RUnlock()
	return fake.registerSignerArgsForCall[i].identity, fake.registerSignerArgsForCall[i].signer, fake.registerSignerArgsForCall[i].verifier
}

func (fake *SigService) RegisterSignerReturns(result1 error) {
	fake.RegisterSignerStub = nil
	fake.registerSignerReturns = struct {
		result1 error
	}{result1}
}

func (fake *SigService) RegisterSignerReturnsOnCall(i int, result1 error) {
	fake.RegisterSignerStub = nil
	if fake.registerSignerReturnsOnCall == nil {
		fake.registerSignerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerSignerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerSignerMutex.RLock()
	defer fake.registerSignerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SigService) recordInvocation(key string, args []interface{}) {
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

var _ id.SigService = new(SigService)