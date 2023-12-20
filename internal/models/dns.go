package models

type Dns struct {
	Timestamp string
	SrcAddr   string
	DstAddr   string
	Question  question
	Answers   []answer
	Response  string
}

func (d *Dns) ToString() string {
	return string(d.SrcAddr + " -> " + d.DstAddr + " " + d.Question.Name + " " + d.Question.Type)
}

func NewDns(timestamp string, srcaddr string, dstaddr string, quest question, response string) *Dns {
	dns := &Dns{
		Timestamp: timestamp,
		SrcAddr:   srcaddr,
		DstAddr:   dstaddr,
		Question:  quest,
		Answers:   nil,
		Response:  response,
	}

	return dns
}
