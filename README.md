# lingjing37
面试题
Go后端笔试题（研发岗）
1. 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。
difference.go
2. 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
poker.go
3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。
cache.go
4. 实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出，下一个人接着从1开始报数，如此反复，求最后的胜利者编号。
例如，n=3，m=2，那么失败者编号依次是1、0，最后的胜利者是2号。
这里考虑m，n都是正常的数据范围，其中：
1 <= n <= 10^5
1 <= m <= 10^6
算法要求考虑时间效率。
findpeople.go

如无不便，请优先使用github.com，gitlab.com或gitee.com提交答案，回复公开代码库地址即可。否则请使用zip压缩包形式提交代码。
！！！请将附上您的简历及答题！！！发送至邮箱，谢谢，邮箱(base64)：YXJrQGxpbmdqaW5nMzcuY29t
