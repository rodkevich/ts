// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/ticket/v1/ticket.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId     string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	NameShort   string                 `protobuf:"bytes,3,opt,name=name_short,json=nameShort,proto3" json:"name_short,omitempty"`
	NameExt     string                 `protobuf:"bytes,4,opt,name=name_ext,json=nameExt,proto3" json:"name_ext,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int32                  `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       float64                `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	Currency    int32                  `protobuf:"varint,8,opt,name=currency,proto3" json:"currency,omitempty"`
	Priority    string                 `protobuf:"bytes,10,opt,name=priority,proto3" json:"priority,omitempty"`
	Published   bool                   `protobuf:"varint,11,opt,name=published,proto3" json:"published,omitempty"`
	Active      bool                   `protobuf:"varint,12,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Deleted     bool                   `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Ticket) GetNameShort() string {
	if x != nil {
		return x.NameShort
	}
	return ""
}

func (x *Ticket) GetNameExt() string {
	if x != nil {
		return x.NameExt
	}
	return ""
}

func (x *Ticket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Ticket) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Ticket) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ticket) GetCurrency() int32 {
	if x != nil {
		return x.Currency
	}
	return 0
}

func (x *Ticket) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *Ticket) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

func (x *Ticket) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Ticket) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId     string  `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	NameShort   string  `protobuf:"bytes,2,opt,name=name_short,json=nameShort,proto3" json:"name_short,omitempty"`
	NameExt     string  `protobuf:"bytes,3,opt,name=name_ext,json=nameExt,proto3" json:"name_ext,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int32   `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Currency    int32   `protobuf:"varint,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Priority    string  `protobuf:"bytes,8,opt,name=priority,proto3" json:"priority,omitempty"`
	Published   bool    `protobuf:"varint,9,opt,name=published,proto3" json:"published,omitempty"`
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateTicketRequest) GetNameShort() string {
	if x != nil {
		return x.NameShort
	}
	return ""
}

func (x *CreateTicketRequest) GetNameExt() string {
	if x != nil {
		return x.NameExt
	}
	return ""
}

func (x *CreateTicketRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTicketRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateTicketRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateTicketRequest) GetCurrency() int32 {
	if x != nil {
		return x.Currency
	}
	return 0
}

func (x *CreateTicketRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *CreateTicketRequest) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

type CreateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreateTicketResponse) Reset() {
	*x = CreateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketResponse) ProtoMessage() {}

func (x *CreateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketResponse.ProtoReflect.Descriptor instead.
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type GetTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketRequest) Reset() {
	*x = GetTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketRequest) ProtoMessage() {}

func (x *GetTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketRequest.ProtoReflect.Descriptor instead.
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *GetTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTicketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reverse   bool     `protobuf:"varint,2,opt,name=reverse,proto3" json:"reverse,omitempty"`
	Search    bool     `protobuf:"varint,3,opt,name=search,proto3" json:"search,omitempty"`
	Paging    bool     `protobuf:"varint,4,opt,name=paging,proto3" json:"paging,omitempty"`
	PageSize  uint32   `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string   `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Fields    []string `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *ListTicketsRequest) Reset() {
	*x = ListTicketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketsRequest) ProtoMessage() {}

func (x *ListTicketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketsRequest.ProtoReflect.Descriptor instead.
func (*ListTicketsRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ListTicketsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListTicketsRequest) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *ListTicketsRequest) GetSearch() bool {
	if x != nil {
		return x.Search
	}
	return false
}

func (x *ListTicketsRequest) GetPaging() bool {
	if x != nil {
		return x.Paging
	}
	return false
}

func (x *ListTicketsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTicketsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTicketsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ListTicketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount    uint64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	HasMore       bool      `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Tickets       []*Ticket `protobuf:"bytes,3,rep,name=tickets,proto3" json:"tickets,omitempty"`
	NextPageToken string    `protobuf:"bytes,4,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTicketsResponse) Reset() {
	*x = ListTicketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketsResponse) ProtoMessage() {}

func (x *ListTicketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketsResponse.ProtoReflect.Descriptor instead.
func (*ListTicketsResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *ListTicketsResponse) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListTicketsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListTicketsResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

func (x *ListTicketsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type UpdateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *UpdateTicketRequest) Reset() {
	*x = UpdateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketRequest) ProtoMessage() {}

func (x *UpdateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketRequest.ProtoReflect.Descriptor instead.
func (*UpdateTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTicketRequest) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hard bool   `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteTicketRequest) GetHard() bool {
	if x != nil {
		return x.Hard
	}
	return false
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_v1_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_v1_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_v1_ticket_proto_rawDescGZIP(), []int{8}
}

var File_proto_ticket_v1_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_v1_ticket_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x03, 0x0a, 0x06, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x90, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x61,
	0x6d, 0x65, 0x45, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x22,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x40, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x61, 0x72, 0x64, 0x22, 0x16, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x99, 0x03, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_v1_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_v1_ticket_proto_rawDescData = file_proto_ticket_v1_ticket_proto_rawDesc
)

func file_proto_ticket_v1_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_v1_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_v1_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_v1_ticket_proto_rawDescData)
	})
	return file_proto_ticket_v1_ticket_proto_rawDescData
}

var file_proto_ticket_v1_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_ticket_v1_ticket_proto_goTypes = []interface{}{
	(*Ticket)(nil),                // 0: ticket.v1.Ticket
	(*CreateTicketRequest)(nil),   // 1: ticket.v1.CreateTicketRequest
	(*CreateTicketResponse)(nil),  // 2: ticket.v1.CreateTicketResponse
	(*GetTicketRequest)(nil),      // 3: ticket.v1.GetTicketRequest
	(*ListTicketsRequest)(nil),    // 4: ticket.v1.ListTicketsRequest
	(*ListTicketsResponse)(nil),   // 5: ticket.v1.ListTicketsResponse
	(*UpdateTicketRequest)(nil),   // 6: ticket.v1.UpdateTicketRequest
	(*DeleteTicketRequest)(nil),   // 7: ticket.v1.DeleteTicketRequest
	(*DeleteTicketResponse)(nil),  // 8: ticket.v1.DeleteTicketResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_proto_ticket_v1_ticket_proto_depIdxs = []int32{
	9,  // 0: ticket.v1.Ticket.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: ticket.v1.Ticket.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: ticket.v1.CreateTicketResponse.ticket:type_name -> ticket.v1.Ticket
	0,  // 3: ticket.v1.ListTicketsResponse.tickets:type_name -> ticket.v1.Ticket
	0,  // 4: ticket.v1.UpdateTicketRequest.ticket:type_name -> ticket.v1.Ticket
	1,  // 5: ticket.v1.TicketService.CreateTicket:input_type -> ticket.v1.CreateTicketRequest
	3,  // 6: ticket.v1.TicketService.GetTicket:input_type -> ticket.v1.GetTicketRequest
	4,  // 7: ticket.v1.TicketService.ListTickets:input_type -> ticket.v1.ListTicketsRequest
	6,  // 8: ticket.v1.TicketService.UpdateTicket:input_type -> ticket.v1.UpdateTicketRequest
	7,  // 9: ticket.v1.TicketService.DeleteTicket:input_type -> ticket.v1.DeleteTicketRequest
	2,  // 10: ticket.v1.TicketService.CreateTicket:output_type -> ticket.v1.CreateTicketResponse
	5,  // 11: ticket.v1.TicketService.GetTicket:output_type -> ticket.v1.ListTicketsResponse
	5,  // 12: ticket.v1.TicketService.ListTickets:output_type -> ticket.v1.ListTicketsResponse
	5,  // 13: ticket.v1.TicketService.UpdateTicket:output_type -> ticket.v1.ListTicketsResponse
	8,  // 14: ticket.v1.TicketService.DeleteTicket:output_type -> ticket.v1.DeleteTicketResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_ticket_v1_ticket_proto_init() }
func file_proto_ticket_v1_ticket_proto_init() {
	if File_proto_ticket_v1_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_v1_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketRequest); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketResponse); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketRequest); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketsRequest); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketsResponse); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketRequest); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketRequest); i {
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
		file_proto_ticket_v1_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketResponse); i {
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
			RawDescriptor: file_proto_ticket_v1_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_v1_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_v1_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_v1_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_v1_ticket_proto = out.File
	file_proto_ticket_v1_ticket_proto_rawDesc = nil
	file_proto_ticket_v1_ticket_proto_goTypes = nil
	file_proto_ticket_v1_ticket_proto_depIdxs = nil
}
