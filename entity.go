package yingyan

import (
	"encoding/json"
)

func (serv *s) EntityAdd(entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
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

func (serv *s) EntityUpdate(entityName, entityDesc string) (r *CommonResp, err error) {
	param := map[string]string{
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

func (serv *s) EntityDelete(entityName string) (r *CommonResp, err error) {
	param := map[string]string{
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
