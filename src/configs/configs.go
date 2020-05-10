package configs

type server struct {
	RpcPort        string `json:"rpcPort"`
	RestPort       string `json:"restPort"`
	RequestTimeout int    `json:"requestTimeout"`
}

type node struct {
	Name    string   `json:"name"`
	Fellows []string `json:"fellows"`
}

type storage struct {
	BaseLocation string `json:"baseLocation"`
}

type events struct {
	RequestArrival        string
	RequestExecutionBegin string
	RequestExecuted       string
}
