package middleware

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type CommonData struct {
	RequestID string      `json:"requestId"`
	Data      interface{} `json:"data"`
}

type CommonResponse struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func ReturnProtoMsg(result CommonResponse, method string) (proto.Message, error) {

	//tempMethod := method[1:]
	name := protoreflect.FullName("tenant.TenantRpcService/TenantList")
	returnData, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		logx.Infof("打印一下是否能够找到返回值的结构体: %v+", returnData)
		return nil, err
	}
	msg := returnData.New().Interface()

	jsonBt, err := json.Marshal(&result)
	if err != nil {
		return nil, err
	}

	if err = protojson.Unmarshal(jsonBt, msg); err != nil {
		return nil, err
	}

	tempMsg := msg.(proto.Message)

	return tempMsg, nil
}
