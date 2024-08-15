//十大排序算法总结


/*
//冒泡排序
void BubbleSort(int num[], int length)
{
    for(int i = 0; i < length; i++){
        for(int j = length-2; j>=i; j--){   //不推荐使用这种排序,j=length-2容易出错
            if(num[j] > num[j+1]) swap(num[j], num[j+1]);
        }
    }

}
*/

/*
//改进的冒泡排序(如果序列已经有序,则不需要再继续后面的循环判断工作了)
//避免因已经有序的情况下的无意义的循环判断
void BubbleSort(int num[], int length)
{
    bool flag = true;

    for(int i = 0; i < length && flag; i++){
        flag = false;
        for(int j = length-2; j>=i; j--){
            if(num[j] > num[j+1]){
                swap(num[j], num[j+1]);
                flag = true;
            }
            
        }
    }

}
*/


//及时终止的冒泡排序(推荐使用这种冒泡排序,好理解)
void BubbleSort(int num[], int length)
{
    bool flag = true;

    for(int i = length-1; i>0 && flag; i--){
        flag = false;
        for(int j = 0; j < i; j++){
            if(num[j] > num[j+1]){
                swap(num[j], num[j+1]);
                flag = true;
            }
        }
    }

}


//简单选择排序
void SelectSort(int num[], int length)
{
    int min;
    for(int i = 0; i < length-1; i++){
        min = i;
        for(int j = i+1; j < length; j++){
            if(num[min] > num[j]) min = j;
        }
        if(min!=i) swap(num[min], num[i]);
    }

}


/*
//直接插入排序
void InsertSort(int num[], int length)
{
    int i, j;
    for(i = 1; i < length; i++){
        if(num[i-1] > num[i]){
            int temp = num[i];
            for( j = i-1; num[j]>temp && j>=0; j--){    //防止下标越界,一定要加j>=0
                num[j+1] = num[j];
            }
            num[j+1] = temp;
        }
    }

}
*/

//插入排序(这种好理解)
void InsertSort(int num[], int len)
{
    for(int i = 1; i < length; i++){
        int tmp = num[i];
        int j = i-1;
        while(j >= 0 && num[j] > tmp){         //temp < num[j]是升序排列,即从小到大
            num[j+1] = num[j];              //temp > num[j]是降序排列,即从大到小
            j--;
        }
        num[j+1] = tmp;
    }

}



//堆排序 注:传入的数组num的第一个元素为0或者-1,让元素计数从1开始
//比如要对{30,50,80,40}这几个数排序
//输入为 num[]={-1,30,50,80,40}, len = sizeof(num)/sizeof(num[0])-1;

void heapSort(int num[], int len)   //这里传进去的len数值要比实际的长度少一个,因为前面的-1不算,是用作哨兵
{
    int i;      
    for(i = len/2; i > 0; i--)
    {     
        heapAdjust(num, i, len);
    }

    for(i = len; i > 1; i--)
    {    
        //注意:是与num[1]交换,不是与num[0]交换
        std::swap(num[1], num[i]);

        //交换之后再构建大顶堆元素个数就少了1个
        heapAdjust(num, 1, i-1);
    }

}

//调整成大顶堆
void heapAdjust(int *num, int s, int m)
{
    int temp = num[s];

    for(int j = 2*s; j <= m; j *= 2)
    {
        if(j < m && num[j] < num[j+1]) j++; 

        if(temp >= num[j]) break;
        
        num[s] = num[j];

        s = j;
    }
    
    num[s] = temp;
}



//希尔排序(在直接插入排序的基础上引入了跨度,即跨度不再是1)
void ShellSort(int num[], int length)
{
    int i, j;
    int gap = length;
    do{
        gap = gap/3 + 1;    //引入跨度
        for(int i = gap; i < length; i+=gap){
            if(num[i] < num[i-gap]){
                int temp = num[i];
                for(j = i-gap; num[j]>temp && j>=0; j-=gap){
                    num[j+gap] = num[j];
                }
                num[j+gap] = temp;
            }
        }

    }while(gap>1);

}


//归并排序(递归实现)
void mergeSort(int *num, int len)
{
    if (len <= 1) return;

    //左半部分
    int *l1 = num;
    int l1_size = len/2;
    
    //右半部分
    int *l2 = num+l1_size;
    int l2_size = len-l1_size;

    mergeSort(l1, l1_size);
    mergeSort(l2, l2_size);
    merge(l1, l1_size, l2, l2_size);

}

void merge(int *l1, int l1_size, int *l2, int l2_size)
{
    int i = 0, j = 0, k = 0;

    //定义一个临时数组存储排序好的结果
    int tmp[l1_size + l2_size];

    while(i < l1_size && j < l2_size)
    {
        l1[i] < l2[j] ? tmp[k++] = l1[i++] : tmp[k++] = l2[j++];
    }

    //这两个while循环只可能有一个执行,不可能两个都执行
    while(i < l1_size) tmp[k++] = l1[i++];     
    while(j < l2_size) tmp[k++] = l2[i++];

    //排序完成的结果存储到数组l1中去
    //注意:不能存储到l2中去,因为在219行就是直接将l1指向了数组num首地址;而将l2指向了num后的l1_size个元素
    for(int m = 0; m < l1_size+l2_size; m++) 
    {
        l1[m] = tmp[m];    
    }

}



//归并排序(迭代实现) 注:传入的数组num的第一个元素为0或者-1,让元素计数从1开始
void mergeSort(int *num, int len)
{
    int *TR = new int[len];
    int k = 1;
    while(k < len){
        mergePass(num, TR, k, len);
        k = 2*k;
        mergePass(TR, num, k ,len);
        k = 2*k;
    }

}

void mergePass(int SR[], int TR[], int s, int n)
{
    int i = 1;
    while(i <= n-2*s+1){
        merge(SR, TR, i, i+s-1, i+2*s-1);
        i = i+2*s;
    }

    if(i < n-s+1) {
        merge(SR, TR, i, i+s-1, n);
    }else{
        for(int j = i; j <= n; j++)  TR[j] = SR[j];
    }

}


void merge(int SR[], int TR[], int i, int m, int n){
    int j,k,l;

    for(j=m+1,k=i; i<=m && j<=n; k++){
        SR[i]<SR[j] ? TR[k] = SR[i++] : TR[k] = SR[j++];
    }

    if(i <= m){
        for( l = 0; l <= m-i; l++)  TR[k+l] = SR[i+l];
    }

    if(j <= n){
        for( l = 0; l <= n-j; l++)  TR[k+l] = SR[j+l];
    }

}


//希尔排序
void shellSort(std::vector<int> &nums)
{
    int gap = nums.size();
    
    while(gap > 1)
    {
        gap = gap/3 + 1;
        for(int i = gap; i < nums.size(); i++)
        {
            if(nums[i] < nums[i-gap]) 
            {
                int tmp = nums[i];
                int j;
                for(j = i-gap; j >= 0 && nums[j] > tmp; j -= gap) {
                    nums[j + gap] = nums[j];
                }
                nums[j+gap] = tmp;
            }
        }
    }

}


void merge(int num[], int left, int mid, int right)
{
    int lIndex = left;
    int rIndex = mid + 1;
    int *team = new int[right-left+1];
    int teamIndex = 0;

    while(lIndex <= mid && rIndex <= right)
    {
        if(num[lIndex] <= num[rIndex]) {
            team[teamIndex++] = num[lIndex++];
        }
        else {
            team[teamIndex++] = num[rIndex++];
        }
    }

    while(lIndex <= mid) {
        team[teamIndex++] = num[lIndex++];
    }

    while(rindex <= right) {
        team[teamIndex++] = num[rIndex++];
    }

    //将排序好的数组传给num
    for(int i = 0; i < teamIndex; i++) {
        num[left+i] = team[i];
    }

    delete[] team;

}

//归并排序(递归实现)
void mergeSort(int num[], int left, int right)
{
    int mid = left + ((right-left)>>1);

    if(left < right)    //记住不要写成了while
    {
        mergeSort(num, left, mid);
        mergeSort(num, mid + 1, right);
        merge(num, left, mid, right);
    }

}


void mergePass(int num[], int k, int len)
{
    int i;
    for(i = 0; i < len-2*k; i += 2*k) {
        merge(num, i, i+k-1, i+2*k-1);
    }

    //归并最后两个序列
    if(i + k < len) {
        merge(num, i, i+k-1, len-1);
    }

}


//归并排序(迭代实现)
void mergeSort(int num[], int n)
{
    int k = 1;
    while(k < n) 
    {
        mergePass(num, k, n);
        k *= 2;
    }

}


//快速排序
void quickSort(int num[], int low, int hig)
{
    if(low < hig)
    {
        //算出枢轴的位置
        int pivot = partition(num, low, hig);

        //对低子表进行递归排序
        quickSort(num, low, pivot-1);   
        
        //对高子表进行递归排序
        quickSort(num, pivot+1, hig); 
    }

}

int partition(int num[], int low, int hig)
{
    int pivotkey = num[low];

    while(low < hig)        //注意:这里不能写成low<=hig,如果是low等于hig则下面任何语句都不会执行,则程序会进入死循环
    {
        while(low < hig && num[hig] >= pivotkey) {
            hig--;
        }
        std::swap(num[low], num[hig]);  //将比枢轴小的元素放在左边

        while(low < hig && num[low] <= pivotkey) {
            low++;
        }
        std::swap(num[low], num[hig]);  //将比枢轴大的元素放在右边
    }

    return low;  //返回枢轴的位置
}



//优化后的快速排序
void quickSort(int *num, int low, int hig)
{
    while(low < hig)
    {
        int pivot = Partition(num, low, hig);

        quickSort(num, low, pivot-1);   //这样少用一次递归
        
        low = pivot + 1;
    }

}


int Partition(int *num, int low, int hig)
{
    //三数取中选取枢纽
    int mid = (low+hig)/2;
    if(num[low] > num[hig]) std::swap(num[low], num[hig]);    
    if(num[mid] > num[hig]) std::swap(num[mid], num[hig]);    //这两步是保证右端hig是最大的
    if(num[mid] > num[low]) std::swap(num[mid], num[low]);    //保证中端mid是最小的,并且左端low是次小的

    int pivotkey = num[low];   

    while(low < hig)
    {
        while(low < hig && num[hig] >= pivotkey) {
            hig--;
        } 
        num[low] = num[hig];    //采用替换而不是交换的方式进行

        while(low < hig && num[low] <= pivotkey) {
            low++;
        }   
        num[hig] = num[low];

    }

    num[low] = pivotkey;

    return low;
}


