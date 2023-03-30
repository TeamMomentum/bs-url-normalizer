all: lint test

.PHONY: assets
assets: assets/n2url_path_depth.js assets/sp_host.js

assets/n2url_path_depth.js: ../resources/norm_host_path.csv
	node tools/create_data_path_depth.js > assets/n2url_path_depth.js

assets/sp_host.js: ../resources/norm_host_sp.csv
	node tools/create_data_sp_host.js > assets/sp_host.js

node_modules: package.json
	npm install

.PHONY: test test_node test_bq
test: assets test_node test_bq

test_node: node_modules 
	npm test

test_bq: node_modules
	npm run test:bq

.PHONY: lint 
lint: node_modules assets
	npm run lint

