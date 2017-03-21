package yingyan

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type PointData struct {
	EntityName     string         `json:"entity_name"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	LocTime        int64          `json:"loc_time"`
	CoordTypeInput CoordTypeInput `json:"coord_type_input"` //坐标类型	 wgs84 GPS坐标 | gcj02 国测局加密坐标 | bd09ll 百度经纬度坐标
	Speed          float64        `json:"speed"`            //非必须
	Direction      int            `json:"direction"`        // 非必须 0 正北 顺时针
	Height         float64        `json:"height"`           // 非必须
	Radius         float64        `json:"radius"`           //非必须 定位精度，GPS或定位SDK返回的数值 定位精度，GPS或定位SDK返回的值
}

type CoordTypeInput string

var (
	GPSCoordType   CoordTypeInput = `wgs84`
	GcjCoordType   CoordTypeInput = `gcj02`
	BaiDuCoordType CoordTypeInput = `bd09ll`
)

func (serv *s) AddPoint(pointData *PointData) (r *CommonResp, err error) {
	param := map[string]string{
		"entity_name":      pointData.EntityName,
		"latitude":         strconv.FormatFloat(pointData.Latitude, 'f', 8, 64),
		"longitude":        strconv.FormatFloat(pointData.Longitude, 'f', 8, 64),
		"loc_time":         strconv.FormatInt(time.Now().Unix(), 10),
		"coord_type_input": string(pointData.CoordTypeInput),
	}
	if pointData.LocTime != 0 {
		param["latitude"] = strconv.FormatInt(pointData.LocTime, 10)
	}
	if pointData.Speed != 0 {
		param["speed"] = strconv.FormatFloat(pointData.Speed, 'f', 2, 64)
	}
	if pointData.Direction != 0 {
		param["direction"] = strconv.Itoa(pointData.Direction)
	}
	if pointData.Direction != 0 {
		param["height"] = strconv.FormatFloat(pointData.Height, 'f', 2, 64)
	}
	if pointData.Radius != 0 {
		param["radius"] = strconv.FormatFloat(pointData.Radius, 'f', 2, 64)
	}
	respByte, err := serv.Post(trackAddPoint, param)
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

func (serv *s) AddPoints(pointDatas []*PointData) (r *CommonResp, err error) {
	if len(pointDatas) > 100 {
		return nil, errors.New("too many points")
	}
	paramsData, err := json.Marshal(pointDatas)
	if err != nil {
		return nil, err
	}
	respByte, err := serv.Post(trackAddPoints, map[string]string{
		"point_list": string(paramsData),
	})
	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	fmt.Println(string(respByte))
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
