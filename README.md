## Creonsoft Backend Developer Test

### Guidelines

Prerequisite:
- Golang 1.18 or higher
- PostgreSQL
- Make (CLI)

Installation
1. Copy file `config.yml.example` to `config.yml` in the directory `/configs` (if not existed)
2. Fill out the DB connection config based on you local Postgres DB
3. At the line 5 in the Makefile, adjust your local DB connection from your `config.yml` file
4. From root directory, then run `make run` or `go run cmd/app/main.go`
5. If you want to run other migrations, you can run `make migrate-up`
6. The service will be available on `localhost:8081` on your local (address & port based in `config.yml` file)


In this source code has refactored and written to adopt Clean Architecture with some modified code and directory styles.
Every file has separated to its functionality. And the flow of this API application started from Router -> Handler -> Service -> Repository (DB).


### API Documentation
https://drive.google.com/file/d/19gRpIkzZoGILx59GR_DGzDnEhsSQMhen/view?usp=sharing

### Main Problem 
1. Daftar Nasabah
``` bash
[POST] http://localhost:8081/daftar
```
``` cURL
curl --location 'localhost:8081/daftar' \
--header 'Content-Type: application/json' \
--data '{
    "nama": "Ryo",
    "nik": "3515182705000019",
    "no_hp": "087754478769"
}'
```
Request Body contoh
```json
{
    "nama": "Ryo",
    "nik": "3515182705000019",
    "no_hp": "087754478769"
}
```

2. Saldo Rekening
``` bash
[POST] http://localhost:8081/saldo/:no_rekening
```
``` cURL
curl --location 'localhost:8081/saldo/:8899084887'
```

3. Tabung
``` bash
[POST] http://localhost:8081/tabung
```
``` cURL
curl --location 'localhost:8081/tabung' \
--header 'Content-Type: application/json' \
--data '{
    "no_rekening": "8899084887",
    "nominal": 2405
}'
```
Request Body contoh
```json
{
  "no_rekening": "8899084887",
  "nominal": 2405
}
```

4. Tarik Saldo
``` bash
[POST] http://localhost:8081/tarik
```
``` cURL
curl --location 'localhost:8081/tarik' \
--header 'Content-Type: application/json' \
--data '{
    "no_rekening": "8899084887",
    "nominal": 790
}'
```
Request Body contoh
```json
{
  "no_rekening": "8899084887",
  "nominal": 790
}
```

### Tantangan
API Mutasi 
``` bash
[GET] http://localhost:8081/mutasi/:no_rekening
```
``` cURL
curl --location 'localhost:8081/mutasi/5118912798'
```

It will return like this
``` json
# Success
{
    "status": "SUCCESS",
    "code": "0000",
    "mutasi": [
        {
            "waktu": "2023-09-29T00:09:50.741348+07:00",
            "no_rekening": "8899084887",
            "kode_transaksi": "C",
            "nominal": "2405"
        },
        {
            "waktu": "2023-09-29T00:10:16.108511+07:00",
            "no_rekening": "8899084887",
            "kode_transaksi": "D",
            "nominal": "2003"
        },
        {
            "waktu": "2023-09-29T00:10:31.926329+07:00",
            "no_rekening": "8899084887",
            "kode_transaksi": "D",
            "nominal": "790"
        }
    ]
}

# Invalid Request no_rekening
{
    "status": "ERROR",
    "code": "0005",
    "data": null,
    "pagination": null,
    "error_message": "User not found",
    "remark": "Nasabah dengan `no_rekening` tersebut tidak dikenali"
}
```