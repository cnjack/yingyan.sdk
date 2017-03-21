package yingyan

import "time"

type Time time.Time

const timeFormart = "2006-01-02 15:04:05"

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

type CommonResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type EntityResp struct {
	EntityName     string       `json:"entity_name"`
	EntityDesc     string       `json:"entity_desc"`
	ModifyTime     Time         `json:"modify_time"`
	CreateTime     Time         `json:"create_time"`
	LatestLocation *LatestPoint `json:"latest_point,omitempty"`
}

type EntitiesListResp struct {
	CommonResp
	Total    int           `json:"total"`
	Size     int64         `json:"size"`
	Entities []*EntityResp `json:"entities,omitempty"`
}

type LatestPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	LocTime   int64   `json:"loc_time"`
	Direction int64   `json:"direction"`
	Speed     float64 `json:"speed"`
	Height    float64 `json:"height"`
	Floor     string  `json:"floor"`
}

type LatestPointResp struct {
	CommonResp
	LatestPoint *LatestPoint `json:"latest_point,omitempty"`
}

type DistanceResp struct {
	CommonResp
	Distance float64 `json:"distance"`
}
