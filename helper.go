package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// this is so clever trick
const null = math.MinInt

func buildList(a []int) *ListNode {
	var prev *ListNode
	for i := len(a) - 1; i >= 0; i-- {
		tp := ListNode{a[i], prev}
		prev = &tp
	}
	return prev
}

func buildTree(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	head := &TreeNode{
		Val: a[0],
	}
	queue := []*TreeNode{head}
	i := 1
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		var left, right *TreeNode
		if i < len(a) && a[i] != null {
			left = &TreeNode{
				Val: a[i],
			}
			queue = append(queue, left)
		}
		i++
		if i < len(a) && a[i] != null {
			right = &TreeNode{
				Val: a[i],
			}
			queue = append(queue, right)
		}
		i++
		top.Left = left
		top.Right = right
	}
	return head
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	tp := findNode(root.Left, val)
	if tp != nil {
		return tp
	}
	return findNode(root.Right, val)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func call(f any, args ...any) {
	typ := reflect.TypeOf(f)
	assertEq(typ.NumIn(), len(args))
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		assertEq(typ.In(i), reflect.TypeOf(arg))
		in[i] = reflect.ValueOf(arg)
	}
	v := reflect.ValueOf(f)

	// print function
	fmt.Printf("%s( ", funcName(f))
	for _, v := range in {
		fmt.Print(v, " ")
	}
	fmt.Print(") => ")
	for _, v := range v.Call(in) {
		fmt.Print(v.Interface(), " ")
	}
	fmt.Println()
}

func assertEq(a any, b any) {
	if a != b {
		log.Fatalf("got: %v, expected: %v", b, a)
	}
}

func funcName(temp any) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
