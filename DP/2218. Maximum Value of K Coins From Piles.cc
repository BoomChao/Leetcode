
// dp[i][k] 代表从(piles[i]...piles[n]) pick k 个元素的总和, dp[0][k] 即为我们的答案
// dp[i][k] = dp[i+1][k-j-1] + sum(piles[i][...j])
// 注意上面是 k-j-1, 这是因为 j=0 表示第一个元素

int maxValuesOfCoins(std::vector<std::vecotr<int>> &piles, int K) {
    int n = piles.size();
    std::vector<std::vector<int>> dp(n+1, std::vector<int>(K+1, 0));

    for(int i = n; i >= 0; i--) {
        for(int k = K; k >= 0; k--) {
            if(k == 0 || i == n) {
                dp[i][k] = 0;
                continue;
            }
            int res = dp[i+1][k], cur = 0;
            for(int j = 0; j < std::min(int(piles[i].size()), k); j++) {
                cur += piles[i][j];
                res = std::max(res, dp[i+1][k-j-1] + cur);
            }
            dp[i][k] = res;
        }
    }

    return dp[0][K];
}