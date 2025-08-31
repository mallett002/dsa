const { ListNode, printList } = require('./list-node');

/*

My approach:
    if head is the target, just set next to null (unlink it)
    set pointer to next's next (skip target's link in the chain)
    if tail is head set prev's next to null

    rm 2:
    1 -> 2 -> 3 -> 4
    1 -> 3 -> 4
    point curr's next to the next one after next

    rm 1:
    1 -> 2 -> 3 -> 4
    2 -> 3 -> 4
    ?? maybe set first one's next to null

    rm 4:
    1 -> 2 -> 3 -> 4
    1 -> 2 -> 3
    set curr's next to null

*/
    
const head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);


function deleteNodeWithTarget(head, targetVal) {
    let curr = head;

    // if head is the target, just unlink it:
    if (curr.val === targetVal) {
        curr.next = null;
    }

    while (curr?.next) {
        if (curr.next.val === targetVal) {

            // if next is the tail:
            if (!curr.next.next) {
                curr.next = null;
            } else {
                // Else not the tail, move pointers around
                curr.next = curr.next.next;
            }
        }

        curr = curr.next;
    }

    console.log('Finished!');
}

deleteNodeWithTarget(head, 4);

printList(head);
