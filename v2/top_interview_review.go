package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	var ask []int

	for i, num := range nums {

		if k, ok := valueMap[target-num]; ok {
			ask = append(ask, k, i)
			return ask
		} else {
			valueMap[num] = i
		}

	}

	return ask
}

func romanToInt(s string) int {
	var symbol2value = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	val := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			symbol := s[i : i+2]
			if value, ok := symbol2value[symbol]; ok {
				val += value
				i++
				continue
			}
		}
		if value, ok := symbol2value[string(s[i])]; ok {
			val += value
		}
	}

	return val
}

//14. Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	for pivotIndex := 0; pivotIndex < len(strs[0]); pivotIndex++ {
		i := 0
		for i < len(strs) {
			if pivotIndex >= len(strs[i]) {
				break
			}
			if strs[i][pivotIndex] != strs[0][pivotIndex] {
				break
			}
			i++
		}
		if i == len(strs) {
			sb.WriteByte(strs[0][pivotIndex])
		} else {
			break
		}
	}
	return sb.String()
}

//20. Valid Parentheses
func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		lth := len(stack)
		if lth > 0 {
			if s[i] == ')' && stack[lth-1] == '(' || s[i] == '}' && stack[lth-1] == '{' || s[i] == ']' && stack[lth-1] == '[' {
				stack = stack[:lth-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

// Definition for singly-linked list

type ListNode struct {
	Val  int
	Next *ListNode
}

//26. Remove Duplicates from Sorted Array
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		head.Next = mergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		head.Next = mergeTwoLists(list1, list2.Next)
	}
	return head
}

//26. Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cmp := nums[0]
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != cmp {
			nums[index] = nums[i]
			cmp = nums[i]
			index++
		}
	}
	return index
}

func removeDuplicatesV2(nums []int) int {
	cur := 0
	for cur < len(nums)-1 {
		pivot := nums[cur]
		j := cur + 1
		for j < len(nums) && nums[j] == pivot {
			j++
		}
		nums = append(nums[:cur+1], nums[j:]...)
		cur++
	}
	return len(nums)

}

//28. Implement strStr()
func strStr(haystack string, needle string) int {
	lthn := len(needle)
	lthh := len(haystack)
	for i := 0; i < lthh; i++ {
		if haystack[i] == needle[0] && i+lthn <= lthh {
			if haystack[i:i+lthn] == needle {
				return i
			}
		}
	}
	return -1
}

//66. Plus One
func plusOne(digits []int) []int {
	var res []int
	lth := len(digits)

	carry := 1

	for i := lth - 1; i >= 0; i-- {
		sum := digits[i] + carry
		num := sum % 10
		carry = sum / 10
		res = append(res, num)
	}
	if carry > 0 {
		res = append(res, carry)
	}
	reverseSlice(res)
	return res
}

func reverseSlice(nums []int) {
	lhs := 0
	rhs := len(nums) - 1
	for lhs < rhs {
		nums[lhs], nums[rhs] = nums[rhs], nums[lhs]
		lhs++
		rhs--
	}
}

func plusOneV2(digits []int) []int {
	lth := len(digits)
	going := 1
	for i := lth - 1; i >= 0; i-- {
		sum := digits[i] + going
		going = sum / 10
		val := sum % 10
		digits[i] = val
	}

	if going == 1 {
		var res []int
		res = append(res, going)
		res = append(res, digits...)
		return res
	} else {
		return digits
	}
}

//69. Sqrt(x)
func mySqrt(x int) int {
	res := 0
	for res*res < x {
		res++
	}
	if res*res == x {
		return res
	}
	return res - 1

}

func mySqrtV2(x int) int {
	if x == 1 {
		return 1
	}
	low, high := 0, x
	var mid int
	var sqr int

	for {
		mid = (low + high) / 2
		if mid == low {
			return mid
		}

		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			high = mid
		}
		if sqr < x {
			low = mid
		}
	}
}

//70. Climbing Stairs
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func testSlice(nums []int) {
	nums[0] = 99
	nums = append(nums, 2)
}

//88. Merge Sorted Array
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	putPosition := m + n - 1
	cmpNums1 := m - 1
	cmpNums2 := n - 1
	for cmpNums1 >= 0 && cmpNums2 >= 0 {
		if nums1[cmpNums1] > nums2[cmpNums2] {
			nums1[putPosition] = nums1[cmpNums1]
			cmpNums1--
		} else {
			nums1[putPosition] = nums2[cmpNums2]
			cmpNums2--
		}
		putPosition--
	}
	if cmpNums1 < 0 && cmpNums2 >= 0 {
		for i := cmpNums2; i >= 0; i, putPosition = i-1, putPosition-1 {
			nums1[putPosition] = nums2[i]
		}
	}

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	recInorderTraversal(root, &res)
	return res
}

func recInorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recInorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	recInorderTraversal(root.Right, res)
}

//101. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	return recIsSymmetric(root.Left, root.Right)
}

func recIsSymmetric(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return recIsSymmetric(node1.Left, node2.Right) && recIsSymmetric(node1.Right, node2.Left)
}

func isSymmetricV2(root *TreeNode) bool {

	var nodeList []*TreeNode
	nodeList = append(nodeList, root)

	for len(nodeList) > 0 {
		var valList []int
		lth := len(nodeList)
		for i := 0; i < lth; i++ {
			if nodeList[i] == nil {
				valList = append(valList, -1000)
				//nodeList = append(nodeList, nil, nil)
			} else {
				valList = append(valList, nodeList[i].Val)
				nodeList = append(nodeList, nodeList[i].Left, nodeList[i].Right)
			}
		}
		lhs := 0
		rhs := len(valList) - 1
		for lhs < rhs {
			if valList[lhs] != valList[rhs] {
				return false
			}
			lhs++
			rhs--
		}
		nodeList = nodeList[lth:]
	}
	return true
}

//104. Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	return recMaxDepth(root, 0)

}
func recMaxDepth(node *TreeNode, depth int) int {
	if node != nil {
		depth++
		return max(recMaxDepth(node.Left, depth), recMaxDepth(node.Right, depth))
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//108. Convert Sorted Array to Binary Search Tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

//118. Pascal's Triangle
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		lhs := 0
		rhs := i
		res[i] = append(res[i], 1)
		for j := lhs + 1; j < rhs; j++ {
			res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
		}
		if rhs > lhs {
			res[i] = append(res[i], 1)
		}
	}
	return res
}

//121. Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	maxP := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			maxP = max(maxP, prices[i]-min)
		}
	}
	return maxP
}

//125. Valid Palindrome
func isPalindrome(s string) bool {
	lhs := 0
	rhs := len(s) - 1
	s = strings.ToLower(s)
	for lhs < rhs {
		for lhs < len(s) && !unicode.IsNumber(rune(s[lhs])) && !unicode.IsLetter(rune(s[lhs])) {
			lhs++
		}
		for rhs >= 0 && !unicode.IsNumber(rune(s[rhs])) && !unicode.IsLetter(rune(s[rhs])) {
			rhs--
		}
		//i := s[lhs]
		//j := s[rhs]
		//fmt.Println(i, j)
		if lhs < rhs {
			if s[lhs] != s[rhs] {
				return false
			}
		}

		lhs++
		rhs--
	}
	return true
}

//136. Single Number
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

//141. Linked List Cycle
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//160. Intersection of Two Linked Lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stepA := 0
	stepB := 0
	curA := headA
	curB := headB
	for curA != nil {
		stepA++
		curA = curA.Next
	}

	for curB != nil {
		stepB++
		curB = curB.Next
	}
	curA = headA
	curB = headB
	if stepA > stepB {
		for i := 1; i <= stepA-stepB; i++ {
			curA = curA.Next
		}
	}
	if stepB > stepA {
		for i := 1; i <= stepB-stepA; i++ {
			curB = curB.Next
		}
	}
	for curB != nil && curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

//169. Majority Element
func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
			continue
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

//171. Excel Sheet Column Number
func titleToNumber(columnTitle string) int {
	acc := 1
	val := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		num := acc * int(columnTitle[i]-'A'+1)
		val += num
		acc *= 26
	}
	return val

}

//190. Reverse Bits
func reverseBits(num uint32) uint32 {
	val := uint32(0)
	for i := 1; i <= 32; i++ {
		tmp := num & 1
		num >>= 1
		val <<= 1
		val += tmp
	}
	return val
}

//191. Number of 1 Bits
func hammingWeight(num uint32) int {
	cnt := 0
	for i := 1; i <= 32; i++ {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}

//202. Happy Number
func isHappy(n int) bool {
	result := make(map[int]bool)
	sum := n
	for {
		nBytes := []byte(strconv.Itoa(sum))
		sum = 0
		for _, i := range nBytes {
			sum += int((i - '0') * (i - '0'))
		}
		if sum == 1 {
			return true
		}

		if _, exist := result[sum]; exist {
			return false
		}
		result[sum] = true
	}

	return false
}

//206. Reverse Linked List
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = nil
	return pre
}

//217. Contains Duplicate
func containsDuplicate(nums []int) bool {
	num2cnt := make(map[int]bool)
	for _, num := range nums {
		if _, ok := num2cnt[num]; ok {
			return true
		}
		num2cnt[num] = true
	}
	return false
}

//234. Palindrome Linked List
func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	cur := head
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	mid := cnt / 2
	otherHead := head
	for i := 1; i <= mid; i++ {
		otherHead = otherHead.Next
	}
	pre := otherHead
	cur = otherHead.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	otherHead.Next = nil

	list1 := head
	list2 := pre
	for list2 != nil {
		if list2.Val != list1.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return true
}

func isPalindromeListV2(head *ListNode) bool {
	slow := head
	fast := head
	var pre *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}

//237. Delete Node in a Linked Lis
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//242. Valid Anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lth := len(s)
	letters := make([]int, 26)
	for _, ch := range s {
		letters[ch-'a']++
	}
	for _, ch := range t {
		if letters[ch-'a'] > 0 {
			letters[ch-'a']--
			lth--
		}

	}

	if lth == 0 {
		return true
	}
	return false
}

//268. Missing Number

func missingNumber(nums []int) int {
	lth := len(nums)
	sum := (1 + lth) * lth / 2
	cmpSum := 0
	for _, num := range nums {
		cmpSum += num
	}
	return sum - cmpSum

}

//283. Move Zeroes
func moveZeroes(nums []int) {
	swapIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			swapIndex++
		}
	}
}

//326. Power of Three
func isPowerOfThree(n int) bool {
	acc := 1
	for {
		if acc == n {
			return true
		} else if acc > n {
			return false
		} else {
			acc *= 3
		}

	}
}

func isPowerOfThreeV2(n int) bool {
	max3 := 1162261467
	if n <= 0 {
		return false
	}
	return max3%n == 0
}

func isPowerOfThreeV3(n int) bool {
	switch {
	case n == 1:
		return true
	case n == 0:
		return false
	case n%3 != 0:
		return false
	default:
		return isPowerOfThree(n / 3)
	}
}

//344. Reverse String
func reverseString(s []byte) {
	lhs := 0
	rhs := len(s) - 1
	for lhs < rhs {
		s[lhs], s[rhs] = s[rhs], s[lhs]
		lhs++
		rhs--
	}

}

//350. Intersection of Two Arrays II
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	num2cnt := make(map[int]int)
	for _, num := range nums1 {
		num2cnt[num]++
	}
	for _, num := range nums2 {
		if _, ok := num2cnt[num]; ok && num2cnt[num] > 0 {
			res = append(res, num)
			num2cnt[num]--
		}
	}
	return res
}

//387. First Unique Character in a String
func firstUniqChar(s string) int {
	hash := make([]int, 26)
	for _, ch := range s {
		hash[ch-'a']++
	}
	for i, ch := range s {
		if hash[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

//412. Fizz Buzz
//answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
//answer[i] == "Fizz" if i is divisible by 3.
//answer[i] == "Buzz" if i is divisible by 5.
//answer[i] == i (as a string) if none of the above conditions are true.
func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

//2. Add Two Numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	mergeHead := new(ListNode)
	cur := mergeHead
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		cur.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	for l1 != nil {
		val := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		cur.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		cur = cur.Next
		l1 = l1.Next

	}

	for l2 != nil {
		val := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		cur.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		cur = cur.Next
		l2 = l2.Next

	}
	if carry > 0 {

		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return mergeHead.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	head.Next = addTwoNumbersRec(l1, l2, 0)
	return head.Next

}

func addTwoNumbersRec(l1 *ListNode, l2 *ListNode, jinWei int) *ListNode {
	node := new(ListNode)
	if l1 == nil && l2 != nil {
		node.Val = (jinWei + l2.Val) % 10
		jinWei = (jinWei + l2.Val) / 10
		node.Next = addTwoNumbersRec(l1, l2.Next, jinWei)
	} else if l1 != nil && l2 == nil {
		node.Val = (jinWei + l1.Val) % 10
		jinWei = (jinWei + l1.Val) / 10
		node.Next = addTwoNumbersRec(l1.Next, l2, jinWei)
	} else if l1 != nil && l2 != nil {
		node.Val = (jinWei + l1.Val + l2.Val) % 10
		jinWei = (jinWei + l1.Val + l2.Val) / 10
		node.Next = addTwoNumbersRec(l1.Next, l2.Next, jinWei)
	} else {
		if jinWei != 0 {
			node.Val = jinWei
		} else {
			node = nil
		}
	}

	return node
}

//3. Longest Substring Without Repeating Character
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ch2index := make(map[int32]int)
	maxLth := 0
	dp := make([]int, len(s))
	dp[0] = 0
	for i, ch := range s {
		if _, ok := ch2index[ch]; !ok {
			ch2index[ch] = i
			if i > 0 {
				dp[i] = dp[i-1]
			}
			if maxLth < (i - dp[i] + 1) {
				maxLth = i - dp[i] + 1
			}
		} else {
			index := ch2index[ch]
			if index < dp[i-1] {
				ch2index[ch] = i
				dp[i] = dp[i-1]
				if maxLth < (i - dp[i] + 1) {
					maxLth = i - dp[i] + 1
				}
			} else {
				dp[i] = ch2index[ch] + 1
				ch2index[ch] = i
			}

		}
	}
	return maxLth

}

func lengthOfLongestSubstringV2(s string) int {
	maxLength := 0
	left := 0
	hash := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if hash[s[i]] == 0 || hash[s[i]] < left {
			if maxLength < i-left+1 {
				maxLength = i - left + 1
			}
		} else {
			left = hash[s[i]]
		}
		hash[s[i]] = i + 1
	}
	return maxLength
}

//5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			if j-i == 1 || i == j {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && len(maxStr) < j-i+1 {
				maxStr = s[i : j+1]
			}
		}
	}
	return maxStr
}

func longestPalindromeV2(s string) string {
	maxSub := s[:1]
	lth := len(s)
	for i := 1; i < lth; i++ {
		lhs := i - 1
		rhs := i + 1
		for lhs >= 0 && s[lhs] == s[i] {
			lhs--
		}
		for rhs < lth && s[rhs] == s[i] {
			rhs++
		}
		for lhs >= 0 && rhs < lth {
			if s[lhs] != s[rhs] {
				break
			}
			lhs--
			rhs++
		}
		if len(s[lhs+1:rhs]) > len(maxSub) {
			maxSub = s[lhs+1 : rhs]
		}
	}
	return maxSub

}

//7. Reverse Integer
func reverse(x int) int {
	num := 0
	for x != 0 {
		remain := x % 10
		x /= 10
		num = num*10 + remain
		if num < math.MinInt32 || num > math.MaxInt32 {
			return 0
		}
	}
	return num
}

//8. String to Integer (atoi)
func myAtoi(s string) int {
	num := 0
	start := 0
	isPositive := 1
	for ; start < len(s) && s[start] == ' '; start++ {
	}
	if start < len(s) && (s[start] == '+' || s[start] == '-') {
		if s[start] == '-' {
			isPositive = -1
		}
		start++
	}
	for i := start; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
		if num*isPositive > math.MaxInt32 {
			return math.MaxInt32
		}
		if num*isPositive < math.MinInt32 {
			return math.MinInt32
		}
	}
	return num * isPositive
}

//11. Container With Most Water
func maxArea(height []int) int {
	lhs := 0
	rhs := len(height) - 1
	maxContain := 0
	for lhs < rhs {
		line := 0
		diff := rhs - lhs
		if height[lhs] < height[rhs] {
			line = height[lhs]
			lhs++
		} else {
			line = height[rhs]
			rhs--
		}
		contain := line * diff
		if contain > maxContain {
			maxContain = contain
		}

	}
	return maxContain
}

func maxAreaV2(height []int) int {
	lth := len(height)
	lhs := 0
	rhs := lth - 1
	maxContain := 0
	for lhs < rhs {
		line := 0
		if height[lhs] < height[rhs] {
			line = height[lhs]
		} else {
			line = height[rhs]
		}
		contain := line * (rhs - lhs)
		if contain > maxContain {
			maxContain = contain
		}
		for lhs < rhs && height[lhs] <= line {
			lhs++
		}
		for lhs < rhs && height[rhs] <= line {
			rhs--
		}

	}
	return maxContain
}

//15. 3Sum
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	lth := len(nums)
	var res [][]int
	for i := 0; i <= lth-3 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lhs := i + 1
		rhs := lth - 1
		target := 0 - nums[i]
		for lhs < rhs {
			if nums[rhs]+nums[lhs] == target {
				res = append(res, []int{nums[i], nums[lhs], nums[rhs]})
				lhs++
				rhs--
				for lhs < rhs && nums[lhs] == nums[lhs-1] {
					lhs++
				}
				for lhs < rhs && nums[rhs] == nums[rhs+1] {
					rhs--
				}
			} else if nums[rhs]+nums[lhs] > target {
				rhs--
			} else {
				lhs++
			}
		}
	}
	return res

}

//17. Letter Combinations of a Phone Number
//TODO string:range方法和取下标的区别
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var digit2letter = map[int32]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var res = []string{""}
	for _, digit := range digits {
		res = letterCombinationsHelp(digit2letter[digit], res)
	}
	return res
}

func letterCombinationsHelp(letters string, cur []string) []string {
	var res []string
	for _, str := range cur {
		for _, letter := range letters {
			tmp := str + string(letter)
			res = append(res, tmp)
		}
	}
	return res
}

//19. Remove Nth Node From End of List
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
	}
	if slow.Next == nil {
		return nil
	}
	slow.Next = slow.Next.Next
	return head
}

//22. Generate Parentheses
func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisHelp(n, 1, 0, "(", &res)
	return res
}

func generateParenthesisHelp(n, left, right int, curParenthesis string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, curParenthesis)
		return
	}
	if left > right && left < n && right < n {
		generateParenthesisHelp(n, left+1, right, curParenthesis+"(", res)
		generateParenthesisHelp(n, left, right+1, curParenthesis+")", res)
	} else if left == right && left < n {
		generateParenthesisHelp(n, left+1, right, curParenthesis+"(", res)
	} else if left == n && right < n {
		generateParenthesisHelp(n, left, right+1, curParenthesis+")", res)
	}
}

//29. Divide Two Integers
func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	isPositive := 1

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		isPositive = -1
	}
	if dividend < 0 {
		dividend *= -1
	}
	if divisor < 0 {
		divisor *= -1
	}
	cnt := 0
	sum := 0
	for sum < dividend {
		sum += divisor
		cnt++
		if cnt*isPositive < math.MinInt32 {
			return math.MinInt32
		}
		if cnt*isPositive > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if sum == dividend {
		return cnt * isPositive
	} else {
		return (cnt - 1) * isPositive
	}
}

func divideV2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	dendSigned := 1
	if dividend < 0 {
		dendSigned = 0
		dividend *= -1
	}
	sorSigned := 1
	if divisor < 0 {
		sorSigned = 0
		divisor *= -1
	}
	signed := 1
	if dendSigned^sorSigned == 1 {
		signed = -1
	}
	ret := 0

	for dividend >= divisor {
		tmp, m := divisor, 1
		for dividend >= tmp<<1 {
			tmp, m = tmp<<1, m<<1
		}
		dividend -= tmp
		ret += m
	}
	return signed * ret

}

//33. Search in Rotated Sorted Array
func search(nums []int, target int) int {
	lhs := 0
	rhs := len(nums) - 1
	for lhs <= rhs {
		mid := (lhs + rhs) / 2
		if target == nums[mid] {
			return mid
			//右半部分排好序
		}
		if nums[mid] < nums[rhs] {
			if target > nums[mid] && target <= nums[rhs] {
				lhs = mid + 1
			} else {
				rhs = mid - 1
			}
		} else if nums[mid] >= nums[lhs] {
			if target >= nums[lhs] && target < nums[mid] {
				rhs = mid - 1
			} else {
				lhs = mid + 1
			}
		}
	}
	return -1
}

//34. Find First and Last Position of Element in Sorted Array
func searchRange(nums []int, target int) []int {
	lhs := 0
	rhs := len(nums) - 1
	index := -1
	for lhs <= rhs {
		mid := (lhs + rhs) / 2
		if target == nums[mid] {
			index = mid
			break
		} else if nums[mid] < target {
			lhs = mid + 1
		} else {
			rhs = mid - 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}
	start, end := index, index
	for ; start >= 0 && nums[start] == target; start-- {
	}
	for ; end < len(nums) && nums[end] == target; end++ {
	}
	return []int{start + 1, end - 1}

}

//36. Valid Sudoku
func isValidSudoku(board [][]byte) bool {
	rowRecord := make([]map[byte]bool, 9)
	columnRecord := make([]map[byte]bool, 9)
	subRecord := make([]map[byte]bool, 9)
	for i := 0; i < len(rowRecord); i++ {
		rowRecord[i] = make(map[byte]bool)
		columnRecord[i] = make(map[byte]bool)
		subRecord[i] = make(map[byte]bool)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] < '1' && board[i][j] > '9' || board[i][j] == '.' {
				continue
			}
			if _, ok := rowRecord[i][board[i][j]]; ok {
				return false
			}

			if _, ok := columnRecord[j][board[i][j]]; ok {
				return false
			}
			rowRecord[i][board[i][j]] = true
			columnRecord[j][board[i][j]] = true
			subRow := i / 3
			subColumn := j/3 + 1
			subIndex := 3*subRow + subColumn
			if _, ok := subRecord[subIndex-1][board[i][j]]; ok {
				return false
			}
			subRecord[subIndex-1][board[i][j]] = true
		}
	}

	return true

}

//38. Count and Say
func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		digit := str[0]
		cnt := 0
		var tmp string
		for j := 0; j < len(str); j++ {
			if str[j] == digit {
				cnt++
			} else {
				tmp = tmp + strconv.Itoa(cnt) + string(digit)
				digit = str[j]
				cnt = 1
			}
		}
		tmp = tmp + strconv.Itoa(cnt) + string(digit)
		str = tmp
	}
	return str
}

//46. Permutations
func permute(nums []int) [][]int {
	var res [][]int
	recPermute(0, nums, &res)
	return res
}

func recPermute(start int, nums []int, res *[][]int) {
	if start == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		recPermute(start+1, nums, res)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

//48. Rotate Image
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for x := 0; x < len(matrix)/2; x++ {
		for y := x; y < n-x; y++ {
			t := matrix[x][y]
			matrix[x][y] = matrix[n-y][x]
			matrix[n-y][x] = matrix[n-x][n-y]
			matrix[n-x][n-y] = matrix[y][n-x]
			matrix[y][n-x] = t
		}
	}
}

//49. Group Anagrams

func groupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string)
	for _, v := range strs {
		ana := [26]int{}
		for _, c := range v {
			ana[c-'a']++
		}

		if _, ok := dict[ana]; !ok {
			dict[ana] = make([]string, 0)
		}
		dict[ana] = append(dict[ana], v)
	}

	res := make([][]string, 0)
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

//50. Pow(x, n)
func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x, n = x*x, n>>1
	}
	return res
}

//53. Maximum Subarray
func maxSubArray(nums []int) int {
	maxSubArr := nums[0]
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
		if dp[i][i] > maxSubArr {
			maxSubArr = dp[i][i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if i == j {
				continue
			} else if j-i == 1 {
				dp[i][j] = nums[i] + nums[j]
			} else {
				dp[i][j] = dp[i+1][j-1] + nums[i] + nums[j]
			}
			if dp[i][j] > maxSubArr {
				maxSubArr = dp[i][j]
			}
		}
	}
	return maxSubArr
}

func maxSubArrayV2(nums []int) int {
	lth := len(nums)
	dp := make([]int, lth)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < lth; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, dp[i])
	}
	return res
}

//54. Spiral Matrix
func spiralOrder(matrix [][]int) []int {
	var res []int
	row := len(matrix)
	column := len(matrix[0])
	iterRows := row / 2
	if row&1 == 0 {
		iterRows -= 1
	}
	isVisited := make([]int, row*column)
	for i := 0; i <= iterRows; i++ {
		curRow := i
		curColumn := i
		//从左向右
		for ; curColumn < column-i; curColumn++ {
			key := curRow*column + curColumn
			if isVisited[key] == 0 {
				res = append(res, matrix[curRow][curColumn])
				isVisited[key] = 1
			}

		}
		curColumn -= 1
		//从上向下
		for curRow += 1; curRow < row-i; curRow++ {
			key := curRow*column + curColumn
			if isVisited[key] == 0 {
				res = append(res, matrix[curRow][curColumn])
				isVisited[key] = 1
			}
		}
		curRow -= 1
		//从右向左
		for curColumn -= 1; curColumn >= i; curColumn-- {
			key := curRow*column + curColumn
			if isVisited[key] == 0 {
				res = append(res, matrix[curRow][curColumn])
				isVisited[key] = 1
			}

		}
		curColumn += 1
		//从下向上
		for curRow -= 1; curRow > i; curRow-- {
			key := curRow*column + curColumn
			if isVisited[key] == 0 {
				res = append(res, matrix[curRow][curColumn])
				isVisited[key] = 1
			}
		}
	}
	return res

}

func spiralOrderV2(matrix [][]int) []int {
	startRow := 0
	startColumn := 0
	endRow := len(matrix) - 1
	endColumn := len(matrix[0]) - 1
	caps := (endRow + 1) * (endColumn + 1)
	var res []int
	var cnt int
	for endRow >= startRow && endColumn >= startColumn {
		for i := startColumn; i <= endColumn && cnt < caps; i++ {
			res = append(res, matrix[startRow][i])
			cnt++
		}
		for i := startRow + 1; i <= endRow && cnt < caps; i++ {
			res = append(res, matrix[i][endColumn])
			cnt++
		}

		for i := endColumn - 1; i >= startColumn && cnt < caps; i-- {
			res = append(res, matrix[endRow][i])
			cnt++
		}

		for i := endRow - 1; i > startRow && cnt < caps; i-- {
			res = append(res, matrix[i][startColumn])
			cnt++
		}
		startRow++
		endRow--
		startColumn++
		endColumn--
	}
	return res
}

func spiralOrderV3(matrix [][]int) []int {
	var res []int
	row := len(matrix)
	column := len(matrix[0])
	iterRows := row / 2
	if row&1 == 0 {
		iterRows -= 1
	}
	caps := row * column
	cnt := 0
	for i := 0; i <= iterRows; i++ {
		curRow := i
		curColumn := i
		//从左向右
		for ; curColumn < column-i && cnt < caps; curColumn++ {
			res = append(res, matrix[curRow][curColumn])
			cnt++
		}
		curColumn -= 1
		//从上向下
		for curRow += 1; curRow < row-i && cnt < caps; curRow++ {
			res = append(res, matrix[curRow][curColumn])
			cnt++
		}
		curRow -= 1
		//从右向左
		for curColumn -= 1; curColumn >= i && cnt < caps; curColumn-- {
			res = append(res, matrix[curRow][curColumn])
			cnt++

		}
		curColumn += 1
		//从下向上
		for curRow -= 1; curRow > i && cnt < caps; curRow-- {
			res = append(res, matrix[curRow][curColumn])
			cnt++
		}
	}
	return res

}

//55. Jump Game
func canJump(nums []int) bool {
	//[2,3,1,1,4]
	var res bool
	recCanJump(nums, 0, &res)
	return res
}

func recCanJump(nums []int, position int, res *bool) {

	if position == len(nums)-1 {
		*res = true
		return
	}

	for i := 1; i <= nums[position]; i++ {
		if *res == true {
			return
		}
		if position+i < len(nums) {
			recCanJump(nums, position+i, res)
		}
	}

}

func canJumpV2(nums []int) bool {
	lth := len(nums) - 1
	if lth <= 0 {
		return true
	}
	res := make([]bool, lth+1)
	recCanJumpV2(lth, 0, nums, res)
	return res[lth]
}

func recCanJumpV2(lth int, startIndex int, nums []int, res []bool) {
	for i := 1; i <= nums[startIndex]; i++ {
		pos := i + startIndex
		if pos < lth {
			if res[pos] {
				continue
			}
			res[pos] = true
			recCanJumpV2(lth, pos, nums, res)
		} else if pos == lth {
			res[pos] = true

		} else {
			break
		}
	}
}

func canJumpV3(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	} else if nums[0] == 0 && length > 1 {
		return false
	}

	maxJump := nums[0]
	for i := 1; i < length-1; i++ {
		next := nums[i] + i
		if next > maxJump {
			maxJump = next
		}

		if i >= maxJump {
			return false
		}
	}

	return true
}

//牛逼
func canJumpV4(nums []int) bool {
	lastPos := len(nums) - 1 // last position can reach the end index
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func canJumpV5(nums []int) bool {
	memo := make([]int, len(nums))
	memo[len(nums)-1] = 1 // when index == len(nums)-1, able to reach the last index
	for i := len(nums) - 2; i >= 0; i-- {
		furthestJump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthestJump; j++ {
			if memo[j] == 1 {
				memo[i] = 1
				break
			}
		}
	}
	return memo[0] == 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//56. Merge Intervals
func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	start := intervals[0][0]
	end := intervals[0][1]
	lth := len(intervals)
	index := 0
	for index < lth {
		cmpStart := intervals[index][0]
		cmpEnd := intervals[index][1]
		if cmpStart <= end {
			if cmpEnd > end {
				end = cmpEnd
			}
		} else {
			res = append(res, []int{start, end})
			start = cmpStart
			end = cmpEnd
		}
		index++
	}
	res = append(res, []int{start, end})

	return res
}

//62. Unique Paths
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

//73. Set Matrix Zeroes
func setZeroes(matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])
	var isCZero bool
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			isCZero = true
		}
		for j := 1; j < column; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < column; i++ {
			matrix[0][i] = 0
		}
	}

	if isCZero {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}

}

//75 Sort Colors
func sortColors(nums []int) {
	lhs := 0
	rhs := len(nums) - 1
	i := 0
	for i <= rhs {
		if nums[i] == 1 {
			i++
		} else if nums[i] < 1 {
			nums[i], nums[lhs] = nums[lhs], nums[i]
			i++
			lhs++
		} else {
			nums[i], nums[rhs] = nums[rhs], nums[i]
			rhs--
		}
	}
}

//[78]Subsets	7,293.2%	Medium	0.0%

func subsets(nums []int) [][]int {
	var res [][]int
	subsetsHelp([]int{}, &res, 0, nums)
	return res
}

func subsetsHelp(baseNums []int, res *[][]int, next int, nums []int) {
	*res = append(*res, baseNums)
	if next == len(nums) {
		return
	}

	for i := next; i < len(nums); i++ {
		tmp := make([]int, len(baseNums))
		copy(tmp, baseNums)
		tmp = append(tmp, nums[i])
		subsetsHelp(tmp, res, i+1, nums)
	}
}

func subsetsV2(nums []int) [][]int {
	res := [][]int{[]int{}}
	for _, v := range nums {
		lth := len(res)
		for i := 0; i < lth; i++ {
			tmp := make([]int, len(res[i]))
			copy(tmp, res[i])
			tmp = append(tmp, v)
			res = append(res, tmp)
		}
	}
	return res
}

//79. Word Search
func existV1(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if existHelper(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func existHelper(board [][]byte, word string, i, j, matchCnt int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] == word[matchCnt] {
		if matchCnt == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = 0
		found := existHelper(board, word, i+1, j, matchCnt+1) || existHelper(board, word, i, j+1, matchCnt+1) || existHelper(board, word, i-1, j, matchCnt+1) || existHelper(board, word, i, j-1, matchCnt+1)
		board[i][j] = temp
		return found
	}
	return false
}

func existV2(board [][]byte, word string) bool {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == word[0] && dfs(board, word[1:], x, y) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, x, y int) bool {
	if len(word) == 0 {
		return true
	}

	tmp := board[x][y]
	board[x][y] = '0'
	if x-1 >= 0 && board[x-1][y] == word[0] && dfs(board, word[1:], x-1, y) {
		return true
	}
	if y-1 >= 0 && board[x][y-1] == word[0] && dfs(board, word[1:], x, y-1) {
		return true
	}
	if x+1 < len(board) && board[x+1][y] == word[0] && dfs(board, word[1:], x+1, y) {
		return true
	}
	if y+1 < len(board[0]) && board[x][y+1] == word[0] && dfs(board, word[1:], x, y+1) {
		return true
	}
	board[x][y] = tmp
	return false
}

//91. Decode Ways
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	lth := len(s)
	dp := make([]int, lth+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < lth; i++ {
		val := (s[i-1]-'0')*10 + (s[i] - '0')
		if s[i] >= '1' && s[i] <= '9' && val >= 10 && val <= 26 {
			dp[i+1] = dp[i] + dp[i-1]
		} else if s[i] >= '1' && s[i] <= '9' && (val < 10 || val > 26) {
			dp[i+1] = dp[i]
		} else if s[i] == '0' && val >= 10 && val <= 26 {
			dp[i+1] = dp[i-1]
		} else {
			return 0
		}
	}
	return dp[lth]

}

func numDecodingsV2(s string) int {
	lth := len(s)
	if string(s[0]) == "0" {
		return 0
	}
	dp := make([]int, lth+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < lth+1; i++ {
		dp[i] = 0
		if s[i-1] > '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			dp[i] += dp[i-2]

		}
	}
	return dp[lth]
}

//98. Validate Binary Search Tree
func isValidBST(root *TreeNode) bool {
	return RecValidate(root, nil, nil)
}

func RecValidate(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}
	return RecValidate(node.Left, min, node) && RecValidate(node.Right, node, max)
}

//102. Binary Tree Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		lth := len(stack)
		var tmp []int
		for i := 0; i < lth; i++ {
			tmp = append(tmp, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res = append(res, tmp)
		stack = stack[lth:]
	}

	return res
}

func levelOrderV2(root *TreeNode) [][]int {
	result := [][]int{}
	return traceTree(root, 0, result)
}

func traceTree(root *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}
	if len(result) < level+1 {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	result = traceTree(root.Left, level+1, result)
	result = traceTree(root.Right, level+1, result)
	return result
}

//103. Binary Tree Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	cnt := 0
	for len(stack) > 0 {
		lth := len(stack)
		var levelVals []int
		for i := 0; i < lth; i++ {
			levelVals = append(levelVals, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		if cnt&1 == 1 {
			lhs := 0
			rhs := len(levelVals) - 1
			for lhs < rhs {
				levelVals[lhs], levelVals[rhs] = levelVals[rhs], levelVals[lhs]
				lhs++
				rhs--
			}
		}
		res = append(res, levelVals)
		stack = stack[lth:]
		cnt++
	}
	return res
}

func zigzagLevelOrderV2(root *TreeNode) [][]int {
	res := [][]int{}
	return zigzagLevelOrderHelp(root, 0, res)
}

func zigzagLevelOrderHelp(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) < level+1 {
		res = append(res, []int{})
	}
	if level&1 == 1 {
		res[level] = append([]int{root.Val}, res[level]...)
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = zigzagLevelOrderHelp(root.Left, level+1, res)
	res = zigzagLevelOrderHelp(root.Right, level+1, res)
	return res
}

//105. Construct Binary Tree from Preorder and Inorder Traversal
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := targetIndex(inorder, preorder[0])
	leftLength := len(inorder[:index])
	//rightLength := len(inorder[index+1:])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:leftLength+1], inorder[:index]),
		Right: buildTree(preorder[leftLength+1:], inorder[index+1:]),
	}
	return root
}

func targetIndex(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//116. Populating Next Right Pointers in Each Node
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0 {
		lth := len(stack)
		for i := 0; i < lth; i++ {
			if i < lth-1 {
				stack[i].Next = stack[i+1]
			}

			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[lth:]
	}
	return root
}

func connectV2(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}

	connectV2(root.Left)
	connectV2(root.Right)

	return root
}

//122. Best Time to Buy and Sell Stock II
func maxProfitII(prices []int) int {
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxP += prices[i] - prices[i-1]
		}
	}
	return maxP
}

func maxProfitIIV2(prices []int) int {
	lth := len(prices)
	dp := make([][]int, lth)
	for i := 0; i < lth; i++ {
		dp[i] = make([]int, 2)
	}
	//持有
	dp[0][0] = -prices[0]
	//未持有
	dp[0][1] = 0
	for i := 1; i < lth; i++ {
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return max(dp[lth-1][0], dp[lth-1][1])
}

func maxProfitIIV3(prices []int) int {
	lth := len(prices)
	hold := -prices[0]
	noHold := 0
	for i := 1; i < lth; i++ {
		hold = max(hold, noHold-prices[i])
		noHold = max(noHold, hold+prices[i])
	}
	return noHold
}

func longestConsecutive(nums []int) int {
	child2father := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		child2father[nums[i]] = nums[i]
	}
	for k, _ := range child2father {
		if _, ok := child2father[k+1]; ok {
			child2father[k] = child2father[k+1]
		}
	}

	for c, f := range child2father {
		father := f
		for {
			if v, ok := child2father[father]; ok && v != father {
				child2father[c] = child2father[father]
				father = v
			} else {
				break
			}
		}
	}
	maxC := 0
	for c, f := range child2father {
		if f-c+1 > maxC {
			maxC = f - c + 1
		}
	}
	return maxC
}

func longestConsecutiveV2(nums []int) int {
	record := make(map[int]bool)
	for _, num := range nums {
		record[num] = true
	}
	maxC := 0
	for _, num := range nums {
		if _, ok := record[num-1]; !ok {
			curNum := num + 1
			tmpCon := 1
			for {
				if _, tmpOk := record[curNum]; tmpOk {
					curNum++
					tmpCon++
				} else {
					break
				}
			}
			if tmpCon > maxC {
				maxC = tmpCon
			}
		}

	}
	return maxC
}

type Dsu struct {
	parent map[int]int
}

func NewDsu(nums []int) *Dsu {
	parent := make(map[int]int)
	for _, num := range nums {
		parent[num] = num
	}
	return &Dsu{parent: parent}
}

func (d *Dsu) find(child int) int {
	if father, ok := d.parent[child]; ok && child != father {
		d.parent[child] = d.find(father)
	}
	return d.parent[child]
}

func (d *Dsu) connect(child int, father int) {
	d.parent[child] = d.find(father)
}

func longestConsecutiveV3(nums []int) int {
	dsu := NewDsu(nums)
	for _, num := range nums {
		if _, ok := dsu.parent[num-1]; ok {
			dsu.connect(num-1, num)
		}
		if _, ok := dsu.parent[num+1]; ok {
			dsu.connect(num, num+1)
		}
	}
	for _, num := range nums {
		dsu.find(num)
	}
	max := 0
	for child, father := range dsu.parent {
		if father-child+1 > max {
			max = father - child + 1
		}
	}
	return max
}

//130. Surrounded Regions
func solve(board [][]byte) {
	row := len(board)
	column := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 || j == 0 || i == row-1 || j == column-1 {
				if board[i][j] == 'O' {
					solveHelp(board, i, j)
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

		}
	}
}

func solveHelp(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '#'
		solveHelp(board, i-1, j)
		solveHelp(board, i+1, j)
		solveHelp(board, i, j-1)
		solveHelp(board, i, j+1)
	}
	return
}

type UF struct {
	Parent []int
	Rank   []int //rank[i] = rank of subtree rooted at i
	Count  int
}

func NewUf(n int) *UF {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	count := n
	return &UF{
		Parent: parent,
		Rank:   rank,
		Count:  count,
	}
}

func (u *UF) find(p int) int {
	for p != u.Parent[p] {
		u.Parent[p] = u.Parent[u.Parent[p]]
		p = u.Parent[p]
	}
	return p
}

func (u *UF) connect(p int, q int) {
	pp := u.find(p)
	qp := u.find(q)
	if pp == qp {
		return
	}
	if u.Rank[pp] < u.Rank[qp] {
		u.Parent[pp] = qp
	} else if u.Rank[pp] > u.Rank[qp] {
		u.Parent[qp] = pp
	} else {
		u.Parent[qp] = pp
		u.Rank[pp]++
	}
	u.Count--
}

func (u *UF) isConnect(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func solveV2(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	column := len(board[0])
	uf := NewUf(row*column + 1)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if (i == 0 || j == 0 || i == row-1 || j == column-1) && board[i][j] == 'O' {
				uf.connect(i*column+j, row*column)
			} else if board[i][j] == 'O' {
				if board[i-1][j] == 'O' {
					uf.connect(i*column+j, (i-1)*column+j)
				}
				if board[i+1][j] == 'O' {
					uf.connect(i*column+j, (i+1)*column+j)
				}
				if board[i][j-1] == 'O' {
					uf.connect(i*column+j, i*column+j-1)
				}
				if board[i][j+1] == 'O' {
					uf.connect(i*column+j, i*column+j+1)
				}

			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == 'O' && !uf.isConnect(i*column+j, row*column) {
				board[i][j] = 'X'
			}
		}
	}
}

//131. Palindrome Partitioning
func partition(s string) [][]string {
	var list []string
	var res [][]string
	start := 0
	for i := start; i < len(s); i++ {
		if isPal(s[start : i+1]) {
			partitionRec(s, s[start:i+1], i, list, &res)
		}

	}

	return res
}

func partitionRec(s string, substr string, end int, list []string, res *[][]string) {
	if !isPal(substr) {
		return
	}
	list = append(list, substr)
	if end == len(s)-1 {
		*res = append(*res, list)
		return
	}

	for i := end + 1; i < len(s); i++ {
		tmp := make([]string, len(list))
		copy(tmp, list)
		partitionRec(s, s[end+1:i+1], i, tmp, res)
	}

}

func partitionV2(s string) [][]string {
	dp := make([][][]string, len(s))
	dp[0] = append(dp[0], []string{string(s[0])})
	for i := 1; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			if isPal(s[j : i+1]) {
				if j-1 >= 0 {
					for _, list := range dp[j-1] {
						temp := make([]string, len(list))
						copy(temp, list)
						temp = append(temp, s[j:i+1])
						dp[i] = append(dp[i], temp)
					}
				}
				if j == 0 {
					dp[i] = append(dp[i], []string{s[j : i+1]})
				}

			}
		}

	}
	return dp[len(s)-1]
}

func isPal(str string) bool {
	lhs := 0
	rhs := len(str) - 1
	for lhs < rhs {
		if str[lhs] != str[rhs] {
			return false
		}
		lhs++
		rhs--
	}
	return true
}

//134. Gas Station
func canCompleteCircuit(gas []int, cost []int) int {
	return 0
}

func main() {
	//board := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	//fmt.Println(isValidSudoku(board))
	//["abbbbbbbbbbb","aaaaaaaaaaab"]
	//fmt.Println(numDecodings("226"))
	//nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(partition("aab"))
}
