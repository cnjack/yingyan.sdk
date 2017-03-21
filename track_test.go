package yingyan_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cnjack/yingyan.sdk"
)

func TestS_AddPoint(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.AddPoint(&yingyan.PointData{
		EntityName: "run_1",
		Latitude:   39.848434,
		Longitude:  116.492943,
		CoordType:  yingyan.BaiDuCoordType,
	})
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_AddPoints(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	params := []*yingyan.PointData{}
	params = append(params, &yingyan.PointData{
		EntityName: "run_1",
		Latitude:   39.845434,
		LocTime:    time.Now().Unix(),
		Longitude:  116.442943,
		CoordType:  yingyan.BaiDuCoordType,
	})
	params = append(params, &yingyan.PointData{
		EntityName: "run_1",
		Latitude:   39.845534,
		LocTime:    time.Now().Unix() - 60,
		Longitude:  116.495343,
		CoordType:  yingyan.BaiDuCoordType,
	})
	body, err := client.AddPoints(params)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}
