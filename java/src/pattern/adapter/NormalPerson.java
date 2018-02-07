package pattern.adapter;

public class NormalPerson {
    private String name;

    public NormalPerson(String name) {
        this.name = name;
    }

    public void speak(String words) {
        System.out.println(name + " speak:" + words);
    }
}
