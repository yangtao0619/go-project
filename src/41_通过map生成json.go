// 41_通过map生成json
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	infoMap := make(map[string]interface{}, 5)
	infoMap["company"] = "iFlytek"
	infoMap["score"] = []float64{87.9, 87.5, 90.5}
	infoMap["name"] = "Mike"

	result, err := json.MarshalIndent(infoMap, "", " ")
	if err != nil {
		fmt.Println("json生成出错, ", err)
		return
	}

	fmt.Println("infomap json is ", string(result))
}
