# jingxiu web服务开发脚手架
version: v0.0.1  

1. 开始一个项目
   ```shell
    jingxiu start ${project-name}
    ```
2. 创建一个应用
    ```shell
    jingxiu create ${controller-name} ${interface-name}
   ```
   目前支持的API接口
   - APIHandler
   - ...待开发
3. 生成数据库模型
    ```shell
    jingxiu model --src ${config-file-path}
    ```
4. 路由导出
    ```shell
    jingxiu route [append ${controller-name}]
    ```