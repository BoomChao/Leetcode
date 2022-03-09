/*

type CombinationIterator struct {
	Words []string
	Iterator int
}

func Constructor(str string, len int) CombinationIterator {

	return CombinationIterator{ Combination(str, len), 0}
}

func (this *CombinationIterator) Next() string {
	if this.Iterator >= len(this.Words) {
		return ""
	} else {
		cur := this.Words[this.Iterator]
		this.Iterator++;

		return cur
	}
}


func(this *CombinationIterator) HasNext() bool {
	return this.Iterator < len(this.Words)
}



func Combination(str string, lenght int) []string {

	res, comb, mask := []string{}, []byte{}, 1 << len(str)

	for no := 1; no < mask; no++ {
		num, i := no, 0

		for num > 0 {
			if num & 1 > 0 {
				comb = append(comb, str[i])
			}
			i++
			num >>= 1
		}

		if len(comb) == lenght {
			res = append(res, string(comb))
		}

		comb = []byte{}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func Combination_2(str string, lenght int) []string {

	var num uint32

	for _, c := range s {
		num |= 1 << uint32(c - 'a')
	}

	var res []string
	for n := num; n > 0; n = (n-1)&num {
		if checkLen(n, lenght) {
			res = append(res, transformToString(n))
		}
	}

	sort.Strings(res)

	return res
}


func transformToString(n uint32) string {
	s := ""
	mask := uint32(1)

	for i := 0; i < 32; i++ {
		if n & (mask << i) > 0 {
			s += string(byte(i) + 'a')
		}
	}

	return s
}


func checkLen(n uint32, lenght int) bool {
	for n > 0 {
		lenght -= int(n&1)
		n >>= 1
	}
	return lenght == 0
}




func findNumOfValidWords(Words []string, puzzles []string) []int {
	mp := make(map[int]int)

	for _, word := range words {
		mask := 0
		for _, c := range word {
			mask |= 1 << (c - 'a')
		}
		mp[mask]++					//记录下每个单词出现的次数
	}

	res := make([]int, len(puzzles))

	for i := 0; i < len(puzzles); i++ {
		mask := 0
		str := puzzles[i]

		for _, c := range str {
			mask |= 1 << (c - 'a')
		}

		c := 0
		var sub int = mask
		first := 1 << (str[0] - 'a')

		for {
			if (sub & first) == first {
				c += mp[sub]
			}
			if sub == 0 {
				break
			}
			sub = (sub - 1) & mask
		}

		res[i] = c
	}

	return res
}


func largestDivisibleSubset(nums []int) []int {

	var res []int
	if len(nums) == 0 {
		return res
	}

	dp, preIndex := make([]int, len(nums)), make([]int, len(nums))

	max_size, max_index := 0, 0

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if (nums[i] % nums[j]) == 0 && dp[i] < dp[j] + 1 {
				dp[i] = dp[j] + 1
				preIndex[i] = j
				if max_size < dp[i] {
					max_size = dp[i]
					max_index = i
				}
			}
		}
	}

	for i := 0; i < max_size; i++ {
		res = append(res, nums[max_index])
		max_index = preIndex[max_index]
	}

	return res
}


func largestDivisibleSubset_2(nums []int) []int {

	sub_sets := make([][]int, len(nums))
	sort.Ints(nums)

	for i, _ := range nums {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(sub_sets[i]) < len(sub_sets[j]) {
				sub_sets[i] = sub_sets[j]
			}
		}

		sub_sets[i] = append(sub_sets[i], nums[i])
	}

	var max_sub []int
	for _, s := range sub_sets {
		if len(max_sub) < len(s) {
			max_sub = s
		}
 	}

	return max_sub
}



type Tuple struct {
	Val int
	Index int
}


func dailyTemperatures(nums []int) []int {

	res := make([]int, len(nums))

	var stack []Tuple

	//Golang没有栈，从后往前模拟栈运算即可
	for i := len(nums)-1; i >= 0; i-- {
		tmp := nums[i]

		for len(stack) > 0 && tmp >= stack[len(stack)-1].Val {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1].Index - i
		}

		stack = append(stack, Tuple{tmp, i})
	}

	return res
}


func findKthNumber(m int, n int, k int) int {

	left, right := 1, m*n

	for left < right {
		mid := left + ((right - left) >> 1)
		c := count(mid, m, n)
		if c >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func count(v, m, n int) int {
	num := 0
	for i := 1; i <= m; i++ {
		num += Min(v/i, n)
	}
	return num
}

func Min(a, b int) {
	if a < b {
		return a
	} else {
		return b
	}
}

*/

package main

import (
	"fmt"
	_ "sort"
)

func main() {
	// str, len := "abc", 2
	// obj := Constructor(str, len)
	// fmt.Println(obj.Next())
	// fmt.Println(obj.HasNext())

	// nums := []int{1,4,7,8}
	// largestDivisibleSubset_2(nums)

	// for i := range nums {
	// 	fmt.Println(i)
	// }

	// fmt.Println(largestDivisibleSubset(nums))

	nums := []int{10, 20, 30, 40, 50}

	if pos, find := findNumberInArray(nums, 50); find {
		fmt.Println(pos)
	}

}

func findNumberInArray(nums []int, target int) (int, bool) {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, false
}

func findDisapperaredNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		if index := abs(nums[i]) - 1; nums[index] > 0 {
			nums[index] *= -1
		}
	}

	res := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preoreder []int, inorder []int) *TreeNode {

	inorderPos := make(map[int]int, len(inorder))

	for i := 0; i < len(inorder); i++ {
		inorderPos[inorder[i]] = i
	}

	return dfs(preoreder, 0, len(preoreder)-1, 0, inorderPos)
}

func dfs(preoreder []int, startPreorder, endPreorder, startInorder int, inorderPos map[int]int) *TreeNode {

	if startPreorder > endPreorder {
		return nil
	}

	root := &TreeNode{preoreder[startPreorder], nil, nil}

	rootIndex := inorderPos[preoreder[startPreorder]]
	leftLen := rootIndex - startInorder

	root.Left = dfs(preoreder, startPreorder+1, startPreorder+leftLen, startInorder, inorderPos)
	root.Right = dfs(preoreder, startPreorder+leftLen+1, endPreorder, rootIndex+1, inorderPos)

	return root
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	dfs(root, &res)

	return res
}

func dfs(root *TreeNode, res *int) int {
	left, right := 0, 0

	if root.Left == nil {
		left = dfs(root.Left, res)
	} else {
		left = 0
	}

	if root.Right == nil {
		right = dfs(root.Right, res)
	} else {
		right = 0
	}

	*res += Abs(left, right)

	return left + right + root.Val
}

func Abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
