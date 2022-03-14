10.2
为 HTTPServer 添加 0-2 秒的随机延时；
为 HTTPServer 项目添加延时 Metric；
将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
从 Promethus 界面中查询延时指标数据；
（可选）创建一个 Grafana Dashboard 展现延时分配情况。


2.1编写一个 HTTP 服务器
介绍
1. 接收客户端 request，并将 request 中带的 header 写入 response header 
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header 
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 
4. 当访问 localhost/healthz 时，应返回 200