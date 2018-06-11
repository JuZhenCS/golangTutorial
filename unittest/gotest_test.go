//一些注意事项
//测试程序文件名以_test结尾
//import测试包testing
//测试函数以Test开头
package main

import (
	"testing"
)

func Test_Division(t *testing.T) { //test.T记录错误或测试状态
	i, err := Division(6, 2)
	if i != 3 || err != nil {
		t.Error("除法测试没通过!")
	} else {
		t.Log("第一个测试通过了!")
	}
}

func Benchmark_Divisionb(b *testing.B) { //testing.B数据很有用
	b.StopTimer() //停止时间计数
	//做一些前置操作
	//例如测试压缩解压
	//在这里打开文件,读取数据,不会记录时间
	b.StartTimer()             //开始时间计数
	for i := 0; i < b.N; i++ { //b.N是一个很大的数,自动生成
		Division(4, 5)
	}
}
