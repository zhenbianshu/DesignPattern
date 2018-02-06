package pattern.command;

public class Soldier {
    private String name;
    private Order command;

    public Soldier(String name) {
        this.name = name;
    }

    public void execCommand() {
        System.out.println(name + " " + command.execute());
    }

    public void undoCommand() {
        System.out.println(name + " " + command.undo());
    }

    public void setCommand(Order command) {
        this.command = command;
    }
}
