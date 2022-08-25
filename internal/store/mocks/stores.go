// Code generated by MockGen. DO NOT EDIT.
// Source: ./stores.go

// Package mock_stores is a generated GoMock package.
package mock_stores

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "gitlab.ozon.dev/krotovkk/homework/internal/model"
	ports "gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Cart mocks base method.
func (m *MockStore) Cart() ports.CartStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cart")
	ret0, _ := ret[0].(ports.CartStore)
	return ret0
}

// Cart indicates an expected call of Cart.
func (mr *MockStoreMockRecorder) Cart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cart", reflect.TypeOf((*MockStore)(nil).Cart))
}

// Product mocks base method.
func (m *MockStore) Product() ports.ProductStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Product")
	ret0, _ := ret[0].(ports.ProductStore)
	return ret0
}

// Product indicates an expected call of Product.
func (mr *MockStoreMockRecorder) Product() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Product", reflect.TypeOf((*MockStore)(nil).Product))
}

// MockProductStore is a mock of ProductStore interface.
type MockProductStore struct {
	ctrl     *gomock.Controller
	recorder *MockProductStoreMockRecorder
}

// MockProductStoreMockRecorder is the mock recorder for MockProductStore.
type MockProductStoreMockRecorder struct {
	mock *MockProductStore
}

// NewMockProductStore creates a new mock instance.
func NewMockProductStore(ctrl *gomock.Controller) *MockProductStore {
	mock := &MockProductStore{ctrl: ctrl}
	mock.recorder = &MockProductStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductStore) EXPECT() *MockProductStoreMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductStore) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, p)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductStoreMockRecorder) CreateProduct(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductStore)(nil).CreateProduct), ctx, p)
}

// DeleteProduct mocks base method.
func (m *MockProductStore) DeleteProduct(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductStoreMockRecorder) DeleteProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductStore)(nil).DeleteProduct), ctx, id)
}

// GetAllProducts mocks base method.
func (m *MockProductStore) GetAllProducts(ctx context.Context, limit, offset uint64) ([]*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductStoreMockRecorder) GetAllProducts(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductStore)(nil).GetAllProducts), ctx, limit, offset)
}

// GetProductOne mocks base method.
func (m *MockProductStore) GetProductOne(ctx context.Context, id int64) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductOne", ctx, id)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductOne indicates an expected call of GetProductOne.
func (mr *MockProductStoreMockRecorder) GetProductOne(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductOne", reflect.TypeOf((*MockProductStore)(nil).GetProductOne), ctx, id)
}

// UpdateProduct mocks base method.
func (m *MockProductStore) UpdateProduct(ctx context.Context, p *model.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductStoreMockRecorder) UpdateProduct(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductStore)(nil).UpdateProduct), ctx, p)
}

// MockCartStore is a mock of CartStore interface.
type MockCartStore struct {
	ctrl     *gomock.Controller
	recorder *MockCartStoreMockRecorder
}

// MockCartStoreMockRecorder is the mock recorder for MockCartStore.
type MockCartStoreMockRecorder struct {
	mock *MockCartStore
}

// NewMockCartStore creates a new mock instance.
func NewMockCartStore(ctrl *gomock.Controller) *MockCartStore {
	mock := &MockCartStore{ctrl: ctrl}
	mock.recorder = &MockCartStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartStore) EXPECT() *MockCartStoreMockRecorder {
	return m.recorder
}

// AddProductToCart mocks base method.
func (m *MockCartStore) AddProductToCart(ctx context.Context, productId, cartId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductToCart", ctx, productId, cartId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProductToCart indicates an expected call of AddProductToCart.
func (mr *MockCartStoreMockRecorder) AddProductToCart(ctx, productId, cartId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductToCart", reflect.TypeOf((*MockCartStore)(nil).AddProductToCart), ctx, productId, cartId)
}

// CreateCart mocks base method.
func (m *MockCartStore) CreateCart(ctx context.Context, c *model.Cart) (*model.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", ctx, c)
	ret0, _ := ret[0].(*model.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartStoreMockRecorder) CreateCart(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartStore)(nil).CreateCart), ctx, c)
}

// GetCartProducts mocks base method.
func (m *MockCartStore) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartProducts", ctx, id)
	ret0, _ := ret[0].([]*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartProducts indicates an expected call of GetCartProducts.
func (mr *MockCartStoreMockRecorder) GetCartProducts(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartProducts", reflect.TypeOf((*MockCartStore)(nil).GetCartProducts), ctx, id)
}