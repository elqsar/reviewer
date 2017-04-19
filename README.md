### Simple application for review

* `goose -dir="./migrations" postgres "user=musatov dbname=reviewer sslmode=disable" up`
* `goose -dir="./migrations" postgres "user=musatov dbname=reviewer sslmode=disable" down`
* `export SECRET=<YOUR_SECRET> go run main.go --username=<DB_USER> --database=<DB_NAME>`