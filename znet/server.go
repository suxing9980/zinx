package znet

import (
	"fmt"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// tcp4 or other
	IPVersion string
	// 服务器绑定的IP地址
	IP string
	// 服务器绑定的端口
	Port int
}

// 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	// 开启一个go程去去做服务器端的Linster业务
	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		// 2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		// 已经监听成功
		fmt.Println("start Zinx server", s.Name, " succ, now listenning...")
		// 3 启动server网络连接业务
		for {
			// 3.1 阻塞等待客户端建立连接
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			// 3.2 TODO Server.Start() 设置服务器最大连接控制，如果超过最大连接，则关闭新的连接
			// 3.3 TODO Server.Start() 处理新请求的业务方法，此时handle和conn应该是绑定的

			// 这里暂时在做一个最大512字节的回显服务
			go func() {
				// 不断循环，从客户端获取数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}
					// 回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}

				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}
