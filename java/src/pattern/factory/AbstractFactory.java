package pattern.factory;

public class AbstractFactory {
    private String brand;

    public AbstractFactory(String brand) {
        this.brand = brand;
    }

    public Shoe product(String type) {
        return new Shoe(type, this.brand);
    }
}
