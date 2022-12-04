package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type ResponseBody struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func BindJSON(r *http.Request, obj interface{}) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, obj)
}
func getByPage(start string, end string, limit string, page string, arr []struct{}) ([]struct{}, error) {
	st, errST := strconv.Atoi(start)
	ed, errED := strconv.Atoi(end)
	lm, errLM := strconv.Atoi(limit)
	pg, errPG := strconv.Atoi(page)
	if errST != nil {
		st = 0
	}
	if errED != nil {
		ed = len(arr)
	}
	if errLM != nil {
		lm = len(arr)
	}
	if errPG != nil {
		pg = 1
	}
	divideResult := len(arr) / lm
	surplus := len(arr) % lm
	if surplus == 0 && pg > divideResult || surplus != 0 && pg > divideResult+1 {
		return nil, errors.New("don't have record")
	}
	if surplus != 0 && pg == divideResult+1 {
		return arr[lm*(pg-1) : ed], nil
	}
	return arr[st:ed], nil
}
