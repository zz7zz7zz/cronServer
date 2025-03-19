package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ServerWebHook struct {
}

func (s ServerWebHook) OnWebHook() {
	//登录管理后台
	token, err := getToken()
	if err != nil {
		fmt.Println("登录失败:", err)
		return
	}
	fmt.Println("Step 1: ", token)

	//添加版本
	code, err := versionAdd(token, "2.21.4")
	if err != nil {
		fmt.Println("添加版本失败:", err)
		return
	}
	fmt.Println("Step 2: ", code)

	//设置状态
	code2, err2 := versioncontrol(token, 58)
	if err2 != nil {
		fmt.Println("设置状态失败:", err2)
		return
	}
	fmt.Println("Step 3: ", code2)
}

func getToken() (string, error) {
	// 构造请求数据
	requestData := map[string]interface{}{
		"username": "admin",
		"password": "123456", // 复杂数据可以序列化为JSON字符串
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	// 发送 POST 请求
	response, err := PostJSON("http://system-api.dev.moyuvedio.com/user/login", headers, requestData)
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
		"verno":         version,
		"Authorization": "Bearer " + token,
	}

	// 构造请求头
	headers := map[string]string{
		// "Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Bearer " + token,
	}
	response, err := PostJSON("http://system-api.dev.moyuvedio.com/app/version/add", headers, requestData)
	if err != nil {
		fmt.Println("请求失败:", err)
		return -1, err
	}

	code, err := GetNestedValue(response, "code")
	if err != nil {
		return -1, fmt.Errorf("字段类型错误")
	}

	fmt.Println("code: ", code)
	ret, ok := code.(int)
	if !ok {
		return -1, fmt.Errorf("最终值类型错误")
	}

	return ret, nil
}

func versioncontrol(token string, id int) (int, error) {
	// 构造请求数据
	requestData := map[string]interface{}{
		"id":            58,
		"control[1][1]": 1,
		"control[1][2]": 0,
		"control[1][3]": 1,
		"control[2][1]": 1,
		"control[2][2]": 0,
		"control[2][3]": 1,
	}

	// 发送 POST 请求
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	response, err := PostJSON2("http://system-api.dev.moyuvedio.com/app/version/control", headers, requestData)
	if err != nil {
		fmt.Println("请求失败:", err)
		return -1, err
	}

	code, err := GetNestedValue(response, "code")
	if err != nil {
		return -1, fmt.Errorf("字段类型错误")
	}

	fmt.Println("code: ", code)
	ret, ok := code.(int)
	if !ok {
		return -1, fmt.Errorf("最终值类型错误")
	}

	return ret, nil
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
func PostJSON(apiUrl string, headers map[string]string, data interface{}) (result map[string]interface{}, err error) {
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
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	// req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value) // 添加自定义请求头
		fmt.Println("key:", key, " value:", value)
	}

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
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return result, nil
}

func PostJSON2(apiUrl string, headers map[string]string, data interface{}) (result map[string]interface{}, err error) {
	// 1. 序列化请求数据为 JSON
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, fmt.Errorf("序列化 JSON 失败: %v", err)
	// }

	// 2. 创建 HTTP 请求

	params := url.Values{}
	for key, value := range data.(map[string]interface{}) {
		params.Add(key, value.(string))
	}

	fmt.Println("req_body: ", params.Encode())
	resp, err := http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))

	// body := bytes.NewBuffer(jsonData)
	// req, err := http.NewRequest("POST", apiUrl, body)
	// if err != nil {
	// 	return nil, fmt.Errorf("创建请求失败: %v", err)
	// }
	// // req.Header.Set("Content-Type", "application/json")
	// for key, value := range headers {
	// 	req.Header.Set(key, value) // 添加自定义请求头
	// 	fmt.Println("key:", key, " value:", value)
	// }

	// // 3. 发送请求（带超时控制）
	// client := &http.Client{Timeout: 10 * time.Second}
	// resp, err := client.Do(req)
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
	fmt.Println("rsp_body: ", string(body))
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return result, nil
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
