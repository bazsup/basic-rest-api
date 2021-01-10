# CRUD with go

## system required
- GO (golang)
- [go migrate](https://github.com/golang-migrate/migrate) : used for migrate database from SQL file


## development guide
1. setup database SQL by yourself

2. run migrate database using
```sh
make migrate
```

3. start application server using
```sh
make run
```

## add migration file
```sh
make make-migration name=<filename>
# e.g.
make make-migration name=create_users_table
```
