# ipy

### One Liner Create PostgreSQL Server And Connect With Client

<br/>

## Install Dependency

```
curl -OL https://github.com/cbot918/ipy/archive/refs/tags/v0.0.1.tar.gz && tar -xzf v0.0.1.tar.gz && sudo mv ipy-0.0.1/ipy /usr/local/bin
```

## Usecase

### create postgres server

```
docker run -it --name [container_name] -e POSTGRES_PASSWORD=12345 postgres
```

### connect to postgres server

```
docker run -it --rm postgres psql -h `ipy [container_name]` -U postgres -W
```

<br/>

## ipy detail

### docker network inspect bridge, search and return target container's ip address for connect purpose

<br/>

## Test and PreRequesties

- linux / docker
