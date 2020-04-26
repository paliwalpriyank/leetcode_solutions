/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode first = new ListNode(0);
        ListNode current = first;
        boolean isfirst = true;
        int carry =0;
        while(l1!=null || l2!=null){
            if(!isfirst){
                current.next = new ListNode(0);
                current = current.next;
            }
            isfirst = false;
            int a=0, b =0;
            if(l1!=null){
                a=l1.val;
                l1=l1.next;
            }
            if(l2!=null){
                b=l2.val;
                l2=l2.next;
            }
            int x =a+b+carry;
            carry=0;
            if(x>9){
                current.val = x%10;
                carry=x/10;
            }else{
                current.val=x;
            }
            
        }
        if(carry!=0){
            current.next = new ListNode(carry);
        }
        return first;
    }
}