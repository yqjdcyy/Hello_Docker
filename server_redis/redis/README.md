# Install
## Search
- `docker search redis`
    ```
    NAME                                 DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
    redis                                Redis is an open source key-value store that…   5306                [OK]                
    bitnami/redis                        Bitnami Redis Docker Image                      78                                      [OK]
    sameersbn/redis                                                                      67                                      [OK]
    tenstartups/redis-commander                                                          32                                      [OK]
    hypriot/rpi-redis                    Raspberry Pi compatible redis image             31                                      
    kubeguide/redis-master               redis-master with "Hello World!"                23                                      
    kubeguide/guestbook-redis-slave      Guestbook redis slave                           18                                      
    redislabs/redis                      Clustered in-memory database engine compatib…   15                                      
    webhippie/redis                      Docker images for redis                         8                                       [OK]
    arm32v7/redis                        Redis is an open source key-value store that…   7                                       
    rediscommander/redis-commander       Alpine image for redis-commander - Redis man…   6                                       [OK]
    rtoma/logspout-redis-logstash        Logspout including Redis adapter for sending…   5                                       
    oliver006/redis_exporter              Prometheus Exporter for Redis Metrics. Supp…   4                                       
    centos/redis-32-centos7              Redis in-memory data structure store, used a…   3                                       
    dynomitedb/redis                     Redis backend for DynomiteDB.                   2                                       [OK]
    frodenas/redis                       A Docker Image for Redis                        2                                       [OK]
    tomesar/redis-arm                    Redis for ARM!                                  2                                       [OK]
    kilsoo75/redis-master                This image is for the redis master of SK Clo…   1                                       
    google/guestbook-python-redis        A simple guestbook example written in Python…   1                                       
    tiredofit/redis                      Redis Server w/ Zabbix monitoring and S6 Ove…   1                                       [OK]
    anchorfree/redis                     redis cache server for logging                  0                                       
    brendangibat/docker-logspout-redis   Docker Logspout container with Logspout-Redi…   0                                       [OK]
    iadvize/redis                                                                        0                                       
    sstarcher/fluent-redis-aws           https://github.com/sstarcher/docker-fluent-r…   0                                       [OK]
    ajmath/fluentd-redis                 Use fluentd logs to send docker logs to redi…   0                                       [OK]
    ```

## Image.pull
- `docker pull redis`
    ```
    Using default tag: latest
    latest: Pulling from library/redis
    4d0d76e05f3c: Pull complete 
    aa39ed64d5b2: Pull complete 
    1593212ab2f2: Pull complete 
    2fc592fc15dc: Pull complete 
    bf7bb0819354: Pull complete 
    6335ae550d25: Pull complete 
    Digest: sha256:0b4b7bd6ab69ed4642102232bc402bea9ef891ba43e0bb14a1f2d14c7ff1bcc9
    Status: Downloaded newer image for redis:latest
    ```


## Redis.Server.run
- `docker run -p 8030:6379 -v /data/redis/data:/data -d redis redis-server --appendonly yes`
    ```
    8be998175aebd506538633c975304435631ee85a5d4527f7690f2282fe1095a2
    ```

- `docker container ls --all`
    ```
    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
    8be998175aeb        redis               "docker-entrypoint.s…"   10 seconds ago      Up 9 seconds        0.0.0.0:8030->6379/tcp   vigorous_torvalds
    ```

## Redis.Client.run
- `docker exec -it 8be998175aeb redis-cli`


# Reference
- [Docker 安装 Redis](http://www.runoob.com/docker/docker-install-redis.html)
- [redis](https://docs.docker.com/samples/library/redis/)
