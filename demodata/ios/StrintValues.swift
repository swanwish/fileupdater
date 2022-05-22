import Foundation

class StringValues: NSObject {
    @objc static let shared = StringValues()

    private override init() {}

    private func getValue(_ key: String, _ defaultValue: String) -> String {
        let nextInt = Int.random(in: 0...1000)
        if nextInt % 2 == 0 {
            return defaultValue
        } else {
            return "value \(nextInt)"
        }
    }

    @objc var key1: String {
        getValue("key1", "default value 1")
    }
    @objc var key2: String {
        getValue("key2", "default value 2")
    }
    @objc var key3: String {
        getValue("key3", "default value 3")
    }
    @objc var key4: String {
        getValue("key4", "default value 4")
    }
    @objc var key5: String {
        getValue("key5", "default value 5")
    }
}
