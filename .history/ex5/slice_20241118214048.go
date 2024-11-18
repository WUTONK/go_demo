package main

func main() {
	var arr []int
	// 请在 arr 添加 1 - 100
	// 并且使用迭代写法打印出所有 arr 的元素
	// code here

	for i := 1; i <= 100; i++ {
		arr = append(arr, i)
	}

	for _, value := range arr {
		print(value)
	}

	var m map[int]int = make(map[int]int) // 初始化映射，防止nil空指针引用错误
	// 请在 m 添加：
	// key: 1, value: 2
	// key: 2, value: 3
	// key: 3, value: 4
	// 直到 key: 100, value: 101
	// 并且使用迭代写法打印出所有的 key 和 value
	// code here

	for i := 1; i < 101; i++ {
		m[i] = i + 1
	}

	for key, value := range m {
		print("key:%d,value%d", key, value)
	}

}
