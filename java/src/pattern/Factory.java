package pattern;

import pattern.factory.*;

public class Factory {
    public static void main(String[] args) {
        SimpleFactory simpleFactory = new SimpleFactory();
        Shoe shoeA = simpleFactory.product("sports", "adidas");
        System.out.println(shoeA.getDetail());

        AbstractFactory liningShoeFactory = new AbstractFactory("lining");
        Shoe shoeB = liningShoeFactory.product("cotton");
        System.out.println(shoeB.getDetail());
    }
}
