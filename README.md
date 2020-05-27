# go-sqlboiler

## sqlboiler

Do it once.
```bash
go get -u -t github.com/volatiletech/sqlboiler
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

If you had changes in your database schemas. Refresh the models with it.
```bash
sqlboiler mysql --wipe
```