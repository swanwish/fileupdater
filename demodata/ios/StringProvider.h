#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

@interface StringProvider : NSObject

+ (StringProvider *)sharedInstance;

- (NSString *)getValue:(NSString *)strKey defaultValue:(NSString *)strDefaultValue;

@end

NS_ASSUME_NONNULL_END
