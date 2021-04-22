package cmd

import (
	"planet/env"
	"planet/env/server"
	"planet/pkg/gcore"
	//"planet/insecure"
	"github.com/spf13/cobra"
	//mq "planet/env/mq"
)

func init() {
	//服务
	RootCmd.AddCommand(serveCmd)
	//网关
	RootCmd.AddCommand(gatewayCmd)
	//单体服务
	RootCmd.AddCommand(RunCmd)
}

//微服务
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the serve ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)>0 {
			servFlag := args[0]
			grpcPort := env.Config.GetString("Server."+servFlag+".GrpcPort")
			//各个rpc服务
			gcore.RunServeGRPC(server.ConfigServer(servFlag),grpcPort)
			//mq.New().RunConsumers(mq.ConfigConsumer)
			/*var rbmqConfig grbmq.ConnConfig
			env.ScanConfig("Rbmq",&rbmqConfig)
			grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/
		}
	},
}

// 网关
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Launches the gateway  ",
	Run: func(cmd *cobra.Command, args []string) {
		gcore.RunServeHTTP(server.ConfigServer("GateWay"),env.Config.GetString("Server.GateWay.HttpPort"))
	},
}

// 单体服务（开发）
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Launches the example webserver  ",
	Run: func(cmd *cobra.Command, args []string) {
		//单体模式，所有服务+网关，全部起来
		go gcore.RunServeGRPC(server.ConfigServer(""), env.Config.GetString("Server.GateWay.GrpcPort"))
		gcore.RunServeHTTP(server.ConfigServer(""),env.Config.GetString("Server.GateWay.HttpPort"))
		//gcore.MakeInsecure(insecure.Key,insecure.Cert)

		/*var rbmqConfig grbmq.ConnConfig
		env.ScanConfig("Rbmq",&rbmqConfig)
		grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/
	},
}