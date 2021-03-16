module github.com/varunamachi/patti/pt

go 1.15

replace github.com/varunamachi/teak => ../../teak

require (
	github.com/lib/pq v1.9.0
	github.com/satori/go.uuid v1.2.0
	github.com/varunamachi/teak v0.0.0-20210311165402-00f2066c65b0
	gopkg.in/urfave/cli.v1 v1.20.0
)
