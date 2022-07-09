### 每天固定密码生成器 

* 目前用来修改VPN的密码，自动双端同时修改每天0点进行修改。（会有误差不过启动VPN这个场景在凌晨几秒钟误差可以接受。）




#### 使用 

```bash
./GenPasswordForEveryDay -i 0 # 初始化superKey 里面的计算顺序随便瞎写的


./GenPasswordForEveryDay # 查看对应SuperKey 今天的密码

```