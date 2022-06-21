package pkg

type Request struct {
	Time int    `json:"time,omitempty"`
	Sign string `json:"sign"`
}

type Returned struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg,omitempty"`
}

type ClientBinding struct {
	Returned
	Request
	Domain         string `json:"domain"`
	Key            string `json:"key"`
	Networklaundry int    `json:"networklaundry"`
}

type BindingRequest struct {
	Request
	Domain         string `json:"domain"`
	Key            string `json:"key"`
	Networklaundry int    `json:"networklaundry"`
}

type BindingReturned struct {
	Returned
	Request
	Networklaundry int `json:"networklaundry,omitempty"`
	Data           struct {
		BindingReturnedConfig
	} `json:"data"`
}

type WireguardReturned struct {
	Returned
	Request
	Data WireguardStatus
}

type PingData struct {
	Address string `json:"address"`
	Result  string `json:"result,omitempty"`
}

type PingRequest struct {
	Returned
	Request
	Interval int        `json:"interval,omitempty"`
	Data     []PingData `json:"data"`
}

type PingReturned struct {
	Ret int    `json:"ret,omitempty"`
	Msg string `json:"msg,omitempty"`
	Request
	Interval int        `json:"interval,omitempty"`
	Data     []PingData `json:"data"`
}

type GetLogger struct {
	Returned
	Request
	LogType string `json:"logtype"`
	NumRows int    `json:"numrows"`
}

type LogReturned struct {
	Returned
	Request
	Data LogData `json:"data"`
}

type LogData struct {
	LogData string `json:"log_data"`
}

type Wireguard struct {
	Request
	PrivateKey          string `json:"privatekey,omitempty" binding:"required"`
	Address             string `json:"address,omitempty" binding:"required"`
	DNS                 string `json:"dns,omitempty" binding:"required"`
	MTU                 string `json:"mtu,omitempty" binding:"required"`
	PublicKey           string `json:"publickey,omitempty" binding:"required"`
	AllowedIPs          string `json:"allowedips,omitempty" binding:"required"`
	Endpoint            string `json:"endpoint,omitempty" binding:"required"`
	PersistentKeepalive string `json:"persistentkeepalive,omitempty" binding:"required"`
	PeerInfo            string `json:"peerinfo,omitempty" binding:"required"`
	WireguardStatus
}

type WireguardStatus struct {
	WireguardConfig  string `json:"wireguard_config,omitempty"`
	WireguardRunning string `json:"wireguard_running,omitempty"`
}

type BindingReturnedConfig struct {
	PeerInfo        string `json:"peerinfo"`
	FrpcConfig      string `json:"frpcconfig"`
	Udp2rawEndpoint int    `json:"udp2rawendpoint,omitempty"`
	Udp2rawServerIp string `json:"udp2rawserver,omitempty"`
	Udp2rawPasswd   string `json:"udp2rawpasswd,omitempty"`
}

type Errors struct {
	Returned
	Request
}

type PostRequest struct {
	Returned
	Request
	Interval int    `json:"interval,omitempty"`
	Domain   string `json:"domain"`
	Data     []struct {
		Address string `json:"address"`
		Result  string `json:"result"`
	} `json:"data,omitempty"`
	Key            string `json:"key,omitempty"`
	Networklaundry int    `json:"networklaundry,omitempty"`
	File           string `json:"file,omitempty"`
	SupRunning     string `json:"sup_running,omitempty"`
}

type StatusByte struct {
	Request
	Frpc      Status `json:"frpc"`
	Udp2Raw   Status `json:"udp2raw"`
	Oneclick  Status `json:"oneclick"`
	Wireguard Status `json:"wireguard"`
}

type Status struct {
	Status int `json:"status"`
}

// 8GczMyIG/emlMkC02Go9FRYtbvUagPDlQUKnsmMOHG8=
