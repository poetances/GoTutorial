package main

import "fmt"

/*
map使用前必须初始化，不然只声明，默认是nil。
初始化方式有两种：
m1 := make(map[string]int)

m2 := map[string]int{}

// 只声明，不初始化会报错
var m3 map[string]int

6. 总结
✅ map 是 Go 内置的数据结构，用于存储键值对，查找快。
✅ 常见初始化方式：make()、字面量 {}，但不能使用 new()。
✅ 增删改查：
	•	m[key] = value 添加/修改
	•	delete(m, key) 删除
	•	val, exists := m[key] 判断 key 是否存在
✅ 遍历：for key, value := range map，但无序，如需排序可先取出 key 排序。
✅ map 是引用类型，赋值时不会拷贝数据，需手动复制。
✅ key 必须是可比较类型，不能是 slice、map、function。
*/
func MapTutorial() {
	// 直接使用make
	m := make(map[string]int)
	m["name"] = 12
	m["age"] = 3
	fmt.Println(m)

	// 使用字面量
	m2 := map[string]int{
		"name": 12,
		"age": 3,
	}
	fmt.Println(m2)

	val, exists := m2["apple"]
	if exists {
		fmt.Println("apple:", val)
	} else {
		fmt.Println("apple 不存在")
	}

	// 删除元素
	delete(m2, "apple")

	for key, value := range m {
		fmt.Println(key, value)
	}

	//（1）map 是引用类型
	m3 := m2
	m3["age"] = 10 // m3的赋值会影响m2
	
	//（2）map 的 key 只能是可比较类型
	// ✅ 可以作为 key：string、int、bool、float、struct（可比较）

	//5. map 与 struct 的区别
	/*
	特性		map	struct
	键的类型	只能是可比较类型	字段名固定
	值的类型	任意类型	字段类型固定
	存储方式	哈希表，查找快	顺序存储，内存紧凑
	是否有序	无序	有序
	用途		适合动态存储键值对	适合固定字段的数据
	*/
}