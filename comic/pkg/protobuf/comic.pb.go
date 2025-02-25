// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.2
// source: comic.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Comic struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug          string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Author        string                 `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Artist        string                 `protobuf:"bytes,6,opt,name=artist,proto3" json:"artist,omitempty"`
	Status        string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Type          string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Genres        []string               `protobuf:"bytes,9,rep,name=genres,proto3" json:"genres,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Comic) Reset() {
	*x = Comic{}
	mi := &file_comic_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comic) ProtoMessage() {}

func (x *Comic) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comic.ProtoReflect.Descriptor instead.
func (*Comic) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{0}
}

func (x *Comic) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Comic) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Comic) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Comic) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Comic) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *Comic) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Comic) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Comic) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

type InsertComicRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Artist        string                 `protobuf:"bytes,4,opt,name=artist,proto3" json:"artist,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Type          string                 `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Genres        []string               `protobuf:"bytes,7,rep,name=genres,proto3" json:"genres,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InsertComicRequest) Reset() {
	*x = InsertComicRequest{}
	mi := &file_comic_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertComicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertComicRequest) ProtoMessage() {}

func (x *InsertComicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertComicRequest.ProtoReflect.Descriptor instead.
func (*InsertComicRequest) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{1}
}

func (x *InsertComicRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InsertComicRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InsertComicRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *InsertComicRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *InsertComicRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InsertComicRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InsertComicRequest) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

type InsertComicResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comic         *Comic                 `protobuf:"bytes,1,opt,name=comic,proto3" json:"comic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InsertComicResponse) Reset() {
	*x = InsertComicResponse{}
	mi := &file_comic_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertComicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertComicResponse) ProtoMessage() {}

func (x *InsertComicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertComicResponse.ProtoReflect.Descriptor instead.
func (*InsertComicResponse) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{2}
}

func (x *InsertComicResponse) GetComic() *Comic {
	if x != nil {
		return x.Comic
	}
	return nil
}

type GetComicBySlugRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slug          string                 `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetComicBySlugRequest) Reset() {
	*x = GetComicBySlugRequest{}
	mi := &file_comic_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetComicBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetComicBySlugRequest) ProtoMessage() {}

func (x *GetComicBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetComicBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetComicBySlugRequest) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{3}
}

func (x *GetComicBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetComicBySlugResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comic         *Comic                 `protobuf:"bytes,1,opt,name=comic,proto3" json:"comic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetComicBySlugResponse) Reset() {
	*x = GetComicBySlugResponse{}
	mi := &file_comic_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetComicBySlugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetComicBySlugResponse) ProtoMessage() {}

func (x *GetComicBySlugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetComicBySlugResponse.ProtoReflect.Descriptor instead.
func (*GetComicBySlugResponse) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{4}
}

func (x *GetComicBySlugResponse) GetComic() *Comic {
	if x != nil {
		return x.Comic
	}
	return nil
}

type Cover struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ComicId       int64                  `protobuf:"varint,2,opt,name=comic_id,json=comicId,proto3" json:"comic_id,omitempty"`
	FileKey       string                 `protobuf:"bytes,3,opt,name=file_key,json=fileKey,proto3" json:"file_key,omitempty"`
	IsCurrent     bool                   `protobuf:"varint,4,opt,name=is_current,json=isCurrent,proto3" json:"is_current,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cover) Reset() {
	*x = Cover{}
	mi := &file_comic_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cover) ProtoMessage() {}

func (x *Cover) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cover.ProtoReflect.Descriptor instead.
func (*Cover) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{5}
}

func (x *Cover) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cover) GetComicId() int64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *Cover) GetFileKey() string {
	if x != nil {
		return x.FileKey
	}
	return ""
}

func (x *Cover) GetIsCurrent() bool {
	if x != nil {
		return x.IsCurrent
	}
	return false
}

type UploadCoverRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*UploadCoverRequest_Chunk
	//	*UploadCoverRequest_Metadata
	Data          isUploadCoverRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadCoverRequest) Reset() {
	*x = UploadCoverRequest{}
	mi := &file_comic_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadCoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCoverRequest) ProtoMessage() {}

func (x *UploadCoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCoverRequest.ProtoReflect.Descriptor instead.
func (*UploadCoverRequest) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{6}
}

func (x *UploadCoverRequest) GetData() isUploadCoverRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadCoverRequest) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*UploadCoverRequest_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

func (x *UploadCoverRequest) GetMetadata() *CoverMetadata {
	if x != nil {
		if x, ok := x.Data.(*UploadCoverRequest_Metadata); ok {
			return x.Metadata
		}
	}
	return nil
}

type isUploadCoverRequest_Data interface {
	isUploadCoverRequest_Data()
}

type UploadCoverRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3,oneof"`
}

type UploadCoverRequest_Metadata struct {
	Metadata *CoverMetadata `protobuf:"bytes,2,opt,name=metadata,proto3,oneof"`
}

func (*UploadCoverRequest_Chunk) isUploadCoverRequest_Data() {}

func (*UploadCoverRequest_Metadata) isUploadCoverRequest_Data() {}

type CoverMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ComicId       int64                  `protobuf:"varint,1,opt,name=comic_id,json=comicId,proto3" json:"comic_id,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoverMetadata) Reset() {
	*x = CoverMetadata{}
	mi := &file_comic_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoverMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoverMetadata) ProtoMessage() {}

func (x *CoverMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoverMetadata.ProtoReflect.Descriptor instead.
func (*CoverMetadata) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{7}
}

func (x *CoverMetadata) GetComicId() int64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *CoverMetadata) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadCoverResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cover         *Cover                 `protobuf:"bytes,1,opt,name=cover,proto3" json:"cover,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadCoverResponse) Reset() {
	*x = UploadCoverResponse{}
	mi := &file_comic_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadCoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCoverResponse) ProtoMessage() {}

func (x *UploadCoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comic_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCoverResponse.ProtoReflect.Descriptor instead.
func (*UploadCoverResponse) Descriptor() ([]byte, []int) {
	return file_comic_proto_rawDescGZIP(), []int{8}
}

func (x *UploadCoverResponse) GetCover() *Cover {
	if x != nil {
		return x.Cover
	}
	return nil
}

var File_comic_proto protoreflect.FileDescriptor

var file_comic_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x6d, 0x69, 0x63, 0x22, 0xd7, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0xc0,
	0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x22, 0x39, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x22, 0x2b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63,
	0x52, 0x05, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x22, 0x6c, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x46, 0x0a, 0x0d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x32, 0xa3, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d,
	0x69, 0x63, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1c, 0x2e, 0x63, 0x6f,
	0x6d, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x79, 0x53, 0x6c,
	0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x69,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x56, 0x0a, 0x0c, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x69, 0x6c, 0x69, 0x73, 0x63, 0x69, 0x74, 0x65, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x63,
	0x6f, 0x6d, 0x69, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_comic_proto_rawDescOnce sync.Once
	file_comic_proto_rawDescData []byte
)

func file_comic_proto_rawDescGZIP() []byte {
	file_comic_proto_rawDescOnce.Do(func() {
		file_comic_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_comic_proto_rawDesc), len(file_comic_proto_rawDesc)))
	})
	return file_comic_proto_rawDescData
}

var file_comic_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_comic_proto_goTypes = []any{
	(*Comic)(nil),                  // 0: comic.Comic
	(*InsertComicRequest)(nil),     // 1: comic.InsertComicRequest
	(*InsertComicResponse)(nil),    // 2: comic.InsertComicResponse
	(*GetComicBySlugRequest)(nil),  // 3: comic.GetComicBySlugRequest
	(*GetComicBySlugResponse)(nil), // 4: comic.GetComicBySlugResponse
	(*Cover)(nil),                  // 5: comic.Cover
	(*UploadCoverRequest)(nil),     // 6: comic.UploadCoverRequest
	(*CoverMetadata)(nil),          // 7: comic.CoverMetadata
	(*UploadCoverResponse)(nil),    // 8: comic.UploadCoverResponse
}
var file_comic_proto_depIdxs = []int32{
	0, // 0: comic.InsertComicResponse.comic:type_name -> comic.Comic
	0, // 1: comic.GetComicBySlugResponse.comic:type_name -> comic.Comic
	7, // 2: comic.UploadCoverRequest.metadata:type_name -> comic.CoverMetadata
	5, // 3: comic.UploadCoverResponse.cover:type_name -> comic.Cover
	1, // 4: comic.ComicService.InsertComic:input_type -> comic.InsertComicRequest
	3, // 5: comic.ComicService.GetComicBySlug:input_type -> comic.GetComicBySlugRequest
	6, // 6: comic.CoverService.UploadCover:input_type -> comic.UploadCoverRequest
	2, // 7: comic.ComicService.InsertComic:output_type -> comic.InsertComicResponse
	4, // 8: comic.ComicService.GetComicBySlug:output_type -> comic.GetComicBySlugResponse
	8, // 9: comic.CoverService.UploadCover:output_type -> comic.UploadCoverResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_comic_proto_init() }
func file_comic_proto_init() {
	if File_comic_proto != nil {
		return
	}
	file_comic_proto_msgTypes[6].OneofWrappers = []any{
		(*UploadCoverRequest_Chunk)(nil),
		(*UploadCoverRequest_Metadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_comic_proto_rawDesc), len(file_comic_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_comic_proto_goTypes,
		DependencyIndexes: file_comic_proto_depIdxs,
		MessageInfos:      file_comic_proto_msgTypes,
	}.Build()
	File_comic_proto = out.File
	file_comic_proto_goTypes = nil
	file_comic_proto_depIdxs = nil
}
