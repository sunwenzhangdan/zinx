package znet

import (
	"fmt"
	"net"
	"zinx_byme/iface"
)

type Server struct {
	Name string
	Port string
	IPversion string
}


func(s *Server) Start(){
	fmt.Printf("start server listener ip:%s,port %s is starting",s.IPversion,s.Port)
	addr,err:=net.ResolveTCPAddr(s.IPversion,fmt.Sprintf("%s:%d",s.IPversion,s.Port))
	if err!=nil{
		fmt.Println("resove tcp error",err)
	}
	net.ListenTCP(s.IPversion,addr)


}

func(s *Server) Stop(){

}

func (s *Server) Server(){

}

func  NewServer(name,port,ipversion string) iface.IServer{
	s:=&Server{
		Port: name,
		Name: port,
		IPversion:ipversion,

	}
	return s
}