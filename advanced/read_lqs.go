package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func json2Map(jsonStr string) {

	// var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic("json format error")
	}
	fmt.Println(js)
	if code, err := js.Get("result").Int(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(code)
	}

	res, err := js.Get("itemList").Array()
	fmt.Println(typeof(res))
	fmt.Println(res)

	// for index := 0; index < len(res); index++ {
	resMap := res[0]
	md, _ := resMap.(map[string]interface{})
	fmt.Println(typeof(resMap))
	fmt.Println(md["creationDate"], md["listType"])
	fmt.Println(typeof(md["creationDate"]))
	listtype, _ := md["listType"].(json.Number).Int64()
	fmt.Println(listtype)
	if listtype == 21 {
		fmt.Println("ok")
	}

}

func main() {

	url := "http://lqs.ppdapi.com/loanService/getLoans"

	payload := strings.NewReader("{\"applicationId\": \"11050001\",\"sceneKey\": \"pata-ding\", \"borrowerIds\": [42631337]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(typeof(body))
	fmt.Println(typeof(string(body)))
	jsonStr := string(body)

	// json2map
	json2Map(jsonStr)

}
