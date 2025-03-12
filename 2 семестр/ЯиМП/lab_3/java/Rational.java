public class Rational implements Comparable<Rational> {
    int n;
    int d;

    public Rational(int n, int d) {
        if (d == 0) {
            System.out.println("cant be 0");
            System.exit(0); // остановка программы
        }
        int gcd = findGCD(n, d);
        this.n = n / gcd;
        this.d = d / gcd;
        if (this.d < 0) {
            this.n = -this.n;
            this.d = -this.d;
        }
    }

    private int findGCD(int a, int b) {
        while (b != 0) {
            int tmp = b;
            b = a % b;
            a = tmp;
        }
        return a;
    }

    public String toString() { return n + "/" + d; }

    public int compareTo(Rational other) {
        if (this.n * other.d < other.n * this.d) {
            return -1;
        } else if (this.n * other.d > other.n * this.d) {
            return 1;
        } else {
            return 0;
        }
    }
}
