package android;

class TestKotlinClass {
    fun testFunction(isAdd: Boolean) {
        val value1 = StringProvider.getValue("key1", "default value 1")
        val value2 = StringProvider.getValue("key2","default value 2")
        val value3 = if (isAdd) { StringProvider.getValue("key3", "default value 3") } else { StringProvider.getValue("key4","default value 4") }
        val value5 = StringProvider.getValue("key5",
        "default value 5")
    }
}