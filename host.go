package go_zabbix

import "fmt"

//创建主机接口参数类型Interfaces
type Interfaces struct {
	Type  int    `json:"type"`
	Main  int    `json:"main"`
	Useip int    `json:"useip"`
	IP    string `json:"ip"`
	Dns   string `json:"dns"`
	Port  string `json:"port"`
}

//创建主机接口参数类型 group数组
type Groups struct {
	Groupid string `json:"groupid"`
}

//创建主机接口参数类型 模板数组
type Templates struct {
	Templateid string `json:"templateid"`
}

//创建主机参数
type HostCreateParames struct {
	Host       string        `json:"host"`
	Interfaces []interface{} `json:"interfaces"`
	Groups     []interface{} `json:"groups"`
	Templates  []interface{} `json:"templates,omitempty"`
}

type HostCreateResp struct {
	HostIds []string `json:"hostids,omitempty"`
}

//创建主机
func (c *Session) CreateHost(params *HostCreateParames) (resp HostCreateResp, err error) {
	//resp := make([]interface{},1)
	err = c.OriginSession.Get("host.create", params, &resp)
	fmt.Printf("params is %v\n", *params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
