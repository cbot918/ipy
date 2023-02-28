psqlc:
	docker run -it --rm postgres psql -h `ipy lbin` -U postgres -W

psqls:
	docker run -it --name testpsql -e POSTGRES_PASSWORD=12345 postgres

test:
	echo $(ipy testpsql) | echo $(</dev/stdin)

.PHONY: psqlc test psqls
.SILENT: psqlc test psqls




curl -OL https://github.com/cbot918/ipy/archive/refs/tags/v0.0.1.tar.gz && tar -xzf v0.0.1.tar.gz && sudo mv ipy-0.0.1/ipy /usr/local/bin