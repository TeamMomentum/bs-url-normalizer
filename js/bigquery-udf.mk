FILENAME=bq-url-normalizer.js
BUCKET_PATH=
DATASET=
TEST_BQ_IN=http://optimized-by.rubiconproject.com/a/xxxxx/xxxxxx/xxxxxxxx.html?&tk_st=1&rf=https%3A//www.m0mentum.co.jp/hello/world%3Fuid%3Dxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx%26nickname%3Dm0mentum&rp_s=c&p_pos=atf&p_screen_res=1366x768&ad_slot=xxxxxxxx
TEST_N1URL_OUT=http://www.m0mentum.co.jp/hello/world/?nickname=m0mentum&uid=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TEST_N2URL_OUT=http://www.m0mentum.co.jp
BQ_QUERY=bq --project_id=$(PROJECT) query --nouse_legacy_sql

define FirstNormalizedURL
CREATE OR REPLACE FUNCTION
  $(DATASET).N1URL(a STRING)
  RETURNS STRING
  LANGUAGE js AS """ return urlnorm.FirstNormalizedURL(a); """
OPTIONS
  ( library = "$(BUCKET_PATH)/$(FILENAME)" ) ;
endef
export FirstNormalizedURL

define SecondNormalizedURL
CREATE OR REPLACE FUNCTION
  $(DATASET).N2URL(a STRING)
  RETURNS STRING
  LANGUAGE js AS """ return urlnorm.SecondNormalizedURL(a); """
OPTIONS
  ( library = "$(BUCKET_PATH)/$(FILENAME)" ) ;
endef
export SecondNormalizedURL

all: deploy bq_test

deploy: build/$(FILENAME)
	gsutil cp build/$(FILENAME) $(BUCKET_PATH)/
	$(BQ_QUERY) "$$(echo $$FirstNormalizedURL)"
	$(BQ_QUERY) "$$(echo $$SecondNormalizedURL)"

build/$(FILENAME): *.js test/*.js
	npx mocha -r ./bq.js test
	npx rollup -c --input bq.js --output build/$(FILENAME)

bq_test: bq_test_n1url bq_test_n2url
bq_test_n1url:
	@echo 'Testing N1URL...'
	$(BQ_QUERY) --format=json "SELECT $(DATASET).N1URL('$(TEST_BQ_IN)')" | sed -e '1d' > bq_test.out
	test $$(cat bq_test.out | jq -r '.[0]["f0_"]') = $(TEST_N1URL_OUT)/
	rm bq_test.out
	@echo 'ok'

bq_test_n2url:
	@echo 'Testing N2URL...'
	$(BQ_QUERY) --format=json "SELECT $(DATASET).N2URL('$(TEST_BQ_IN)')" | sed -e '1d' > bq_test.out
	test $$(cat bq_test.out | jq -r '.[0]["f0_"]') = $(TEST_N2URL_OUT)
	rm bq_test.out
	@echo 'ok'

query:
	echo "$$(echo $$SecondNormalizedURL)"

