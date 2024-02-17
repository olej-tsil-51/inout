TASK = inputf inputs inputs2 inputr 

all: $(TASK)

%: %.go
	go build -o $@ $<

clean:
	rm -f $(TASK)

