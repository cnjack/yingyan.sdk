package yingyan_test

import (
	"testing"
	"github.com/cnjack/yingyan.sdk"
	"fmt"
)

func TestS_GetLatestPoint(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.GetLatestPoint("run_1", &yingyan.ProcessOption{}, yingyan.GcjCoordType)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}