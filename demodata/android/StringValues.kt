package android

import kotlin.random.Random

object StringValues {
    fun getValue(key: String, defaultValue: String): String {
        val nextInt = Random.nextInt()
        return if (nextInt % 2 == 0) {
            defaultValue
        } else {
            String.format("value %d", nextInt);
        }
    }

    val key1: String
        get() = getValue("key1", "default value 1")
    val key2: String
        get() = getValue("key2", "default value 2")
    val key3: String
        get() = getValue("key3", "default value 3")
    val key4: String
        get() = getValue("key4", "default value 4")
    val key5: String
        get() = getValue("key5", "default value 5")
}