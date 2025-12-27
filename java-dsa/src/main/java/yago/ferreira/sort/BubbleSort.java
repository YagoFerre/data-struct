package yago.ferreira.sort;

import java.util.Arrays;

public class BubbleSort {
    public void bubble(int[] nums) {
        int size = nums.length;
        for (int ignored : nums) { // percorre a List com os numeros
            System.out.println(Arrays.toString(nums));
            for (int i = 0; i < size - 1; i++) {
                if (nums[i] > nums[i + 1]) { // [4 > 3]
                    int temp =  nums[i];

                    nums[i] = nums[i + 1];
                    nums[i + 1] = temp;
                }
            }
        }
    }
}
