
OC_OUT = 'objective-c.out'
SWIFT_FILE = 'swift/example.swift'

all:
	gcc -framework Foundation ./objective-c/example.m -o $(OC_OUT)

run:
	./$(OC_OUT)

run.swift:
	swift $(SWIFT_FILE)

clean:
	rm ./$(OC_OUT)
