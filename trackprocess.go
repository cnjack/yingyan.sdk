package yingyan

import (
	"fmt"
	"encoding/json"
)

type Denoise int

var (
	NotNeedDenoise Denoise = 0
	NeedDenoise    Denoise = 1
)

type MapMatch int

var (
	NotNeedMapMatch MapMatch = 0
	NeedMapMatch    MapMatch = 1
)

type TransportMode string

var (
	DrivingTransportMode TransportMode = `driving`
	RidingTransportMode  TransportMode = `riding`
	WalkingTransportMode TransportMode = `walking`
)

type ProcessOption struct {
	Denoise         Denoise
	MapMatch        MapMatch
	RadiusThreshold int
	TransportMode   TransportMode
}

func (p *ProcessOption) ToData() string {
	if p.TransportMode == "" {
		return fmt.Sprintf("need_denoise=%d,need_mapmatch=%d,radius_threshold=%d", p.Denoise, p.MapMatch, p.RadiusThreshold)
	}
	return fmt.Sprintf("need_denoise=%d,need_mapmatch=%d,radius_threshold=%d,transport_mode=%s", p.Denoise, p.MapMatch, p.RadiusThreshold, p.TransportMode)
}

func (serv *s) GetLatestPoint(entityName string, po *ProcessOption, coordType CoordTypeInput) (r *LatestPointResp, err error) {
	fmt.Println(po.ToData())
	param := map[string]string{
		"entity_name": entityName,
		"process_option": po.ToData(),
		"coord_type_output":string(coordType),
	}
	respByte, err := serv.Get(trackGetLatestPoint, param)
	if err != nil {
		return nil, err
	}
	r = &LatestPointResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

