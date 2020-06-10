integration_examples/basic : 
	mkdir -p ./integration_examples/integration/basic
	cd ./integration_examples/integration/basic;\
	git init;\
	echo "test" > index.txt;\
	git add index.txt && git commit -am "test commit";\
	git tag v0.0.1;\
	git tag v0.0.2;

untest : 
	rm ./integration_examples/integration/basic -rf

.PHONY: untest
