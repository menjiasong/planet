package server

import (
	"planet/env"
	"planet/pkg/grbmq"
)

var Rbmq *grbmq.Rbmq

func New() *grbmq.Rbmq {
	if  Rbmq.Priority == "" {
		var rbmqConfig grbmq.ConnConfig
		env.ScanConfig("Rbmq",&rbmqConfig)
		Rbmq = grbmq.New(rbmqConfig)
		Rbmq.Priority = "10"
	}
	return Rbmq
}