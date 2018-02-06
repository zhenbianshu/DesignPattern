import pattern.command.Order;
import pattern.command.Soldier;

public class Command {
    public static void main(String[] args) {
        Order command = new Order("run", "mountain", "river");
        Soldier lisi = new Soldier("lisi");
        lisi.setCommand(command);

        lisi.execCommand(); // lisi run from mountain to river
        lisi.undoCommand();// lisi run from river to mountain
    }
}
