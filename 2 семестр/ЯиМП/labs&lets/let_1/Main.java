class Vector3D {
    private double x, y, z;

    public Vector3D() {
        this.x = 0.0;
        this.y = 0.0;
        this.z = 0.0;
    }

    public void setCoordinates(double x, double y, double z) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    public String getCoordinates() {
        return "(" + x + ", " + y + ", " + z + ")";
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println("Яннаев Александр, ИУ9-21Б");

        Vector3D vector = new Vector3D();
        System.out.println("дефолтное значение: " + vector.getCoordinates());

        vector.setCoordinates(3.5, -2.0, 4.7);
        System.out.println("измененное значение: " + vector.getCoordinates());
    }
}
