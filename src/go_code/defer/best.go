package main 

import (
	"fmt"
)

// 当go 执行一个defer时 不会立即执行defer后的语句 而是压入一个栈中
// 然后继续执行后面的语句 
// 当函数执行完毕后	再从defer栈中依次取出执行

// 在defer将语句放入到栈时， 也会将相关的值拷贝入栈 这个特性很关键

