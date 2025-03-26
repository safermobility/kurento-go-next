.PHONY: all clean build

all: clean prep build

clean:
	rm -rf core elements
	rm build/kms-core-valid-json/*.json
	rm build/kms-elements-valid-json/*.json

prep:
	mkdir core
	mkdir elements

build:
	cd build/fix-kurento-json && ./gradlew run
	go generate ./...
