# How to Use

```
$ docker-compose up
```

# Migrate
DBのマイグレーションには `migrate` を使用する。  

[golang-migrate/migrate: Database migrations. CLI and Golang library.](https://github.com/golang-migrate/migrate)

## up
```
$ docker run \
  --network=echo_default \
  -v `pwd`/db/migrations:/migrations \
  migrate/migrate \
  -path=/migrations \
  -database mysql://root:@tcp(db:3306)/echodb \
  up
```

## create
```
$ docker run \
  -v `pwd`/db/migrations:/migrations \
  migrate/migrate \
  create \
  -ext sql \
  -dir /migrations \
  -seq create_user_table
```






