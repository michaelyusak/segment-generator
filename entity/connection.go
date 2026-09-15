package entity

type PortConnection struct {
	Name   string `json:"name"`
	Source Port   `json:"source"`
	Target Port   `json:"target"`
}

func (pc *PortConnection) WriteName() {
	pc.Name = pc.Target.ID + " - " + pc.Source.ID
}
