package pattern;

import pattern.decorator.*;

public class Decorator {
    public static void main(String[] args){
        Toy toyCar = new Toy("toy car");
        System.out.println(toyCar.sold());

        Pack box = new Pack("box");
        box.packet(toyCar);
        System.out.println(box.sold());

        Pack redRibbon = new Pack("red ribbon");
        redRibbon.packet(box);
        System.out.println(redRibbon.sold());
    }
}
