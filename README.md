# golang_longest_consecutive_sequence

Given an unsorted array of integers `nums`, return *the length of the longest consecutive elements sequence.*

You must write an algorithm that runs in `O(n)` time.

## Examples

**Example 1:**

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is[1, 2, 3, 4]. Therefore its length is 4.

```

**Example 2:**

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

```

**Constraints:**

- `0 <= nums.length <= 105`
- `109 <= nums[i] <= 109`

## 解析

給定一個整數陣列 nums

要求寫一個演算法找出最長的連續遞增數列長度

最直覺的想法就是把 nums 做排序 

然後最依序從最小的數開始檢查後面的數是否有連續的

如果有連續的就加1 發現不連續就重新設定為 1

這樣做的時間複雜度是 O(nlogn)

如果需要用更小的時間複雜度 O(n)

需要用到 HashSet 來紀錄每個拜訪過的數字

因為只需要找遞增數列所以只需要收集唯一出現的數字

對於每個唯一出現的數字

可以發現如果從數列最小值往上逐步加一 如果數列有其他值一定會在 HashSet內

所以可以做以下策略找數列長度

當該數字是數列中最小值時 開始檢查 HashSet 中有沒有該數字 +1 如果有則把累計長度加一

直到遇到沒有當下數字加一的數在 HashSet 這樣所累計的長度就是該數列的長度

當不是數列最小數值時, 則跳過不做運算

關鍵是如何檢查是數列最小數值

假設該數值是 num 檢查 num-1 是否有在 HashSet

如果有則不是

如果沒有則是

具體作法如下圖

![](https://i.imgur.com/82wItBC.png)

因為每次 loop 只有在最小值才需要做累計 時間複雜度是 O(n), 其中 n 代表 len(nums)

## 程式碼
```go
package sol

func longestConsecutive(nums []int) int {
	hash := make(map[int]struct{})
	unique := []int{}
	maxLen := 0
	// accumulate
	for _, num := range nums {
		if _, ok := hash[num]; !ok {
			hash[num] = struct{}{}
			unique = append(unique, num)
		}
	}
	// find
	for _, num := range unique {
		if _, ok := hash[num-1]; !ok { // num is smallest
			currentNum := num
			currentLen := 1
			for _, exist := hash[currentNum+1]; exist; _, exist = hash[currentNum+1] {
				currentNum += 1
				currentLen += 1
			}
			if maxLen < currentLen {
				maxLen = currentLen
			}
		}
	}
	return maxLen
}
```
## 困難點

1. 要想出由最小值 num 逐步往上遞加檢查是否在 HashSet 的方式並不直觀
2. 需要想出檢查最小值的方法

## Solve Point

- [x]  建立一個 HashSet hash 還有一個整數陣列 unique = [] 用來紀錄只出現一次的值
- [x]  初始化 maxLen := 0, currentLen := 0
- [x]  for i := 0.. len(nums) 當 nums[i] 沒出現在 hash 內時把 nums[i] 加入 hash 還有 unique
- [x]  if len(unique) == 1 return 1
- [x]  for i := 0..len(unique) 做以下運算
- [x]  if unique[i]-1 沒出現在 hash 時 做以下運算算
- [x]  初始化 currentNum = unique[i] , currentLen = 1
- [x]  當 currentNum+1 有在 hash 時, 更新 currentNum +=1, currentLen += 1
- [x]  if maxLen < currentLen , 更新 maxLen = currentLen
- [x]  回傳 maxLen