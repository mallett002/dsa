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

    // if head is the target, just unlink it (first node)
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

//deleteNodeWithTarget(head, 4);

//printList(head);

// if head is target, just return next's head (essentially skips it)
// while current node has a value
//      if it's the target, point prev's next to curr's next (skips it in the chain)
//      move curr along: prev is now curr and curr is now next
// return head
function deleteNodeAnswer(head, target) {
    if (head.val === target) return head.next;

    let prev = null;
    let curr = head;

    while (curr) {
        if (curr.val === target) {
            prev.next = curr.next; // skip curr in the chain
            // prev -> curr -> next
            // prev -> next
            break;
        }

        // move it along:
        // p -> c -> n
        // 1 -> 2 -> 3
        // 2 -> 3 -> 4?
        prev = curr;
        curr = curr.next;
    }

    return head;
}

deleteNodeAnswer(head, 4);
printList(head);

