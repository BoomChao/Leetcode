/*
字符串乘法，输出 2 的 1000 次方
*/

// 基本思路
// 肯定不能直接转化成int来做pow,因为会溢出
// 使用字符串乘法
// 1. 2^1000 就等价于 字符串*2 然后循环1000次
// 2. 怎么计算字符串*2呢,使用栈的方式
// 3. 比如249*2,从末尾向前计算, 9*2=18,18/10=1,所以进位为1，当前位为8
// 4. 将当前位入栈,然后下次计算继续计算带上进位的数即可

std::string twoTimes(std::string s) {
    std::string res;
    if(s.empty()) return res;

    int n = s.size() - 1;
    std::stack<std::string> sta;

    int carry = 0;
    for(int i = n; i >= 0; i--) {
        carry += (s[i] - '0') * 2;
        sta.push(std::to_string(carry%10));
        carry /= 10;
    }
    
    if(carry != 0) {
        sta.push(std::to_string(carry));
    }

    while(!sta.empty()) {
        res += sta.top();
        sta.pop();
    }

    return res;
}

std::string powN(std::string str, int n) {
    std::string res;
    for(int i = 0; i < n; i++) {
        res = twoTimes(str);
        str = res;
    }
    return res;
}


/*
给定一个文件路径的列表，将列表渲染成一颗文件树
输入:
"/root/path_a/1.txt"
"/root/path_b/3.txt"
"/root1/4.txt"
"/root/path_a/2.txt"
输出:
- root
    - path_a
        - 1.txt
        - 2.txt
    - path_b
        - 3.txtk
- root1
    -4 .txt
*/