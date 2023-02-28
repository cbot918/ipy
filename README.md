# ipy

## One Liner Create PostgreSQL Server And Connect With Client

<br/>

# Install

```
curl -OL && sudo
```

# Usecase

create postgres server

```
docker run -it --name [container_name] -e POSTGRES_PASSWORD=12345 postgres
```

connect to postgres server

```
docker run -it --rm postgres psql -h `ipy [container_name]` -U postgres -W
```

<br/>

# Test and PreRequesties

## linux / docker
