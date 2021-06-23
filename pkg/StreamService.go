package pkg

import (
	pb "github.com/EDDYCJY/go-grpc-example/proto"
	"io"
	"log"
)

type StreamService struct{}


//List：服务器端流式 RPC
//Record：客户端流式 RPC
//Route：双向流式 RPC

/*
通过阅读源码，可得知是 protoc 在生成时，根据定义生成了各式各样符合标准的接口方法。最终再统一调度内部的 SendMsg 方法，该方法涉及以下过程:
1、消息体（对象）序列化
2、压缩序列化后的消息体
3、对正在传输的消息体增加 5 个字节的 header
4、判断压缩+序列化后的消息体总字节长度是否大于预设的 maxSendMessageSize（预设值为 math.MaxInt32），若超出则提示错误
5、写入给流的数据集
 */

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}

