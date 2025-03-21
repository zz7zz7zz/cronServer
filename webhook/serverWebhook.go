package webhook

import (
	"bytes"
	"cronServer/config"
	"cronServer/constant"
	"cronServer/models"
	"encoding/json"
	"fmt"
	_ "go/constant"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ServerWebHook struct {
}

func (s ServerWebHook) OnWebHook(appReviewRecord *models.AppReviewRecord) {
	//1.登录管理后台
	token, err := getToken()
	if err != nil {
		fmt.Println("登录失败:", err)
		return
	}
	fmt.Println("Step 1: ", token)

	//2.1添加版本
	var ver = appReviewRecord.Ver
	code, err := versionAdd(token, ver)
	if err != nil || code != 200 && code != 1021 {
		fmt.Println("添加版本失败:", err)
		return
	}
	fmt.Println("Step 2: ", code)

	//2.2拉取列表，根据版本获取id
	code2, err2 := versionList(token)
	if err2 != nil {
		fmt.Println("拉取列表失败:", err2)
		return
	}
	fmt.Println("Step 3: ", code2)

	id := -1
	if code2.Code == 200 {
		fmt.Println("Total: ", code2.Total)
		for _, item := range code2.Data {
			if item.Verno == ver {
				id = item.Id
				break
			}
			fmt.Println("Id: ", item.Id)
			fmt.Println("Verno: ", item.Verno)
			fmt.Println("VerVal: ", item.VerVal)
			fmt.Println("Created_at: ", item.Created_at)
			fmt.Println("Updated_at: ", item.Updated_at)
		}
	}
	fmt.Println("Id: ", id)

	//2.3拉取状态
	vd, err3 := versionDetail(token, id)
	if err3 != nil {
		fmt.Println("设置状态失败:", err3)
		return
	}
	fmt.Println("Step 4: ", vd)

	//2.3设置状态
	code4, err4 := versioncontrol(token, id, vd, appReviewRecord.Platform)
	if err4 != nil {
		fmt.Println("设置状态失败:", err4)
		return
	}
	fmt.Println("Step 5: ", code4)

	//2.5发送消息
	if code4 == 200 {
		hook := EnterpriseWechat{}
		hook.OnWebHook(appReviewRecord)
	}
}

func getToken() (string, error) {
	// 构造请求数据
	requestData := map[string]interface{}{
		"username": config.GConfig.Webhook.OurServer.Username,
		"password": config.GConfig.Webhook.OurServer.Password, // 复杂数据可以序列化为JSON字符串
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	// 发送 POST 请求
	response, err := PostJSON("POST", config.GConfig.Webhook.OurServer.Referer+"/user/login", headers, requestData)
	if err != nil {
		fmt.Println("请求失败:", err)
		return "", err
	}

	// 1. 获取键对应的值
	// data, ok1 := response["data"].(map[string]interface{})
	// if !ok1 {
	// 	panic("data 字段类型错误")
	// }

	// // 2. 类型断言转换为 map[string]interface{}
	// token, ok3 := data["token"].(string)
	// if !ok3 {
	// 	panic("email 字段类型错误")
	// }

	token, err := GetNestedValue(response, "data", "token")
	if err != nil {
		panic(err)
	}

	ret, ok := token.(string)
	if !ok {
		panic("最终值类型错误")
	}

	return ret, nil
}

func versionAdd(token string, version string) (int, error) {
	// 构造请求数据
	requestData := map[string]interface{}{
		"verno": version,
	}

	// 构造请求头
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}

	response, err := PostJSON("POST", config.GConfig.Webhook.OurServer.URL+"/mq/app/version/add", headers, requestData)
	if err != nil {
		fmt.Println("请求失败:", err)
		return -1, err
	}

	code, err := GetNestedValue(response, "code")
	if err != nil {
		return -1, fmt.Errorf("字段类型错误")
	}

	ret, ok := code.(int)
	if ok {
		return ret, nil
	}

	ret2, ok2 := code.(float64)
	if ok2 {
		ret = int(ret2) // 显式转换为int
	} else {
		return -1, fmt.Errorf("最终值类型错误")
	}

	return ret, nil
}

func versionDetail(token string, id int) (*VersionDetail, error) {

	// 发送 POST 请求
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}
	body, statusCode, err := GetJSON(config.GConfig.Webhook.OurServer.URL+"/mq/app/version/detail?id="+strconv.Itoa(id), headers)
	fmt.Printf("Status: %d\nResponse: %s\n", statusCode, body)
	if err != nil {
		fmt.Println("请求失败:", statusCode, err)
		return nil, err
	}

	var versionDetailet VersionDetail
	if err := json.Unmarshal(body, &versionDetailet); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return &versionDetailet, nil
}

func versioncontrol(token string, id int, vd *VersionDetail, platform string) (int, error) {
	//TODO 要与原来的状态进行合并
	// 构造请求数据
	requestData := map[string]interface{}{
		"id": id,
	}
	if platform == constant.Ios {
		requestData["control"] = map[string]map[string]int{
			"1": {
				"1": vd.Data.Control["1"]["1"],
				"2": 1,
				"3": vd.Data.Control["1"]["3"],
			},
			"2": {
				"1": vd.Data.Control["2"]["1"],
				"2": 1,
				"3": vd.Data.Control["2"]["3"],
			},
		}
	} else if platform == constant.Android {
		requestData["control"] = map[string]map[string]int{
			"1": {
				"1": 1,
				"2": vd.Data.Control["1"]["2"],
				"3": 1,
			},
			"2": {
				"1": 1,
				"2": vd.Data.Control["2"]["2"],
				"3": 1,
			},
		}
	}

	// 发送 POST 请求
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}
	response, err := PostJSON("POST", config.GConfig.Webhook.OurServer.URL+"/mq/app/version/control", headers, requestData)
	if err != nil {
		fmt.Println("请求失败:", err)
		return -1, err
	}

	code, err := GetNestedValue(response, "code")
	if err != nil {
		return -1, fmt.Errorf("字段类型错误")
	}

	ret, ok := code.(int)
	if ok {
		return ret, nil
	}

	ret2, ok2 := code.(float64)
	if ok2 {
		ret = int(ret2) // 显式转换为int
	} else {
		return -1, fmt.Errorf("最终值类型错误")
	}

	return ret, nil
}

func versionList(token string) (*VersionList, error) {

	// 发送 POST 请求
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}
	body, statusCode, err := GetJSON(config.GConfig.Webhook.OurServer.URL+"/mq/app/version/lists?page=1&size=20", headers)
	fmt.Printf("Status: %d\nResponse: %s\n", statusCode, body)
	if err != nil {
		fmt.Println("请求失败:", statusCode, err)
		return nil, err
	}

	var version VersionList
	if err := json.Unmarshal(body, &version); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return &version, nil
}

// PostJSON 发送 JSON 格式的 POST 请求，并解析返回的 JSON 数据
// 参数:
//
//	apiUrl: 请求的 URL
//	data:   要发送的 JSON 数据 (可以是 map、结构体等)
//
// 返回值:
//
//	result: 解析后的 JSON 数据 (map[string]interface{})
//	err:    错误信息
func PostJSON(method string, apiUrl string, headers map[string]string, data interface{}) (result map[string]interface{}, err error) {
	fmt.Printf("\n---------------%v\n", apiUrl)
	// 1. 序列化请求数据为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	// 2. 创建 HTTP 请求
	// if headers["Content-Type"] == "application/x-www-form-urlencoded" {
	// 	params := url.Values{}
	// 	for key, value := range data.(map[string]interface{}) {
	// 		params.Add(key, value.(string))
	// 	}
	// 	http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	// }
	req, err := http.NewRequest(method, apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value) // 添加自定义请求头
		fmt.Println("key:", key, " value:", value)
	}

	fmt.Printf("req: %v\n", req)

	// 3. 发送请求（带超时控制）
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求发送失败: %v", err)
	}
	defer resp.Body.Close()

	// 4. 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("服务器返回非 200 状态码: %d", resp.StatusCode)
	}

	// 5. 读取并解析响应 JSON
	body, err := io.ReadAll(resp.Body)

	fmt.Printf("\n")
	fmt.Printf("resp: %v\n", string(body))
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return result, nil
}

func GetJSON(apiUrl string, headers map[string]string) ([]byte, int, error) {
	fmt.Printf("\n---------------%v\n", apiUrl)
	// 1. 创建请求对象
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("创建请求失败: %w", err)
	}

	// 2. 设置请求头
	setHeaders(req, headers)

	// 3. 创建带超时的 Client
	client := &http.Client{
		Timeout: 30 * time.Second, // 可自定义超时时间
	}

	// 4. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("请求发送失败: %w", err)
	}
	defer resp.Body.Close()

	// 5. 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("读取响应失败: %w", err)
	}

	return body, resp.StatusCode, nil
}

// 设置请求头私有方法
func setHeaders(req *http.Request, headers map[string]string) {
	if headers == nil {
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// 安全提取嵌套值 (支持任意层级的 map[string]interface{})
func GetNestedValue(data map[string]interface{}, keys ...string) (interface{}, error) {
	current := data
	for i, key := range keys {
		value, exists := current[key]
		if !exists {
			return nil, fmt.Errorf("键 '%s' 不存在", key)
		}

		// 如果是最后一层，直接返回值
		if i == len(keys)-1 {
			return value, nil
		}

		// 继续深入下一层
		nextMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("键 '%s' 的值不是 map[string]interface{}", key)
		}
		current = nextMap
	}
	return nil, fmt.Errorf("无效的键路径")
}
