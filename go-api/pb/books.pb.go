// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: books.proto

package booklight_v4

import (
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

type BooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *BooksRequest) Reset() {
	*x = BooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksRequest) ProtoMessage() {}

func (x *BooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksRequest.ProtoReflect.Descriptor instead.
func (*BooksRequest) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

func (x *BooksRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type BooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksResponse) Reset() {
	*x = BooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponse) ProtoMessage() {}

func (x *BooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponse.ProtoReflect.Descriptor instead.
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *BooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url      string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	ImageUrl string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Rating   float32  `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Reviews  int32    `protobuf:"varint,5,opt,name=reviews,proto3" json:"reviews,omitempty"`
	Authors  []string `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty"`
	Score    float32  `protobuf:"fixed32,7,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Book) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Book) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Book) GetReviews() int32 {
	if x != nil {
		return x.Reviews
	}
	return 0
}

func (x *Book) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Book) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x36,
	0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x43, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12,
	0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x67, 0x64, 0x61, 0x6e,
	0x59, 0x61, 0x72, 0x6f, 0x74, 0x73, 0x6b, 0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2d, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_books_proto_goTypes = []interface{}{
	(*BooksRequest)(nil),  // 0: booklight.BooksRequest
	(*BooksResponse)(nil), // 1: booklight.BooksResponse
	(*Book)(nil),          // 2: booklight.Book
}
var file_books_proto_depIdxs = []int32{
	2, // 0: booklight.BooksResponse.books:type_name -> booklight.Book
	0, // 1: booklight.Books.Get:input_type -> booklight.BooksRequest
	1, // 2: booklight.Books.Get:output_type -> booklight.BooksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksRequest); i {
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
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksResponse); i {
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
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
