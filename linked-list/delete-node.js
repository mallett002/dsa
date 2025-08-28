const { ListNode } = require('./list-node');

    
const head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);

function deleteNodeWithTarget(head, targetVal) {
    let curr = head;
    let nth = 1;

    while (curr?.next && curr.val !== targetVal) {
        curr = curr.next;
        nth++;
    }

    if (curr.val === targetVal) {
        console.log('found it!')
    }

    return nth;
}

console.log(deleteNodeWithTarget(head, 2));
