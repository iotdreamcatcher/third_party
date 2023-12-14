package response

type CommonResponse struct {
	Code      int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RequestID string      `protobuf:"bytes,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Path      string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Data      interface{} `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}
