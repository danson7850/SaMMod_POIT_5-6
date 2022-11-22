package service

import "log"

func (m *Machine) OutputMain(L float64) {
	log.Println("=========================================================")
	log.Printf("|| Общее число заявок: %31d ||", m.AllCount)
	log.Printf("|| Обработанные заявки: %30d ||", m.DoneCount)
	log.Printf("|| Отказанные заявки: %32d ||", m.AllCount-m.DoneCount)
	log.Println("=========================================================")
	log.Printf("|| Вероятность отказа (Pотк): %24.5f ||", float64(m.AllCount-m.DoneCount)/float64(m.AllCount))
	log.Printf("|| Относительная пропускная способность (Q): %9.5f ||", float64(m.DoneCount)/float64(m.AllCount))
	log.Printf("|| Абсолютная пропускная способность (A): %12.5f ||", float64(m.DoneCount)/float64(m.AllCount)*L)
	log.Println("=========================================================")
}

func (m *Machine) OutputOther(Count int) {
	log.Printf("|| P00 : %.5f ||", float64(m.States.P00)/float64(Count))
	log.Printf("|| P01 : %.5f ||", float64(m.States.P01)/float64(Count))
	log.Printf("|| P11 : %.5f ||", float64(m.States.P11)/float64(Count))
	log.Printf("|| P21 : %.5f ||", float64(m.States.P21)/float64(Count))
	log.Printf("|| P31 : %.5f ||", float64(m.States.P31)/float64(Count))
	log.Printf("|| P41 : %.5f ||", float64(m.States.P41)/float64(Count))
	log.Printf("|| P51 : %.5f ||", float64(m.States.P51)/float64(Count))
	log.Printf("|| P1X : %.5f ||", float64(m.States.P1X)/float64(Count))
	log.Printf("|| P2X : %.5f ||", float64(m.States.P2X)/float64(Count))
	log.Printf("|| P3X : %.5f ||", float64(m.States.P3X)/float64(Count))
	log.Printf("|| P4X : %.5f ||", float64(m.States.P4X)/float64(Count))
	log.Printf("|| P5X : %.5f ||", float64(m.States.P5X)/float64(Count))
	log.Println("===================")
}
