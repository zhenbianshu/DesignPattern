package pattern.factory;

public class SimpleFactory {
    public Shoe product(String type, String brand) {
        return new Shoe(type, brand);
    }
}
