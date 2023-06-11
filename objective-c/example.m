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
    // string

    NSString *str = @"Peter";

    // list that allow any type

    NSArray *arr = @[@27, @"hello world", @56.7, @YES];
    NSLog(@"%@", arr);

    // nullish values

    id x = nil;
    NSString *y = NULL;
    NSLog(@"%@", x);
    NSLog(@"%@", y);

    // simple predicate

    NSPredicate *p = [NSPredicate predicateWithFormat: @"name LIKE 'Johnson'"];
    NSLog(@"%@", p);

    // filter by list predicate

    NSArray *arr2 = @[@"March", @"June", @"September", @"December"];
    NSPredicate *p2 = [NSPredicate predicateWithFormat: @"month IN %@", arr2];
    NSLog(@"%@", p2);

    // negate a filter

    NSArray *arr3 = @[@"Spring", @"Winter"];
    // unlike SQL, NOT IN is a compile error
    NSPredicate *p3 = [NSPredicate predicateWithFormat: @"NOT (season IN %@)", arr3];
    NSLog(@"%@", p3);

}
