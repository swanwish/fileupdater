#import "StringProvider.h"

@implementation StringProvider

+ (StringProvider *)sharedInstance {
    static StringProvider *sharedInstance;
    @synchronized (self) {
        if (!sharedInstance) {
            StringProvider *instance =[[StringProvider alloc] init];
            sharedInstance = instance;
        }
        return sharedInstance;
    }
}

- (NSString *)getValue:(NSString *)strKey defaultValue:(NSString *)defaultValue {
    int nextInt = arc4random_uniform(10000);
    if (nextInt % 2 == 0) {
        return defaultValue;
    } else {
        return [NSString stringWithFormat:@"value %d", nextInt];
    }
}

@end
