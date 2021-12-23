// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package storagemock is a generated GoMock package.
package storagemock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/itiky/practicum-examples/04_pgsql/model"
	input "github.com/itiky/practicum-examples/04_pgsql/pkg/input"
)

// MockUserWriter is a mock of UserWriter interface.
type MockUserWriter struct {
	ctrl     *gomock.Controller
	recorder *MockUserWriterMockRecorder
}

// MockUserWriterMockRecorder is the mock recorder for MockUserWriter.
type MockUserWriterMockRecorder struct {
	mock *MockUserWriter
}

// NewMockUserWriter creates a new mock instance.
func NewMockUserWriter(ctrl *gomock.Controller) *MockUserWriter {
	mock := &MockUserWriter{ctrl: ctrl}
	mock.recorder = &MockUserWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserWriter) EXPECT() *MockUserWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockUserWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUserWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUserWriter)(nil).Close))
}

// CreateUser mocks base method.
func (m *MockUserWriter) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserWriterMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserWriter)(nil).CreateUser), ctx, user)
}

// MockUserReader is a mock of UserReader interface.
type MockUserReader struct {
	ctrl     *gomock.Controller
	recorder *MockUserReaderMockRecorder
}

// MockUserReaderMockRecorder is the mock recorder for MockUserReader.
type MockUserReaderMockRecorder struct {
	mock *MockUserReader
}

// NewMockUserReader creates a new mock instance.
func NewMockUserReader(ctrl *gomock.Controller) *MockUserReader {
	mock := &MockUserReader{ctrl: ctrl}
	mock.recorder = &MockUserReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserReader) EXPECT() *MockUserReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockUserReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUserReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUserReader)(nil).Close))
}

// GetUserByEmail mocks base method.
func (m *MockUserReader) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserReaderMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserReader)(nil).GetUserByEmail), ctx, email)
}

// MockOrderWriter is a mock of OrderWriter interface.
type MockOrderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockOrderWriterMockRecorder
}

// MockOrderWriterMockRecorder is the mock recorder for MockOrderWriter.
type MockOrderWriterMockRecorder struct {
	mock *MockOrderWriter
}

// NewMockOrderWriter creates a new mock instance.
func NewMockOrderWriter(ctrl *gomock.Controller) *MockOrderWriter {
	mock := &MockOrderWriter{ctrl: ctrl}
	mock.recorder = &MockOrderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderWriter) EXPECT() *MockOrderWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockOrderWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockOrderWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOrderWriter)(nil).Close))
}

// CreateOrder mocks base method.
func (m *MockOrderWriter) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderWriterMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderWriter)(nil).CreateOrder), ctx, order)
}

// MockOrderReader is a mock of OrderReader interface.
type MockOrderReader struct {
	ctrl     *gomock.Controller
	recorder *MockOrderReaderMockRecorder
}

// MockOrderReaderMockRecorder is the mock recorder for MockOrderReader.
type MockOrderReaderMockRecorder struct {
	mock *MockOrderReader
}

// NewMockOrderReader creates a new mock instance.
func NewMockOrderReader(ctrl *gomock.Controller) *MockOrderReader {
	mock := &MockOrderReader{ctrl: ctrl}
	mock.recorder = &MockOrderReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderReader) EXPECT() *MockOrderReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockOrderReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockOrderReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOrderReader)(nil).Close))
}

// GetOrderByID mocks base method.
func (m *MockOrderReader) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrderReaderMockRecorder) GetOrderByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrderReader)(nil).GetOrderByID), ctx, id)
}

// GetOrdersForUser mocks base method.
func (m *MockOrderReader) GetOrdersForUser(ctx context.Context, userID uuid.UUID, createdAtRangeStart, createdAtRangeEnd *time.Time, pageParams input.PageParams) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersForUser", ctx, userID, createdAtRangeStart, createdAtRangeEnd, pageParams)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersForUser indicates an expected call of GetOrdersForUser.
func (mr *MockOrderReaderMockRecorder) GetOrdersForUser(ctx, userID, createdAtRangeStart, createdAtRangeEnd, pageParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersForUser", reflect.TypeOf((*MockOrderReader)(nil).GetOrdersForUser), ctx, userID, createdAtRangeStart, createdAtRangeEnd, pageParams)
}