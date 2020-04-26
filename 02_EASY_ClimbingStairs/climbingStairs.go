//dp solution
func climbStairs(n int) int {
    //base case
    if(n==0){
        return 0
    }
    if(n==1){
        return 1
    }
    aa:=make([]int, n)
    aa[0]=1
    aa[1]=2
    for i:=2;i<len(aa);i++{
        aa[i]=aa[i-1]+aa[i-2]
    }
    return aa[n-1]
}