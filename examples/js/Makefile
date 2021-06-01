all: install run

install: node_modules

node_moduels: package.json
	npm install

run:
	@echo running CommonJS version...
	node commonjs
	@echo running ES module version...
	node -r esm commonjs
