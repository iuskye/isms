![isms](https://ossnote.iuskye.com/notesimg/isms.png)

## 工具说明

通知类短信命令行发送工具，支持阿里云、华为云、腾讯云三家公有云厂商的短信网关。

## 使用说明

### 申请短信模板

#### 阿里云

必要信息：

- AccessKey ID
- AccessKey Secret
- 签名名称
- 模板 ID

模板内容：

```shell
XXX业务告警
告警时间：${time}
告警级别：${level}
告警信息：${content}
```

变量说明：

- ${time}：日期时间，例如：2024-01-09 10:00:01
- ${level}：告警级别，例如：L1
- ${content}：告警内容，例如：服务 MySQL 异常，请及时处理

#### 华为云

必要信息：

- AccessKey ID
- AccessKey Secret
- 签名名称
- 签名通道号
- 模板 ID

模板内容：

```shell
XXX业务告警
告警日期：${1}
告警时间：${2}
告警级别：${3}
告警信息：${4}
```

变量说明：

> 注意：变量值第一个字符不能是英文，不能包含 “.”、“。”、“ ' ”、“<”、“>”、“{”、“}”、“-” 等字符，可以有中文逗号

- ${1}：日期，例如：2024-01-09，华为云要求将日期和时间拆分为两个变量
- ${2}：时间，例如：10:00:01，华为云要求将日期和时间拆分为两个变量
- ${3}：告警级别，例如：L1
- ${4}：告警内容，例如：服务 MySQL 异常，请及时处理

错误码参考：https://support.huaweicloud.com/api-msgsms/sms_05_0050.html 。

#### 腾讯云

必要信息：

- AccessKey ID
- AccessKey Secret
- 应用 ID
- 签名名称
- 模板 ID

模板内容（注意不要有 `$` 符号）：

```shell
XXX业务告警
告警时间：{1}
告警级别：{2}
告警信息：{3}
```

变量说明：

- {1}：日期时间，例如：2024-01-09 10:00:01
- {2}：告警级别，例如：L1
- {3}：告警内容，例如：服务 MySQL 异常，请及时处理

### 修改配置文件

```shell
vi conf/isms.env

## i.1 短信服务商，支持 aliyun, huawei, tencent
export SMS_PROVIDER=''
## i.2 短信服务商提供的 AccessKey ID
export SMS_ACCESS_KEY_ID=''
## i.3 短信服务商提供的 AccessKey Secret
export SMS_ACCESS_KEY_SECRET=''
## i.4 短信接收者，全局号码格式(包含国家码)，多个号码使用英文逗号隔开，例如：'+8613051830000,+8613153260000'
export SMS_PHONES=''
## a.1 阿里云短信签名名称
export SMS_ALIYUN_SIGN=''
## a.2 阿里云短信模板 Code 码
export SMS_ALIYUN_TEMPLATE=''
## h.1 华为云短信签名名称
export SMS_HUAWEI_SIGN=''
## h.2 华为云国内短信签名通道号
export SMS_HUAWEI_SENDER=''
## h.3 华为云短信模板 ID
export SMS_HUAWEI_TEMPLATE=''
## t.1 腾讯云短信签名名称
export SMS_TENCENT_SIGN=''
## t.2 腾讯云短信应用 ID
export SMS_TENCENT_APPID=''
## t.3 腾讯云短信模板 ID
export SMS_TENCENT_TEMPLATE=''
```

### 发送短信

```shell
source conf/isms.env
go run main.go "params1" "params2"
```

### 编译可执行程序

```shell
cd scripts/
chmod 755 build.sh
./build.sh
```

## 实现效果

![image-20240111170311603](https://ossnote.iuskye.com/notesimg/image-20240111170311603.png)
