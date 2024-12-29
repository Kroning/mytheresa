set -e -u

go get -v ./...

export MODE=TESTS

go test  $(go list ./... | grep -v "cmd\|mock\|/app\|/database\|config\|/logger\|/repository$\|/service\|/transport/http$" ) \
   -race -coverprofile cover.out -covermode atomic

go tool cover -html=cover.out -o cover.html

perc=`go tool cover -func=cover.out | tail -n 1 | sed -Ee 's!^[^[:digit:]]+([[:digit:]]+(\.[[:digit:]]+)?)%$!\1!'`
echo "Total coverage: $perc %"
res=`echo "$perc >= 22.0" | bc`
test "$res" -eq 1 && exit 0
echo "Insufficient coverage: $perc" >&2
exit 1


