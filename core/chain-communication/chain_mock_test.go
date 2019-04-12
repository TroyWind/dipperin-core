// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/dipperin/dipperin-core/core/chain-communication (interfaces: Chain)

package chain_communication

import (
	common "github.com/dipperin/dipperin-core/common"
	model "github.com/dipperin/dipperin-core/core/model"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Chain interface
type MockChain struct {
	ctrl     *gomock.Controller
	recorder *_MockChainRecorder
}

// Recorder for MockChain (not exported)
type _MockChainRecorder struct {
	mock *MockChain
}

func NewMockChain(ctrl *gomock.Controller) *MockChain {
	mock := &MockChain{ctrl: ctrl}
	mock.recorder = &_MockChainRecorder{mock}
	return mock
}

func (_m *MockChain) EXPECT() *_MockChainRecorder {
	return _m.recorder
}

func (_m *MockChain) CurrentBlock() model.AbstractBlock {
	ret := _m.ctrl.Call(_m, "CurrentBlock")
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

func (_mr *_MockChainRecorder) CurrentBlock() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CurrentBlock")
}

func (_m *MockChain) GetBlockByHash(_param0 common.Hash) model.AbstractBlock {
	ret := _m.ctrl.Call(_m, "GetBlockByHash", _param0)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

func (_mr *_MockChainRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockByHash", arg0)
}

func (_m *MockChain) GetBlockByNumber(_param0 uint64) model.AbstractBlock {
	ret := _m.ctrl.Call(_m, "GetBlockByNumber", _param0)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

func (_mr *_MockChainRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockByNumber", arg0)
}

func (_m *MockChain) GetSeenCommit(_param0 uint64) []model.AbstractVerification {
	ret := _m.ctrl.Call(_m, "GetSeenCommit", _param0)
	ret0, _ := ret[0].([]model.AbstractVerification)
	return ret0
}

func (_mr *_MockChainRecorder) GetSeenCommit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSeenCommit", arg0)
}

func (_m *MockChain) GetSlot(_param0 model.AbstractBlock) *uint64 {
	ret := _m.ctrl.Call(_m, "GetSlot", _param0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

func (_mr *_MockChainRecorder) GetSlot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSlot", arg0)
}

func (_m *MockChain) IsChangePoint(_param0 model.AbstractBlock, _param1 bool) bool {
	ret := _m.ctrl.Call(_m, "IsChangePoint", _param0, _param1)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockChainRecorder) IsChangePoint(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsChangePoint", arg0, arg1)
}

func (_m *MockChain) SaveBlock(_param0 model.AbstractBlock, _param1 []model.AbstractVerification) error {
	ret := _m.ctrl.Call(_m, "SaveBlock", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChainRecorder) SaveBlock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveBlock", arg0, arg1)
}
