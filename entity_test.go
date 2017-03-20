package yingyan_test

import (
	"fmt"
	"github.com/cnjack/yingyan.sdk"
	"testing"
)

func TestS_EntityAdd(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntityAdd("run_1", "")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_EntityUpdate(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntityUpdate("run_1", "我是一些描述")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_EntityDelete(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.EntityDelete("run_1")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)

}
