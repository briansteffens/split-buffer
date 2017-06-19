default: build

build:
	go build main.go gap.go split.go

clean:
	rm main
