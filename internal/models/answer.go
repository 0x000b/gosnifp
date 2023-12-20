package models

type answer struct {
	Name   string
	IpAddr string
	TTL    string
}

func NewAnswer(name, ip, ttl string) answer {
	return answer{
		Name:   name,
		IpAddr: ip,
		TTL:    ttl,
	}
}
