// Code generated by mockery v2.46.0. DO NOT EDIT.

package contractmocks

import (
	context "context"

	core "github.com/hyperledger/firefly/pkg/core"

	ffapi "github.com/hyperledger/firefly-common/pkg/ffapi"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddContractAPIListener provides a mock function with given fields: ctx, apiName, eventPath, listener
func (_m *Manager) AddContractAPIListener(ctx context.Context, apiName string, eventPath string, listener *core.ContractListener) (*core.ContractListener, error) {
	ret := _m.Called(ctx, apiName, eventPath, listener)

	if len(ret) == 0 {
		panic("no return value specified for AddContractAPIListener")
	}

	var r0 *core.ContractListener
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *core.ContractListener) (*core.ContractListener, error)); ok {
		return rf(ctx, apiName, eventPath, listener)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *core.ContractListener) *core.ContractListener); ok {
		r0 = rf(ctx, apiName, eventPath, listener)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractListener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *core.ContractListener) error); ok {
		r1 = rf(ctx, apiName, eventPath, listener)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddContractListener provides a mock function with given fields: ctx, listener
func (_m *Manager) AddContractListener(ctx context.Context, listener *core.ContractListenerInput) (*core.ContractListener, error) {
	ret := _m.Called(ctx, listener)

	if len(ret) == 0 {
		panic("no return value specified for AddContractListener")
	}

	var r0 *core.ContractListener
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListenerInput) (*core.ContractListener, error)); ok {
		return rf(ctx, listener)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListenerInput) *core.ContractListener); ok {
		r0 = rf(ctx, listener)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractListener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.ContractListenerInput) error); ok {
		r1 = rf(ctx, listener)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConstructContractListenerSignature provides a mock function with given fields: ctx, listener
func (_m *Manager) ConstructContractListenerSignature(ctx context.Context, listener *core.ContractListenerInput) (*core.ContractListenerSignatureOutput, error) {
	ret := _m.Called(ctx, listener)

	if len(ret) == 0 {
		panic("no return value specified for ConstructContractListenerSignature")
	}

	var r0 *core.ContractListenerSignatureOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListenerInput) (*core.ContractListenerSignatureOutput, error)); ok {
		return rf(ctx, listener)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListenerInput) *core.ContractListenerSignatureOutput); ok {
		r0 = rf(ctx, listener)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractListenerSignatureOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.ContractListenerInput) error); ok {
		r1 = rf(ctx, listener)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContractAPI provides a mock function with given fields: ctx, apiName
func (_m *Manager) DeleteContractAPI(ctx context.Context, apiName string) error {
	ret := _m.Called(ctx, apiName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContractAPI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, apiName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteContractListenerByNameOrID provides a mock function with given fields: ctx, nameOrID
func (_m *Manager) DeleteContractListenerByNameOrID(ctx context.Context, nameOrID string) error {
	ret := _m.Called(ctx, nameOrID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContractListenerByNameOrID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFFI provides a mock function with given fields: ctx, id
func (_m *Manager) DeleteFFI(ctx context.Context, id *fftypes.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFFI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployContract provides a mock function with given fields: ctx, req, waitConfirm
func (_m *Manager) DeployContract(ctx context.Context, req *core.ContractDeployRequest, waitConfirm bool) (interface{}, error) {
	ret := _m.Called(ctx, req, waitConfirm)

	if len(ret) == 0 {
		panic("no return value specified for DeployContract")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractDeployRequest, bool) (interface{}, error)); ok {
		return rf(ctx, req, waitConfirm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractDeployRequest, bool) interface{}); ok {
		r0 = rf(ctx, req, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.ContractDeployRequest, bool) error); ok {
		r1 = rf(ctx, req, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateFFI provides a mock function with given fields: ctx, generationRequest
func (_m *Manager) GenerateFFI(ctx context.Context, generationRequest *fftypes.FFIGenerationRequest) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, generationRequest)

	if len(ret) == 0 {
		panic("no return value specified for GenerateFFI")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIGenerationRequest) (*fftypes.FFI, error)); ok {
		return rf(ctx, generationRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIGenerationRequest) *fftypes.FFI); ok {
		r0 = rf(ctx, generationRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.FFIGenerationRequest) error); ok {
		r1 = rf(ctx, generationRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractAPI provides a mock function with given fields: ctx, httpServerURL, apiName
func (_m *Manager) GetContractAPI(ctx context.Context, httpServerURL string, apiName string) (*core.ContractAPI, error) {
	ret := _m.Called(ctx, httpServerURL, apiName)

	if len(ret) == 0 {
		panic("no return value specified for GetContractAPI")
	}

	var r0 *core.ContractAPI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*core.ContractAPI, error)); ok {
		return rf(ctx, httpServerURL, apiName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *core.ContractAPI); ok {
		r0 = rf(ctx, httpServerURL, apiName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractAPI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, httpServerURL, apiName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractAPIInterface provides a mock function with given fields: ctx, apiName
func (_m *Manager) GetContractAPIInterface(ctx context.Context, apiName string) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, apiName)

	if len(ret) == 0 {
		panic("no return value specified for GetContractAPIInterface")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*fftypes.FFI, error)); ok {
		return rf(ctx, apiName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.FFI); ok {
		r0 = rf(ctx, apiName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, apiName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractAPIListeners provides a mock function with given fields: ctx, apiName, eventPath, filter
func (_m *Manager) GetContractAPIListeners(ctx context.Context, apiName string, eventPath string, filter ffapi.AndFilter) ([]*core.ContractListener, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, apiName, eventPath, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetContractAPIListeners")
	}

	var r0 []*core.ContractListener
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ffapi.AndFilter) ([]*core.ContractListener, *ffapi.FilterResult, error)); ok {
		return rf(ctx, apiName, eventPath, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ffapi.AndFilter) []*core.ContractListener); ok {
		r0 = rf(ctx, apiName, eventPath, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.ContractListener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, apiName, eventPath, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, apiName, eventPath, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetContractAPIs provides a mock function with given fields: ctx, httpServerURL, filter
func (_m *Manager) GetContractAPIs(ctx context.Context, httpServerURL string, filter ffapi.AndFilter) ([]*core.ContractAPI, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, httpServerURL, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetContractAPIs")
	}

	var r0 []*core.ContractAPI
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) ([]*core.ContractAPI, *ffapi.FilterResult, error)); ok {
		return rf(ctx, httpServerURL, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) []*core.ContractAPI); ok {
		r0 = rf(ctx, httpServerURL, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.ContractAPI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, httpServerURL, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, httpServerURL, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetContractListenerByNameOrID provides a mock function with given fields: ctx, nameOrID
func (_m *Manager) GetContractListenerByNameOrID(ctx context.Context, nameOrID string) (*core.ContractListener, error) {
	ret := _m.Called(ctx, nameOrID)

	if len(ret) == 0 {
		panic("no return value specified for GetContractListenerByNameOrID")
	}

	var r0 *core.ContractListener
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.ContractListener, error)); ok {
		return rf(ctx, nameOrID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.ContractListener); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractListener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameOrID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractListenerByNameOrIDWithStatus provides a mock function with given fields: ctx, nameOrID
func (_m *Manager) GetContractListenerByNameOrIDWithStatus(ctx context.Context, nameOrID string) (*core.ContractListenerWithStatus, error) {
	ret := _m.Called(ctx, nameOrID)

	if len(ret) == 0 {
		panic("no return value specified for GetContractListenerByNameOrIDWithStatus")
	}

	var r0 *core.ContractListenerWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.ContractListenerWithStatus, error)); ok {
		return rf(ctx, nameOrID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.ContractListenerWithStatus); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ContractListenerWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameOrID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractListeners provides a mock function with given fields: ctx, filter
func (_m *Manager) GetContractListeners(ctx context.Context, filter ffapi.AndFilter) ([]*core.ContractListener, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetContractListeners")
	}

	var r0 []*core.ContractListener
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*core.ContractListener, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.ContractListener); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.ContractListener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetFFI provides a mock function with given fields: ctx, name, version
func (_m *Manager) GetFFI(ctx context.Context, name string, version string) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, name, version)

	if len(ret) == 0 {
		panic("no return value specified for GetFFI")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*fftypes.FFI, error)); ok {
		return rf(ctx, name, version)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.FFI); ok {
		r0 = rf(ctx, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIByID provides a mock function with given fields: ctx, id
func (_m *Manager) GetFFIByID(ctx context.Context, id *fftypes.UUID) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIByID")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) (*fftypes.FFI, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.FFI); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIByIDWithChildren provides a mock function with given fields: ctx, id
func (_m *Manager) GetFFIByIDWithChildren(ctx context.Context, id *fftypes.UUID) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIByIDWithChildren")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) (*fftypes.FFI, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.FFI); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIEvents provides a mock function with given fields: ctx, id
func (_m *Manager) GetFFIEvents(ctx context.Context, id *fftypes.UUID) ([]*fftypes.FFIEvent, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIEvents")
	}

	var r0 []*fftypes.FFIEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) ([]*fftypes.FFIEvent, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) []*fftypes.FFIEvent); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.FFIEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIMethods provides a mock function with given fields: ctx, id
func (_m *Manager) GetFFIMethods(ctx context.Context, id *fftypes.UUID) ([]*fftypes.FFIMethod, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIMethods")
	}

	var r0 []*fftypes.FFIMethod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) ([]*fftypes.FFIMethod, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) []*fftypes.FFIMethod); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.FFIMethod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIWithChildren provides a mock function with given fields: ctx, name, version
func (_m *Manager) GetFFIWithChildren(ctx context.Context, name string, version string) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, name, version)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIWithChildren")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*fftypes.FFI, error)); ok {
		return rf(ctx, name, version)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.FFI); ok {
		r0 = rf(ctx, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFFIs provides a mock function with given fields: ctx, filter
func (_m *Manager) GetFFIs(ctx context.Context, filter ffapi.AndFilter) ([]*fftypes.FFI, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIs")
	}

	var r0 []*fftypes.FFI
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*fftypes.FFI, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*fftypes.FFI); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvokeContract provides a mock function with given fields: ctx, req, waitConfirm
func (_m *Manager) InvokeContract(ctx context.Context, req *core.ContractCallRequest, waitConfirm bool) (interface{}, error) {
	ret := _m.Called(ctx, req, waitConfirm)

	if len(ret) == 0 {
		panic("no return value specified for InvokeContract")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractCallRequest, bool) (interface{}, error)); ok {
		return rf(ctx, req, waitConfirm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractCallRequest, bool) interface{}); ok {
		r0 = rf(ctx, req, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.ContractCallRequest, bool) error); ok {
		r1 = rf(ctx, req, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeContractAPI provides a mock function with given fields: ctx, apiName, methodPath, req, waitConfirm
func (_m *Manager) InvokeContractAPI(ctx context.Context, apiName string, methodPath string, req *core.ContractCallRequest, waitConfirm bool) (interface{}, error) {
	ret := _m.Called(ctx, apiName, methodPath, req, waitConfirm)

	if len(ret) == 0 {
		panic("no return value specified for InvokeContractAPI")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *core.ContractCallRequest, bool) (interface{}, error)); ok {
		return rf(ctx, apiName, methodPath, req, waitConfirm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *core.ContractCallRequest, bool) interface{}); ok {
		r0 = rf(ctx, apiName, methodPath, req, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *core.ContractCallRequest, bool) error); ok {
		r1 = rf(ctx, apiName, methodPath, req, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Manager) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrepareOperation provides a mock function with given fields: ctx, op
func (_m *Manager) PrepareOperation(ctx context.Context, op *core.Operation) (*core.PreparedOperation, error) {
	ret := _m.Called(ctx, op)

	if len(ret) == 0 {
		panic("no return value specified for PrepareOperation")
	}

	var r0 *core.PreparedOperation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) (*core.PreparedOperation, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) *core.PreparedOperation); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PreparedOperation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Operation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveContractAPI provides a mock function with given fields: ctx, httpServerURL, api
func (_m *Manager) ResolveContractAPI(ctx context.Context, httpServerURL string, api *core.ContractAPI) error {
	ret := _m.Called(ctx, httpServerURL, api)

	if len(ret) == 0 {
		panic("no return value specified for ResolveContractAPI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.ContractAPI) error); ok {
		r0 = rf(ctx, httpServerURL, api)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResolveFFI provides a mock function with given fields: ctx, ffi
func (_m *Manager) ResolveFFI(ctx context.Context, ffi *fftypes.FFI) error {
	ret := _m.Called(ctx, ffi)

	if len(ret) == 0 {
		panic("no return value specified for ResolveFFI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFI) error); ok {
		r0 = rf(ctx, ffi)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResolveFFIReference provides a mock function with given fields: ctx, ref
func (_m *Manager) ResolveFFIReference(ctx context.Context, ref *fftypes.FFIReference) error {
	ret := _m.Called(ctx, ref)

	if len(ret) == 0 {
		panic("no return value specified for ResolveFFIReference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIReference) error); ok {
		r0 = rf(ctx, ref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunOperation provides a mock function with given fields: ctx, op
func (_m *Manager) RunOperation(ctx context.Context, op *core.PreparedOperation) (fftypes.JSONObject, core.OpPhase, error) {
	ret := _m.Called(ctx, op)

	if len(ret) == 0 {
		panic("no return value specified for RunOperation")
	}

	var r0 fftypes.JSONObject
	var r1 core.OpPhase
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation) (fftypes.JSONObject, core.OpPhase, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation) fftypes.JSONObject); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.JSONObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.PreparedOperation) core.OpPhase); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Get(1).(core.OpPhase)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *core.PreparedOperation) error); ok {
		r2 = rf(ctx, op)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
