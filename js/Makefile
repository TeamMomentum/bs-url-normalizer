all: lint test

.PHONY: assets
assets: assets/path_depth.js assets/sp_host.js

assets/path_depth.js: ../resources/norm_host_path.csv
	node tools/create_data_path_depth.js > assets/path_depth.js

assets/sp_host.js: ../resources/norm_host_sp.csv
	node tools/create_data_sp_host.js > assets/sp_host.js

.PHONY: test
test: assets test_node test_bq

test_node: node_modules 
	npx mocha test --ignore=norm_ipv6.json

test_bq: node_modules
	npx mocha -r ./bq.js test

.PHONY: lint 
lint: node_modules
	npx prettier --check '**/*'
	npx eslint --fix *.js adframe assets test tools

node_modules: package.json
	npm install
