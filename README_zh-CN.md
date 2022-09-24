[English](./README.md) **简体中文**

# X

🔨 一个如 `Laravel Collection` 或 `Lodash.js` 的工具集

## 安装

```shell
go get github.com/bygo/x
```

## 开始

```go
package main

import (
	"fmt"
	"github.com/bygo/x"
)

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	nums = x.Collect(nums).
		Filter(func(val int, k int) bool {
			return val%2 == 1 // []int{1, 3, 5}
		}).
		Diff([]int{1}). // []int{3, 5}
		Map(func(val int, k int) int {
			if val == 3 {
				return val
			}
			return val * 2 // []int{3, 10}
		}).
		Replace(3, 5, 1). // Replace(old, new, n) => []int(5, 10)
		ForEach(func(val int, k int) {
			println(val) // Output 5,10
		}).
		ToSlice()

	fmt.Printf("%+v", nums) // Output []int{5,10}
}
```



