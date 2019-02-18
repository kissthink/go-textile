// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cafe.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CafeChallenge struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeChallenge) Reset()         { *m = CafeChallenge{} }
func (m *CafeChallenge) String() string { return proto.CompactTextString(m) }
func (*CafeChallenge) ProtoMessage()    {}
func (*CafeChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{0}
}
func (m *CafeChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeChallenge.Unmarshal(m, b)
}
func (m *CafeChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeChallenge.Marshal(b, m, deterministic)
}
func (dst *CafeChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeChallenge.Merge(dst, src)
}
func (m *CafeChallenge) XXX_Size() int {
	return xxx_messageInfo_CafeChallenge.Size(m)
}
func (m *CafeChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CafeChallenge proto.InternalMessageInfo

func (m *CafeChallenge) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CafeNonce struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeNonce) Reset()         { *m = CafeNonce{} }
func (m *CafeNonce) String() string { return proto.CompactTextString(m) }
func (*CafeNonce) ProtoMessage()    {}
func (*CafeNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{1}
}
func (m *CafeNonce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeNonce.Unmarshal(m, b)
}
func (m *CafeNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeNonce.Marshal(b, m, deterministic)
}
func (dst *CafeNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeNonce.Merge(dst, src)
}
func (m *CafeNonce) XXX_Size() int {
	return xxx_messageInfo_CafeNonce.Size(m)
}
func (m *CafeNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeNonce.DiscardUnknown(m)
}

var xxx_messageInfo_CafeNonce proto.InternalMessageInfo

func (m *CafeNonce) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CafeRegistration struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sig                  []byte   `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRegistration) Reset()         { *m = CafeRegistration{} }
func (m *CafeRegistration) String() string { return proto.CompactTextString(m) }
func (*CafeRegistration) ProtoMessage()    {}
func (*CafeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{2}
}
func (m *CafeRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRegistration.Unmarshal(m, b)
}
func (m *CafeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRegistration.Marshal(b, m, deterministic)
}
func (dst *CafeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRegistration.Merge(dst, src)
}
func (m *CafeRegistration) XXX_Size() int {
	return xxx_messageInfo_CafeRegistration.Size(m)
}
func (m *CafeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRegistration proto.InternalMessageInfo

func (m *CafeRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CafeRegistration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CafeRegistration) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *CafeRegistration) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *CafeRegistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeRefreshSession struct {
	Access               string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              string   `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRefreshSession) Reset()         { *m = CafeRefreshSession{} }
func (m *CafeRefreshSession) String() string { return proto.CompactTextString(m) }
func (*CafeRefreshSession) ProtoMessage()    {}
func (*CafeRefreshSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{3}
}
func (m *CafeRefreshSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRefreshSession.Unmarshal(m, b)
}
func (m *CafeRefreshSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRefreshSession.Marshal(b, m, deterministic)
}
func (dst *CafeRefreshSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRefreshSession.Merge(dst, src)
}
func (m *CafeRefreshSession) XXX_Size() int {
	return xxx_messageInfo_CafeRefreshSession.Size(m)
}
func (m *CafeRefreshSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRefreshSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRefreshSession proto.InternalMessageInfo

func (m *CafeRefreshSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeRefreshSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

type CafePublishContact struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Contact              *Contact `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafePublishContact) Reset()         { *m = CafePublishContact{} }
func (m *CafePublishContact) String() string { return proto.CompactTextString(m) }
func (*CafePublishContact) ProtoMessage()    {}
func (*CafePublishContact) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{4}
}
func (m *CafePublishContact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePublishContact.Unmarshal(m, b)
}
func (m *CafePublishContact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePublishContact.Marshal(b, m, deterministic)
}
func (dst *CafePublishContact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePublishContact.Merge(dst, src)
}
func (m *CafePublishContact) XXX_Size() int {
	return xxx_messageInfo_CafePublishContact.Size(m)
}
func (m *CafePublishContact) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePublishContact.DiscardUnknown(m)
}

var xxx_messageInfo_CafePublishContact proto.InternalMessageInfo

func (m *CafePublishContact) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafePublishContact) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type CafePublishContactAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafePublishContactAck) Reset()         { *m = CafePublishContactAck{} }
func (m *CafePublishContactAck) String() string { return proto.CompactTextString(m) }
func (*CafePublishContactAck) ProtoMessage()    {}
func (*CafePublishContactAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{5}
}
func (m *CafePublishContactAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePublishContactAck.Unmarshal(m, b)
}
func (m *CafePublishContactAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePublishContactAck.Marshal(b, m, deterministic)
}
func (dst *CafePublishContactAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePublishContactAck.Merge(dst, src)
}
func (m *CafePublishContactAck) XXX_Size() int {
	return xxx_messageInfo_CafePublishContactAck.Size(m)
}
func (m *CafePublishContactAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePublishContactAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafePublishContactAck proto.InternalMessageInfo

func (m *CafePublishContactAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Deprecated: Do not use.
type CafeContactQuery struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	FindId               string   `protobuf:"bytes,2,opt,name=findId,proto3" json:"findId,omitempty"`
	FindAddress          string   `protobuf:"bytes,3,opt,name=findAddress,proto3" json:"findAddress,omitempty"`
	FindUsername         string   `protobuf:"bytes,4,opt,name=findUsername,proto3" json:"findUsername,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Wait                 int32    `protobuf:"varint,6,opt,name=wait,proto3" json:"wait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeContactQuery) Reset()         { *m = CafeContactQuery{} }
func (m *CafeContactQuery) String() string { return proto.CompactTextString(m) }
func (*CafeContactQuery) ProtoMessage()    {}
func (*CafeContactQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{6}
}
func (m *CafeContactQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeContactQuery.Unmarshal(m, b)
}
func (m *CafeContactQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeContactQuery.Marshal(b, m, deterministic)
}
func (dst *CafeContactQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeContactQuery.Merge(dst, src)
}
func (m *CafeContactQuery) XXX_Size() int {
	return xxx_messageInfo_CafeContactQuery.Size(m)
}
func (m *CafeContactQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeContactQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CafeContactQuery proto.InternalMessageInfo

func (m *CafeContactQuery) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeContactQuery) GetFindId() string {
	if m != nil {
		return m.FindId
	}
	return ""
}

func (m *CafeContactQuery) GetFindAddress() string {
	if m != nil {
		return m.FindAddress
	}
	return ""
}

func (m *CafeContactQuery) GetFindUsername() string {
	if m != nil {
		return m.FindUsername
	}
	return ""
}

func (m *CafeContactQuery) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CafeContactQuery) GetWait() int32 {
	if m != nil {
		return m.Wait
	}
	return 0
}

// Deprecated: Do not use.
type CafeContactQueryResult struct {
	Contacts             []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CafeContactQueryResult) Reset()         { *m = CafeContactQueryResult{} }
func (m *CafeContactQueryResult) String() string { return proto.CompactTextString(m) }
func (*CafeContactQueryResult) ProtoMessage()    {}
func (*CafeContactQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{7}
}
func (m *CafeContactQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeContactQueryResult.Unmarshal(m, b)
}
func (m *CafeContactQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeContactQueryResult.Marshal(b, m, deterministic)
}
func (dst *CafeContactQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeContactQueryResult.Merge(dst, src)
}
func (m *CafeContactQueryResult) XXX_Size() int {
	return xxx_messageInfo_CafeContactQueryResult.Size(m)
}
func (m *CafeContactQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeContactQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_CafeContactQueryResult proto.InternalMessageInfo

func (m *CafeContactQueryResult) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type CafeStore struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStore) Reset()         { *m = CafeStore{} }
func (m *CafeStore) String() string { return proto.CompactTextString(m) }
func (*CafeStore) ProtoMessage()    {}
func (*CafeStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{8}
}
func (m *CafeStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStore.Unmarshal(m, b)
}
func (m *CafeStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStore.Marshal(b, m, deterministic)
}
func (dst *CafeStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStore.Merge(dst, src)
}
func (m *CafeStore) XXX_Size() int {
	return xxx_messageInfo_CafeStore.Size(m)
}
func (m *CafeStore) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStore.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStore proto.InternalMessageInfo

func (m *CafeStore) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStore) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObjectList struct {
	Cids                 []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObjectList) Reset()         { *m = CafeObjectList{} }
func (m *CafeObjectList) String() string { return proto.CompactTextString(m) }
func (*CafeObjectList) ProtoMessage()    {}
func (*CafeObjectList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{9}
}
func (m *CafeObjectList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObjectList.Unmarshal(m, b)
}
func (m *CafeObjectList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObjectList.Marshal(b, m, deterministic)
}
func (dst *CafeObjectList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObjectList.Merge(dst, src)
}
func (m *CafeObjectList) XXX_Size() int {
	return xxx_messageInfo_CafeObjectList.Size(m)
}
func (m *CafeObjectList) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObjectList.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObjectList proto.InternalMessageInfo

func (m *CafeObjectList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObject struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Node                 []byte   `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObject) Reset()         { *m = CafeObject{} }
func (m *CafeObject) String() string { return proto.CompactTextString(m) }
func (*CafeObject) ProtoMessage()    {}
func (*CafeObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{10}
}
func (m *CafeObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObject.Unmarshal(m, b)
}
func (m *CafeObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObject.Marshal(b, m, deterministic)
}
func (dst *CafeObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObject.Merge(dst, src)
}
func (m *CafeObject) XXX_Size() int {
	return xxx_messageInfo_CafeObject.Size(m)
}
func (m *CafeObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObject.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObject proto.InternalMessageInfo

func (m *CafeObject) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeObject) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *CafeObject) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CafeObject) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

type CafeStoreThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreThread) Reset()         { *m = CafeStoreThread{} }
func (m *CafeStoreThread) String() string { return proto.CompactTextString(m) }
func (*CafeStoreThread) ProtoMessage()    {}
func (*CafeStoreThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{11}
}
func (m *CafeStoreThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreThread.Unmarshal(m, b)
}
func (m *CafeStoreThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreThread.Marshal(b, m, deterministic)
}
func (dst *CafeStoreThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreThread.Merge(dst, src)
}
func (m *CafeStoreThread) XXX_Size() int {
	return xxx_messageInfo_CafeStoreThread.Size(m)
}
func (m *CafeStoreThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreThread proto.InternalMessageInfo

func (m *CafeStoreThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStoreThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeStoreThread) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type CafeClientThread struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeClientThread) Reset()         { *m = CafeClientThread{} }
func (m *CafeClientThread) String() string { return proto.CompactTextString(m) }
func (*CafeClientThread) ProtoMessage()    {}
func (*CafeClientThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{12}
}
func (m *CafeClientThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeClientThread.Unmarshal(m, b)
}
func (m *CafeClientThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeClientThread.Marshal(b, m, deterministic)
}
func (dst *CafeClientThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeClientThread.Merge(dst, src)
}
func (m *CafeClientThread) XXX_Size() int {
	return xxx_messageInfo_CafeClientThread.Size(m)
}
func (m *CafeClientThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeClientThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeClientThread proto.InternalMessageInfo

func (m *CafeClientThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeClientThread) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *CafeClientThread) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type CafeStored struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStored) Reset()         { *m = CafeStored{} }
func (m *CafeStored) String() string { return proto.CompactTextString(m) }
func (*CafeStored) ProtoMessage()    {}
func (*CafeStored) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{13}
}
func (m *CafeStored) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStored.Unmarshal(m, b)
}
func (m *CafeStored) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStored.Marshal(b, m, deterministic)
}
func (dst *CafeStored) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStored.Merge(dst, src)
}
func (m *CafeStored) XXX_Size() int {
	return xxx_messageInfo_CafeStored.Size(m)
}
func (m *CafeStored) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStored.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStored proto.InternalMessageInfo

func (m *CafeStored) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeDeliverMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeliverMessage) Reset()         { *m = CafeDeliverMessage{} }
func (m *CafeDeliverMessage) String() string { return proto.CompactTextString(m) }
func (*CafeDeliverMessage) ProtoMessage()    {}
func (*CafeDeliverMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{14}
}
func (m *CafeDeliverMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeliverMessage.Unmarshal(m, b)
}
func (m *CafeDeliverMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeliverMessage.Marshal(b, m, deterministic)
}
func (dst *CafeDeliverMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeliverMessage.Merge(dst, src)
}
func (m *CafeDeliverMessage) XXX_Size() int {
	return xxx_messageInfo_CafeDeliverMessage.Size(m)
}
func (m *CafeDeliverMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeliverMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeliverMessage proto.InternalMessageInfo

func (m *CafeDeliverMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeDeliverMessage) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type CafeCheckMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeCheckMessages) Reset()         { *m = CafeCheckMessages{} }
func (m *CafeCheckMessages) String() string { return proto.CompactTextString(m) }
func (*CafeCheckMessages) ProtoMessage()    {}
func (*CafeCheckMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{15}
}
func (m *CafeCheckMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeCheckMessages.Unmarshal(m, b)
}
func (m *CafeCheckMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeCheckMessages.Marshal(b, m, deterministic)
}
func (dst *CafeCheckMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeCheckMessages.Merge(dst, src)
}
func (m *CafeCheckMessages) XXX_Size() int {
	return xxx_messageInfo_CafeCheckMessages.Size(m)
}
func (m *CafeCheckMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeCheckMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeCheckMessages proto.InternalMessageInfo

func (m *CafeCheckMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeMessage struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId               string               `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeMessage) Reset()         { *m = CafeMessage{} }
func (m *CafeMessage) String() string { return proto.CompactTextString(m) }
func (*CafeMessage) ProtoMessage()    {}
func (*CafeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{16}
}
func (m *CafeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessage.Unmarshal(m, b)
}
func (m *CafeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessage.Marshal(b, m, deterministic)
}
func (dst *CafeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessage.Merge(dst, src)
}
func (m *CafeMessage) XXX_Size() int {
	return xxx_messageInfo_CafeMessage.Size(m)
}
func (m *CafeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessage proto.InternalMessageInfo

func (m *CafeMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeMessage) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *CafeMessage) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type CafeMessages struct {
	Messages             []*CafeMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CafeMessages) Reset()         { *m = CafeMessages{} }
func (m *CafeMessages) String() string { return proto.CompactTextString(m) }
func (*CafeMessages) ProtoMessage()    {}
func (*CafeMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{17}
}
func (m *CafeMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessages.Unmarshal(m, b)
}
func (m *CafeMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessages.Marshal(b, m, deterministic)
}
func (dst *CafeMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessages.Merge(dst, src)
}
func (m *CafeMessages) XXX_Size() int {
	return xxx_messageInfo_CafeMessages.Size(m)
}
func (m *CafeMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessages proto.InternalMessageInfo

func (m *CafeMessages) GetMessages() []*CafeMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CafeDeleteMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessages) Reset()         { *m = CafeDeleteMessages{} }
func (m *CafeDeleteMessages) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessages) ProtoMessage()    {}
func (*CafeDeleteMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{18}
}
func (m *CafeDeleteMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessages.Unmarshal(m, b)
}
func (m *CafeDeleteMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessages.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessages.Merge(dst, src)
}
func (m *CafeDeleteMessages) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessages.Size(m)
}
func (m *CafeDeleteMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessages proto.InternalMessageInfo

func (m *CafeDeleteMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeleteMessagesAck struct {
	More                 bool     `protobuf:"varint,1,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessagesAck) Reset()         { *m = CafeDeleteMessagesAck{} }
func (m *CafeDeleteMessagesAck) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessagesAck) ProtoMessage()    {}
func (*CafeDeleteMessagesAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_e72e04ed39309702, []int{19}
}
func (m *CafeDeleteMessagesAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessagesAck.Unmarshal(m, b)
}
func (m *CafeDeleteMessagesAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessagesAck.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessagesAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessagesAck.Merge(dst, src)
}
func (m *CafeDeleteMessagesAck) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessagesAck.Size(m)
}
func (m *CafeDeleteMessagesAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessagesAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessagesAck proto.InternalMessageInfo

func (m *CafeDeleteMessagesAck) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func init() {
	proto.RegisterType((*CafeChallenge)(nil), "CafeChallenge")
	proto.RegisterType((*CafeNonce)(nil), "CafeNonce")
	proto.RegisterType((*CafeRegistration)(nil), "CafeRegistration")
	proto.RegisterType((*CafeRefreshSession)(nil), "CafeRefreshSession")
	proto.RegisterType((*CafePublishContact)(nil), "CafePublishContact")
	proto.RegisterType((*CafePublishContactAck)(nil), "CafePublishContactAck")
	proto.RegisterType((*CafeContactQuery)(nil), "CafeContactQuery")
	proto.RegisterType((*CafeContactQueryResult)(nil), "CafeContactQueryResult")
	proto.RegisterType((*CafeStore)(nil), "CafeStore")
	proto.RegisterType((*CafeObjectList)(nil), "CafeObjectList")
	proto.RegisterType((*CafeObject)(nil), "CafeObject")
	proto.RegisterType((*CafeStoreThread)(nil), "CafeStoreThread")
	proto.RegisterType((*CafeClientThread)(nil), "CafeClientThread")
	proto.RegisterType((*CafeStored)(nil), "CafeStored")
	proto.RegisterType((*CafeDeliverMessage)(nil), "CafeDeliverMessage")
	proto.RegisterType((*CafeCheckMessages)(nil), "CafeCheckMessages")
	proto.RegisterType((*CafeMessage)(nil), "CafeMessage")
	proto.RegisterType((*CafeMessages)(nil), "CafeMessages")
	proto.RegisterType((*CafeDeleteMessages)(nil), "CafeDeleteMessages")
	proto.RegisterType((*CafeDeleteMessagesAck)(nil), "CafeDeleteMessagesAck")
}

func init() { proto.RegisterFile("cafe.proto", fileDescriptor_cafe_e72e04ed39309702) }

var fileDescriptor_cafe_e72e04ed39309702 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6f, 0x6b, 0x13, 0x4f,
	0x10, 0xe6, 0x2e, 0x69, 0x9a, 0x4c, 0xf2, 0xeb, 0xaf, 0x2e, 0x1a, 0x8e, 0x22, 0x1a, 0x97, 0x82,
	0xa9, 0xc2, 0x15, 0x2a, 0x82, 0xf8, 0xca, 0xb6, 0x22, 0x08, 0x5a, 0xf5, 0x5a, 0x11, 0x44, 0x84,
	0xcd, 0xde, 0x24, 0x59, 0x7b, 0x7f, 0xc2, 0xee, 0xa6, 0xea, 0x3b, 0x3f, 0x94, 0x1f, 0x50, 0xf6,
	0xcf, 0x5d, 0xae, 0xad, 0xa1, 0xf8, 0x6e, 0x9e, 0xd9, 0xe7, 0x9e, 0x99, 0x7b, 0x76, 0x76, 0x00,
	0x38, 0x9b, 0x62, 0xbc, 0x90, 0xa5, 0x2e, 0x77, 0xee, 0xcf, 0xca, 0x72, 0x96, 0xe1, 0xbe, 0x45,
	0x93, 0xe5, 0x74, 0x5f, 0x8b, 0x1c, 0x95, 0x66, 0xf9, 0xc2, 0x13, 0xfa, 0x79, 0x99, 0x62, 0xe6,
	0x00, 0xdd, 0x83, 0xff, 0x8e, 0xd9, 0x14, 0x8f, 0xe7, 0x2c, 0xcb, 0xb0, 0x98, 0x21, 0x89, 0x60,
	0x93, 0xa5, 0xa9, 0x44, 0xa5, 0xa2, 0x60, 0x14, 0x8c, 0x7b, 0x49, 0x05, 0xe9, 0x03, 0xe8, 0x19,
	0xea, 0x49, 0x59, 0x70, 0x24, 0xb7, 0x61, 0xe3, 0x82, 0x65, 0x4b, 0xf4, 0x24, 0x07, 0xe8, 0xaf,
	0x00, 0xb6, 0x0d, 0x27, 0xc1, 0x99, 0x50, 0x5a, 0x32, 0x2d, 0xca, 0x62, 0xbd, 0xe2, 0x4a, 0x24,
	0x6c, 0x88, 0x98, 0x6c, 0x61, 0x6a, 0x44, 0x2d, 0x97, 0xb5, 0x80, 0x6c, 0x43, 0x4b, 0x89, 0x59,
	0xd4, 0x1e, 0x05, 0xe3, 0x41, 0x62, 0x42, 0xc3, 0xd3, 0xe5, 0x39, 0x16, 0xd1, 0x86, 0xe3, 0x59,
	0x40, 0x5f, 0x01, 0x71, 0x1d, 0x4c, 0x25, 0xaa, 0xf9, 0x29, 0x2a, 0x65, 0x7a, 0x18, 0x42, 0x87,
	0x71, 0xbe, 0x6a, 0xc1, 0x23, 0xd3, 0x9b, 0x74, 0x4c, 0xdf, 0x43, 0x05, 0xe9, 0x89, 0xd3, 0x79,
	0xbf, 0x9c, 0x64, 0x42, 0xcd, 0x8f, 0xcb, 0x42, 0x33, 0xae, 0x57, 0x35, 0x83, 0x46, 0x4d, 0x42,
	0x61, 0x93, 0x3b, 0x82, 0x55, 0xe9, 0x1f, 0x74, 0x63, 0xff, 0x41, 0x52, 0x1d, 0xd0, 0x87, 0x70,
	0xe7, 0xba, 0xde, 0x21, 0x3f, 0x27, 0x5b, 0x10, 0x8a, 0xd4, 0xeb, 0x85, 0x22, 0xa5, 0xbf, 0xbd,
	0x87, 0x9e, 0xf2, 0x61, 0x89, 0xf2, 0xe7, 0x9a, 0xba, 0x43, 0xe8, 0x4c, 0x45, 0x91, 0xbe, 0x4e,
	0x7d, 0xf3, 0x1e, 0x91, 0x11, 0xf4, 0x4d, 0x74, 0xe8, 0x5d, 0x77, 0x3e, 0x36, 0x53, 0x84, 0xc2,
	0xc0, 0xc0, 0x8f, 0x0a, 0x65, 0xc1, 0x72, 0xb4, 0xb6, 0xf6, 0x92, 0x4b, 0x39, 0x53, 0x33, 0x13,
	0xb9, 0xd0, 0xd6, 0xdf, 0x8d, 0xc4, 0x01, 0x42, 0xa0, 0xfd, 0x9d, 0x09, 0x1d, 0x75, 0x6c, 0xd2,
	0xc6, 0xcf, 0xc3, 0x28, 0xa0, 0x47, 0x30, 0xbc, 0xda, 0x75, 0x82, 0x6a, 0x99, 0x69, 0xb2, 0x0b,
	0x5d, 0x6f, 0x82, 0x71, 0xbf, 0x75, 0xc9, 0x9e, 0xfa, 0xc4, 0x6a, 0x3c, 0x75, 0x13, 0x76, 0xaa,
	0x4b, 0x89, 0x6b, 0x7e, 0x99, 0x40, 0x9b, 0x8b, 0x54, 0x45, 0xe1, 0xa8, 0x35, 0xee, 0x25, 0x36,
	0xa6, 0xbb, 0xb0, 0x65, 0x3e, 0x7b, 0x37, 0xf9, 0x86, 0x5c, 0xbf, 0x11, 0x4a, 0xd7, 0xac, 0xa0,
	0xc1, 0xfa, 0x02, 0xb0, 0x62, 0xad, 0x51, 0xdf, 0x86, 0x16, 0x17, 0x95, 0x9b, 0x26, 0x34, 0x4a,
	0x29, 0xd3, 0xcc, 0x7a, 0x38, 0x48, 0x6c, 0x6c, 0x72, 0x45, 0x99, 0xa2, 0x9f, 0x45, 0x1b, 0xd3,
	0x4f, 0xf0, 0x7f, 0xdd, 0xfa, 0xd9, 0x5c, 0x22, 0x4b, 0xd7, 0x94, 0x70, 0xd7, 0x1d, 0x56, 0xd7,
	0x4d, 0xee, 0x01, 0x70, 0xb1, 0x98, 0xa3, 0xd4, 0xf8, 0x43, 0xfb, 0x32, 0x8d, 0x0c, 0xfd, 0xea,
	0xa7, 0x21, 0x13, 0x58, 0x68, 0xaf, 0x7c, 0x65, 0x64, 0xc8, 0x0e, 0x74, 0xb9, 0x3d, 0xaf, 0x27,
	0xa1, 0xc6, 0x37, 0xea, 0xdf, 0x75, 0xb6, 0xd8, 0xc6, 0xaf, 0x29, 0xd3, 0x17, 0xee, 0x15, 0xbc,
	0xc4, 0x4c, 0x5c, 0xa0, 0x7c, 0x8b, 0x4a, 0xb1, 0x19, 0xfe, 0x4b, 0x7d, 0xba, 0x07, 0xb7, 0xdc,
	0x82, 0x41, 0x7e, 0xee, 0xbf, 0x57, 0x7f, 0xb7, 0x86, 0x22, 0xf4, 0x0d, 0x75, 0x5d, 0x95, 0x21,
	0x74, 0x16, 0x88, 0x72, 0x35, 0xed, 0x0e, 0x91, 0xd8, 0x5e, 0x91, 0x5b, 0x17, 0xfd, 0x83, 0x9d,
	0xd8, 0xed, 0xbf, 0xb8, 0xda, 0x7f, 0xf1, 0x59, 0xb5, 0xff, 0xec, 0xf5, 0x21, 0x7d, 0x06, 0x83,
	0x46, 0x19, 0x45, 0xc6, 0xd0, 0xcd, 0x7d, 0xec, 0xe7, 0x73, 0x10, 0x37, 0x08, 0x49, 0x7d, 0x4a,
	0x1f, 0xd5, 0x6e, 0xa0, 0xc6, 0x1b, 0x7e, 0xe6, 0xb1, 0x7b, 0xef, 0x97, 0xb9, 0xe6, 0xbd, 0x13,
	0x68, 0xe7, 0xa5, 0x74, 0x8b, 0xb3, 0x9b, 0xd8, 0xf8, 0xa8, 0xfd, 0x39, 0x5c, 0x4c, 0x26, 0x1d,
	0xdb, 0xf2, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xb7, 0x8f, 0x61, 0xce, 0x05, 0x00,
	0x00,
}
