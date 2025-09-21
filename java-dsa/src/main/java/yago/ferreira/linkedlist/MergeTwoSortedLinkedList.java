package yago.ferreira.linkedlist;

public class MergeTwoSortedLinkedList {
    public ListNode mergeTwoLists(ListNode listNode1, ListNode listNode2) {
        ListNode listNode = new ListNode(Integer.MIN_VALUE); // NodeVazio = 0x80000000
        ListNode result = listNode; // NodeVazio = 0x80000000

        while (listNode1 != null && listNode2 != null) { // nao itera sozinho!
            if (listNode1.val <= listNode2.val) {
                listNode.next = listNode1; // menor da primeira lista
                listNode1 = listNode1.next; // listNode.next++ | 1 -> 2
            } else {
                listNode.next = listNode2; // menor da segunda lista adiciona Node = 1
                listNode2 = listNode2.next; // apenas lista 2 avanca list2++ = list2.next
            }

            listNode = listNode.next; // sortedList -> ponteiro | 0x80000000 -> 1 -> Node
        }

        if (listNode1 == null) {
            listNode.next = listNode2; // ultimo em diante
        }

        if (listNode2 == null) {
            listNode.next = listNode1; // 4 -> ultimo em diante
        }

        return result.next;
    }
}
