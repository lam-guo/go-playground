1.先对go文件执行go build -o [生成路径] [packages]
2.对编译后文件执行go tool objdump [xxx.o] 可以获得汇编结果