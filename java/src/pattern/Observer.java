package pattern;

import pattern.observer.*;

public class Observer {
    public static void main(String[] args) {
        TvStation station = new TvStation();

        Reporter zhangsan = new Reporter("zhangsan");
        station.register(zhangsan);
        Reporter lisi = new Reporter("lisi");
        station.register(lisi);
        Reporter wangwu = new Reporter("wangwu");
        station.register(wangwu);

        station.eventHappen("a man murdered!");

        station.cancel(lisi);

        station.eventHappen("a car accident...");

    }
}
