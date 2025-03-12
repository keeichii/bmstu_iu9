public class Calculator {
    public static double getDistance(Universe u1, Universe u2) {
        double[] c1 = u1.getCenter();
        double[] c2 = u2.getCenter();
        return Math.sqrt(Math.pow(c1[0] - c2[0], 2) + Math.pow(c1[1] - c2[1], 2) + Math.pow(c1[2] - c2[2], 2));
    }
}
