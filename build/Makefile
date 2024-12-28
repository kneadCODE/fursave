golib-codacy-coverage: SERVICE=golib
golib-codacy-coverage: codacy-coverage-go

ledgersvc-codacy-coverage: SERVICE=ledgersvc
ledgersvc-codacy-coverage: codacy-coverage-go

codacy-coverage-go:
	curl -Ls -o codacy-cov.sh https://coverage.codacy.com/get.sh
	chmod +x codacy-cov.sh
	./codacy-cov.sh report --force-coverage-parser go --partial -r src/${SERVICE}/coverage.out
