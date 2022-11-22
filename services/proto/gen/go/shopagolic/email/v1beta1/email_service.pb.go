// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shopagolic/email/v1beta1/email_service.proto

package v1beta1

import (
	v1 "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cost *v1.Money `protobuf:"bytes,2,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *OrderItem) GetCost() *v1.Money {
	if x != nil {
		return x.Cost
	}
	return nil
}

type OrderResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId            string       `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ShippingTrackingId string       `protobuf:"bytes,2,opt,name=shipping_tracking_id,json=shippingTrackingId,proto3" json:"shipping_tracking_id,omitempty"`
	ShippingCost       *v1.Money    `protobuf:"bytes,3,opt,name=shipping_cost,json=shippingCost,proto3" json:"shipping_cost,omitempty"`
	ShippingAddress    *Address     `protobuf:"bytes,4,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	Items              []*OrderItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderResult) Reset() {
	*x = OrderResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResult) ProtoMessage() {}

func (x *OrderResult) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResult.ProtoReflect.Descriptor instead.
func (*OrderResult) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{3}
}

func (x *OrderResult) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderResult) GetShippingTrackingId() string {
	if x != nil {
		return x.ShippingTrackingId
	}
	return ""
}

func (x *OrderResult) GetShippingCost() *v1.Money {
	if x != nil {
		return x.ShippingCost
	}
	return nil
}

func (x *OrderResult) GetShippingAddress() *Address {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

func (x *OrderResult) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type SendOrderConfirmationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Order *OrderResult `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SendOrderConfirmationRequest) Reset() {
	*x = SendOrderConfirmationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderConfirmationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderConfirmationRequest) ProtoMessage() {}

func (x *SendOrderConfirmationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderConfirmationRequest.ProtoReflect.Descriptor instead.
func (*SendOrderConfirmationRequest) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{4}
}

func (x *SendOrderConfirmationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendOrderConfirmationRequest) GetOrder() *OrderResult {
	if x != nil {
		return x.Order
	}
	return nil
}

type SendOrderConfirmationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendOrderConfirmationResponse) Reset() {
	*x = SendOrderConfirmationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderConfirmationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderConfirmationResponse) ProtoMessage() {}

func (x *SendOrderConfirmationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_email_v1beta1_email_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderConfirmationResponse.ProtoReflect.Descriptor instead.
func (*SendOrderConfirmationResponse) Descriptor() ([]byte, []int) {
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP(), []int{5}
}

var File_shopagolic_email_v1beta1_email_service_proto protoreflect.FileDescriptor

var file_shopagolic_email_v1beta1_email_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67,
	0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x08, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x73, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x36, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c,
	0x69, 0x63, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0xa4, 0x02, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x71,
	0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x9b, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c,
	0x69, 0x63, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x6d, 0x65, 0x65, 0x72, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c,
	0x69, 0x63, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopagolic_email_v1beta1_email_service_proto_rawDescOnce sync.Once
	file_shopagolic_email_v1beta1_email_service_proto_rawDescData = file_shopagolic_email_v1beta1_email_service_proto_rawDesc
)

func file_shopagolic_email_v1beta1_email_service_proto_rawDescGZIP() []byte {
	file_shopagolic_email_v1beta1_email_service_proto_rawDescOnce.Do(func() {
		file_shopagolic_email_v1beta1_email_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopagolic_email_v1beta1_email_service_proto_rawDescData)
	})
	return file_shopagolic_email_v1beta1_email_service_proto_rawDescData
}

var file_shopagolic_email_v1beta1_email_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shopagolic_email_v1beta1_email_service_proto_goTypes = []interface{}{
	(*Address)(nil),                       // 0: shopagolic.email.v1beta1.Address
	(*CartItem)(nil),                      // 1: shopagolic.email.v1beta1.CartItem
	(*OrderItem)(nil),                     // 2: shopagolic.email.v1beta1.OrderItem
	(*OrderResult)(nil),                   // 3: shopagolic.email.v1beta1.OrderResult
	(*SendOrderConfirmationRequest)(nil),  // 4: shopagolic.email.v1beta1.SendOrderConfirmationRequest
	(*SendOrderConfirmationResponse)(nil), // 5: shopagolic.email.v1beta1.SendOrderConfirmationResponse
	(*v1.Money)(nil),                      // 6: shopagolic.money.v1.Money
}
var file_shopagolic_email_v1beta1_email_service_proto_depIdxs = []int32{
	1, // 0: shopagolic.email.v1beta1.OrderItem.item:type_name -> shopagolic.email.v1beta1.CartItem
	6, // 1: shopagolic.email.v1beta1.OrderItem.cost:type_name -> shopagolic.money.v1.Money
	6, // 2: shopagolic.email.v1beta1.OrderResult.shipping_cost:type_name -> shopagolic.money.v1.Money
	0, // 3: shopagolic.email.v1beta1.OrderResult.shipping_address:type_name -> shopagolic.email.v1beta1.Address
	2, // 4: shopagolic.email.v1beta1.OrderResult.items:type_name -> shopagolic.email.v1beta1.OrderItem
	3, // 5: shopagolic.email.v1beta1.SendOrderConfirmationRequest.order:type_name -> shopagolic.email.v1beta1.OrderResult
	4, // 6: shopagolic.email.v1beta1.EmailService.SendOrderConfirmation:input_type -> shopagolic.email.v1beta1.SendOrderConfirmationRequest
	5, // 7: shopagolic.email.v1beta1.EmailService.SendOrderConfirmation:output_type -> shopagolic.email.v1beta1.SendOrderConfirmationResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shopagolic_email_v1beta1_email_service_proto_init() }
func file_shopagolic_email_v1beta1_email_service_proto_init() {
	if File_shopagolic_email_v1beta1_email_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderConfirmationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopagolic_email_v1beta1_email_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderConfirmationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shopagolic_email_v1beta1_email_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopagolic_email_v1beta1_email_service_proto_goTypes,
		DependencyIndexes: file_shopagolic_email_v1beta1_email_service_proto_depIdxs,
		MessageInfos:      file_shopagolic_email_v1beta1_email_service_proto_msgTypes,
	}.Build()
	File_shopagolic_email_v1beta1_email_service_proto = out.File
	file_shopagolic_email_v1beta1_email_service_proto_rawDesc = nil
	file_shopagolic_email_v1beta1_email_service_proto_goTypes = nil
	file_shopagolic_email_v1beta1_email_service_proto_depIdxs = nil
}
