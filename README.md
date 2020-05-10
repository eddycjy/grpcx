# packx

[TODO] Quick use of grpc and feature expansion.

## Demo

```go
func main() {
	lis1, _ := net.Listen("tcp", ":9001")
	lis2, _ := net.Listen("tcp", ":9002")

	grpcS := driver.NewGRPCServer()
	grpcS.SetListener(lis1)
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world."))
	})
	grpc2Gateway := driver.NewGRPC2GatewayServer(grpcS.GetEngine(), httpMux, runtime.NewServeMux())
	grpc2Gateway.SetListener(lis1)

	ginS := driver.NewGinServer()
	ginS.SetListener(lis2)
	ginS.GetEngine().GET("/hello_world", func(c *gin.Context) {
		c.JSON(200, "hello world.")
	})

	engine := packx.New()
	engine.Use(grpc2Gateway, ginS)
	engine.Run()
}
```

## feature

- 服务端运行
    - grpc
    - grpc-gateway
    - cmux
    - gin
    - net/http
    - swagger
- 常用拦截器
- 服务注册和发现
    - etcd
    - consul
- 链路追踪
    - opentracing
    - jaeger
- 自定义功能支持
    - callback
    - logger
- 工具链（grpcx-tools）
    - 交互式生成
    - 自动生成模板
    - 快速生成 proto
    - 快速生成 swagger
 - 单元测试