package models

type TimeFrame int

const (
	M1  TimeFrame = 1
	M5  TimeFrame = 5
	M15 TimeFrame = 15
	M30 TimeFrame = 30
	M60 TimeFrame = 60
	H4  TimeFrame = 240
	D1  TimeFrame = 1440
)
