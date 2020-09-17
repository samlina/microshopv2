### Docker-compose用法

```go
//构建所有服务容器
docker-compose build

//查看构建的所有镜像
docker image ls

//运行刚才构建的容器
docker-compose up -d   //-d代表后台模式运行

//指定容器运行。
docker-compose up -d gxc-user-service

//一次性调用某个容器的服务
docker-compose run gxc-user-service

//查看所有已经启动容器
docker container ls
```



Git 常用方法

```
git init   // 初始化版本库

git add .   // 添加文件到版本库（只是添加到缓存区），.代表添加文件夹下所有文件 

git commit -m "first commit" // 把添加的文件提交到版本库，并填写提交备注

git remote remove origin  //移除已经关联的远程库

git remote add origin 你的远程库地址  // 把本地库与远程库关联

git push -u origin master    // 第一次推送时

git push origin master  // 第一次推送后，直接使用该命令即可推送修改
```

