import java.util.Scanner;

public class Test {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);


        int n;

        while (true) {

            System.out.print("введите размер матрицы (n): ");
            n = scanner.nextInt();

            if (n > 0) break;
            System.out.println("размер должен быть > 0.");

        }

        double[] firstRow = new double[n];

        double[] firstColumn = new double[n];
        double[] mainDiagonal = new double[n];


        while (true) {

            System.out.println("\nвведите первую строку матрицы:");
            for (int i = 0; i < n; i++) {
                System.out.print("firstRow[" + i + "]: ");

                firstRow[i] = scanner.nextDouble();
            }


            System.out.println("\nвведите первый столбец матрицы:");

            for (int i = 0; i < n; i++) {
                System.out.print("firstColumn[" + i + "]: ");

                firstColumn[i] = scanner.nextDouble();

            }

            System.out.println("\nвведите главную диагональ матрицы:");

            for (int i = 1; i < n; i++) {
                System.out.print("mainDiagonal[" + i + "]: ");

                mainDiagonal[i] = scanner.nextDouble();


            }
            ArrowMatrix matrix = new ArrowMatrix(n, firstRow, firstColumn, mainDiagonal);

            System.out.println("\nвведённая матрица:");
            System.out.println(matrix);
            System.out.print("правильно? (да/нет): ");
            String answer = scanner.next();

            if (answer.equals("да")) {
                System.out.println("\nопределитель матрицы: " + matrix.determinant());
                break;

            }
            System.out.println("попробуйте ввести заново.");
        }


        scanner.close();
    }

}
