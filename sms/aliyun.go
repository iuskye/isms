// Copyright 2024 Simon Liu <iuskye@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache License Version 2.0 style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iuskye/isms.

package sms

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func SendSmsAliyun() {
	err := sendSmsAliyun(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

func sendSmsAliyun(args []*string) (_err error) {
	client, _err := CreateClient(tea.String(os.Getenv("SMS_ACCESS_KEY_ID")), tea.String(os.Getenv("SMS_ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	TemplateParamStr := "{\"time\":\"" + nowTime + "\",\"level\":\"" + os.Args[1] + "\",\"content\":\"" + os.Args[2] + "\"}"

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(os.Getenv("SMS_PHONES")),
		SignName:      tea.String(os.Getenv("SMS_ALIYUN_SIGN")),
		TemplateCode:  tea.String(os.Getenv("SMS_ALIYUN_TEMPLATE")),
		TemplateParam: tea.String(TemplateParamStr),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		// if m, ok := data.(map[string]interface{}); ok {
		// 	recommend, _ := m["Recommend"]
		// 	fmt.Println(recommend)
		// }
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
