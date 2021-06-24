.PHONY: build clen gen

build:
	go build -o add_person add_person.go
	go build -o list_people list_people.go

clean:
	rm -f add_person list_people

gen:
	protoc -I=$$GOPATH/src -I=. --go_out=. addressbook.proto
