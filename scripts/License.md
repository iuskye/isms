## License 自动创建

```shell
## 安装 license 工具
go install github.com/nishanths/license/v5@latest
## 查看支持的代码协议
license -list
# 在 isms 项目根目录下执行
license -n 'Simon Liu <iuskye@foxmail.com>' -o LICENSE apache-2.0
ls LICENSE

## 安装 addlicense 工具
go install github.com/marmotedu/addlicense@latest
## 添加版权头信息
addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,vendor,_output . added license
```
