import Foundation

let name = "Benjamin"
let p = NSPredicate(format: "name == %@", name)
print(p)

// filter by list

let nums = [2, 3, 5, 7]
let p2 = NSPredicate.init(format: "count IN %@", nums)
print(p2)

// special values

let p3 = NSPredicate.init(value: true)
print(p3)

let p4 = NSPredicate.init(format: "SELF != %@", [nil])
print(p4)
