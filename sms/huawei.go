// Copyright 2024 Simon Liu <iuskye@foxmail.com>. All rights reserved.
// Use of this source code is governed by a Apache License Version 2.0 style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iuskye/isms.

package sms

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	hcore "huaweicloud.com/apig/signer"
)

func SendSmsHuawei() {
	appInfo := hcore.Signer{
		Key:    os.Getenv("SMS_ACCESS_KEY_ID"),     //App Key
		Secret: os.Getenv("SMS_ACCESS_KEY_SECRET"), //App Secret
	}
	// APP接入地址(在控制台"应用管理"页面获取)+接口访问 URI
	apiAddress := "https://smsapi.cn-north-4.myhuaweicloud.com:443/sms/batchSendSms/v1"
	
	// 国内短信签名通道号
	sender := os.Getenv("SMS_HUAWEI_SENDER")
	//模板 ID
	templateId := os.Getenv("SMS_HUAWEI_TEMPLATE")
	//签名名称
	signature := os.Getenv("SMS_HUAWEI_SIGN")

	//必填,全局号码格式(包含国家码),示例:+86151****6789,多个号码之间用英文逗号分隔
	receiver := os.Getenv("SMS_PHONES") //短信接收人号码

	//选填,短信状态报告接收地址,推荐使用域名,为空或者不填表示不接收状态报告
	statusCallBack := ""

	/*
	 * 选填,使用无变量模板时请赋空值 string templateParas = "";
	 * 单变量模板示例:模板内容为"您的验证码是${1}"时,templateParas可填写为"[\"369751\"]"
	 * 双变量模板示例:模板内容为"您有${1}件快递请到${2}领取"时,templateParas可填写为"[\"3\",\"人民公园正门\"]"
	 * 模板中的每个变量都必须赋值，且取值不能为空
	 * 查看更多模板规范和变量规范:产品介绍>短信模板须知和短信变量须知
	 */
	//模板变量，此处以单变量验证码短信为例，请客户自行生成6位验证码，并定义为字符串类型，以杜绝首位0丢失的问题（例如：002569变成了2569）。

	// 注意，华为云短信网关模板变量如果同时包含日期和时间，需要设置为两个变量
	nowDate := time.Now().Format("2006-01-02")
	nowTime := time.Now().Format("15:04:05")
	templateParas := "[\"" + nowDate + "\",\"" + nowTime + "\",\"" + os.Args[1] + "\",\"" + os.Args[2] + "\"]"

	body := buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature)
	resp, err := post(apiAddress, []byte(body), appInfo)

	if err != nil {
		return
	}
	fmt.Println(resp)
}

/**
 * sender,receiver,templateId不能为空
 */
func buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature string) string {
	param := "from=" + url.QueryEscape(sender) + "&to=" + url.QueryEscape(receiver) + "&templateId=" + url.QueryEscape(templateId)
	if templateParas != "" {
		param += "&templateParas=" + url.QueryEscape(templateParas)
	}
	if statusCallBack != "" {
		param += "&statusCallback=" + url.QueryEscape(statusCallBack)
	}
	if signature != "" {
		param += "&signature=" + url.QueryEscape(signature)
	}
	return param
}

func post(url string, param []byte, appInfo hcore.Signer) (string, error) {
	if param == nil || appInfo == (hcore.Signer{}) {
		return "", nil
	}

	// 代码样例为了简便，设置了不进行证书校验，请在商用环境自行开启证书校验。
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if err != nil {
		return "", err
	}

	// 对请求增加内容格式，固定头域
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 对请求进行HMAC算法签名，并将签名结果设置到Authorization头域。
	appInfo.Sign(req)

	fmt.Println(req.Header)
	// 发送短信请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// 获取短信响应
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
