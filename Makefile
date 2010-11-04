all: benchmark.8 helloworld.8 
	8l helloworld.8
	
%.8: %.go
	8g $<