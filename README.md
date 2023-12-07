# NanoService aka MicroService

* Loosely coupled

* Can be owned by one person or small team

* Other layers (like "transport" ) can be added later 


### TODO

* remove dependency from gorm


## Usage

### If cmd/server/main.go

put `test.txt` contents to database

 `curl -s -X POST --data @test.txt "http://localhost:9999/api/v1/records?category=test&domain=root.toor"`

 list all records

 `curl "http://localhost:9999/api/v1/records?category=abcde&domain=root.toor"`

 read one record

 `curl "http://localhost:9999/api/v1/records/{uuid}?category=abcde&domain=root.toor"`

 
* replace {uuid} with uuid of record from `list all records`

update record

` curl -s -X PUT --data @x.sh "http://localhost:9999/api/v1/records/{uuid}?category=abcde&domain=root.toor"`

* replace {uuid} with uuid of record from `list all records`

delete record

`curl -s -X DELETE "http://localhost:9999/api/v1/records/{uuid}?category=abcde&domain=root.toor"`

* replace {uuid} with uuid of record from `list all records`


### If cmd/cli/main.go

put `test.txt` contents to database

`cat testas.txt | go run cmd/cli/main.go --action create --category beef -`

list all records in category 

`go run cmd/cli/main.go -action list -category beef`

read one record by uuid

`go run cmd/cli/main.go -action read -category abcde {uuid}`

* replace {uuid} with uuid of record from `list all records`

update {uuid} with contents fo `test.txt`

`cat test.txt | go run cmd/cli/main.go --action update --category beef {uuid} -`

* replace {uuid} with uuid of record from `list all records`

delete {uuid}

`go run cmd/cli/main.go --action delete --category beef {uuid}`


