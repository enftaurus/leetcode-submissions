/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode temp1,temp2,result=null,head=null;
        temp1=l1;
        temp2=l2;
        int a,b,c=0;

        while(temp1!=null && temp2!=null){
            a=temp1.val;
            b=temp2.val;
            ListNode nn=new ListNode();
            nn.val=(a+b+c)%10;
            if(result==null){
                result=nn;
                head=result;
                 
            }
            else{
                result.next=nn;
                result=result.next;
                
            } 
            c=(a+b+c)/10;           
            temp1=temp1.next;
            temp2=temp2.next;
        }
        while(temp1!=null){
            a=temp1.val;
         ListNode nn=new ListNode();
            nn.val=(a+c)%10;
            if(result==null){
                result=nn;
                head=result;
                 
            }
            else{
                result.next=nn;
                result=result.next;
                result.next=null;
                
            } 
            c=(a+c)/10;           
            temp1=temp1.next;
            
        }
        while(temp2!=null){
            b=temp2.val;
         ListNode nn=new ListNode();
            nn.val=(b+c)%10;
            if(result==null){
                result=nn;
                head=result;
                 
            }
            else{
                result.next=nn;
                result=result.next;
                
            } 
            c=(b+c)/10;           
            temp2=temp2.next;
            
        }
        if(c>0){
          ListNode nn=new ListNode();
          nn.val=c;
           if(result==null){
                result=nn;
                head=result;
                 
            }
            else{
                result.next=nn;
                result=result.next;
                
            } 

        }
       
       return head; 
    }
}