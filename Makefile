clean: 
	rm -rf out/

build: 
	go build -o out/demo-plugin

install:
	cf install-plugin out/demo-plugin -f
