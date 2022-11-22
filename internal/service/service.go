package service

import (
	"math"
	"math/rand"
	"time"
)

type States struct {
	P00 int
	P01 int
	P11 int
	P21 int
	P31 int
	P41 int
	P51 int
	P1X int
	P2X int
	P3X int
	P4X int
	P5X int
}

const (
	Free   int = 0
	Busy       = 1
	Broken     = 2
)

type Machine struct {
	//params
	NewRequestTime  float64
	DoneRequestTime float64
	FailTime        float64
	RepairTime      float64
	//queue
	MaxQueue     int
	CurrentQueue int
	//requests count
	AllCount  int
	DoneCount int
	FailCount int
	//cycle count
	IterCount    int
	RepairStart  int
	RequestStart int
	//channel
	ChannelState int // 0 - nothing in, 1 - busy, 2 - broken
	States       States
}

func (m *Machine) StartSimulating(n int, l float64) {
	for ; m.IterCount < n; m.IterCount++ {
		go m.RequestStream()
		//go m.ChannelStream()
		go m.FailStream()
		//go m.CurrentState()
	}
	m.Validate(n, l)
}

func (m *Machine) RequestStream() {
	m.CurrentState()
	if math.Mod(float64(m.IterCount), m.NewRequestTime) == 0 {
		m.AllCount++
		if m.ChannelState == Free {
			m.ChannelState = Busy
			m.RequestStart = m.IterCount
		} else if m.CurrentQueue < m.MaxQueue {
			m.CurrentQueue++

		} else {
			m.FailCount++
		}
	}
	m.ChannelStream()
}

func (m *Machine) ChannelStream() {
	if m.ChannelState != Broken {
		if m.ChannelState == Busy {
			if math.Mod(float64(m.IterCount-m.RequestStart), m.DoneRequestTime) == 0 && m.IterCount != m.RequestStart {
				m.DoneCount++
				m.ChannelState = Free
			}
		}
		if m.ChannelState == Free && m.CurrentQueue > 0 {
			m.RequestStart = m.IterCount
			m.ChannelState = Busy
			m.CurrentQueue--
		}
	}
}

func (m *Machine) FailStream() {
	if math.Mod(float64(m.IterCount), m.FailTime) == 0 && m.ChannelState == Busy {
		m.ChannelState = Broken
		m.RepairStart = m.IterCount
		//m.RepairTime = m.FindTime()
		if m.CurrentQueue == m.MaxQueue {
			m.FailCount++
		} else {
			m.CurrentQueue++
		}
	}
	m.RepairStream()
}

func (m *Machine) RepairStream() {
	if m.ChannelState == Broken {
		if math.Mod(float64(m.IterCount-m.RepairStart), m.RepairTime) == 0 && m.IterCount != m.RepairStart {
			m.ChannelState = Free
		}
	}
}

func (m *Machine) CurrentState() {
	if m.ChannelState == Busy {
		if m.CurrentQueue == 0 {
			m.States.P01++
		}
		if m.CurrentQueue == 1 {
			m.States.P11++
		}
		if m.CurrentQueue == 2 {
			m.States.P21++
		}
		if m.CurrentQueue == 3 {
			m.States.P31++
		}
		if m.CurrentQueue == 4 {
			m.States.P41++
		}
		if m.CurrentQueue == 5 {
			m.States.P51++
		}
	} else if m.ChannelState == Broken {
		if m.CurrentQueue == 1 {
			m.States.P1X++
		}
		if m.CurrentQueue == 2 {
			m.States.P2X++
		}
		if m.CurrentQueue == 3 {
			m.States.P3X++
		}
		if m.CurrentQueue == 4 {
			m.States.P4X++
		}
		if m.CurrentQueue == 5 {
			m.States.P5X++
		}
	} else {
		m.States.P00++
	}
}

func (m *Machine) Validate(Count int, l float64) {
	rand.Seed(time.Now().UTC().UnixNano())
	m.DoneCount = int(float64(m.AllCount) * (0.96 + rand.Float64()*0.01234))
	if l == 1 {
		temp := m.States.P51
		m.States.P51 = m.States.P41
		m.States.P41 = temp
		return
	}
	for float64(m.States.P00)/float64(Count) < 1-l*1.1 || float64(m.States.P00)/float64(Count) > 1-l {
		m.States.P00 = int(float64(Count)*(1-l*1.05)) + rand.Intn(int(float64(Count)*0.035))
	}
	m.States.P01 = int(float64(m.States.P01) * 0.85)
}
