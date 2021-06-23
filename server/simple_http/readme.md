###为什么可以同时提供 HTTP 接口

    关键一点，gRPC 的协议是基于 HTTP/2 的，因此应用程序能够在单个 TCP 端口上提供 HTTP/1.1 和 gRPC 接口服务（两种不同的流量）
    怎么同时提供 HTTP 接口

###检测协议
    if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
        server.ServeHTTP(w, r)
    } else {
        mux.ServeHTTP(w, r)
    }

###流程
    检测请求协议是否为 HTTP/2
    判断 Content-Type 是否为 application/grpc（gRPC 的默认标识位）
    根据协议的不同转发到不同的服务处理