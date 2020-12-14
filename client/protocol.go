package client

type protocol struct {
	Status string              `json:"status"`
	Data   []map[string]string `json:"data"`
}

// Service ...
type Service struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Group     string `json:"group"`
	Type      string `json:"type"`
	PrivateIP string `json:"pip"`
	Tag       string `json:"tag"`
	Path      string `json:"path"`
}

// GetPrivateHost ...
func (s *Service) GetPrivateHost() string {
	if len(s.Port) == 0 {
		return s.PrivateIP
	}

	return s.PrivateIP + ":" + s.Port
}

// GetPublicHost ...
func (s *Service) GetPublicHost() string {
	if len(s.Port) == 0 {
		return s.IP
	}

	return s.IP + ":" + s.Port
}

func (p *protocol) extractServices() []Service {
	ss := []Service{}

	for i := 0; i < len(p.Data); i++ {
		ss = append(ss, Service{
			Name:      p.Data[i]["name"],
			IP:        p.Data[i]["ip"],
			Port:      p.Data[i]["prot"],
			Group:     p.Data[i]["group"],
			Type:      p.Data[i]["type"],
			PrivateIP: p.Data[i]["pip"],
			Tag:       p.Data[i]["tag"],
			Path:      p.Data[i]["path"],
		})
	}

	return ss
}
