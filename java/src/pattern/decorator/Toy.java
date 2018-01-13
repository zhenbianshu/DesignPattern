package pattern.decorator;

public class Toy extends Goods {
    private String name;

    public Toy(String name) {
        this.name = name;
    }

    @Override
    public String sold() {
        return this.name;
    }
}
