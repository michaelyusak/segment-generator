package entity

type Node struct {
	ID    int64  `json:"id"`
	Ports []Port `json:"ports"`
}
