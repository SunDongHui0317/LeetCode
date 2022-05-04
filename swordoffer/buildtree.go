package swordoffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历确定根结点，中序遍历确定左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	i := 0

	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	root := new(TreeNode)
	root.Val = preorder[0]

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
