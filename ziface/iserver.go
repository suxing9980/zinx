package ziface

type IServer interface {
	// 启动服务器方法
	Start()
	// 停止服务器方法
	Stop()
	// 开启服务器方法
	Serve()
}
