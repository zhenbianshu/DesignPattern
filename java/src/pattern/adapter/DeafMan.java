package pattern.adapter;

public class DeafMan {
    private String name;

    public DeafMan(String name) {
        this.name = name;
    }

    public String write(String words) {
        return name + " speak:" + words;
    }
}
