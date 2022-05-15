package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//判断平衡二叉树
func IsBalanced_Solution(pRoot *TreeNode) bool {
	return pRoot == nil || IsBalanced_Solution(pRoot.Left) &&
		IsBalanced_Solution(pRoot.Right) && math.Abs(height(pRoot.Left)-height(pRoot.Right)) < 2
}
func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(height(root.Left), height(root.Right)) + 1
}


func main(){
	//读取输入
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	str = strings.Replace(str,"\n", " ", -1)
	str = strings.Replace(str,"\n", " ", -1)

	//读取输入
	var str2 string
	fmt.Scanln(&str2)
	str2 = strings.Replace(str2,"\n"," ",-1)

	//fmt.Printf("str:%s",str)
	//fmt.Printf("str2:%s",str2)
	//统计
	num := 0
	for _,v := range str{
		if strings.ToLower(string(v)) == strings.ToLower(str2) {
			num++
		}
	}
	fmt.Println(num)

}
