package go_zabbix

import (
	"fmt"
	"github.com/cavaliercoder/go-zabbix"
)

// 查询监控项
type Item struct {
	// HostID is the unique ID of the Host.
	HostID int

	// ItemID is the unique ID of the Item.
	ItemID int

	// Itemname is the technical name of the Item.
	ItemName string

	//Key_ is the key_ of the Item
	Key_ string

	// ItemDescr is the description of the Item.
	ItemDescr string

	// LastClock is the last Item epoh time.
	LastClock int

	// LastValue is the last value of the Item.
	LastValue string

	// LastValueType is the type of LastValue
	// 0 - float; 1 - text; 3 - int;
	LastValueType int
}

func (c *Session) GetItems(params zabbix.ItemGetParams) ([]Item, error) {
	items := make([]jItem, 0)
	err := c.InerSession.Get("item.get", params, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, zabbix.ErrNotFound
	}
	// map JSON Events to Go Events
	out := make([]Item, len(items))
	for i, jitem := range items {
		item, err := jitem.Item()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Item %d in response: %v", i, err)
		}
		out[i] = *item
	}

	return out, nil
}

type ItemCreateParams struct {
	zabbix.GetParameters

	//监控项名称
	Name string `json:"name,omitempty"`

	Key_ string `json:"key_,omitempty"`

	// HostID filters search results to items belong to the
	// given Host ID's.
	HostID string `json:"hostid,omitempty"`

	//数据类型
	Type int `json:"type,omitempty"`

	//数据值类型
	ValueType int `json:"value_type"`

	// InterfaceID filters search results to items that use
	// the given host Interface ID's.
	InterfaceID string `json:"interfaceids,omitempty"`

	Delay string `json:"delay,omitempty"`

	// GraphIDs filters search results to items that are used
	// in the given graph ID's.
	GraphIDs []string `json:"graphids,omitempty"`

	// TriggerIDs filters search results to items that are used
	// in the given Trigger ID's.
	TriggerIDs []string `json:"triggerids,omitempty"`

	// ApplicationIDs filters search results to items that
	// belong to the given Applications ID's.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// WebItems flag includes web items in the result.
	WebItems bool `json:"webitems,omitempty"`

	// Inherited flag return only items inherited from a template
	// if set to 'true'.
	Inherited bool `json:"inherited,omitempty"`

	// Templated flag return only items that belong to templates
	// if set to 'true'.
	Templated bool `json:"templated,omitempty"`

	// Monitored flag return only enabled items that belong to
	// monitored hosts if set to 'true'.
	Monitored bool `json:"monitored,omitempty"`

	// Group filters search results to items belong to a group
	// with the given name.
	Group string `json:"group,omitempty"`

	// Host filters search results to items that belong to a host
	// with the given name.
	Host string `json:"host,omitempty"`

	// Application filters search results to items that belong to
	// an application with the given name.
	Application string `json:"application,omitempty"`

	// WithTriggers flag return only items that are used in triggers
	WithTriggers bool `json:"with_triggers,omitempty"`
}

type ItemResponse struct {
	ItemIds []string `json:"itemids,omitempty"`
}

//创建监控项函数
func (c *Session) ItemCreate(params []ItemCreateParams) (resp ItemResponse, err error) {
	err = c.InerSession.Get("item.create", params, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

//删除监控项
func (c *Session) ItemDelete(itemids []string) (resp ItemResponse, err error) {
	params := map[string][]string{
		"params": itemids,
	}
	err = c.InerSession.Get("item.delete", params, &resp)
	if err != nil {
		return resp, nil
	}
	return resp, nil
}
