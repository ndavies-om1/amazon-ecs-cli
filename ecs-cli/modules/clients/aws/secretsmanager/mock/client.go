// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/secretsmanager (interfaces: SMClient)

// Package mock_secretsmanager is a generated GoMock package.
package mock_secretsmanager

import (
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSMClient is a mock of SMClient interface
type MockSMClient struct {
	ctrl     *gomock.Controller
	recorder *MockSMClientMockRecorder
}

// MockSMClientMockRecorder is the mock recorder for MockSMClient
type MockSMClientMockRecorder struct {
	mock *MockSMClient
}

// NewMockSMClient creates a new mock instance
func NewMockSMClient(ctrl *gomock.Controller) *MockSMClient {
	mock := &MockSMClient{ctrl: ctrl}
	mock.recorder = &MockSMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSMClient) EXPECT() *MockSMClientMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method
func (m *MockSMClient) CreateSecret(arg0 secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error) {
	ret := m.ctrl.Call(m, "CreateSecret", arg0)
	ret0, _ := ret[0].(*secretsmanager.CreateSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockSMClientMockRecorder) CreateSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSMClient)(nil).CreateSecret), arg0)
}

// DescribeSecret mocks base method
func (m *MockSMClient) DescribeSecret(arg0 string) (*secretsmanager.DescribeSecretOutput, error) {
	ret := m.ctrl.Call(m, "DescribeSecret", arg0)
	ret0, _ := ret[0].(*secretsmanager.DescribeSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecret indicates an expected call of DescribeSecret
func (mr *MockSMClientMockRecorder) DescribeSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecret", reflect.TypeOf((*MockSMClient)(nil).DescribeSecret), arg0)
}

// GetSecretValue mocks base method
func (m *MockSMClient) GetSecretValue(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetSecretValue", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue
func (mr *MockSMClientMockRecorder) GetSecretValue(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSMClient)(nil).GetSecretValue), arg0)
}

// ListSecrets mocks base method
func (m *MockSMClient) ListSecrets(arg0 *string) (*secretsmanager.ListSecretsOutput, error) {
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].(*secretsmanager.ListSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockSMClientMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSMClient)(nil).ListSecrets), arg0)
}

// PutSecretValue mocks base method
func (m *MockSMClient) PutSecretValue(arg0 secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	ret := m.ctrl.Call(m, "PutSecretValue", arg0)
	ret0, _ := ret[0].(*secretsmanager.PutSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSecretValue indicates an expected call of PutSecretValue
func (mr *MockSMClientMockRecorder) PutSecretValue(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSecretValue", reflect.TypeOf((*MockSMClient)(nil).PutSecretValue), arg0)
}
