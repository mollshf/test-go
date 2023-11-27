package go_test_module

import (
	"encoding/json"
	"fmt"
)

func saveMyData(id, data string) (Data, error) {
	fmt.Println(id == "" && data == "")
	if id == "" && data == "" {
		return Data{}, ValidationError{"ValidationError", 400, "validasi gk gtu cok", Data{}}
	}
	if id == data {
		return Data{}, &NotFoundError{"NotFoundError", 404, "gk ada datanya cui", Data{}}
	}
	return Data{id, data}, nil
}
func main() {
	result, err := saveMyData("12", "12")
	fmt.Println("hasil:", result)
	fmt.Println("err:", err)
}

// Struct Data
type Data struct {
	Id   string
	Nama string
}

// Struct ValidationError and Method Error
type ValidationError struct {
	Name    string `json:"name"`
	Code    int16  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"Data"`
}

func (v ValidationError) Error() string {
	errMsg, _ := json.Marshal(v)
	return string(errMsg)
}

// Struct NotFoundError and Method Error
type NotFoundError struct {
	Name    string `json:"name"`
	Code    int16  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"Data"`
}

func (n *NotFoundError) Error() string {
	errMsg, _ := json.Marshal(n)
	return string(errMsg)
}
