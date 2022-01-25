wrk-ping:
	wrk -t 4 -c 100 http://0.0.0.0:8000/ping

wrk-ping-config:
	wrk -t 4 -c 100 http://0.0.0.0:8000/ping/config

wrk-tickets-list:
	wrk -t 4 -c 100 http://0.0.0.0:8000/ticket/list

wrk-sparseFields-test:
	wrk -t 4 -c 100 'http://0.0.0.0:8000/ticket/list?id=632f3dd1-28a1-4a5d-86cc-1b09ebc0ed9e&paging=false&page_size=855&page_token=faa9c6ed-a538-456e-88a8-70d850750a93&fields[ticket]=some_field_to_be_returned%2Cticket_photos%2Cticket_description'

go-list:
	go list -m -u all

go-tickets-run-local:
	go run cmd/api.go

go-mod:
	go mod tidy -compat=1.17
