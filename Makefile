
OC_OUT = 'objective-c.out'

all:
	gcc -framework Foundation ./objective-c/example.m -o $(OC_OUT)

run:
	./$(OC_OUT)

clean:
	rm ./$(OC_OUT)
