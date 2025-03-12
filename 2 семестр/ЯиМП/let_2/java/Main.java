import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.print("введите количество частиц в первой вселенной: ");
        int n1 = sc.nextInt();
        Universe u1 = new Universe(n1);
        for (int i = 0; i < n1; i++) {
            System.out.print("введите координаты x, y, z для частицы " + (i + 1) + ": ");
            double x = sc.nextDouble();
            double y = sc.nextDouble();
            double z = sc.nextDouble();
            u1.addParticle(i, x, y, z);
        }

        System.out.print("введите количество частиц во второй вселенной: ");
        int n2 = sc.nextInt();
        Universe u2 = new Universe(n2);
        for (int i = 0; i < n2; i++) {
            System.out.print("введите координаты x, y, z для частицы " + (i + 1) + ": ");
            double x = sc.nextDouble();
            double y = sc.nextDouble();
            double z = sc.nextDouble();
            u2.addParticle(i, x, y, z);
        }

        double distance = Calculator.getDistance(u1, u2);
        System.out.println("расстояние между центрами вселенных: " + distance);
        System.out.println("общее количество частиц: " + Particle.getCount());

        sc.close();
    }
}
