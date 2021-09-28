# jwt_refreshToken

#### Install redis
```
 Using redis on docker
 Setup redis :
   - docker pull redis
   - docker run --name my-redis-container -d redis
 Connect to redis cli(if you want to chack the keys):
   - docker exec -it your_container_id sh ( "docker ps" for check container id )
   - redis-cli
```

#### Add .env file 
```
 ACCESS_SECRET = YOUR_ACCESS_SECRET_KEY
 REFRESH_SECRET = YOUR_REFRESH_SECRET_KEY
```
#### URL Paths

```
 /api/v1/login (login)
 /api/v1/logout (logout)
 /api/v1/refresh (get refresh access token)
 /api/v1/todo (test bearer token)
```

#### Start 

```
go run main.go
```
