/*
	Day11.9
	Leetcode第1178题: Number of Valid Words for Each Puzzle
*/

func findNumOfValidWords(words []string, puzzles []string) []int {
	mp := make(map[int]int)

	for _, word := range words {
		mask := 0
		for _, c := range word {
			mask |= 1 << (c - 'a')
		}
		mp[mask]++ //记录下每个单词出现的次数
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

/*
	Day 11.13
	Leetcode第739题：Daily Temperature 下一日的最高温度
*/

type Tuple struct {
	Val   int
	Index int
}

func dailyTemperatures(nums []int) []int {

	res := make([]int, len(nums))

	var stack []Tuple

	//Golang没有栈，从后往前模拟栈运算即可
	for i := len(nums) - 1; i >= 0; i-- {
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

/*
	Day11.14
	Leetcode第1286题: Iterator for Combination
*/

type CombinationIterator struct {
	Words    []string
	Iterator int
}

func Constructor(str string, len int) CombinationIterator {

	return CombinationIterator{Combination(str, len), 0}
}

func (this *CombinationIterator) Next() string {
	if this.Iterator >= len(this.Words) {
		return ""
	} else {
		cur := this.Words[this.Iterator]
		this.Iterator++

		return cur
	}
}

func (this *CombinationIterator) HasNext() bool {
	return this.Iterator < len(this.Words)
}

func Combination(str string, lenght int) []string {

	res, comb, mask := []string{}, []byte{}, 1<<len(str)

	for no := 1; no < mask; no++ {
		num, i := no, 0

		for num > 0 {
			if num&1 > 0 {
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

/*
	上面是对 byte[] 排序
	Combination可以采用上面的方法，也可以采用这个下面的方法
	下面是对 string[] 排序
*/

func Combination_2(str string, lenght int) []string {

	var num uint32

	for _, c := range s {
		num |= 1 << uint32(c-'a')
	}

	var res []string
	for n := num; n > 0; n = (n - 1) & num {
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
		if n&(mask<<i) > 0 {
			s += string(byte(i) + 'a')
		}
	}

	return s
}

func checkLen(n uint32, lenght int) bool {
	for n > 0 {
		lenght -= int(n & 1)
		n >>= 1
	}
	return lenght == 0
}

/*
	Day 11.15
	Leetcode第368题: Largest Divisible Subset
*/

func largestDivisibleSubset(nums []int) []int {

	var res []int
	if len(nums) == 0 {
		return res
	}

	//dp[i]存储到nums[i]结尾的满足条件的最大子数组的元素个数
	//preIndex[i]存储nums[i]之前的一个元素的下标
	dp, preIndex := make([]int, len(nums)), make([]int, len(nums))

	max_size, max_index := 0, 0

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if (nums[i]%nums[j]) == 0 && dp[i] < dp[j]+1 {
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

//法二：暴力破解
//排好序，两层for循环遍历即可

func largestDivisibleSubset(nums []int) []int {

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

/*
	Day 11.16
	Leetcode第668题：第k个最小元素在 Multiplication Table
*/

//思路：二分查找

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

/*
	Day 11.17
	Leetcode第62题：Unique Paths 到达终点的唯一路径
*/

//思路：动态规划

func uniquePaths(m, n int) int {

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for j := range dp[0] {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

/*
	Day 11.18
	Leetcode第448题：找到数组中没出现的数
*/

//思路：一次循环标记即可

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

//注：Golang无三目运算符
func abs(n int) int {
	// return n > 0 ? n : -1
	if n < 0 {
		return -n
	}

	return n
}

/*
	Day 11.21
	Leetcode第106题：按照后序和中序序列构建二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	inorderPos = make(map[int]int, len(inorder))

	for i := 0; i < len(inorder); i++ {
		inorderPos[inorder[i]] = i
	}

	return dfs(postorder, 0, len(postorder)-1, 0, inorderPos)
}

func dfs(postorder []int, startPostorder, endPostorder, startInorder int, inorderPos map[int]int) *TreeNode {

	if startPostorder > endPostorder {
		return nil
	}

	root := &TreeNode{postorder[endPostorder], nil, nil}

	rootIndex := inorderPos[postorder[endPostorder]]
	leftLen := rootIndex - startInorder

	root.Left = dfs(postorder, startPostorder, startPostorder+leftLen-1, startInorder, inorderPos)
	root.Right = dfs(postorder, startPostorder+leftLen, endPostorder-1, rootIndex+1, inorderPos)

	return root
}

/*
	Day 11.22
	Leetcode第450题: 删除BST中的某个节点
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		pNode := getRightMin(root.Right)
		root.Val = pNode.Val
		root.Right = deleteNode(root.Right, pNode.Val)

	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	return root
}

func getRightMin(root *TreeNode) *TreeNode {

	for root.Left != nil {
		root = root.Left
	}

	return root
}

/*
	Day 11.26
	Leetcode第35题：寻找插入位置
*/

// 暴力破解
func searchInsert(nums []int, target int) int {

	for i, val := range nums {
		if val >= target {
			return i
		}
	}

	return len(nums)
}

// 思路：二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

/*
	Day 11.27
	Leetcode第238题:数组中元素除去自身的所有元素的乘积
*/

//思路：从左往右遍历一次，再从右往左遍历一次

func productExceptSelf(nums []int) []int {

	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[0] = 1
		} else {
			res[i] = res[i-1] * nums[i-1]
		}
	}

	prod := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= prod
		prod *= nums[i]
	}

	return res
}

/*
	Day 11.28
	Leetcode第779题：打印所有的路径
*/

//写法一
func allPathsSourceTarget(graph [][]int) [][]int {
	res, path := [][]int{}, []int{}

	dfs(graph, 0, &res, &path)

	return res
}

func dfs(graph [][]int, cur int, res *[][]int, path *[]int) {

	*path = append(*path, cur)

	if cur == len(graph)-1 {
		x := make([]int, len(*path))
		copy(x, *path)
		*res = append(*res, x)
	} else {
		for _, n := range graph[cur] {
			dfs(graph, n, res, path)
		}
	}

	*path = (*path)[:len(*path)-1]
}

//写法二
var res = [][]int{}
var path = []int{}

func allPathsSourceTarget(graph [][]int) [][]int {
	res = [][]int{}
	path = []int{}

	dfs(graph, 0)

	return res
}

func dfs(graph [][]int, cur int) {

	path = append(path, cur)

	if cur == len(graph)-1 {
		x := make([]int, len(path))
		copy(x, path)
		res = append(res, x)
	} else {
		for _, n := range graph[cur] {
			dfs(graph, n)
		}
	}

	path = path[:len(path)-1]
}

/*
	Day 11.29
	Leetcode第721题：Accounts Merge
*/

type DSU struct {
	parent []int
}

func constructor() DSU {
	parent := make([]int, 10000)
	for k := range parent {
		parent[k] = k
	}

	return DSU{parent: parent}
}

func (ds DSU) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds DSU) union(x, y int) {
	ds.parent[ds.find(x)] = ds.find(y)
}

func accountsMerge(accounts [][]string) [][]string {
	ds := constructor()
	emialToName := make(map[string]string)
	emialToID := make(map[string]int)
	id := 0

	for _, account := range accounts {
		name := ""
		for i, v := range account {
			if i == 0 {
				name = v //获得第一个字符串(就是名字)
				continue
			}

			emialToName[v] = name
			if _, ok := emialToID[v]; !ok {
				emialToID[v] = id
				id++
			}
			ds.union(emialToID[account[1]], emialToID[v])
		}
	}

	//这个ans的key为int, value为切片[]string
	ans := make(map[int][]string)
	for k := range emialToName {
		index := ds.find(emialToID[k])
		ans[index] = append(ans[index], k)
	}

	var result [][]string
	var arr []string
	for _, v := range ans {
		sort.Strings(v)
		arr = nil
		name := emialToName[v[0]]
		arr = append(arr, name)
		arr = append(arr, v...)
		result = append(result, arr)
	}

	fmt.Println(result)

	return result
}

/*
	Day 11.30
	Leetcode第84题：柱状图的最大矩形面积
*/

func maximalRectange(matrix [][]byte) int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	nums := make([]int, len(matrix[0]))
	res := 0

	for i := range matrix {
		for i, val := range matrix[i] {
			if val == '0' {
				nums[i] = 0
			} else {
				nums[i] += int(val - '0')
			}
		}
		res = Max(res, largetsetRectangleArea(nums))
	}

	return res
}

func largetsetRectangleArea(nums []int) int {

	sta := []int{}
	nums = append(nums, 0)
	res, left := 0, 0

	for i, val := range nums {
		for len(sta) > 0 && val < nums[sta[len(sta)-1]] {
			hei := nums[sta[len(sta)-1]]
			sta = sta[:len(sta)-1]
			if len(sta) == 0 {
				left = -1
			} else {
				left = sta[len(sta)-1]
			}
			res = Max(res, hei*(i-left-1))
		}
		sta = append(sta, i)
	}

	return res
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
	Day 12.1
	Leetcode第198题：抢劫犯
*/

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		res[i] = Max(res[i-1], res[i-2]+nums[i])
	}

	return res[len(res)-1]
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
	Day12.3
	LeetCode第152题：乘积最大子数组乘积
*/

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func maxProduct(nums []int) int {
	minProd, maxProd := 1, 1
	res := INT_MIN

	for _, n := range nums {
		if n < 0 {
			minProd, maxProd = maxProd, minProd
		}

		maxProd = Max(maxProd*n, n)
		minProd = Min(minProd*n, n)

		res = Max(res, maxProd)
	}

	return res
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/*
	Day12.4
	Leetcode第1032题：Stream of Characters
*/

//使用字典树来解决

/*
class TrieNode
{
public:
    bool is_leaf;
    TrieNode *childNode[26];

    TrieNode() {
        is_leaf = false;
        for(int i = 0; i < 26; i++) {
            childNode[i] = nullptr;
        }
    }

    void insert_reversed(std::string word)
    {
        reverse(word.begin(), word.end());
        TrieNode *root = this;

        for(int i = 0; i < word.size(); i++) {
            int index = word[i] - 'a';
            if(root->childNode[index] == nullptr) {
                root->childNode[index] = new TrieNode();
            }
            root = root->childNode[index];
        }

        root->is_leaf = true;
    }

};


class StreamChecker
{
public:
    TrieNode *trie;
    std::vector<char> queries;
    int longest_word;

    StreamChecker(std::vector<std::string> &words) {
        trie = new TrieNode();
        longest_word = 0;

        for(auto &word : words) {
            trie->insert_reversed(word);
            longest_word = std::max(int(word.size()), longest_word);
        }
    }

    bool query(char letter)
    {
        //头插法,累加的queries超过最大的单词长度就不用再保存末尾的字符了
        queries.insert(queries.begin(), letter);
        if(queries.size() > longest_word) queries.pop_back();

        TrieNode *cur = trie;
        for(auto it = queries.begin(); it != queries.end(); it++) {
            if(cur->is_leaf) return true;
            if(cur->childNode[*it-'a'] == nullptr) return false;
            cur = cur->childNode[*it-'a'];
        }

        return cur->is_leaf;
    }

};
*/

type trieNode struct {
	children [26]*trieNode
	isWord   bool
}

type StreamChecker struct {
	queryWord    []byte
	trieNodeRoot *trieNode
	longest_word int //记录最大的字符串长度
}

func Constructor(words []string) StreamChecker {

	root := &trieNode{
		children: [26]*trieNode{},
		isWord:   false,
	}

	maxLen := 0

	for _, word := range words {
		maxLen = Max(maxLen, len(word))

		cur := root
		for i := len(word) - 1; i >= 0; i-- {
			ch := int(word[i] - 'a')
			if cur.children[ch] == nil {
				cur.children[ch] = &trieNode{
					children: [26]*trieNode{},
					isWord:   false,
				}
			}
			cur = cur.children[ch]
		}
		cur.isWord = true
	}

	return StreamChecker{
		trieNodeRoot: root,
		queryWord:    []byte{},
		longest_word: maxLen,
	}

}

func (this *StreamChecker) Query(letter byte) bool {

	this.queryWord = append(this.queryWord, letter)
	cur := this.trieNodeRoot

	if len(this.queryWord) > this.longest_word {
		this.queryWord = this.queryWord[1:len(this.queryWord)]
	}

	for i := len(this.queryWord) - 1; i >= 0; i-- {
		ch := int(this.queryWord[i] - 'a')
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
		if cur.isWord {
			return true
		}
	}

	return false
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
	Day 12.5
	Leetcode第337题：打家劫舍
*/

func rob(root *TreeNode) int {
	mp := make(map[*TreeNode]int)

	return dfs(root, mp)
}

func dfs(root *TreeNode, mp map[*TreeNode]int) {
	if root == nil {
		return 0
	}

	if v, ok := mp[root]; ok {
		return v
	}

	money1, money2 := root.Val, 0

	if root.Left != nil {
		money1 += dfs(root.Left.Left, mp) + dfs(root.Left.Right, mp)
	}
	if root.Right != nil {
		money1 += dfs(root.Right.Left, mp) + dfs(root.Right.Right, mp)
	}

	money2 += dfs(root.Left, mp) + dfs(root.Right, mp)

	mp[root] = Max(money1, money2)

	return Max(money1, money2)
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
	Day 12.7
	Leetcode第1290题：二进制表达
*/

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = (res << 1) | head.Val
		head = head.Next
	}

	return res
}

/*
	Day 12.8
	Leetcode第563题
*/

//思路：自下而上递归即可

func findTilt(root *TreeNode) int {
	res := 0
	dfs(root, &res)

	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left, right := dfs(root.Left, res), dfs(root.Right, res)

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

/*
	Day 12.9
	Leetcode第1306题：Jump Game III
*/
func canReach(arr []int, start int) bool {
	mp := make(map[int]int, len(arr))

	return dfs(arr, start, mp)
}

func dfs(arr []int, start int, mp map[int]int) bool {
	if start >= 0 && start < len(arr) && mp[start] == 0 {
		mp[start] = 1
		if arr[start] == 0 {
			return true
		} else {
			return dfs(arr, start+arr[start], mp) || dfs(arr, start-arr[start], mp)
		}
	} else {
		return false
	}
}

/*
	Day 12.12
	Leetcode第416题：将数组切分成两个和相同的子数组
*/

//思路：动态规划

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}

	sum >>= 1

	n := len(nums)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			val := nums[i-1]
			if j < val {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-val]
			}
		}
	}

	return dp[n][sum]
}

//降维成一维写法
func canPartition2(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1

	dp := make([]bool, sum+1)
	dp[0] = true

	for _, n := range nums {
		for j := sum; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}

	return dp[sum]
}

//Leetcode第703题：Golang实现的优先队列

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(h)
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	return this.h.Peek().(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } //实现小顶堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Peek() interface{} {
	return (*h)[0]
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

// Leetcode第131题 (关于切片slice的大坑--一定要注意)

func partition(s string) [][]string {
	res := [][]string{}

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1] == 1) {
				dp[i][j] = 1
			}
		}
	}

	path := []string{}

	dfs(s, 0, dp, path, &res)

	return res
}

func dfs(s string, pos int, dp [][]int, path []string, res *[][]string) {
	if pos == len(s) {
		//fmt.Println(path)
		*res = append(*res, path)                        //1.这种写法会有bug
		*res = append(*res, append([]string{}, path...)) //2.这种写法没有bug
		//fmt.Println(*res)
		return
	}

	for i := pos; i < len(s); i++ {
		if dp[pos][i] == 1 {
			path = append(path, s[pos:i+1])
			dfs(s, i+1, dp, path, res)
			path = path[:len(path)-1]
		}
	}

}

//Leetcode第382题：蓄水池抽样问题(着重看这代码的写法)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	n, res := 2, this.head.Val
	cur := this.head.Next

	for cur != nil {
		if rand.Intn(n) == 0 {
			res = cur.Val
		}
		n++
		cur = cur.Next
	}

	return res
}