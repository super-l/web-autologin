## 基于GO语言与chromedp库实现三维家(网站)自动登录


基于GO语言与chromedp库实现网站自动登录，以长期定时任务，实现三维家网站自动登录获取cookie,并更新到mysql数据库为例。

### 配置文件说明

```
<?xml version="1.0" encoding="UTF-8" ?>
<Config>
    <task_hour>1</task_hour>
    <show_window>1</show_window>
    <username>test</username>
    <password>123456</password>

    <!-- MYSQL信息 -->
    <mysql_host>192.168.1.100</mysql_host>
    <mysql_port>3306</mysql_port>
    <mysql_user>mysql_user</mysql_user>
    <mysql_password>dbpassword</mysql_password>
    <mysql_database>dbname</mysql_database>
</Config>
```

```
task_hour: 任务每x小时执行一次
show_window： 0表示不显示浏览器窗口，1表示显示(方便测试)
username: 三维家账号
password： 三维家密码

mysql_host： mysql主机链接地址
mysql_port： mysql端口号
mysql_user: mysql数据库的用户名
mysql_password：mysql数据库的密码
mysql_database： mysql数据库名称

```

### 使用方法

```
1. 修改config.xml文件，更改mysql信息，以及三维家登录账号信息；
2. 安装chrome浏览器到系统。同时支持windows,mac,linux等操作系统(linux可用无头chrome浏览器)。
```

### 测试方法

运行编译后的可执行文件，运行成功会有如下信息：

```
|----------------------------------------|
|         三维家cookie自动获取             |
|----------------------------------------|
|  自动模拟登录，获取cookie信息             |
|  Author: superl www.xiao6.net (小六博客)|
|----------------------------------------|

08/10 16:15:44 [INFO] 标题:登录-三维家3D云设计
08/10 16:15:44 [INFO] 输入账号完成!
08/10 16:15:44 [INFO] 输入密码完成!
08/10 16:15:45 [INFO] Cookie:JSESSIONID=XXXXXXXXXXXXXXXXXXXXXXXXX;is_authed=1;

```

### 跨平台编译方案



