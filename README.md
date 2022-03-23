# Simple JVM

SimpleJVM是出于学习目的的简单JVM实现，因此在功能实现上将着重简单实用。

# 模块

- cmd
- classpath
- classloader
- classfile

# cmd

cmd是命令行模块，命令行程序是整个SimpleJVM的入口，它能够接收参数并使用接收到的参数启用JVM。

使用`sj`代表SimpleJVM的命令行程序，执行`sj [option...] mainClass`启动SimpleJVM。

选项：

| 选项      | 简称 | 环境变量       | 描述                                                         |
| --------- | ---- | -------------- | ------------------------------------------------------------ |
| Xjre      |      | $JAVA_HOME/jre | 指定jre所在的目录路径，如果未指定将会读取环境变量中的指定的路径，如果读取失败则停止并报告异常。 |
| classpath | cp   | $CLASSPATH     | 指定应用类加载器的加载路径，多个类路径间使用操作系统相关的路径列表分隔符（window是`;`，linux是`:`）分隔。 |
| version   |      |                | 显示当前SimpleJVM的版本信息。                                |
| help      |      |                | 显示当前SimpleJVM命令行的使用方法。                          |

# classpath

classpath模块用于将资源从某个位置加载到程序内存中，其加载规则使用了className格式，例如`java/lang/Object.class`用于将`java.lang.Object`类的字节码流加载到内存中。

当前支持三种加载条目：

- DirEntry：加载以指定文件夹为根的所有子路径的资源。
- ZipEntry：加载指定压缩包（jar、war、zip）中的资源。
- WildcardEntry：加载指定文件夹下的所有ZipEntry。

# classloader

类加载器将负责加载某一位置的字节码流，并将其转换为JVM运行时的类对象。

# classfile

解析字节码文件为运行时结构。

# 参考

- 自己动手写Java虚拟机
- Java虚拟机规范（Java SE 8）
- 深入理解Java虚拟机