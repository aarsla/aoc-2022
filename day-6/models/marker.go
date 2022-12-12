package models

type Marker int

const (
	Message Marker = 14
	Packet         = 4
)

func (m Marker) EnumIndex() int {
	return int(m)
}
