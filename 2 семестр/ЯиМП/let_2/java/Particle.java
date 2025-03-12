public class Particle {
    private double x, y, z;
    private static int count = 0;

    public Particle(double x, double y, double z) {
        this.x = x;
        this.y = y;
        this.z = z;
        count++;
    }

    public static int getCount() {
        return count;
    }

    public double[] getPosition() {
        return new double[]{x, y, z};
    }
}
