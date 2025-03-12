public class Universe {
    private Particle[] particles;

    public Universe(int size) {
        particles = new Particle[size];
    }

    public void addParticle(int index, double x, double y, double z) {
        particles[index] = new Particle(x, y, z);
    }

    public double[] getCenter() {
        double sumX = 0, sumY = 0, sumZ = 0;
        for (Particle p : particles) {
            double[] pos = p.getPosition();
            sumX += pos[0];
            sumY += pos[1];
            sumZ += pos[2];
        }
        int n = particles.length;
        return new double[]{sumX / n, sumY / n, sumZ / n};
    }
}
