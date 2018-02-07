import pattern.adapter.DeafMan;
import pattern.adapter.NormalPerson;
import pattern.adapter.SpeakAdapter;

public class Adapter {
    public static void main(String[] args) {
        NormalPerson zhangsan = new NormalPerson("zhangsan");
        zhangsan.speak("hi"); // zhangsan speak:hi

        DeafMan lisi = new DeafMan("lisi");
        SpeakAdapter adapter = new SpeakAdapter();
        adapter.setMan(lisi);
        adapter.speak("hello"); // lisi speak:hello
    }
}
