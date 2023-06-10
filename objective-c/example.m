#import <Foundation/Foundation.h>


@interface Person: NSObject {
    NSString *name;
    int age;
    float height;
    BOOL vegetarian;
    NSDate *dateOfBirth;
}

@end


int main() {
    NSString *str = @"Peter";

    NSArray *arr = @[@27, @"hello world", @56.7, @YES];
    NSLog(@"%@", arr);

    id x = nil;
    NSString *y = NULL;
    NSLog(@"%@", x);
    NSLog(@"%@", y);

    NSPredicate *p = [NSPredicate predicateWithFormat: @"name LIKE 'Johnson'"];
    NSLog(@"%@", p);

}
