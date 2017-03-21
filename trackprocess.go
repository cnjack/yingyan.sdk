package yingyan

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Denoise int

const (
	NotNeedDenoise Denoise = 0
	NeedDenoise    Denoise = 1
)

type MapMatch int

const (
	NotNeedMapMatch MapMatch = 0
	NeedMapMatch    MapMatch = 1
)

type TransportMode string

const (
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
	param := map[string]string{
		"entity_name":       entityName,
		"process_option":    po.ToData(),
		"coord_type_output": string(coordType),
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

type SupplementMode string

const (
	NoSupplement       SupplementMode = `no_supplement`
	StraightSupplement SupplementMode = `straight`
	DrivingSupplement  SupplementMode = `driving`
	RidingSupplement   SupplementMode = `riding`
	WalkingSupplement  SupplementMode = `walking`
)

func (serv *s) GetDistance(entityName string, isProcessed bool, startTime, endTime int64, po *ProcessOption, supplementMode SupplementMode) (r *DistanceResp, err error) {
	if supplementMode == "" {
		supplementMode = NoSupplement
	}
	isProcessedString := "0"
	if isProcessed {
		isProcessedString = "1"
	}
	param := map[string]string{
		"entity_name":     entityName,
		"is_processed":    isProcessedString,
		"start_time":      strconv.FormatInt(startTime, 10),
		"end_time":        strconv.FormatInt(endTime, 10),
		"process_option":  po.ToData(),
		"supplement_mode": string(supplementMode),
	}
	respByte, err := serv.Get(trackGetLatestPoint, param)
	if err != nil {
		return nil, err
	}
	r = &DistanceResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
