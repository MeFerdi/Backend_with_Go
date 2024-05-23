package piscine

func isBSTUtil(root *TreeNode, min, max string) bool {
	if root == nil {
		return true
	}

	if min != "" && root.Data <= min || max != "" && root.Data >= max {
		return false
	}

	return isBSTUtil(root.Left, min, root.Data) && isBSTUtil(root.Right, root.Data, max)
}

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTUtil(root, "", "")
}
