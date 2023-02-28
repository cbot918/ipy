psqlc:
	docker run -it --rm postgres psql -h `ipy lbin` -U postgres -W

psqls:
	docker run -it --name testpsql -e POSTGRES_PASSWORD=12345 postgres

test:
	echo `ipy lbin` | echo $(</dev/stdin) >> d.txt

.PHONY: psqlc test psqls
.SILENT: psqlc test psqls