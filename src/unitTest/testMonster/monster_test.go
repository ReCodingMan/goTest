package monster

import (
	"fmt"
	"testing"
)

//测试用例
func TestStore(t *testing.T) {
	//先创建 Monster 实例
	monster := Monster{
		Name: "红孩儿",
		Age: 12,
		Skill: "吐火",
	}

	err := monster.Store()
	if !err {
		fmt.Printf("monster.Store() 错误，期望为=%v 实际为=%v\n", true, err)
	}
	t.Logf("monster.Store() 测试成功")
}

func TestReStore(t *testing.T) {

	var monster Monster
	err := monster.ReStore()
	if !err {
		fmt.Printf("monster.ReStore() 错误，期望为=%v 实际为=%v\n", true, err)
	}
	t.Logf("monster.ReStore() 测试成功")
	fmt.Println(monster)
}