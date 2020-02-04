# configJava
Java 的配置文件 application.properties 批量修改参数较为麻烦,这是一个修改参数文件的小工具.

## 下载地址
https://github.com/Waitfish/configJava/releases
## 使用方法
```bash
Usage of configjava
  -n string
        需要设置的参数名称 (default "name")
  -s string
        用来修改参数文件的json文件路径 (default "config.json")
  -t string
        操作类型,有get|set|file (default "file")
  -u string
        需要设置的配置文件路径 (default "application.properties")
  -v string
        需要设置的参数值 (default "configJava")
```
### 修改单个参数值
```bash
configJava -t set -n host -u config.properties;
```

### 根据`json`文件修改多个参数值
```bash
configjava -t file -s test.json -u config.properties  
```

### 示例`json`文件
`所有参数全部用双引号包含起来`
```json
{
  "rocketmq.config.payment.namesrvAddr": "x.x.x.x:9879",
  "spring.redis.timeout": "10000"
}
```

## build for linux64
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build configJava.go   
```
