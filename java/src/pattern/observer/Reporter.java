package pattern.observer;

public class Reporter {
    private String name;

    public Reporter(String name) {
        this.name = name;
    }

    public void notified(String event) {
        System.out.println(this.name + " report a news: " + event);
    }
}
