package pattern.command;

public class Order {
    private String action;
    private String source;
    private String dest;

    public Order(String action, String source, String dest) {
        this.action = action;
        this.source = source;
        this.dest = dest;
    }

    public String execute() {
        return action + " from " + source + " to " + dest;
    }

    public String undo() {
        return action + " from " + dest + " to " + source;
    }
}
