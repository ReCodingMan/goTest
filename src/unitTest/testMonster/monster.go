package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool {
	dir,_ := os.Getwd()
	fmt.Println("当前路径：",dir)
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}

	//保存到文件
	filepath := "../jsonFile/monster.ser"
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return false
	}

	return true
}

func (this *Monster) ReStore() bool {
	//1、先从文件中读取序列化的字符串
	filepath := "../jsonFile/monster.ser"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read file err = ", err)
		return false
	}

	//2、使用切片反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err = ", err)
		return false
	}

	return true
}