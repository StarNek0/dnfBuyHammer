package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	var p1, p2, p3 Point

	startMsg := robotgo.ShowAlert("提示", "点击确定开始录制3个按键", "确定", "取消")
	fmt.Println(startMsg) //确定0，取消1
	if startMsg == 0 {
		fmt.Println("---请按鼠标左键---")

		if m1 := robotgo.AddEvent("mleft"); m1 {
			p1.x, p1.y = robotgo.GetMousePos()
			fmt.Printf("---你按下左键, 坐标m1为(%d, %d)---\n", p1.x, p1.y)

			time.Sleep(time.Millisecond * 300)

			if m2 := robotgo.AddEvent("mleft"); m2 {
				p2.x, p2.y = robotgo.GetMousePos()
				fmt.Printf("---你按下左键, 坐标m2为(%d, %d)---\n", p2.x, p2.y)

				time.Sleep(time.Millisecond * 300)

				if m3 := robotgo.AddEvent("mleft"); m3 {
					p3.x, p3.y = robotgo.GetMousePos()
					fmt.Printf("---你按下左键, 坐标m3为(%d, %d)---\n", p3.x, p3.y)
				}
			}
		}
		endMsg := robotgo.ShowAlert("提示", "录制完毕, 点击确定开始播放，开始后单击右键退出", "确定", "取消")
		fmt.Println(endMsg) //确定0，取消1
		if endMsg == 0 {
			quit := false
			go func(q *bool) {
				var i int
				for quit == false {
					robotgo.MovesClick(p1.x, p1.y, "left", false)
					time.Sleep(time.Millisecond * 500)

					robotgo.MovesClick(p2.x, p2.y, "left", false)
					time.Sleep(time.Millisecond * 500)

					robotgo.MovesClick(p3.x, p3.y, "left", false)
					time.Sleep(time.Millisecond * 500)

					i++
					fmt.Println("执行", i, "次")
				}
			}(&quit)
			quit = robotgo.AddEvent("mright")

		}
	}
	fmt.Println("感谢使用")
}
