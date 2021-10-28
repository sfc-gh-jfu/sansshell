// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proxy.proto

package proxy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*ProxyRequest_StartStream
	//	*ProxyRequest_StreamData
	//	*ProxyRequest_ClientClose
	//	*ProxyRequest_ClientCancel
	Request isProxyRequest_Request `protobuf_oneof:"request"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (m *ProxyRequest) GetRequest() isProxyRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ProxyRequest) GetStartStream() *StartStream {
	if x, ok := x.GetRequest().(*ProxyRequest_StartStream); ok {
		return x.StartStream
	}
	return nil
}

func (x *ProxyRequest) GetStreamData() *StreamData {
	if x, ok := x.GetRequest().(*ProxyRequest_StreamData); ok {
		return x.StreamData
	}
	return nil
}

func (x *ProxyRequest) GetClientClose() *ClientClose {
	if x, ok := x.GetRequest().(*ProxyRequest_ClientClose); ok {
		return x.ClientClose
	}
	return nil
}

func (x *ProxyRequest) GetClientCancel() *ClientCancel {
	if x, ok := x.GetRequest().(*ProxyRequest_ClientCancel); ok {
		return x.ClientCancel
	}
	return nil
}

type isProxyRequest_Request interface {
	isProxyRequest_Request()
}

type ProxyRequest_StartStream struct {
	// A request to open a new stream to a target.
	StartStream *StartStream `protobuf:"bytes,1,opt,name=start_stream,json=startStream,proto3,oneof"`
}

type ProxyRequest_StreamData struct {
	// Additional client data for one or more established streams.
	StreamData *StreamData `protobuf:"bytes,2,opt,name=stream_data,json=streamData,proto3,oneof"`
}

type ProxyRequest_ClientClose struct {
	// A ClientClose indicates that no more data will be sent
	// for one or more established streams.
	ClientClose *ClientClose `protobuf:"bytes,3,opt,name=client_close,json=clientClose,proto3,oneof"`
}

type ProxyRequest_ClientCancel struct {
	// A ClientCancel indicates that the client wants to cancel
	// one or more established streams.
	ClientCancel *ClientCancel `protobuf:"bytes,4,opt,name=client_cancel,json=clientCancel,proto3,oneof"`
}

func (*ProxyRequest_StartStream) isProxyRequest_Request() {}

func (*ProxyRequest_StreamData) isProxyRequest_Request() {}

func (*ProxyRequest_ClientClose) isProxyRequest_Request() {}

func (*ProxyRequest_ClientCancel) isProxyRequest_Request() {}

type ProxyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reply:
	//	*ProxyReply_StartStreamReply
	//	*ProxyReply_StreamData
	//	*ProxyReply_ServerClose
	Reply isProxyReply_Reply `protobuf_oneof:"reply"`
}

func (x *ProxyReply) Reset() {
	*x = ProxyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyReply) ProtoMessage() {}

func (x *ProxyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyReply.ProtoReflect.Descriptor instead.
func (*ProxyReply) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (m *ProxyReply) GetReply() isProxyReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ProxyReply) GetStartStreamReply() *StartStreamReply {
	if x, ok := x.GetReply().(*ProxyReply_StartStreamReply); ok {
		return x.StartStreamReply
	}
	return nil
}

func (x *ProxyReply) GetStreamData() *StreamData {
	if x, ok := x.GetReply().(*ProxyReply_StreamData); ok {
		return x.StreamData
	}
	return nil
}

func (x *ProxyReply) GetServerClose() *ServerClose {
	if x, ok := x.GetReply().(*ProxyReply_ServerClose); ok {
		return x.ServerClose
	}
	return nil
}

type isProxyReply_Reply interface {
	isProxyReply_Reply()
}

type ProxyReply_StartStreamReply struct {
	// A reply to a request to initiate a stream.
	StartStreamReply *StartStreamReply `protobuf:"bytes,1,opt,name=start_stream_reply,json=startStreamReply,proto3,oneof"`
}

type ProxyReply_StreamData struct {
	// Additional server data for one or more established streams.
	StreamData *StreamData `protobuf:"bytes,2,opt,name=stream_data,json=streamData,proto3,oneof"`
}

type ProxyReply_ServerClose struct {
	// An end of stream message for one or more established streams.
	ServerClose *ServerClose `protobuf:"bytes,3,opt,name=server_close,json=serverClose,proto3,oneof"`
}

func (*ProxyReply_StartStreamReply) isProxyReply_Reply() {}

func (*ProxyReply_StreamData) isProxyReply_Reply() {}

func (*ProxyReply_ServerClose) isProxyReply_Reply() {}

// A request to start a stream to a target host.
// The supplied `nonce` is an arbitrary client-chosen value
// that will be echoed in the returned reply to allow clients
// to correlate this request with the associated stream id.
type StartStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stream target, as accepted by grpc.Dial.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The fully-qualified method name (e.g. "/Package.Service/Method")
	MethodName string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// A nonce value which will be echoed in the reply
	// to allow the client to correlate this stream
	// request with the server-assigned stream ID.
	Nonce uint32 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// True if the client will send a stream of requests.
	// if both client_stream and server_stream are false,
	// the opened stream will be automatically closed after
	// the first message is processed.
	ClientStreams bool `protobuf:"varint,4,opt,name=client_streams,json=clientStreams,proto3" json:"client_streams,omitempty"`
	// True if the server will send a stream of replies.
	ServerStreams bool `protobuf:"varint,5,opt,name=server_streams,json=serverStreams,proto3" json:"server_streams,omitempty"`
}

func (x *StartStream) Reset() {
	*x = StartStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStream) ProtoMessage() {}

func (x *StartStream) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStream.ProtoReflect.Descriptor instead.
func (*StartStream) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *StartStream) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *StartStream) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *StartStream) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *StartStream) GetClientStreams() bool {
	if x != nil {
		return x.ClientStreams
	}
	return false
}

func (x *StartStream) GetServerStreams() bool {
	if x != nil {
		return x.ServerStreams
	}
	return false
}

type StartStreamReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The target string originally supplied to StartStream
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The nonce value supplied by the client in StartStream.
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Types that are assignable to Reply:
	//	*StartStreamReply_StreamId
	//	*StartStreamReply_ErrorStatus
	Reply isStartStreamReply_Reply `protobuf_oneof:"reply"`
}

func (x *StartStreamReply) Reset() {
	*x = StartStreamReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStreamReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStreamReply) ProtoMessage() {}

func (x *StartStreamReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStreamReply.ProtoReflect.Descriptor instead.
func (*StartStreamReply) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *StartStreamReply) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *StartStreamReply) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (m *StartStreamReply) GetReply() isStartStreamReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *StartStreamReply) GetStreamId() uint64 {
	if x, ok := x.GetReply().(*StartStreamReply_StreamId); ok {
		return x.StreamId
	}
	return 0
}

func (x *StartStreamReply) GetErrorStatus() *Status {
	if x, ok := x.GetReply().(*StartStreamReply_ErrorStatus); ok {
		return x.ErrorStatus
	}
	return nil
}

type isStartStreamReply_Reply interface {
	isStartStreamReply_Reply()
}

type StartStreamReply_StreamId struct {
	// The server-assigned stream identifier, which should be included
	// in all future messages for this stream.
	// stream_ids are only guaranteed to be unique within the
	// context of a single proxy stream.
	StreamId uint64 `protobuf:"varint,3,opt,name=stream_id,json=streamId,proto3,oneof"`
}

type StartStreamReply_ErrorStatus struct {
	// Status carries an error if the stream could not be
	// established.
	ErrorStatus *Status `protobuf:"bytes,4,opt,name=error_status,json=errorStatus,proto3,oneof"`
}

func (*StartStreamReply_StreamId) isStartStreamReply_Reply() {}

func (*StartStreamReply_ErrorStatus) isStartStreamReply_Reply() {}

// ClientClose is sent by the proxy client to indicate
// that no more messages will be sent to the given stream(s)
// Note that clients do not need to send a ClientClose for
// streams where client_streams is false.
type ClientClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-asssigned stream id(s) to close.
	StreamIds []uint64 `protobuf:"varint,1,rep,packed,name=stream_ids,json=streamIds,proto3" json:"stream_ids,omitempty"`
}

func (x *ClientClose) Reset() {
	*x = ClientClose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientClose) ProtoMessage() {}

func (x *ClientClose) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientClose.ProtoReflect.Descriptor instead.
func (*ClientClose) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *ClientClose) GetStreamIds() []uint64 {
	if x != nil {
		return x.StreamIds
	}
	return nil
}

// ClientCancel is sent by the proxy client to request
// cancellation of the given stream(s).
type ClientCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-assigned stream id(s) to cancel.
	StreamIds []uint64 `protobuf:"varint,1,rep,packed,name=stream_ids,json=streamIds,proto3" json:"stream_ids,omitempty"`
}

func (x *ClientCancel) Reset() {
	*x = ClientCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCancel) ProtoMessage() {}

func (x *ClientCancel) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCancel.ProtoReflect.Descriptor instead.
func (*ClientCancel) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *ClientCancel) GetStreamIds() []uint64 {
	if x != nil {
		return x.StreamIds
	}
	return nil
}

// StreamData is used by both clients and servers to transmit
// data for an established stream.
type StreamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stream identifier, as returned in StartStreamReply
	// This can be repeated, to indicate that the same data is relevant
	// to multiple established streams.
	StreamIds []uint64 `protobuf:"varint,1,rep,packed,name=stream_ids,json=streamIds,proto3" json:"stream_ids,omitempty"`
	// The message payload
	Payload *anypb.Any `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamData) Reset() {
	*x = StreamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamData) ProtoMessage() {}

func (x *StreamData) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamData.ProtoReflect.Descriptor instead.
func (*StreamData) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{6}
}

func (x *StreamData) GetStreamIds() []uint64 {
	if x != nil {
		return x.StreamIds
	}
	return nil
}

func (x *StreamData) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

// A server end-of-stream response, containing the final status
// of the stream.
type ServerClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stream identifier, as returned in StartStreamReply
	// This can be repeated, to indicate that the same status is
	// applicable to multiple streams.
	StreamIds []uint64 `protobuf:"varint,1,rep,packed,name=stream_ids,json=streamIds,proto3" json:"stream_ids,omitempty"`
	// The final status of the stream.
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ServerClose) Reset() {
	*x = ServerClose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerClose) ProtoMessage() {}

func (x *ServerClose) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerClose.ProtoReflect.Descriptor instead.
func (*ServerClose) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{7}
}

func (x *ServerClose) GetStreamIds() []uint64 {
	if x != nil {
		return x.StreamIds
	}
	return nil
}

func (x *ServerClose) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// A wire-compatible version of google.rpc.Status
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code (one of google.rpc.Code)
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-targeted error message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// List of messages carrying error details.
	Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{8}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfd, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x34, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x37, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xcd, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47,
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0xaa, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x9c, 0x01, 0x0a,
	0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x0a, 0x0b, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x0c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x66, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x32, 0x3e, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x13, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f,
	0x73, 0x61, 0x6e, 0x73, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proxy_proto_goTypes = []interface{}{
	(*ProxyRequest)(nil),     // 0: Proxy.ProxyRequest
	(*ProxyReply)(nil),       // 1: Proxy.ProxyReply
	(*StartStream)(nil),      // 2: Proxy.StartStream
	(*StartStreamReply)(nil), // 3: Proxy.StartStreamReply
	(*ClientClose)(nil),      // 4: Proxy.ClientClose
	(*ClientCancel)(nil),     // 5: Proxy.ClientCancel
	(*StreamData)(nil),       // 6: Proxy.StreamData
	(*ServerClose)(nil),      // 7: Proxy.ServerClose
	(*Status)(nil),           // 8: Proxy.Status
	(*anypb.Any)(nil),        // 9: google.protobuf.Any
}
var file_proxy_proto_depIdxs = []int32{
	2,  // 0: Proxy.ProxyRequest.start_stream:type_name -> Proxy.StartStream
	6,  // 1: Proxy.ProxyRequest.stream_data:type_name -> Proxy.StreamData
	4,  // 2: Proxy.ProxyRequest.client_close:type_name -> Proxy.ClientClose
	5,  // 3: Proxy.ProxyRequest.client_cancel:type_name -> Proxy.ClientCancel
	3,  // 4: Proxy.ProxyReply.start_stream_reply:type_name -> Proxy.StartStreamReply
	6,  // 5: Proxy.ProxyReply.stream_data:type_name -> Proxy.StreamData
	7,  // 6: Proxy.ProxyReply.server_close:type_name -> Proxy.ServerClose
	8,  // 7: Proxy.StartStreamReply.error_status:type_name -> Proxy.Status
	9,  // 8: Proxy.StreamData.payload:type_name -> google.protobuf.Any
	8,  // 9: Proxy.ServerClose.status:type_name -> Proxy.Status
	9,  // 10: Proxy.Status.details:type_name -> google.protobuf.Any
	0,  // 11: Proxy.Proxy.Proxy:input_type -> Proxy.ProxyRequest
	1,  // 12: Proxy.Proxy.Proxy:output_type -> Proxy.ProxyReply
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRequest); i {
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
		file_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyReply); i {
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
		file_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStream); i {
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
		file_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStreamReply); i {
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
		file_proxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientClose); i {
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
		file_proxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCancel); i {
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
		file_proxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamData); i {
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
		file_proxy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerClose); i {
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
		file_proxy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
	file_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProxyRequest_StartStream)(nil),
		(*ProxyRequest_StreamData)(nil),
		(*ProxyRequest_ClientClose)(nil),
		(*ProxyRequest_ClientCancel)(nil),
	}
	file_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ProxyReply_StartStreamReply)(nil),
		(*ProxyReply_StreamData)(nil),
		(*ProxyReply_ServerClose)(nil),
	}
	file_proxy_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StartStreamReply_StreamId)(nil),
		(*StartStreamReply_ErrorStatus)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
