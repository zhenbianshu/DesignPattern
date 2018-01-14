package pattern.singleton;

import redis.clients.jedis.Jedis;
import java.util.HashMap;

public class Redis {
    private static volatile HashMap<String, Jedis> INSTANCES = new HashMap<>();

    private Redis() {
    }

    public static Jedis getInstance(String host, int port) {
        String hashKey = host + "_" + port;

        if (Redis.INSTANCES.get(hashKey) == null) {
            synchronized (Redis.class) {
                if (Redis.INSTANCES.get(hashKey) == null) {
                    Jedis instance = new Jedis(host, port);
                    Redis.INSTANCES.put(hashKey, instance);
                }
            }
        }

        return Redis.INSTANCES.get(hashKey);
    }
}
