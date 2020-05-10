package configs

type server struct {
	RpcPort        string `json:"rpcPort"`
	RestPort       string `json:"restPort"`
	RequestTimeout int    `json:"requestTimeout"`
}

type node struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Fellows  []string `json:"fellows"`
}

type storage struct {
	BaseLocation string `json:"baseLocation"`
}