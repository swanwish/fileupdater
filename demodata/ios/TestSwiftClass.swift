import Foundation

class TestSwiftClass: NSObject {
    @objc func testFunction(isAdd: Bool) {
        let value1 = StringProvider.sharedInstance().getValue("key1", defaultValue: "default value 1")
        let value2 = StringProvider.sharedInstance().getValue("key2",defaultValue: "default value 2")
        let value3 = isAdd ? StringProvider.sharedInstance().getValue("key3", defaultValue: "default value 3") : StringProvider.sharedInstance().getValue("key4",defaultValue: "default value 4")
        let value5 = StringProvider.sharedInstance().getValue("key5",
                                                              defaultValue: "default value 5")
        print("The values from swift tester are value1: \(value1), value2: \(value2), value3: \(value3), value5: \(value5)")
    }
}
