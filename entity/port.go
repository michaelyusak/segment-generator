package entity

type Port struct {
	ID     string `json:"id"`
	Value  *int64 `json:"value"`
	NodeID int64  `json:"node_id,omitempty"`
}
