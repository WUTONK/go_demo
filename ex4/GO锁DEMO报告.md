### 互斥锁

一种独占锁，即仅仅允许一个goroutine使用可以获取到锁

### 读写锁

一种并发控制机制，用于在多个读操作和写操作之间提供互斥访问。允许多个goroutine同时进行读操作，但只允许一个写操作进行；使用读锁时，写操作停止，读操作并发。使用写锁时，所有其他操作停止。
