package pattern.adapter;

public class SpeakAdapter {
    private DeafMan man;

    public void setMan(DeafMan man) {
        this.man = man;
    }

    public void speak(String words) {
        System.out.println(this.man.write(words));
    }
}
