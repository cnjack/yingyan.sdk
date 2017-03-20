package yingyan_test

import (
	"fmt"

	"github.com/cnjack/yingyan.sdk"
)
import (
	"testing"
)

func TestS_EntityAdd(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ")
	body, err := client.EntityAdd(135928, "run_1", "")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_EntityUpdate(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ")
	body, err := client.EntityUpdate(135928, "run_1", "我是一些描述")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_EntityDelete(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ")
	body, err := client.EntityDelete(135928, "run_1")
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)

}