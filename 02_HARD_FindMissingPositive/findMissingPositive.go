func firstMissingPositive(nums []int) int {
    //check if 1 exist in the array
    //return if not exist.
    //we will use 1 as a default value for negative numbers, zeros-
    //and numbers bigger than length of nums
    f:= true
    for i:=0;i<len(nums);i++{
        if(nums[i]==1){
            f=false
            break
        }
    }
    //if 1 doesn't exist return 1
    if f {
        return 1
    }
    //edge case where nums contain 1 element with value 1
    if nums[0]==1 && len(nums)==1{
        return 2
    }
    //set all negative numbers, numbers bigger than length of nums-
    //and zeros to 1
    for i:=0;i<len(nums);i++{
        if nums[i]>len(nums) || nums[i]<=0{
            nums[i]=1
        }
    }
    //set the state of number exist in the array by setting
    //the abs(num[i])-1th index in array as negative
    for i:=0;i<len(nums);i++{
        val:=abs(nums[i])-1
        if(nums[val]>0 ){
            nums[val]=-nums[val]
        }
    }
    //get the first positive number. i+1the element is the answer
    for i:=1;i<len(nums);i++{
        if(nums[i]>0){
            return i+1
        }
    }
    //if array is continuous sequence from 1 then return len(nums)+1
    return len(nums)+1
}
func abs(i int)int{
    if(i>0){
        return i
    }
    return -i
}