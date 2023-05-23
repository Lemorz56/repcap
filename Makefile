.RECIPEPREFIX = >

test:
> @go test -race ./...

guirun:
> go run . -i "Ethernet" -r --pcap C:\code\python\ReplayCapturedData\test.pcap --gui

bench:
> @go test -bench=. -benchmem ./...