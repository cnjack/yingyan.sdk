package yingyan_test

import (
	"encoding/json"
	"fmt"
	"github.com/cnjack/yingyan.sdk"
	"testing"
	"time"
)

func TestS_Staypoint(t *testing.T) {
	start := time.Now().Unix() - 6300
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.StayPoint("run_1", start, time.Now().Unix(), 100, &yingyan.ProcessOption{}, yingyan.BaiDuCoordType)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	bodyByte, _ := json.Marshal(body)
	fmt.Println(string(bodyByte))
}
