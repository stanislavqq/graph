// Code generated by MockGen. DO NOT EDIT.
// Source: parser.go

// Package mocks is a generated GoMock package.
package mocks

import (
	parser "graph/parser"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockParser) Parse(fileName string) (map[int][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", fileName)
	ret0, _ := ret[0].(map[int][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserMockRecorder) Parse(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), fileName)
}

// SetSkipRows mocks base method.
func (m *MockParser) SetSkipRows(num int) parser.Parser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSkipRows", num)
	ret0, _ := ret[0].(parser.Parser)
	return ret0
}

// SetSkipRows indicates an expected call of SetSkipRows.
func (mr *MockParserMockRecorder) SetSkipRows(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSkipRows", reflect.TypeOf((*MockParser)(nil).SetSkipRows), num)
}
