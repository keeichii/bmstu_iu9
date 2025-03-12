import java.util.Arrays;
import java.util.Scanner;

public class Test {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("num: ");
        int count = sc.nextInt();

        Rational[] fractions = new Rational[count];

        for (int i = 0; i < count; i++) {
            System.out.print("fraction №" + (i + 1) + "; format: 'n d': ");
            int n = sc.nextInt();
            int d = sc.nextInt();

            fractions[i] = new Rational(n, d);
        }

        Arrays.sort(fractions);
        System.out.println("\nsorted: ");
        for (Rational r : fractions) {
            System.out.println(r);
        }
        sc.close();
    }
}
