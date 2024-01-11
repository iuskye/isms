// Copyright 2024 Simon Liu <iuskye@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache License Version 2.0 style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iuskye/isms.

package main

import (
	"fmt"
	"os"

	"github.com/iuskye/isms/sms"
)

func main() {
	// 设置短信服务提供商
	provider := os.Getenv("SMS_PROVIDER")

	// 调用短信发送函数
	switch provider {
	// https://api.aliyun.com/api/Dysmsapi/2017-05-25/SendSms?tab=DEMO&lang=GO
	case "aliyun":
		sms.SendSmsAliyun()
	// https://support.huaweicloud.com/devg-msgsms/sms_04_0017.html
	case "huawei":
		sms.SendSmsHuawei()
	// https://cloud.tencent.com/document/product/382/43199
	case "tencent":
		sms.SendSmsTencent()
	// 添加其他短信服务提供商的处理...
	default:
		fmt.Println("Unsupported SMS provider:", provider)
	}
}
