DB_NAME=dbtest \
DB_HOST=hosttest \
DB_PORT=porttest \
DB_USER=usertest \
DB_PASSWORD=passtest \
DB_SSL_MODE=disable \
APP_PORT=4000 \
go test $(go list ./... | grep -v /ent | grep -v /staticdata | grep -v /extra) -cover
