## PUBLISH

一个一键将代码打包部署到服务器上的命令行工具

### 使用步骤
1. 将publish命令行工具添加至环境变量下
2. 在.gitignore中添加一条publish.yaml,并使用git先进行提交
3. 在当前项目目录下新建文件publish.yaml，内容如下
   ```yaml
    server:
        host: 192.168.36.232       #远程服务器ip
        port: 22                   #服务器端口
        user: test                 #链接账户
        password: 123456           #链接密码

    file:
        local: example        #编译后生成可执行文件的路径    
        remote: /home/test    #将可执行文件上传至远程服务器的该目录下

    localCmd:                         #上传远程之前在本地执行的命令，一般为编译和构建
        - set GOOS=linux&&go build .  #go语言举例：该命令为将当前项目编译为Linux下的可执行文件

    remoteBeforeCmd:            #将文件上传至远程服务器之前在服务器上执行的命令
        - rm /home/test/example #判断是否由/home/ubuntu/example文件，若有则删除
        - port-kill 8080        #若8080端口被占用，则结束占用端口的进程（一般为服务启动端口）

    remoteAfterCmd:                         #将文件上传至远程服务器之后在服务器上执行的命令
      - chmod 777 example                   #将文件在服务器上赋予可执行权限
      - nohup ./example > nohup.out 2>&1 &  #在后台启动服务，标准日志和错误日志均打印在nohup.out文件下
4. 将上述内容改为自己项目的内容后，在该文件的同级目录下执行命令publish，即可将项目发布至测试环境