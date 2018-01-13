package pattern.observer;

import java.util.ArrayList;
import java.util.Iterator;

public class TvStation {
    private ArrayList reporters;

    public TvStation() {
        this.reporters = new ArrayList();
    }

    public void register(Reporter reporter) {
        this.reporters.add(reporter);
    }

    public void cancel(Reporter reporter) {
        Iterator iterator = this.reporters.iterator();

        while (iterator.hasNext()) {
            Reporter item = (Reporter) iterator.next();
            if (item.equals(reporter)) {
                iterator.remove();
            }
        }
    }

    public void eventHappen(String event) {
        Iterator iterator = this.reporters.iterator();

        while (iterator.hasNext()) {
            Reporter reporter = (Reporter) iterator.next();
            reporter.notified(event);
        }
    }
}
