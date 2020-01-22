package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/magiconair/properties"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//将 json 文件读到 map 中来
func readJsonByPath(filePath string) (res map[string]string) {
	res = make(map[string]string)
	buf, err := ioutil.ReadFile(filePath)
	check(err)
	err = json.Unmarshal([]byte(string(buf)), &res)
	check(err)
	for key, value := range res {
		fmt.Println("key", key, "value", value)
	}
	return res
}

//根据 json 文件修改 properties 文件的内容
func setByFilePath(sourcePath, filePath string) (ok bool) {
	p := properties.MustLoadFile(filePath, properties.UTF8)
	//jsonRes :=make(map[string]string)
	jsonRes := readJsonByPath(sourcePath)
	for key, value := range jsonRes {
		p.SetValue(key, value)
	}
	f, err := os.Create(filePath)
	check(err)
	//p.Write(f, properties.UTF8)
	p.WriteComment(f,"#",properties.UTF8)
	return true

}

//根据参数名设置新的参数
func setByName(name, newValue, filePath string) (ok bool) {
	p := properties.MustLoadFile(filePath, properties.UTF8)

	p.SetValue(name, newValue)
	f, err := os.Create(filePath)
	check(err)
	//p.Write(f, properties.UTF8)
	p.WriteComment(f,"#",properties.UTF8)
	fmt.Println(true)
	return true
}

//根据参数获取某个配置的值,返回的是字典
func getByName(name, filePath string) (res string) {
	p := properties.MustLoadFile(filePath, properties.UTF8)
	res = p.MustGetString(name)
	//fmt.Println(res)
	return res
}
func main() {
	var Name = flag.String("n", "name", "需要设置的参数名称")
	var FilePath = flag.String("u", "application.properties", "需要设置的配置文件路径")
	var Value = flag.String("v", "configJava", "需要设置的参数值")
	var Action = flag.String("t", "file", "操作类型,有get|set|file")
	var SourcePath = flag.String("s", "config.json", "用来修改参数文件的json文件路径")

	flag.Parse()
	//fmt.Println(flag.NFlag())
	if flag.NFlag() < 3 {
		//usage := "使用示例:  configJava -t set -n host -u config.properties"
		usage :="举例:\n设置单个值:  configJava -t set -n host -u config.properties; \n利用json文件设置多个值: configjava -t file -s test.json -u config.properties  "
		fmt.Println(usage)
		return
	}

	sourcePath := *SourcePath
	action := *Action
	name := *Name
	filePath := *FilePath
	value := *Value

	fmt.Println(action, name, filePath, value)

	switch {
	case action == "get":
		{
			fmt.Println(getByName(name, filePath))
		}
	case action == "set":
		{
			setByName(name, value, filePath)
		}
	case action == "file":
		{
			setByFilePath(sourcePath, filePath)
		}

	}
}
