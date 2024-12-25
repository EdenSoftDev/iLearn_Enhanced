package client

import (
	"encoding/json"
	"fmt"
	"iLearn_Enhanced/model"
	"iLearn_Enhanced/utils"
	"io"
	"net/http"
	"strings"
)

func GetILearnLt() (string, string, error) {
	time := utils.GetTime()
	resp, err := http.Get(fmt.Sprintf("https://ilearn.jlu.edu.cn/cas-server/login?service=https://ilearntec.jlu.edu.cn/&get-lt=true&n=%s&callback=jsonpcallback&_=%s", time, time))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	lt, execution, err := parseILearnLt(resp.Body)
	if err != nil {
		return "", "", err
	}

	return lt, execution, nil
}

func parseILearnLt(body io.Reader) (string, string, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return "", "", err
	}

	content := string(data)

	content = strings.TrimPrefix(content, "jsonpcallback(")
	content = strings.TrimSuffix(content, ");")

	var response model.LtResponse
	err = json.Unmarshal([]byte(content), &response)
	if err != nil {
		return "", "", err
	}

	return response.Lt, response.Execution, nil
}
