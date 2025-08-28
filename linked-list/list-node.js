class ListNode {
    constructor(val, next = null) {
        this.val = val;
        this.next = next;
    }

    print() {
        console.log(`Curr: ${this.val}; Next: ${this.next?.val ?? null}`);
    }
}

// 1 -> 2 -> 3 -> 4
// 1    2    3    4
function findLength(head) {
    let count = 0;
    let curr = head;

    while (curr) {
        count++;
        curr = curr.next;
    }

    return count;
}



const head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);
//head.next.next?.print();


console.log(findLength(head));

