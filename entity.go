package yingyan

import (
	"encoding/json"
	"strconv"
)

func (serv *s) EntityAdd(serviceID int, entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
		"service_id":  strconv.Itoa(serviceID),
		"entity_name": entityName,
		"entity_desc": entityDesc,
	}
	respByte, err := serv.Post(entityAddPath, param)
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

func (serv *s) EntityUpdate(serviceID int, entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
		"service_id":  strconv.Itoa(serviceID),
		"entity_name": entityName,
		"entity_desc": entityDesc,
	}
	respByte, err := serv.Post(entityUpdatePath, param)
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

func (serv *s) EntityDelete(serviceID int, entityName string) (r *CommonResp, err error) {
	param := map[string]string{
		"service_id":  strconv.Itoa(serviceID),
		"entity_name": entityName,
	}
	respByte, err := serv.Post(entityDeletePath, param)
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
