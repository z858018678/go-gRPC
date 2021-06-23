package main

import (
	"gRPC/pkg"
	pb "github.com/EDDYCJY/go-grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const PORT = "9001"
//gRPC Server 就建立起需证书认证的服务啦
func main() {
	//生成证书，加密、解密传输内容 根据服务端输入的证书文件和密钥构造 TLS 凭证
	c,err:=credentials.NewClientTLSFromFile("../../conf/server.pem","../../conf/server.key")
	if err!=nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}

	server := grpc.NewServer(grpc.Creds(c))

	pb.RegisterSearchServiceServer(server, &pkg.SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
