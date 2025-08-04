public class ArrowMatrix {
    private int n;

    private double[][] matrix;


    public ArrowMatrix(int n, double[] firstRow, double[] firstColumn, double[] mainDiagonal) {
        this.n = n;

        this.matrix = new double[n][n];

        for (int i = 0; i < n; i++) {

            matrix[0][i] = firstRow[i];
            matrix[i][0] = firstColumn[i];


            if (i > 0) { matrix[i][i] = mainDiagonal[i]; }


        }

    }

    public double determinant() {
        return computeDeterminant(matrix, n);}

    private double computeDeterminant(double[][] mat, int size) {
        if (size == 1) { return mat[0][0]; }
        if (size == 2) { return mat[0][0] * mat[1][1] - mat[0][1] * mat[1][0]; }


        double det = 0;

        for (int i = 0; i < size; i++) {
            double[][] subMatrix = getSubMatrix(mat, size, i);

            double cofactor = (i % 2 == 0 ? 1 : -1) * mat[0][i] * computeDeterminant(subMatrix, size - 1);
            det += cofactor;


        }
        return det;

    }

    private double[][] getSubMatrix(double[][] mat, int size, int column) {

        double[][] subMatrix = new double[size - 1][size - 1];
        for (int i = 1; i < size; i++) {
            int subCol = 0;

            for (int j = 0; j < size; j++) {
                if (j == column) continue;
                subMatrix[i - 1][subCol] = mat[i][j];
                subCol++;


            }
        }

        return subMatrix;
    }

    public String toString() {

        StringBuilder result = new StringBuilder();

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {

                result.append(String.format("%10.1f ", matrix[i][j]));
            }
            result.append("\n");
        }

        return result.toString();


    }
}


