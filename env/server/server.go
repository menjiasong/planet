package server

import (
	"planet/env"
	"planet/pb"
	"planet/pkg/gcore"
	"planet/service"
)

func ConfigServer(servFlag string)  []gcore.ServeSetting {
	var serveSettings  []gcore.ServeSetting
	switch servFlag {
		case "Bas":
			//基础
			serveSettings =  []gcore.ServeSetting{
				//测试demo
				{":"+env.Config.GetString("Server.Bas.GrpcPort"),&service.TestServer{},pb.RegisterTestServer,pb.RegisterTestHandlerFromEndpoint},
			}
			break
		case "Pro":
			//商品
			serveSettings =  []gcore.ServeSetting{
				//商品
				{":"+env.Config.GetString("Server.Pro.GrpcPort"),&service.ProServer{},pb.RegisterProServer,pb.RegisterProHandlerFromEndpoint},
			}
			break
		case "Usr":
			//用户
			serveSettings =  []gcore.ServeSetting{
				//用户
				{":"+env.Config.GetString("Server.Usr.GrpcPort"),&service.UsrServer{},pb.RegisterUsrServer,pb.RegisterUsrHandlerFromEndpoint},
			}
			break
		case "GateWay":
			//网关
			serveSettings =  []gcore.ServeSetting{
				//测试demo
				{"bas:"+env.Config.GetString("Server.Bas.GrpcPort"),&service.TestServer{},pb.RegisterTestServer,pb.RegisterTestHandlerFromEndpoint},
				{"pro:"+env.Config.GetString("Server.Pro.GrpcPort"),&service.ProServer{},pb.RegisterProServer,pb.RegisterProHandlerFromEndpoint},
				{"usr:"+env.Config.GetString("Server.Usr.GrpcPort"),&service.UsrServer{},pb.RegisterUsrServer,pb.RegisterUsrHandlerFromEndpoint},
			}
			break
		default:
		//单体模式，所有服务+网关，全部起来
		serveSettings =  []gcore.ServeSetting{
			//测试demo
			{":"+env.Config.GetString("Server.GateWay.GrpcPort"),&service.TestServer{},pb.RegisterTestServer,pb.RegisterTestHandlerFromEndpoint},
			{":"+env.Config.GetString("Server.GateWay.GrpcPort"),&service.ProServer{},pb.RegisterProServer,pb.RegisterProHandlerFromEndpoint},
			{":"+env.Config.GetString("Server.GateWay.GrpcPort"),&service.UsrServer{},pb.RegisterUsrServer,pb.RegisterUsrHandlerFromEndpoint},
		}
		break
	}

	return serveSettings

}
