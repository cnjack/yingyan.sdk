package yingyan_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cnjack/yingyan.sdk"
)

func TestS_EntitySearch(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntitySearch("", &yingyan.EntityListFilter{}, yingyan.BaiDuCoordType, 1, 100)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	bodyByte, _ := json.Marshal(body)
	fmt.Println(string(bodyByte))
}

func TestS_EntityBoundSearch(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntityBoundSearch(&yingyan.Bounds{Points: []yingyan.Point{
		{Latitude: 39.838434, Longitude: 116.472943},
		{Latitude: 39.848434, Longitude: 116.492943},
	}}, &yingyan.EntityListFilter{}, yingyan.BaiDuCoordType, 1, 100)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	bodyByte, _ := json.Marshal(body)
	fmt.Println(string(bodyByte))
}

func TestS_EntityRoundSearch(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntityRoundSearch(yingyan.Point{Latitude: 39.838434, Longitude: 116.472943}, 5000, &yingyan.EntityListFilter{}, yingyan.BaiDuCoordType, 1, 100)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	bodyByte, _ := json.Marshal(body)
	fmt.Println(string(bodyByte))
}
