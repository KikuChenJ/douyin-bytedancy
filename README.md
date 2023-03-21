# 抖音项目使用

导入项目后，在/define/define.go中修改相应参数适配本地信息

## ffmpeg安装

用于截取发布视频的封面图

下载地址 https://github.com/BtbN/FFmpeg-Builds/releases

配置教程 https://blog.csdn.net/HYEHYEHYE/article/details/122000352

## redis安装

用于和mysql一起建立二级缓存(未完成)

windows版本下载地址 https://github.com/tporadowski/redis/releases

下载教程 https://blog.csdn.net/qq_26373925/article/details/109269459

无需配置 只需要安装好后打开对应文件夹，开启`redis-server.exe`和`redis-cli.exe`即可使用！

测试的时候需要在`util/redisCache.go`中更改为自己的mysql密码

# 功能说明

抖音接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

app更新动态
https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg


# 待优化

## user模块
个人信息中作品数（目前前端未提供该属性接口）

## video模块
video封面url未获取

前端请求参数中latest_time属性未使用(已解决)

next_time属性未使用(已解决)

## comment模块
没找到怎么删评论(已解决，长按评论，该功能已实现)

-----------------------------------

退出登录之后，那个点赞还显示(红的)这是bug吗？然后我登录新的账号，那个点赞还在（app问题，最新版已解决）

发现视频无法播放，但是有声音。

视频流中每个视频的作者信息拿不到  具体就是在app里点作者头像，里面是空的（已解决）