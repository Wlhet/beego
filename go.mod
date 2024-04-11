module github.com/wlhet/beego

require (
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.1.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/pkg/errors v0.9.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)

require (
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	golang.org/x/text v0.3.0 // indirect
)

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85

go 1.19
