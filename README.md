# sertif-validator
sertif-validator

### Migrate DB Up

```
migrate -database "mysql://root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/sv" -path migration up
```

### Migrate DB Down

```
migrate -database "mysql://root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/sv" -path migration down
```

### Generate migration file command

```
migrate create -ext sql -dir {directory_path} -seq {migration_name}