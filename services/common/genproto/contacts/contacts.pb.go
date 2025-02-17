// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: contacts.proto

package contacts

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

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"` // PENDING, ACCEPTED,REJECTED
}

func (x *Contact) Reset() {
	*x = Contact{}
	mi := &file_contacts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Contact) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type AddContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContactUserId int32 `protobuf:"varint,2,opt,name=contact_user_id,json=contactUserId,proto3" json:"contact_user_id,omitempty"`
}

func (x *AddContactRequest) Reset() {
	*x = AddContactRequest{}
	mi := &file_contacts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContactRequest) ProtoMessage() {}

func (x *AddContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContactRequest.ProtoReflect.Descriptor instead.
func (*AddContactRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{1}
}

func (x *AddContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddContactRequest) GetContactUserId() int32 {
	if x != nil {
		return x.ContactUserId
	}
	return 0
}

type AddContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddContactResponse) Reset() {
	*x = AddContactResponse{}
	mi := &file_contacts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContactResponse) ProtoMessage() {}

func (x *AddContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContactResponse.ProtoReflect.Descriptor instead.
func (*AddContactResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{2}
}

func (x *AddContactResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RemoveContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContactUserId int32 `protobuf:"varint,2,opt,name=contact_user_id,json=contactUserId,proto3" json:"contact_user_id,omitempty"`
}

func (x *RemoveContactRequest) Reset() {
	*x = RemoveContactRequest{}
	mi := &file_contacts_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveContactRequest) ProtoMessage() {}

func (x *RemoveContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveContactRequest.ProtoReflect.Descriptor instead.
func (*RemoveContactRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RemoveContactRequest) GetContactUserId() int32 {
	if x != nil {
		return x.ContactUserId
	}
	return 0
}

type RemoveContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemoveContactResponse) Reset() {
	*x = RemoveContactResponse{}
	mi := &file_contacts_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveContactResponse) ProtoMessage() {}

func (x *RemoveContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveContactResponse.ProtoReflect.Descriptor instead.
func (*RemoveContactResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveContactResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AcceptContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContactUserId int32 `protobuf:"varint,2,opt,name=contact_user_id,json=contactUserId,proto3" json:"contact_user_id,omitempty"`
}

func (x *AcceptContactRequest) Reset() {
	*x = AcceptContactRequest{}
	mi := &file_contacts_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptContactRequest) ProtoMessage() {}

func (x *AcceptContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptContactRequest.ProtoReflect.Descriptor instead.
func (*AcceptContactRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{5}
}

func (x *AcceptContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AcceptContactRequest) GetContactUserId() int32 {
	if x != nil {
		return x.ContactUserId
	}
	return 0
}

type AcceptContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AcceptContactResponse) Reset() {
	*x = AcceptContactResponse{}
	mi := &file_contacts_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptContactResponse) ProtoMessage() {}

func (x *AcceptContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptContactResponse.ProtoReflect.Descriptor instead.
func (*AcceptContactResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptContactResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetContactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetContactsRequest) Reset() {
	*x = GetContactsRequest{}
	mi := &file_contacts_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsRequest) ProtoMessage() {}

func (x *GetContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsRequest.ProtoReflect.Descriptor instead.
func (*GetContactsRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{7}
}

func (x *GetContactsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetContactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *GetContactsResponse) Reset() {
	*x = GetContactsResponse{}
	mi := &file_contacts_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsResponse) ProtoMessage() {}

func (x *GetContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsResponse.ProtoReflect.Descriptor instead.
func (*GetContactsResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{8}
}

func (x *GetContactsResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type GetPendingSentContactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPendingSentContactsRequest) Reset() {
	*x = GetPendingSentContactsRequest{}
	mi := &file_contacts_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingSentContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingSentContactsRequest) ProtoMessage() {}

func (x *GetPendingSentContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingSentContactsRequest.ProtoReflect.Descriptor instead.
func (*GetPendingSentContactsRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{9}
}

func (x *GetPendingSentContactsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPendingSentContactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *GetPendingSentContactsResponse) Reset() {
	*x = GetPendingSentContactsResponse{}
	mi := &file_contacts_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingSentContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingSentContactsResponse) ProtoMessage() {}

func (x *GetPendingSentContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingSentContactsResponse.ProtoReflect.Descriptor instead.
func (*GetPendingSentContactsResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{10}
}

func (x *GetPendingSentContactsResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type GetPendingReceivedContactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPendingReceivedContactsRequest) Reset() {
	*x = GetPendingReceivedContactsRequest{}
	mi := &file_contacts_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingReceivedContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingReceivedContactsRequest) ProtoMessage() {}

func (x *GetPendingReceivedContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingReceivedContactsRequest.ProtoReflect.Descriptor instead.
func (*GetPendingReceivedContactsRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{11}
}

func (x *GetPendingReceivedContactsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPendingReceivedContactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *GetPendingReceivedContactsResponse) Reset() {
	*x = GetPendingReceivedContactsResponse{}
	mi := &file_contacts_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPendingReceivedContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingReceivedContactsResponse) ProtoMessage() {}

func (x *GetPendingReceivedContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingReceivedContactsResponse.ProtoReflect.Descriptor instead.
func (*GetPendingReceivedContactsResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{12}
}

func (x *GetPendingReceivedContactsResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type RejectContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContactUserId int32 `protobuf:"varint,2,opt,name=contact_user_id,json=contactUserId,proto3" json:"contact_user_id,omitempty"`
}

func (x *RejectContactRequest) Reset() {
	*x = RejectContactRequest{}
	mi := &file_contacts_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectContactRequest) ProtoMessage() {}

func (x *RejectContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectContactRequest.ProtoReflect.Descriptor instead.
func (*RejectContactRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{13}
}

func (x *RejectContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RejectContactRequest) GetContactUserId() int32 {
	if x != nil {
		return x.ContactUserId
	}
	return 0
}

type RejectContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RejectContactResponse) Reset() {
	*x = RejectContactResponse{}
	mi := &file_contacts_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectContactResponse) ProtoMessage() {}

func (x *RejectContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectContactResponse.ProtoReflect.Descriptor instead.
func (*RejectContactResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{14}
}

func (x *RejectContactResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetContactsNotInGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId int32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetContactsNotInGroupRequest) Reset() {
	*x = GetContactsNotInGroupRequest{}
	mi := &file_contacts_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsNotInGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsNotInGroupRequest) ProtoMessage() {}

func (x *GetContactsNotInGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsNotInGroupRequest.ProtoReflect.Descriptor instead.
func (*GetContactsNotInGroupRequest) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{15}
}

func (x *GetContactsNotInGroupRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetContactsNotInGroupRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetContactsNotInGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *GetContactsNotInGroupResponse) Reset() {
	*x = GetContactsNotInGroupResponse{}
	mi := &file_contacts_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsNotInGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsNotInGroupResponse) ProtoMessage() {}

func (x *GetContactsNotInGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contacts_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsNotInGroupResponse.ProtoReflect.Descriptor instead.
func (*GetContactsNotInGroupResponse) Descriptor() ([]byte, []int) {
	return file_contacts_proto_rawDescGZIP(), []int{16}
}

func (x *GetContactsNotInGroupResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

var File_contacts_proto protoreflect.FileDescriptor

var file_contacts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x54,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x57, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x57, 0x0a, 0x14,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x4a, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x57,
	0x0a, 0x14, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x1d,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x32, 0xdb, 0x04, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x15,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x15,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x73, 0x12, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x65, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x4e,
	0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x75, 0x61, 0x6e, 0x62, 0x69, 0x6e, 0x32, 0x37, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x57,
	0x65, 0x62, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contacts_proto_rawDescOnce sync.Once
	file_contacts_proto_rawDescData = file_contacts_proto_rawDesc
)

func file_contacts_proto_rawDescGZIP() []byte {
	file_contacts_proto_rawDescOnce.Do(func() {
		file_contacts_proto_rawDescData = protoimpl.X.CompressGZIP(file_contacts_proto_rawDescData)
	})
	return file_contacts_proto_rawDescData
}

var file_contacts_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_contacts_proto_goTypes = []any{
	(*Contact)(nil),                            // 0: Contact
	(*AddContactRequest)(nil),                  // 1: AddContactRequest
	(*AddContactResponse)(nil),                 // 2: AddContactResponse
	(*RemoveContactRequest)(nil),               // 3: RemoveContactRequest
	(*RemoveContactResponse)(nil),              // 4: RemoveContactResponse
	(*AcceptContactRequest)(nil),               // 5: AcceptContactRequest
	(*AcceptContactResponse)(nil),              // 6: AcceptContactResponse
	(*GetContactsRequest)(nil),                 // 7: GetContactsRequest
	(*GetContactsResponse)(nil),                // 8: GetContactsResponse
	(*GetPendingSentContactsRequest)(nil),      // 9: GetPendingSentContactsRequest
	(*GetPendingSentContactsResponse)(nil),     // 10: GetPendingSentContactsResponse
	(*GetPendingReceivedContactsRequest)(nil),  // 11: GetPendingReceivedContactsRequest
	(*GetPendingReceivedContactsResponse)(nil), // 12: GetPendingReceivedContactsResponse
	(*RejectContactRequest)(nil),               // 13: RejectContactRequest
	(*RejectContactResponse)(nil),              // 14: RejectContactResponse
	(*GetContactsNotInGroupRequest)(nil),       // 15: GetContactsNotInGroupRequest
	(*GetContactsNotInGroupResponse)(nil),      // 16: GetContactsNotInGroupResponse
}
var file_contacts_proto_depIdxs = []int32{
	0,  // 0: GetContactsResponse.contacts:type_name -> Contact
	0,  // 1: GetPendingSentContactsResponse.contacts:type_name -> Contact
	0,  // 2: GetPendingReceivedContactsResponse.contacts:type_name -> Contact
	0,  // 3: GetContactsNotInGroupResponse.contacts:type_name -> Contact
	1,  // 4: ContactService.AddContact:input_type -> AddContactRequest
	3,  // 5: ContactService.RemoveContact:input_type -> RemoveContactRequest
	5,  // 6: ContactService.AcceptContact:input_type -> AcceptContactRequest
	7,  // 7: ContactService.GetContacts:input_type -> GetContactsRequest
	9,  // 8: ContactService.GetPendingSentContacts:input_type -> GetPendingSentContactsRequest
	11, // 9: ContactService.GetPendingReceivedContacts:input_type -> GetPendingReceivedContactsRequest
	13, // 10: ContactService.RejectContact:input_type -> RejectContactRequest
	15, // 11: ContactService.GetContactsNotInGroup:input_type -> GetContactsNotInGroupRequest
	2,  // 12: ContactService.AddContact:output_type -> AddContactResponse
	4,  // 13: ContactService.RemoveContact:output_type -> RemoveContactResponse
	6,  // 14: ContactService.AcceptContact:output_type -> AcceptContactResponse
	8,  // 15: ContactService.GetContacts:output_type -> GetContactsResponse
	10, // 16: ContactService.GetPendingSentContacts:output_type -> GetPendingSentContactsResponse
	12, // 17: ContactService.GetPendingReceivedContacts:output_type -> GetPendingReceivedContactsResponse
	14, // 18: ContactService.RejectContact:output_type -> RejectContactResponse
	16, // 19: ContactService.GetContactsNotInGroup:output_type -> GetContactsNotInGroupResponse
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_contacts_proto_init() }
func file_contacts_proto_init() {
	if File_contacts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contacts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contacts_proto_goTypes,
		DependencyIndexes: file_contacts_proto_depIdxs,
		MessageInfos:      file_contacts_proto_msgTypes,
	}.Build()
	File_contacts_proto = out.File
	file_contacts_proto_rawDesc = nil
	file_contacts_proto_goTypes = nil
	file_contacts_proto_depIdxs = nil
}
