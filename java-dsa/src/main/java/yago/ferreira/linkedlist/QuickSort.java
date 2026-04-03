package yago.ferreira.linkedlist;

public class QuickSort {
    public void quickSortArray(int[] arr, int left, int right) {
        if (left < right) {
            partition(arr, left, right);
        }
    }

    private void partition(int[] arr, int left, int right) {
        int pivot = arr[right];
    }
}
