gen-proto:
	buf generate
	cp -r gen/go/lotus/* . && rm -rf gen
