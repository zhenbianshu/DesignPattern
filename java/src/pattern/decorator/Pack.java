package pattern.decorator;

public class Pack extends Goods {
    private Goods goods;
    private String name;

    public Pack(String name) {
        this.name = name;
    }

    public void packet(Goods goods) {
        this.goods = goods;
    }

    @Override
    public String sold() {
        return this.name + "( " + this.goods.sold() + " )";
    }
}
