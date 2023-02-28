# IPY

### A Toy project To Wrap One Liner To Create PostgreSQL Server Container And Connect With PSQL Client Container

<br/>

## Install Dependency

```
curl -OL https://github.com/cbot918/ipy/archive/refs/tags/v0.0.1.tar.gz && tar -xzf v0.0.1.tar.gz && sudo chmod +x ipy-0.0.1/ipy &&sudo mv ipy-0.0.1/ipy /usr/local/bin
```

## Usecase

### create postgres server

```
sudo docker run -dit --name [container_name] -e POSTGRES_PASSWORD=12345 postgres
```

### connect to postgres server

```
sudo docker run -it --rm postgres psql -h `ipy [container_name]` -U postgres -W
```

<br/>

## Ipy Executable Detail

### docker network inspect bridge, search and return target container's ip address for connect purpose

<br/>

## Test and PreRequesties

- linux / docker
