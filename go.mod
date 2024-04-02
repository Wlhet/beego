module github.com/wlhet/beego

require (
	github.com/Knetic/govaluate v3.0.0+incompatible // indirect
	github.com/bradfitz/gomemcache v0.0.0-20180710155616-bc664df96737
	github.com/casbin/casbin v1.7.0
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.1.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/golang-lru v0.5.4
	github.com/lib/pq v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644
	github.com/ssdb/gossdb v0.0.0-20180723034631-88f6b59b84ec
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
)

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85

go 1.13
