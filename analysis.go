package yingyan

import (
	"strconv"
	"encoding/json"
)

func (c *Client) StayPoint(entityName string, startTime, endTime, stayTime int64, po *ProcessOption, coordType CoordType) (r *StayPointResp, err error) {
	if coordType == "" {
		coordType = BaiDuCoordType
	}
	if stayTime == 0 {
		stayTime = 600
	}
	param := map[string]string{
		"entity_name":       entityName,
		"start_time":        strconv.FormatInt(startTime, 10),
		"end_time":          strconv.FormatInt(endTime, 10),
		"stay_time":         strconv.FormatInt(stayTime, 10),
		"process_option":    po.ToData(),
		"coord_type_output": string(coordType),
	}
	respByte, err := c.Get(trackGetLatestPoint, param)
	if err != nil {
		return nil, err
	}
	r = &StayPointResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
