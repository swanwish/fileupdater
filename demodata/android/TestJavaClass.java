package android;

//  This is just a test file demonstrate the java codes

public class TestJavaClass {
    private void testFunction(boolean isAdd) {
        String value1 = StringProvider.getValue("key1", "default value 1");
        String value2 = StringProvider.getValue("key2", "default value 2");
        String value3 = isAdd ? StringProvider.getValue("key3", "default value 3") : StringProvider.getValue("key4","default value 4");
        String value4 = StringProvider.getValue("key5",
                "default value 5");
    }
}