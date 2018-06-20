[TOC]

## 教程
- 参见 [Hello.Docker](http://domain.yqjdcyy.com/post/Hello.Docker/)


## 运行
### Hello World
- `docker container run hello-world`
    ```
    Hello from Docker!
    This message shows that your installation appears to be working correctly.

    To generate this message, Docker took the following steps:
    1. The Docker client contacted the Docker daemon.
    2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
        (amd64)
    3. The Docker daemon created a new container from that image which runs the
        executable that produces the output you are currently reading.
    4. The Docker daemon streamed that output to the Docker client, which sent it
        to your terminal.

    To try something more ambitious, you can run an Ubuntu container with:
    $ docker run -it ubuntu bash

    Share images, automate workflows, and more with a free Docker ID:
    https://hub.docker.com/

    For more examples and ideas, visit:
    https://docs.docker.com/engine/userguide/
    ```

### Empty
- 描述
    - 空指令
- 操作
    - `docker image build -t yqjdcyy/hello-world-empty .`
        ```
        Sending build context to Docker daemon  2.048kB
        Step 1/2 : FROM centos
        ---> 49f7960eb7e4
        Step 2/2 : RUN /bin/sh -c echo "hello world"
        ---> Running in e96051ef783e

        Removing intermediate container e96051ef783e
        ---> 61217d10b966
        Successfully built 61217d10b966
        Successfully tagged yqjdcyy/hello-world-empty:latest        
        ```
    - `docker run --rm -it yqjdcyy/hello-world-empty /bin/sh`
        ```
        sh-4.2#
        ```

### Command
- 描述
    - 打印请求服务的参数信息
- 操作
    - `docker image build -t yqjdcyy/hello-world-resp .`
        ```
        Sending build context to Docker daemon  1.569MB
        Step 1/6 : FROM centos
        ---> 49f7960eb7e4
        Step 2/6 : MAINTAINER yqjdcyy <yqjdcyy@gmail.com>
        ---> Running in cb973e40f304
        Removing intermediate container cb973e40f304
        ---> e43d8d71bc4b
        Step 3/6 : COPY . /
        ---> 6b5c3b186504
        Step 4/6 : RUN echo "hello"
        ---> Running in af8f43cc164c
        hello
        Removing intermediate container af8f43cc164c
        ---> f3fb78ded699
        Step 5/6 : RUN chmod +x resp
        ---> Running in d6e11e7d4310
        Removing intermediate container d6e11e7d4310
        ---> 97cf0b42cb8f
        Step 6/6 : CMD ./resp
        ---> Running in 2f2f1e4f5619
        Removing intermediate container 2f2f1e4f5619
        ---> f1ebcf08f7b9
        Successfully built f1ebcf08f7b9
        Successfully tagged yqjdcyy/hello-world-resp:latest
        ```
    - `docker container run yqjdcyy/hello-world-resp`
        ```
        args= [./resp]
        ```

### Server
- 描述
    - 指定端口（默认8888）开启服务，提供 `/call` 固定返回 `CALL.BACK`
- 操作
    - `docker imagee build -t yqjdcyy/hello-world-server .`
        ```
        Sending build context to Docker daemon  6.224MB
        Step 1/4 : FROM centos
        ---> 49f7960eb7e4
        Step 2/4 : COPY . /
        ---> 6524ff9b562c
        Step 3/4 : RUN chmod +x server
        ---> Running in 647c3414ef68
        Removing intermediate container 647c3414ef68
        ---> bedcdeb9f6be
        Step 4/4 : CMD ./server -port=3000
        ---> Running in e86274084244
        Removing intermediate container e86274084244
        ---> 97ab267a7baf
        Successfully built 97ab267a7baf
        Successfully tagged yqjdcyy/hello-world-server:latest
        ```
    - `docker run -p 8030:3000 yqjdcyy/hello-world-server`
        ```
        Server start at :3000
        ```
    - `http://<ip>:8030/call`
        ```
        CALL.BACK
        ```


### Server-Redis
- 描述
    - 支持 `Redis` 的服务
- 操作
    - 请先按照 `../redis/README.md` 启动 `Reids`
    - `docker [image] build -t yqjdcyy/hello-docker-server-redis .`
        ```
        Sending build context to Docker daemon  7.194MB
        Step 1/5 : FROM centos
        ---> 49f7960eb7e4
        Step 2/5 : MAINTAINER yqjdcyy <yqjdcyy@gmail.com>
        ---> Running in 4b932b06bf20
        Removing intermediate container 4b932b06bf20
        ---> 97bff887c61a
        Step 3/5 : COPY . /
        ---> 033030de61fd
        Step 4/5 : RUN chmod +x server
        ---> Running in 398732e1bf86
        Removing intermediate container 398732e1bf86
        ---> c2c4c5007f20
        Step 5/5 : CMD ./server -port=3000 -addr=localhost:6379
        ---> Running in 39f6279ebf76
        Removing intermediate container 39f6279ebf76
        ---> 719139c2b17d
        Successfully built 719139c2b17d
        Successfully tagged yqjdcyy/hello-docker-server-redis:latest
        ```

    - `docker run -d -p:8031:3000 yqjdcyy/hello-docker-server-redis`
        ```
        c117703a542519dbb9c7b51761883c96ebdbe89d5d7facf7839d0b2cac255e1e
        ```