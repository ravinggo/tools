# keepalive-gen

Record all memory data allocated for a single request, roll back when an error occurs, otherwise assign the modified data to the resident memory data after processing is completed</br>
把单次请求所分配的内存数据全部记录, 当发生错误时回滚，否则处理完成之后会把修改的数据赋值给常驻内存数据

## ***NOTE:***</BR>
Must be used with deepcopy-gen </br>
必须配合 deepcopy-gen 使用

It cannot be used in single-threaded direct memory operation mode because the data of the ***function DoDB cannot cross goroutine***</br>
不能用于单线程直接操作内存模式,因为 ***函数 DoDB 的数据不能跨goroutine***

## sample:
### [output_tests](./output_tests) </br>
 How to use sync.Pool to optimize</br>
 使用 sync.Pool 优化的方式

 
### [output_tests_no_ka](./output_tests_no_ka)
No sync.Pool optimization is used, all memory is allocated directly</br>
未使用sync.Pool优化，所有内存直接分配