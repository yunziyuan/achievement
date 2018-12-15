package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "ShowIdiom",
            Router: `/achievement/showidiom`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "S",
            Router: `/S`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "AddExam",
            Router: `/achievement/addexam`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "GetInformation",
            Router: `/achievement/exam`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/achievement/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:ItemController"] = append(beego.GlobalControllerRouter["achievement/backstage:ItemController"],
        beego.ControllerComments{
            Method: "ShowItem",
            Router: `/achievement/showitem`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/backlogin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "JudgeLogin",
            Router: `/backlogin/judge`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:ScroeController"] = append(beego.GlobalControllerRouter["achievement/backstage:ScroeController"],
        beego.ControllerComments{
            Method: "ShowScore",
            Router: `/achievement/showscore`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
