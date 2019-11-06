## return后面的defer到底会不会执行？

### return后面的defer会执行

defer函数会在整个函数的生命周期的结束时执行，一般用于回收垃圾，回收资源等。
