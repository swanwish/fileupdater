package android;

import java.util.Random;

public class StringProvider {
    public static String getValue(String key, String defaultValue) {
        Random rand = new Random();
        int nextInt = rand.nextInt();
        if (nextInt % 2 == 0) {
            return defaultValue;
        } else {
            return String.format("value %d", nextInt);
        }
    }
}
