package yingyan

type CommonResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LatestPointResp struct {
	CommonResp
	LatestPoint struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		LocTime   int64   `json:"loc_time"`
		Direction int64   `json:"direction"`
		Speed     float64 `json:"speed"`
		Height    float64 `json:"height"`
		Floor     string  `json:"floor"`
	} `json:"latest_point,omitempty"`
}
