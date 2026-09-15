package entity

type Node struct {
	ID    string `json:"id"`
	Ports []Port `json:"ports"`
}
