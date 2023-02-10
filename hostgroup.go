package go_zabbix

import "fmt"

type HostgroupCreateParams struct {
	Name string `json:"name"`
}

type HostgroupCreateResponse struct {
	GroupIds []string `json:"groupids"`
}

//创建hostGroup
func (c *Session) CreateHostgroup(hostgroupName string) (resp HostgroupCreateResponse, err error) {
	params := HostgroupCreateParams{
		Name: hostgroupName,
	}

	err = c.OriginSession.Get("hostgroup.create", params, resp)
	if err != nil {
		fmt.Printf("创建hostgroup 出错，错误原因为%s\n", err.Error())
		return HostgroupCreateResponse{}, err
	}
	return resp, nil
}
