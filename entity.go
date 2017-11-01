package yingyan

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (c *Client) EntityAdd(entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
		"entity_name": entityName,
		"entity_desc": entityDesc,
	}
	respByte, err := c.Post(entityAddPath, param)
	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) EntityUpdate(entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
		"entity_name": entityName,
		"entity_desc": entityDesc,
	}
	respByte, err := c.Post(entityUpdatePath, param)
	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) EntityDelete(entityName string) (r *CommonResp, err error) {
	param := map[string]string{
		"entity_name": entityName,
	}
	respByte, err := c.Post(entityDeletePath, param)
	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type EntityListFilter struct {
	EntityNames  []string
	ActiveTime   int64 //该时间之后活跃的用户
	InactiveTime int64 //该时间之后不活跃的用户
}

func (f *EntityListFilter) ToData() string {
	param := map[string]string{}
	names := strings.Join(f.EntityNames, ",")
	if names != "" {
		param["entity_names"] = names
	}
	//有且只有一个Time
	if f.ActiveTime != 0 {
		param["active_time"] = strconv.FormatInt(f.ActiveTime, 10)
	} else if f.InactiveTime != 0 {
		param["inactive_time"] = strconv.FormatInt(f.InactiveTime, 10)
	}
	var params []string
	for k, v := range param {
		params = append(params, k+"="+v)
	}
	return strings.Join(params, "|")
}

func (c *Client) EntityList(filter *EntityListFilter, coordType CoordType, pageIndex, pageSize int) (r *EntitiesListResp, err error) {
	if coordType == "" {
		coordType = BaiDuCoordType
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 100
	}
	param := map[string]string{
		"coord_type_output": string(coordType),
		"page_index":        strconv.Itoa(pageIndex),
		"page_size":         strconv.Itoa(pageSize),
	}
	filterString := filter.ToData()
	if filterString != "" {
		param["filter"] = filterString
	}
	respByte, err := c.Get(entityList, param)
	if err != nil {
		return nil, err
	}
	r = &EntitiesListResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
