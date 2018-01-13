package pattern;

import pattern.strategy.*;

public class Strategy {
    public static void main(String[] args) {
        Alg bubble = new Bubble();
        Alg quickSort = new QuickSort();

        Strategy.sort(bubble);
        Strategy.sort(quickSort);
    }

    private static void sort(Alg alg) {
        alg.cal();
    }
}
