import pattern.singleton.Redis;
import redis.clients.jedis.Jedis;

public class Singleton {
    public static void main(String[] argv) {
        Jedis client = Redis.getInstance("localhost", 6379);
        client.set("test", "hello");
    }
}
