package ursa_orm

/********************************************************
* @author: Ihc
* @date: 2022/6/10 0010 17:49
* @version: 1.0
* @description: uorm模块

内存
	引用-数据

表管理器(接口)
	1.从数据存储介质中加载数据到内存中
	2.对加载的内存进行操作
	3.清除加载到内存中的数据
	4.将内存中的数据存储到指定的存储介质中
	5.第1条的扩展;从存储介质中恢复数据到内存

	加载数据(接口)
		不同数据来源可有不同的实现；如何来自Json，来自数据库，来自自定义文件
	写出数据(接口)
		同理不同存储介质有不同的实现
	如何修改加载出的数据？
		1.查询内存，返回引用
		2.通过引用修改数据
	如何删除数据？
		方式一：
			1.查询内存，返回引用
			2.将引用传递给管理器，管理器解除其引用
		方式二：
			1.查询内存，返回引用
			2.引用调用删除方法；这需要将管理器引用放在返回引用的数据结构中
	如何提高查询效率？
		1.添加索引
	对于不存在索引的数据如何查询？


*********************************************************/