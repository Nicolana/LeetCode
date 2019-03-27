class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *hn=NULL,*t,*p,*q;
        int cin=0,sum;

        //只要没有完全遍历完，就继续遍历
        for(p=l1,q=l2;p || q || cin;)
        {

            sum=(p?p->val:0) + (q?q->val:0)+cin;
            cin=sum/10; // 取进位
			
            if(!hn)
            {
                hn=t=(ListNode*)malloc(sizeof(ListNode));
            }else{
				// 相当于 
				// t->next = new node
				// t = new node
                t=t->next=(ListNode*)malloc(sizeof(ListNode));
            }
			
			// t->next 置空
			// 取个位
            t->next=NULL;
            t->val=sum%10;
			
			// 移动指针到下一个数字
            p=p?p->next:p;
            q=q?q->next:q;
        }

        return hn;
    }
};
