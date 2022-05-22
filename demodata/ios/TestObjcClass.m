#import "TestObjcClass.h"

#import "StringProvider.h"

@implementation TestObjcClass

- (void)testFunction:(BOOL)isAdd {
    NSString *value1 = [[StringProvider sharedInstance] getValue:@"key1" defaultValue:@"default value 1"];
    NSString *value2 = [[StringProvider sharedInstance] getValue:@"key2"defaultValue:@"default value 2"];
    NSString *value3 = isAdd ? [[StringProvider sharedInstance] getValue:@"key3" defaultValue:@"default value 3"] : [[StringProvider sharedInstance] getValue:@"key4"defaultValue:@"default value 4"];
    NSString *value5 = [[StringProvider sharedInstance] getValue:@"key5"
                                                    defaultValue:@"default value 5"];
    NSLog(@"The values from objc tester are value1: %@, value2: %@, value3: %@, value5: %@", value1, value2, value3, value5);
}

@end
