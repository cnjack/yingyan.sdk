package yingyan

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (serv *s) EntitySearch(query string, filter *EntityListFilter, coordType CoordTypeInput, pageIndex, pageSize int) (r *EntitiesListResp, err error) {
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
	if query != "" {
		param["query"] = query
	}
	filterString := filter.ToData()
	if filterString != "" {
		param["filter"] = filterString
	}
	respByte, err := serv.Get(entitySearch, param)
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

type Point struct {
	Latitude  float64
	Longitude float64
}

type Bounds struct {
	Points []Point
}

func (b *Bounds) ToData() (string, error) {
	if len(b.Points) != 2 {
		return "", errors.New("point is not enough")
	}
	return fmt.Sprintf("%f,%f;%f,%f", b.Points[0].Latitude, b.Points[0].Longitude, b.Points[1].Latitude, b.Points[1].Longitude), nil
}

// 根据矩形地理范围搜索entity，并返回实时位置
func (serv *s) EntityBoundSearch(bounds *Bounds, filter *EntityListFilter, coordType CoordTypeInput, pageIndex, pageSize int) (r *EntitiesListResp, err error) {
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
	param["bounds"], err = bounds.ToData()
	if err != nil {
		return
	}
	filterString := filter.ToData()
	if filterString != "" {
		param["filter"] = filterString
	}
	respByte, err := serv.Get(entityBoundSearch, param)
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

func (serv *s) EntityRoundSearch(center Point, radius int, filter *EntityListFilter, coordType CoordTypeInput, pageIndex, pageSize int) (r *EntitiesListResp, err error) {
	if coordType == "" {
		coordType = BaiDuCoordType
	}
	if radius < 1 || radius > 5000 {
		return nil, errors.New("param error")
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
		"center":            fmt.Sprintf("%f,%f", center.Latitude, center.Longitude),
		"radius":            strconv.Itoa(radius),
	}
	filterString := filter.ToData()
	if filterString != "" {
		param["filter"] = filterString
	}
	respByte, err := serv.Get(entityAroundSearch, param)
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
