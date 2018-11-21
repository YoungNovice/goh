package monster

import (
	"testing"
)
func TestStore(t *testing.T) {
	monster := Monster{
		Name:"yangxuan",
		Age : 18,
		Skill: "吃东西",
	}
	err := monster.Store()
	if err != nil {
		t.Fatalf("保存失败 testing err")
	}
	t.Logf("Monster.Store测试通过")
}

func TestRestore(t *testing.T) {
	var monster Monster
	err := monster.Restore("/Users/yangxuan/monster.json")
	if err != nil {
		t.Fatalf("反序列化失败 testing err")
	}
	t.Logf("Monster.Restore测试通过 monster=%v", monster)

}
