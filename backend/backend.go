/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-01-20
 */
package backend
import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/nilhost/overnet/backend/site/routers"
)

func RunService() {
	go beego.Run()
}
