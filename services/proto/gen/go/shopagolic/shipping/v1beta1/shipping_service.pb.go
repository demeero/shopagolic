// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shopagolic/shipping/v1beta1/shipping_service.proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items   []*Item  `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetQuoteRequest) Reset() {
	*x = GetQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteRequest) ProtoMessage() {}

func (x *GetQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuoteRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetQuoteRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cost *v1.Money `protobuf:"bytes,1,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *GetQuoteResponse) Reset() {
	*x = GetQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteResponse) ProtoMessage() {}

func (x *GetQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteResponse.ProtoReflect.Descriptor instead.
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetQuoteResponse) GetCost() *v1.Money {
	if x != nil {
		return x.Cost
	}
	return nil
}

type ShipOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items   []*Item  `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ShipOrderRequest) Reset() {
	*x = ShipOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderRequest) ProtoMessage() {}

func (x *ShipOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderRequest.ProtoReflect.Descriptor instead.
func (*ShipOrderRequest) Descriptor() ([]byte, []int) {
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{3}
}

func (x *ShipOrderRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ShipOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ShipOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingId string `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
}

func (x *ShipOrderResponse) Reset() {
	*x = ShipOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderResponse) ProtoMessage() {}

func (x *ShipOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderResponse.ProtoReflect.Descriptor instead.
func (*ShipOrderResponse) Descriptor() ([]byte, []int) {
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{4}
}

func (x *ShipOrderResponse) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

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
		mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[5]
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
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP(), []int{5}
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

var File_shopagolic_shipping_v1beta1_shipping_service_proto protoreflect.FileDescriptor

var file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61,
	0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69,
	0x63, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xea, 0x01, 0x0a,
	0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x69, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x09, 0x53,
	0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61,
	0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67,
	0x6f, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x65, 0x65, 0x72, 0x6f, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x61, 0x67, 0x6f, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescOnce sync.Once
	file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescData = file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDesc
)

func file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescGZIP() []byte {
	file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescOnce.Do(func() {
		file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescData)
	})
	return file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDescData
}

var file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shopagolic_shipping_v1beta1_shipping_service_proto_goTypes = []interface{}{
	(*Item)(nil),              // 0: shopagolic.shipping.v1beta1.Item
	(*GetQuoteRequest)(nil),   // 1: shopagolic.shipping.v1beta1.GetQuoteRequest
	(*GetQuoteResponse)(nil),  // 2: shopagolic.shipping.v1beta1.GetQuoteResponse
	(*ShipOrderRequest)(nil),  // 3: shopagolic.shipping.v1beta1.ShipOrderRequest
	(*ShipOrderResponse)(nil), // 4: shopagolic.shipping.v1beta1.ShipOrderResponse
	(*Address)(nil),           // 5: shopagolic.shipping.v1beta1.Address
	(*v1.Money)(nil),          // 6: shopagolic.money.v1.Money
}
var file_shopagolic_shipping_v1beta1_shipping_service_proto_depIdxs = []int32{
	5, // 0: shopagolic.shipping.v1beta1.GetQuoteRequest.address:type_name -> shopagolic.shipping.v1beta1.Address
	0, // 1: shopagolic.shipping.v1beta1.GetQuoteRequest.items:type_name -> shopagolic.shipping.v1beta1.Item
	6, // 2: shopagolic.shipping.v1beta1.GetQuoteResponse.cost:type_name -> shopagolic.money.v1.Money
	5, // 3: shopagolic.shipping.v1beta1.ShipOrderRequest.address:type_name -> shopagolic.shipping.v1beta1.Address
	0, // 4: shopagolic.shipping.v1beta1.ShipOrderRequest.items:type_name -> shopagolic.shipping.v1beta1.Item
	1, // 5: shopagolic.shipping.v1beta1.ShippingService.GetQuote:input_type -> shopagolic.shipping.v1beta1.GetQuoteRequest
	3, // 6: shopagolic.shipping.v1beta1.ShippingService.ShipOrder:input_type -> shopagolic.shipping.v1beta1.ShipOrderRequest
	2, // 7: shopagolic.shipping.v1beta1.ShippingService.GetQuote:output_type -> shopagolic.shipping.v1beta1.GetQuoteResponse
	4, // 8: shopagolic.shipping.v1beta1.ShippingService.ShipOrder:output_type -> shopagolic.shipping.v1beta1.ShipOrderResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_shopagolic_shipping_v1beta1_shipping_service_proto_init() }
func file_shopagolic_shipping_v1beta1_shipping_service_proto_init() {
	if File_shopagolic_shipping_v1beta1_shipping_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteRequest); i {
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
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteResponse); i {
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
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderRequest); i {
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
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderResponse); i {
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
		file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopagolic_shipping_v1beta1_shipping_service_proto_goTypes,
		DependencyIndexes: file_shopagolic_shipping_v1beta1_shipping_service_proto_depIdxs,
		MessageInfos:      file_shopagolic_shipping_v1beta1_shipping_service_proto_msgTypes,
	}.Build()
	File_shopagolic_shipping_v1beta1_shipping_service_proto = out.File
	file_shopagolic_shipping_v1beta1_shipping_service_proto_rawDesc = nil
	file_shopagolic_shipping_v1beta1_shipping_service_proto_goTypes = nil
	file_shopagolic_shipping_v1beta1_shipping_service_proto_depIdxs = nil
}
