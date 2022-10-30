package main

func main() {
    buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    builder := &Builder{inorder, postorder}

    return builder.build(0, len(builder.inorder)-1, 0, len(builder.postorder)-1)
}

type Builder struct {
    inorder []int
    postorder []int
}

func (b *Builder) build(l1, r1, l2, r2 int) *TreeNode {
    if l1 > r1 {
        return nil
    }
    root := &TreeNode{Val: b.postorder[r2]}
    var i = l1
    for b.inorder[i] != root.Val {
        i++
    }
    root.Left = b.build(l1, i-1, l2, l2+i-l1-1)

    root.Right = b.build(i+1, r1, l2+i-l1, r2-1)


    return root
}
