// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/product/v1/product.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProductBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductBody) Reset() {
	*x = ProductBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBody) ProtoMessage() {}

func (x *ProductBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBody.ProtoReflect.Descriptor instead.
func (*ProductBody) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductBody) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductBody) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateProductBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Price       *int64  `protobuf:"varint,4,opt,name=price,proto3,oneof" json:"price,omitempty"`
}

func (x *UpdateProductBody) Reset() {
	*x = UpdateProductBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductBody) ProtoMessage() {}

func (x *UpdateProductBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductBody.ProtoReflect.Descriptor instead.
func (*UpdateProductBody) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProductBody) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateProductBody) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateProductBody) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *ProductBody `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductRequest) GetBody() *ProductBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type CreateProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateProductReply) Reset() {
	*x = CreateProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReply) ProtoMessage() {}

func (x *CreateProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReply.ProtoReflect.Descriptor instead.
func (*CreateProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProductReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReply) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *UpdateProductBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetBody() *UpdateProductBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type UpdateProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateProductReply) Reset() {
	*x = UpdateProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductReply) ProtoMessage() {}

func (x *UpdateProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductReply.ProtoReflect.Descriptor instead.
func (*UpdateProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProductReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductReply) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductReply) Reset() {
	*x = DeleteProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReply) ProtoMessage() {}

func (x *DeleteProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReply.ProtoReflect.Descriptor instead.
func (*DeleteProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProductReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{8}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetProductReply) Reset() {
	*x = GetProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReply) ProtoMessage() {}

func (x *GetProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReply.ProtoReflect.Descriptor instead.
func (*GetProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{9}
}

func (x *GetProductReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProductReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetProductReply) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ListProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListProductRequest) Reset() {
	*x = ListProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductRequest) ProtoMessage() {}

func (x *ListProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductRequest.ProtoReflect.Descriptor instead.
func (*ListProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{10}
}

func (x *ListProductRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type ListProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*CreateProductReply `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProductReply) Reset() {
	*x = ListProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductReply) ProtoMessage() {}

func (x *ListProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductReply.ProtoReflect.Descriptor instead.
func (*ListProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{11}
}

func (x *ListProductReply) GetProducts() []*CreateProductReply {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_api_product_v1_product_proto protoreflect.FileDescriptor

var file_api_product_v1_product_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x70, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x70, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xbb, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x72, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x09, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x77, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x1a, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x71, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x48, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x6e, 0x61, 0x69, 0x6e, 0x67, 0x74,
	0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_product_v1_product_proto_rawDescOnce sync.Once
	file_api_product_v1_product_proto_rawDescData = file_api_product_v1_product_proto_rawDesc
)

func file_api_product_v1_product_proto_rawDescGZIP() []byte {
	file_api_product_v1_product_proto_rawDescOnce.Do(func() {
		file_api_product_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_product_v1_product_proto_rawDescData)
	})
	return file_api_product_v1_product_proto_rawDescData
}

var file_api_product_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_product_v1_product_proto_goTypes = []interface{}{
	(*ProductBody)(nil),          // 0: api.product.v1.ProductBody
	(*UpdateProductBody)(nil),    // 1: api.product.v1.UpdateProductBody
	(*CreateProductRequest)(nil), // 2: api.product.v1.CreateProductRequest
	(*CreateProductReply)(nil),   // 3: api.product.v1.CreateProductReply
	(*UpdateProductRequest)(nil), // 4: api.product.v1.UpdateProductRequest
	(*UpdateProductReply)(nil),   // 5: api.product.v1.UpdateProductReply
	(*DeleteProductRequest)(nil), // 6: api.product.v1.DeleteProductRequest
	(*DeleteProductReply)(nil),   // 7: api.product.v1.DeleteProductReply
	(*GetProductRequest)(nil),    // 8: api.product.v1.GetProductRequest
	(*GetProductReply)(nil),      // 9: api.product.v1.GetProductReply
	(*ListProductRequest)(nil),   // 10: api.product.v1.ListProductRequest
	(*ListProductReply)(nil),     // 11: api.product.v1.ListProductReply
}
var file_api_product_v1_product_proto_depIdxs = []int32{
	0,  // 0: api.product.v1.CreateProductRequest.body:type_name -> api.product.v1.ProductBody
	1,  // 1: api.product.v1.UpdateProductRequest.body:type_name -> api.product.v1.UpdateProductBody
	3,  // 2: api.product.v1.ListProductReply.products:type_name -> api.product.v1.CreateProductReply
	2,  // 3: api.product.v1.Product.CreateProduct:input_type -> api.product.v1.CreateProductRequest
	4,  // 4: api.product.v1.Product.UpdateProduct:input_type -> api.product.v1.UpdateProductRequest
	6,  // 5: api.product.v1.Product.DeleteProduct:input_type -> api.product.v1.DeleteProductRequest
	8,  // 6: api.product.v1.Product.GetProduct:input_type -> api.product.v1.GetProductRequest
	10, // 7: api.product.v1.Product.ListProduct:input_type -> api.product.v1.ListProductRequest
	3,  // 8: api.product.v1.Product.CreateProduct:output_type -> api.product.v1.CreateProductReply
	5,  // 9: api.product.v1.Product.UpdateProduct:output_type -> api.product.v1.UpdateProductReply
	7,  // 10: api.product.v1.Product.DeleteProduct:output_type -> api.product.v1.DeleteProductReply
	9,  // 11: api.product.v1.Product.GetProduct:output_type -> api.product.v1.GetProductReply
	11, // 12: api.product.v1.Product.ListProduct:output_type -> api.product.v1.ListProductReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_product_v1_product_proto_init() }
func file_api_product_v1_product_proto_init() {
	if File_api_product_v1_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_product_v1_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBody); i {
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
		file_api_product_v1_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductBody); i {
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
		file_api_product_v1_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductReply); i {
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
		file_api_product_v1_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductReply); i {
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
		file_api_product_v1_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductReply); i {
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
		file_api_product_v1_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReply); i {
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
		file_api_product_v1_product_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductReply); i {
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
	file_api_product_v1_product_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_product_v1_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_product_v1_product_proto_goTypes,
		DependencyIndexes: file_api_product_v1_product_proto_depIdxs,
		MessageInfos:      file_api_product_v1_product_proto_msgTypes,
	}.Build()
	File_api_product_v1_product_proto = out.File
	file_api_product_v1_product_proto_rawDesc = nil
	file_api_product_v1_product_proto_goTypes = nil
	file_api_product_v1_product_proto_depIdxs = nil
}
