
// 迭代返回每个 [left, right, res]
// left is the maximum length in direction of root.left
// right is the maximum length in direction of root.right
// result is the maximum length in the whole sub tree.

// left = dfs(root->left), right = dfs(root->right)
// res = max(max(left[1], right[0]) + 1, max(left[2], right[2]))

int longestZigZag(TreeNode* root) {
    return dfs(root)[2];
}

std::vector<int> dfs(TreeNode *root) {
    if(!root) return {-1, -1, -1};
    auto left = dfs(root->left), right = dfs(root->right);
    int res = std::max(max(left[1], right[0]) + 1, std::max(left[2], right[2]));
    return {left[1]+1, right[0]+1, res};
}

