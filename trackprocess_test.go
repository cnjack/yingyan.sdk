package yingyan_test

import (
	"testing"
	"github.com/cnjack/yingyan.sdk"
	"fmt"
	"time"
)

func TestS_GetLatestPoint(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.GetLatestPoint("run_1", &yingyan.ProcessOption{}, yingyan.GcjCoordType)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_GetDistance(t *testing.T) {
	start := time.Now().Unix()-6300
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.GetDistance("run_1",true,start,time.Now().Unix(), &yingyan.ProcessOption{}, yingyan.NoSupplement)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}