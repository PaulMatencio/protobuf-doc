// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: src/document/document.proto

package documentpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId      string `protobuf:"bytes,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	PageNumber  int32  `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	Size        int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`        // page size
	Metadata    string `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"` // page metadata
	ContentType string `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Object      []byte `protobuf:"bytes,6,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_document_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_src_document_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_src_document_document_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *Page) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Page) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Page) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

type Pdf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf      []byte `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Pdf) Reset() {
	*x = Pdf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_document_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pdf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pdf) ProtoMessage() {}

func (x *Pdf) ProtoReflect() protoreflect.Message {
	mi := &file_src_document_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pdf.ProtoReflect.Descriptor instead.
func (*Pdf) Descriptor() ([]byte, []int) {
	return file_src_document_document_proto_rawDescGZIP(), []int{1}
}

func (x *Pdf) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

func (x *Pdf) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocId         string                 `protobuf:"bytes,1,opt,name=docId,proto3" json:"docId,omitempty"`
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`             // document size
	PageNumber    int32                  `protobuf:"varint,3,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"` //  must be 0
	NumberOfPages int32                  `protobuf:"varint,4,opt,name=numberOfPages,proto3" json:"numberOfPages,omitempty"`
	S3Meta        string                 `protobuf:"bytes,5,opt,name=s3Meta,proto3" json:"s3Meta,omitempty"`     //  s3  document user  metadata  ( stored in S3)
	Metadata      string                 `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"` // document user metadata ( stored with  the object)
	Object        []byte                 `protobuf:"bytes,7,opt,name=object,proto3" json:"object,omitempty"`     //  page 0
	Pdf           *Pdf                   `protobuf:"bytes,8,opt,name=pdf,proto3" json:"pdf,omitempty"`
	Clip          bool                   `protobuf:"varint,9,opt,name=clip,proto3" json:"clip,omitempty"` // this document has fpClipping ( page 0)
	Page          []*Page                `protobuf:"bytes,10,rep,name=page,proto3" json:"page,omitempty"`
	LastUpdated   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_document_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_src_document_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_src_document_document_proto_rawDescGZIP(), []int{2}
}

func (x *Document) GetDocId() string {
	if x != nil {
		return x.DocId
	}
	return ""
}

func (x *Document) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Document) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Document) GetNumberOfPages() int32 {
	if x != nil {
		return x.NumberOfPages
	}
	return 0
}

func (x *Document) GetS3Meta() string {
	if x != nil {
		return x.S3Meta
	}
	return ""
}

func (x *Document) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Document) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *Document) GetPdf() *Pdf {
	if x != nil {
		return x.Pdf
	}
	return nil
}

func (x *Document) GetClip() bool {
	if x != nil {
		return x.Clip
	}
	return false
}

func (x *Document) GetPage() []*Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *Document) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

var File_src_document_document_proto protoreflect.FileDescriptor

var file_src_document_document_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x72, 0x63, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x33, 0x0a, 0x03, 0x50, 0x64, 0x66, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x64, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x64, 0x66, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe2, 0x02, 0x0a, 0x08, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x33, 0x4d, 0x65, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x33, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x70, 0x64, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x64,
	0x66, 0x52, 0x03, 0x70, 0x64, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x63, 0x6c, 0x69, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x19, 0x5a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_src_document_document_proto_rawDescOnce sync.Once
	file_src_document_document_proto_rawDescData = file_src_document_document_proto_rawDesc
)

func file_src_document_document_proto_rawDescGZIP() []byte {
	file_src_document_document_proto_rawDescOnce.Do(func() {
		file_src_document_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_document_document_proto_rawDescData)
	})
	return file_src_document_document_proto_rawDescData
}

var file_src_document_document_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_document_document_proto_goTypes = []interface{}{
	(*Page)(nil),                  // 0: documentpb.Page
	(*Pdf)(nil),                   // 1: documentpb.Pdf
	(*Document)(nil),              // 2: documentpb.Document
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_src_document_document_proto_depIdxs = []int32{
	1, // 0: documentpb.Document.pdf:type_name -> documentpb.Pdf
	0, // 1: documentpb.Document.page:type_name -> documentpb.Page
	3, // 2: documentpb.Document.last_updated:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_document_document_proto_init() }
func file_src_document_document_proto_init() {
	if File_src_document_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_document_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_src_document_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pdf); i {
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
		file_src_document_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
			RawDescriptor: file_src_document_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_document_document_proto_goTypes,
		DependencyIndexes: file_src_document_document_proto_depIdxs,
		MessageInfos:      file_src_document_document_proto_msgTypes,
	}.Build()
	File_src_document_document_proto = out.File
	file_src_document_document_proto_rawDesc = nil
	file_src_document_document_proto_goTypes = nil
	file_src_document_document_proto_depIdxs = nil
}
