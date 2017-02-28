all:
	go-bindata -pkg binhtml -o binhtml/assert.go template/
	go build -o airdisk main.go
