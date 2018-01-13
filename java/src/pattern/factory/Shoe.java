package pattern.factory;

public class Shoe {
    private String type;
    private String brand;

    public Shoe(String type, String brand) {
        this.type = type;
        this.brand = brand;
    }

    public String getDetail() {
        return "a " + this.brand + " " + this.type + " shoe.";
    }
}
