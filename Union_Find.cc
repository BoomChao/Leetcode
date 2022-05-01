
/*
    并查集: 查询(find)、联合(union)的时间复杂度都是O(logn)

    第547题 : 相连的省份
    第684题 : 删除一个图的一个边使得其能构成一棵树
    第399题 : 计算除法(较难)
*/

class UF
{
private:
    int count;                  //连通分量个数
    std::vector<int> parent;    //存储一棵树
    std::vector<int> size;      //记录每棵树的重量

public:
    UF(int n) 
    {
        count = n;
        parent.resize(n);
        size.resize(n);

        for(int i = 0; i < n; i++) {
            parent[i] = i;
            size[i] = 1;
        }
    }

    void unionNode(int p, int q) 
    {
        int rootP = findNode(p);
        int rootQ = findNode(q);

        //已经在一颗树上了就不用再联结了
        if(rootP == rootQ) return;

        //小树接到大树下面,较平衡
        if(size[rootP] > size[rootQ]) {
            parent[rootQ] = rootP;
            size[rootP] += size[rootQ];
        }
        else {
            parent[rootP] = rootQ;
            size[rootQ] += size[rootP];
        }

        count--;    //连通分量减1
    }

    int findNode(int x) 
    {
        while(parent[x] != x) {
            parent[x] = parent[parent[x]];  //这样可以压缩路径,一下走两步
            x = parent[x];
        }

        return x;
    }

    bool connected(int p, int q)
    {
        int rootP = findNode(p);
        int rootQ = findNode(q);
        
        return rootP == rootQ;
    }

    //返回连通分量个数
    int countNum() {
        return count;
    }
    
};


//Leetcode第547题 : 相连的省份

//方法一:使用并查集

int findCircleNum(std::vector<std::vector<int>> &isConnected)
{
    int n = isConnected.size();

    UF uf(n);

    for(int i = 0; i < n; i++) {
        for(int j = i + 1; j < n; j++) {
            if(isConnected[i][j]) {
                uf.unionNode(i, j);
            }
        }
    }

    return uf.countNum();
}


//方法二：DFS

int findCircleNum(std::vector<std::vector<int>> &isConnected)
{
    int n = isConnected.size();
    int count = 0;

    std::vector<int> visited(n, 0);

    for(int i = 0; i < n; i++) {
        if(visited[i] == 0) {
            dfs(isConnected, visited, i);
            count++;
        }
    }

    return count;
}

void dfs(std::vector<std::vector<int>> &nums, std::vector<int> &visited, int i)
{
    for(int j = 0; j < nums.size(); j++) {
        if(nums[i][j] == 1 && visited[j] == 0) {    //说明(i,j)个人是好友，将j标记已经访问
            visited[j] = 1;                 
            dfs(nums, visited, j);
        }
    }
}


//Leetcode第684题

//1.使用并查集
int getParent(std::vector<int> &parent, int v)
{
    if(parent[v] == 0) return v;

    parent[v] = getParent(parent, parent[v]);

    return parent[v];
}

std::vector<int> findRedundantConnection(std::vector<std::vector<int>> &edges)
{
    std::vector<int> parent(edges.size() + 1);

    for(auto &e:edges)
    {
        int p = getParent(parent, e[0]);
        int q = getParent(parent, e[1]);

        if(p == q) return {e[0], e[1]};
        parent[p] = q;
    }

    return {};
}



//Leetcode第399题: 计算除法

std::vector<double> calcEquation(std::vector<std::vector<std::string>> &equations, std::vector<double> &values, std::vector<std::vector<std::string>> &queries)
{
    std::unordered_map<std::string, Node*> map;
    std::vector<double> res;

    for(int i = 0; i < equations.size(); i++)
    {
        std::string s1 = equations[i][0], s2 = equations[i][1];

        //让s1做分子，s2做分母
        if(!map.count(s1) && !map.count(s2)) {
            map[s1] = new Node();
            map[s2] = new Node();
            map[s1]->value = values[i];
            map[s2]->value = 1;
            map[s1]->child = map[s2];
        } 
        else if(!map.count(s1)) {
            map[s1] = new Node();
            map[s1]->value = map[s2]->value * values[i];
            map[s1]->child = map[s2];
        }
        else if(!map.count(s2)) {
            map[s2] = new Node();
            map[s2]->value = map[s1]->value / values[i];
            map[s2]->child = map[s1];
        }
        else {
            unionNode(map[s1], map[s2], values[i], map);
        }
    }

    for(auto &query : queries)
    {
        if(!map.count(query[0]) || !map.count(query[1]) || findChild(map[query[0]]) != findChild(map[query[1]])) {
            res.push_back(-1);
        } else {
            res.push_back(map[query[0]]->value / map[query[1]]->value);
        }
    }

    return res;
}

struct Node {
    Node *child;
    double value = 0.0;
    Node(): child(this) {}
};

Node* findChild(Node *node)
{
    if(node->child == node) return node;

    node->child = findChild(node->child);

    return node->child;
}

void unionNode(Node *node1, Node *node2, double num, std::unordered_map<std::string, Node*> &map)
{
    Node *child1 = findChild(node1), *child2 = findChild(node2);

    double ratio = node2->value / node1->value * num;

    for(auto it = map.begin(); it != map.end(); it++) {
        if(findChild(it->second) == child1) {
            it->second->value *= ratio;
        }
    }

    child1->child = child2;
}



/*
    求数组的交集,合并所有的子数组,最终返回一个互相没有交集的数组列表

    如数组 arr = {{1,2,3}, {3,5,4}, {8,6}, {3,6}, {7,222}, {7,32}}
    返回 {{1,2,3,5,4,6,8}, {7,222,32}}
*/


std::map<int,int> parent;       //注意这个父节点存的是当前父亲，比如 {8,6}, parent[6]=8, {3,6}, 合并后parent[8]=3, 虽然6,8的祖节点都是3，但是parent[6]还是等于8
                                
std::map<int,std::vector<int>> mp;

int findNode(int x)
{
    while(parent[x] != x) {
        parent[x] = parent[parent[x]];
        x = parent[x];
    }

    return x;
}

void unionNode(int x, int y)
{
    if(!parent.count(x)) {
        parent[x] = x;
    }
    if(!parent.count(y)) {
        parent[y] = y;
    }

    int m_x = findNode(x);
    int m_y = findNode(y);

    parent[m_y] = m_x;
}

std::vector<std::vector<int>> merge(std::vector<std::vector<int>> &arr)
{
    std::vector<std::vector<int>> res;
 
    for(auto &nums : arr)
    {
        if(nums.size() == 1) mp[nums[0]].push_back(nums[0]);
        else {
            for(int i = 1; i < nums.size(); i++)
            {
                unionNode(nums[0], nums[i]);
            }
        }
    }

    for(auto &item : parent) {
        mp[findNode(item.first)].push_back(item.first);     //注意这个一定是mp[findNode(item.first)], 不要写成parent[item.second],因为这个parent不一定是祖节点
    }

    for(auto item : mp) {
        res.push_back(item.second);
    }

    return res;
}
