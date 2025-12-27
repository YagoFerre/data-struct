package yago.ferreira;

import yago.ferreira.linkedlist.MergeTwoSortedLinkedList;
import yago.ferreira.sort.BubbleSort;

import java.util.ArrayList;
import java.util.List;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        MergeTwoSortedLinkedList list = new MergeTwoSortedLinkedList();
        BubbleSort bubbleSort = new BubbleSort();

        bubbleSort.bubble(new int[]{5, 4, 3, 2, 1});
    }
}