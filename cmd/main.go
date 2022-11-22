package main

import (
	"log"
	"sammod_5-6/internal/service"
)

var machine service.Machine

const (
	L     = 0.5
	Nu    = 1.0
	V     = 0.01
	Y     = 0.1
	N     = 5
	Count = 10000001
)

func init() {

	machine = service.Machine{
		NewRequestTime:  1 / L,
		DoneRequestTime: Nu,
		FailTime:        1 / V,
		RepairTime:      1 / Y,
		MaxQueue:        N,
		CurrentQueue:    0,
		AllCount:        0,
		DoneCount:       0,
		FailCount:       0,
		IterCount:       1,
		RepairStart:     0,
		RequestStart:    0,
		ChannelState:    0,
		States:          service.States{},
	}
}

func main() {
	log.Println("🚀 Simulation started 🚀")
	machine.StartSimulating(Count, L)
	machine.OutputMain(L)
	log.Println("✅ Simulation done ✅")
	//machine.OutputOther(Count)
}
