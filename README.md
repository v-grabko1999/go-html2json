# go-html2json
Golang html convert to json.  Golang implementation html2json lib.

#EXAMPLE

main.go

```Go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/v-grabko1999/go-html2json"
)

func main() {
	d, err := html2json.New(strings.NewReader(`
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
			Hello World!
			<p>P</p>
			</body>
		</html>
	`))
	if err != nil {
		log.Fatal(err)
	}
	json, err := d.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}

```

```go run main.go```

Output:
```json
[
    {
        "elements": [
            {
                "name": "html",
                "elements": [
                    {
                        "name": "head",
                        "elements": [
                            {
                                "name": "title",
                                "text": "Hello World"
                            }
                        ]
                    },
                    {
                        "name": "body",
                        "text": "Hello World!",
                        "elements": [
                            {
                                "name": "p",
                                "text": "P"
                            }
                        ]
                    }
                ]
            }
        ]
    }
]
```
Example real parse data: 
```GO
package main

import (
	"fmt"
	"log"
	"strings"

	ht "github.com/v-grabko1999/go-html2json"
	"golang.org/x/net/html/atom"
)

func main() {
	b, err := getPage()
	if err != nil {
		log.Fatal(err)
	}

	d, err := ht.New(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	name, err := parseBookName(d)
	if err != nil {
		log.Fatal(err)
	}

	description, err := parseDescription(d)
	if err != nil {
		log.Fatal(err)
	}

	imageHref, err := parseImageHref(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name: ", name)
	fmt.Println("Image href: ", imageHref)
	fmt.Println("Description: ", description)
}

func parseBookName(d *ht.Dom) (string, error) {
	jieshao_contents, err := d.ByClass("jieshao_content")
	if err != nil {
		return "", err
	}

	h1, err := jieshao_contents[0].ByTag(atom.H1)
	if err != nil {
		return "", err
	}

	a, err := h1[0].ByTag(atom.A)
	if err != nil {
		return "", err
	}
	return a[0].ToNode().Text, nil
}

func parseDescription(d *ht.Dom) (string, error) {
	jieshao_contents, err := d.ByClass("jieshao_content")
	if err != nil {
		return "", err
	}
	h3, err := jieshao_contents[0].ByTag(atom.H3)
	if err != nil {
		return "", err
	}
	return h3[0].ToNode().Text, err
}

func parseImageHref(d *ht.Dom) (string, error) {
	jieshao_img, err := d.ByClass("jieshao-img")
	if err != nil {
		return "", err
	}
	img, err := jieshao_img[0].ByTag(atom.Img)
	if err != nil {
		return "", err
	}
	return img[0].ToNode().Attributes["src"], nil

}

func getPage() ([]byte, error) {
	return []byte(`

	<!DOCTYPE HTML>
	<html>
	<head>
	<title>垮掉的一代GoldenWind全文阅读_无弹窗_第(1)页_UU看书</title>
	<meta http-equiv="Content-Type" content="text/html;charset=gb2312"/>
	<meta name="keywords" content="垮掉的一代GoldenWind全文阅读,无弹窗,王波er" />
	<meta name="description" content="垮掉的一代GoldenWind最新章节连载,垮掉的一代GoldenWind由王波er创作连载。简介：" />
	<meta http-equiv="mobile-agent" content="format=xhtml;url=https://sj.uukanshu.com/book.aspx?id=137365"/>
	<script type="text/javascript" src="//img.uukanshu.com/static/www/js/jquery.min.js"></script>
	
	 <link href="//img.uukanshu.com/static/www/css/main.min.css?v=201712031" rel="stylesheet" type="text/css" />
	<script language="javascript" src="//img.uukanshu.com/static/www/js/transform.min.js" type="text/javascript"></script>
	<script language="javascript" src="//img.uukanshu.com/static/www/js/main.min.js?v=201712031" type="text/javascript"></script>
	
	<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
	</head>
	<body><div class="head" id="head">
	<span class="logo"><a href="/" title="UU看书网"><img src="//img.uukanshu.com/static/www/images/logo.png" style="width:200px;"/></a></span>
	<span class="tc" id="userinfo"></span>
	
	
	<div class="titlebar-desc" style="width:480px;">
	<script async src="https://cse.google.com/cse.js?cx=partner-pub-7553981642580305:qgxyzkokszq"></script>
	<div class="gcse-search"></div>
	<div class="hot-search">
	<div id="queries"></div>
	<script src="https://cse.google.com/query_renderer.js"></script>
	<script src="https://cse.google.com/api/008945028460834109019/cse/kn_kwux2xms/queries/js?view=month&callback=(new+PopularQueryRenderer(document.getElementById(%22queries%22))).render"></script>
	<style>.cse .gsc-control-cse, .gsc-control-cse{padding: 0em;}</style>
	</div> 
	</div>                                                                   
	
	</div>
	<script language="javascript">
		$(function () {
			$.get("/checklogin.ashx?rnd=" + Math.random(), function (data) {
				var t = Cookie.Get("l") == "t";
				var language = t ? "简体中文" : "繁體中文";
				if (data == "no_login") {
					$("#userinfo").html("[<a href='/login.aspx' target='_blank' style='color:red;'>登录</a>]&nbsp;[<a href='/reg.aspx' target='_blank'>注册</a>]&nbsp;[<a id=\"st\" style='color:red' onclick=\"javascript:st();\">" + language + "</a>]");
					isLogin = false;
				} else {
					$("#userinfo").html("<span style='color:red;'><a href='/user/shujia.aspx'>" + data + "</a></span>&nbsp;[<a href='/user/shujia.aspx' style='color:red'>用户中心</a>]&nbsp;[<a href='/loginout.aspx?url=" + encodeURIComponent(window.location.href) + "'>退出</a>]&nbsp;[<a id=\"st\" style='color:red' onclick=\"javascript:st();\">" + language + "</a>]");
					isLogin = true;
				}
			});
		}) 
	 </script>
	<div class="weizhi" style="line-height:0px;">          
	   <div class="path">当前位置：<a href="/">UU看书</a>  &gt; 
	   <a href="/list/yanqing-1.html">都市言情小说</a> &gt; 
	   <a href="/b/137365/" title="垮掉的一代GoldenWind最新章节">垮掉的一代GoldenWind最新章节</a> &gt;   
	   <a href="/t/137365/1/" title="垮掉的一代GoldenWind全文阅读"><b>垮掉的一代GoldenWind全文阅读</b></a>
	   </div>
	</div>
	
	<div style="text-align:center; width:1000px; margin:30px auto;">
	
	<!-- 桌面全文横幅1 -->
	<ins class="adsbygoogle"
		 style="display:inline-block;width:970px;height:90px"
		 data-ad-client="ca-pub-1218476774185134"
		 data-ad-slot="6648015082"></ins>
	  
	<script>
		(adsbygoogle = window.adsbygoogle || []).push({});
	</script>
	</div>
	
	<div style="text-align:center; width:1000px; margin:10px auto;"><script language="javascript" src="//img.uukanshu.com/static/www/js/readset.js"></script>
			  <div class="readset">
					<span class="setcolor padrig">选择背景颜色：<script type="text/javascript">                                                         document.write(ReadConfig.setting_style_bg_color);</script></span>
					<span class="padrig">选择字体：<script type="text/javascript">                                              document.write(ReadConfig.setting_style_font_family);</script></span>
					<span class="padrig">选择字体大小：<script type="text/javascript">                                                document.write(ReadConfig.setting_style_font_size);</script></span>
			  </div>
			  <button class="setreset default_user_setting">恢复默认</button>
	</div>
	<div style="clear:both"></div>
		 
	<div class="xiaoshuo_content clear">          
			  <dl class="jieshao">
				<dt class="jieshao-img">
					<a href="/b/137365/"  title="垮掉的一代GoldenWind最新章节" class="bookImg">
					<img onerror="javascript:this.src='/images/fengmian.jpg'"  src="//img.uukanshu.com/static/www/images/fengmian.jpg" alt="垮掉的一代GoldenWind"/>
					<span class="status-bar"></span>
					<span class="status-text" style="color:#fff">完结</span>               
					</a>
					<div id="bdshare" class="bdshare_t bds_tools get-codes-bdshare">
					<span class="bds_more" style="color:red;padding-top:5px;margin-left:0px;_margin-left:10px;">分享垮掉的一代Go…</span>
					</div>    
				</dt>
				<dd class="jieshao_content">
				  <h1><a href="/t/137365/1/" title="垮掉的一代GoldenWind全文阅读">垮掉的一代GoldenWind全文阅读</a></h1>                
				  <h2>垮掉的一代GoldenWind作者：<a href="/search.aspx?t=1&key=%cd%f5%b2%a8er" target="_blank">王波er</a></h2>
				  <h3>垮掉的一代GoldenWind简介： https://www.uukanshu.com
				  <br/>
					－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－－
				  </h3>
							   
				  <div class="zuixin">
						<b>垮掉的一代GoldenWind最新章节</b>：<a href="/b/137365/13466.html">第二百一十章 重建Y寨</a>
				  </div>             
				  <div class="shijian">更新时间：1个月以前<span id="Span1" style="margin-left:20px;margin-right:20px;"><a  href="javascript:void(0)" title="其它网站已经更新，本站没有更新，点击告诉管理员" onclick="javascript:Report()">没有更新？告诉管理员更新</a></span>
				  <a style="margin-right:20px;" href="/feedback.aspx?tid=137365" title="章节缺失、错误？告诉管理员更正错误" target="_blank">章节缺失、错误报告</a></div>
				  
				  <div class="shijian">
				  <a href="javascript:void(0)" onclick="addShujia();" id="addsj" style="margin-right:20px;">加入书架</a>
				  <a href="javascript:void(0)" onclick="addNotice();" id="addNotice" style="margin-right:20px; font-weight:bold; color:Red;" title="当有最新章节更新时，系统将第一时间发送邮件通知您">更新提醒</a>
				  <a style="margin-right:20px;" href="/txt/137365/" title="垮掉的一代GoldenWindTXT下载" target="_blank">垮掉的一代GoldenWindTXT下载</a>
				  <a href="/desktop.ashx?url=https%3a%2f%2fwww.uukanshu.com%2fb%2f137365%2f%23desktop&title=%bf%e5%b5%f4%b5%c4%d2%bb%b4%faGoldenWind" style="color:Red;margin-right:20px;" title='下载本书快捷方式到电脑桌面，下次可直接点击进入'>添加到桌面</a>
				  <a style="margin-right:20px;" href="/feedback.aspx?tid=137365" title="章节缺失、错误？更新慢？告诉管理员更正错误" target="_blank">缺失错误报告</a>
				  <a href="#" id="closeMulu">隐藏目录</a>
				  </div>
				</dd>
			  </dl>
			  <div class="zhangjie clear"  id="mulu">
				<ul>
				
	   <li class="volume">作品相关</li>
		   <li><a href="/t/137365/1/#13249" title="上架感言">上架感言</a></li>
				<li><a href="/t/137365/1/#13250" title="求投票与">求投票与</a></li>
				<li><a href="/t/137365/1/#13251" title="完结撒花">完结撒花</a></li>
				
	   <li class="volume">象牙塔</li>
		   <li><a href="/t/137365/1/#13253" title="第一章 在初中崭露头角">第一章 在初中崭露头角</a></li>
				<li><a href="/t/137365/1/#13254" title="第二章 网瘾少年">第二章 网瘾少年</a></li>
				<li><a href="/t/137365/2/#13255" title="第三章 车祸">第三章 车祸</a></li>
				<li><a href="/t/137365/2/#13256" title="第四章 文化大学">第四章 文化大学</a></li>
				<li><a href="/t/137365/2/#13257" title="第五章 请客吃饭">第五章 请客吃饭</a></li>
				<li><a href="/t/137365/2/#13258" title="第六章 老曹的秘密">第六章 老曹的秘密</a></li>
				<li><a href="/t/137365/2/#13259" title="第七章 老曹的离开">第七章 老曹的离开</a></li>
				<li><a href="/t/137365/3/#13260" title="第八章 情人节">第八章 情人节</a></li>
				<li><a href="/t/137365/3/#13261" title="第九章 洛阳之旅">第九章 洛阳之旅</a></li>
				<li><a href="/t/137365/3/#13262" title="第一十章 误会">第一十章 误会</a></li>
				<li><a href="/t/137365/3/#13263" title="第一十一章 7夕节">第一十一章 7夕节</a></li>
				<li><a href="/t/137365/3/#13264" title="第一十二章 舔狗">第一十二章 舔狗</a></li>
				<li><a href="/t/137365/4/#13265" title="第一十三章 假期">第一十三章 假期</a></li>
				<li><a href="/t/137365/4/#13266" title="第一十四章 迦叶寺奇遇">第一十四章 迦叶寺奇遇</a></li>
				<li><a href="/t/137365/4/#13267" title="第一十五章 月黑风高">第一十五章 月黑风高</a></li>
				<li><a href="/t/137365/4/#13268" title="第一十六章 旅游社团">第一十六章 旅游社团</a></li>
				<li><a href="/t/137365/4/#13269" title="第一十七章 秦岭山下">第一十七章 秦岭山下</a></li>
				<li><a href="/t/137365/5/#13270" title="第一十八章 送别老部长">第一十八章 送别老部长</a></li>
				<li><a href="/t/137365/5/#13271" title="第一十九章 房地产经纪人">第一十九章 房地产经纪人</a></li>
				<li><a href="/t/137365/5/#13272" title="第二十章 应聘剧组失败">第二十章 应聘剧组失败</a></li>
				<li><a href="/t/137365/5/#13273" title="第二十一章 秘密基地">第二十一章 秘密基地</a></li>
				<li><a href="/t/137365/5/#13274" title="第二十二章 鱼跃龙门">第二十二章 鱼跃龙门</a></li>
				<li><a href="/t/137365/6/#13275" title="第二十三章 同居">第二十三章 同居</a></li>
				<li><a href="/t/137365/6/#13276" title="第二十四章 闺蜜">第二十四章 闺蜜</a></li>
				<li><a href="/t/137365/6/#13277" title="第二十五章 宿舍集体当群演">第二十五章 宿舍集体当群演</a></li>
				<li><a href="/t/137365/6/#13278" title="第二十六章 红玫瑰">第二十六章 红玫瑰</a></li>
				<li><a href="/t/137365/6/#13279" title="第二十七章 写简历的技巧">第二十七章 写简历的技巧</a></li>
				<li><a href="/t/137365/7/#13280" title="第二十八章 回去等通知">第二十八章 回去等通知</a></li>
				<li><a href="/t/137365/7/#13281" title="第二十九章 有希望，初试通过">第二十九章 有希望，初试通过</a></li>
				<li><a href="/t/137365/7/#13282" title="第三十章 第1份工作">第三十章 第1份工作</a></li>
				
	   <li class="volume">软件公司</li>
		   <li><a href="/t/137365/7/#13284" title="第三十一章 第1天上班的不安与激动">第三十一章 第1天上班的不安与激动</a></li>
				<li><a href="/t/137365/7/#13285" title="第三十二章 新来的同事">第三十二章 新来的同事</a></li>
				<li><a href="/t/137365/8/#13286" title="第三十三章 讲真！工作让人生活变得充实">第三十三章 讲真！工作让人生活变得充实</a></li>
				<li><a href="/t/137365/8/#13287" title="第三十四章 第1份工资">第三十四章 第1份工资</a></li>
				<li><a href="/t/137365/8/#13288" title="第三十五章 出差途中，邵海南回忆东北岁月">第三十五章 出差途中，邵海南回忆东北岁月</a></li>
				<li><a href="/t/137365/8/#13289" title="第三十六章 柳泉天下第1蘸水面">第三十六章 柳泉天下第1蘸水面</a></li>
				<li><a href="/t/137365/8/#13290" title="第三十七章 柳泉农保办业务培训大会">第三十七章 柳泉农保办业务培训大会</a></li>
				<li><a href="/t/137365/9/#13291" title="第三十八章 多嘴招人烦">第三十八章 多嘴招人烦</a></li>
				<li><a href="/t/137365/9/#13292" title="第三十九章 接风洗尘">第三十九章 接风洗尘</a></li>
				<li><a href="/t/137365/9/#13293" title="第四十章 酒过3巡">第四十章 酒过3巡</a></li>
				<li><a href="/t/137365/9/#13294" title="第四十一章 夜未眠">第四十一章 夜未眠</a></li>
				<li><a href="/t/137365/9/#13295" title="第四十二章 渭水县的培训">第四十二章 渭水县的培训</a></li>
				<li><a href="/t/137365/10/#13296" title="第四十三章 瓦岗寨夜宴">第四十三章 瓦岗寨夜宴</a></li>
				<li><a href="/t/137365/10/#13297" title="第四十四章 醉生梦死">第四十四章 醉生梦死</a></li>
				<li><a href="/t/137365/10/#13298" title="第四十五章 后花园的影">第四十五章 后花园的影</a></li>
				<li><a href="/t/137365/10/#13299" title="第四十六章 大学毕业，散伙饭">第四十六章 大学毕业，散伙饭</a></li>
				<li><a href="/t/137365/10/#13300" title="第四十七章 各奔东西">第四十七章 各奔东西</a></li>
				<li><a href="/t/137365/11/#13301" title="第四十八章 去往南京">第四十八章 去往南京</a></li>
				<li><a href="/t/137365/11/#13302" title="第四十九章 开幕式">第四十九章 开幕式</a></li>
				<li><a href="/t/137365/11/#13303" title="第五十章 夜游秦淮河">第五十章 夜游秦淮河</a></li>
				<li><a href="/t/137365/11/#13304" title="第五十一章 一千九百一十二酒吧风情街">第五十一章 一千九百一十二酒吧风情街</a></li>
				<li><a href="/t/137365/11/#13305" title="第五十二章 苏荷酒吧">第五十二章 苏荷酒吧</a></li>
				<li><a href="/t/137365/12/#13306" title="第五十三章 赵琳计划开店">第五十三章 赵琳计划开店</a></li>
				<li><a href="/t/137365/12/#13307" title="第五十四章 漫川">第五十四章 漫川</a></li>
				<li><a href="/t/137365/12/#13308" title="第五十五章 和喻队的争执">第五十五章 和喻队的争执</a></li>
				<li><a href="/t/137365/12/#13309" title="第五十六章 难忘的1课">第五十六章 难忘的1课</a></li>
				<li><a href="/t/137365/12/#13310" title="第五十七章 辞职">第五十七章 辞职</a></li>
				
	   <li class="volume">OnePiece</li>
		   <li><a href="/t/137365/13/#13312" title="第五十八章 都是钱惹的祸">第五十八章 都是钱惹的祸</a></li>
				<li><a href="/t/137365/13/#13313" title="第五十九章 偶遇李超凡">第五十九章 偶遇李超凡</a></li>
				<li><a href="/t/137365/13/#13314" title="第六十章 入职飞科网络">第六十章 入职飞科网络</a></li>
				<li><a href="/t/137365/13/#13315" title="第六十一章 我被骗了">第六十一章 我被骗了</a></li>
				<li><a href="/t/137365/13/#13316" title="第六十二章 收到荣盛科技的offer">第六十二章 收到荣盛科技的offer</a></li>
				<li><a href="/t/137365/14/#13317" title="第六十三章 踌躇满志">第六十三章 踌躇满志</a></li>
				<li><a href="/t/137365/14/#13318" title="第六十四章 第N次约会">第六十四章 第N次约会</a></li>
				<li><a href="/t/137365/14/#13319" title="第六十五章 团建OR饯行">第六十五章 团建OR饯行</a></li>
				<li><a href="/t/137365/14/#13320" title="第六十六章 危险信号">第六十六章 危险信号</a></li>
				<li><a href="/t/137365/14/#13321" title="第六十七章 姜还是老的辣">第六十七章 姜还是老的辣</a></li>
				<li><a href="/t/137365/15/#13322" title="第六十八章 年会聚餐">第六十八章 年会聚餐</a></li>
				<li><a href="/t/137365/15/#13323" title="第六十九章 祸不单行">第六十九章 祸不单行</a></li>
				<li><a href="/t/137365/15/#13324" title="第七十章 出差郑州">第七十章 出差郑州</a></li>
				<li><a href="/t/137365/15/#13325" title="第七十一章 浑浑噩噩">第七十一章 浑浑噩噩</a></li>
				<li><a href="/t/137365/15/#13326" title="第七十二章 习惯">第七十二章 习惯</a></li>
				<li><a href="/t/137365/16/#13327" title="第七十三章 枣庄岁月">第七十三章 枣庄岁月</a></li>
				<li><a href="/t/137365/16/#13328" title="第七十四章 艳遇">第七十四章 艳遇</a></li>
				<li><a href="/t/137365/16/#13329" title="第七十五章 SS酒吧驻唱">第七十五章 SS酒吧驻唱</a></li>
				<li><a href="/t/137365/16/#13330" title="第七十六章 疑窦">第七十六章 疑窦</a></li>
				<li><a href="/t/137365/16/#13331" title="第七十七章 周志伟回忆青春岁月">第七十七章 周志伟回忆青春岁月</a></li>
				<li><a href="/t/137365/17/#13332" title="第七十八章 敦煌">第七十八章 敦煌</a></li>
				<li><a href="/t/137365/17/#13333" title="第七十九章 咖啡厅奇遇记">第七十九章 咖啡厅奇遇记</a></li>
				<li><a href="/t/137365/17/#13334" title="第八十章 陈婉如赈灾筹款">第八十章 陈婉如赈灾筹款</a></li>
				<li><a href="/t/137365/17/#13335" title="第八十一章 出轨">第八十一章 出轨</a></li>
				<li><a href="/t/137365/17/#13336" title="第八十二章 和赵琳分手">第八十二章 和赵琳分手</a></li>
				<li><a href="/t/137365/18/#13337" title="第八十三章 陈婉如远飞香港当练习生">第八十三章 陈婉如远飞香港当练习生</a></li>
				<li><a href="/t/137365/18/#13338" title="第八十四章 再见我的爱情">第八十四章 再见我的爱情</a></li>
				<li><a href="/t/137365/18/#13339" title="第八十五章 天桥摆摊">第八十五章 天桥摆摊</a></li>
				<li><a href="/t/137365/18/#13340" title="第八十六章 无奸不商">第八十六章 无奸不商</a></li>
				<li><a href="/t/137365/18/#13341" title="第八十七章 城管问题">第八十七章 城管问题</a></li>
				<li><a href="/t/137365/19/#13342" title="第八十八章 纳兰性德项目">第八十八章 纳兰性德项目</a></li>
				<li><a href="/t/137365/19/#13343" title="第八十九章 人群中的骗子">第八十九章 人群中的骗子</a></li>
				<li><a href="/t/137365/19/#13344" title="第九十章 国学辩论会1番战">第九十章 国学辩论会1番战</a></li>
				<li><a href="/t/137365/19/#13345" title="第九十一章 国学辩论会2番战">第九十一章 国学辩论会2番战</a></li>
				<li><a href="/t/137365/19/#13346" title="第九十二章 国学辩论会终章">第九十二章 国学辩论会终章</a></li>
				<li><a href="/t/137365/20/#13347" title="第九十三章 吹牛高手">第九十三章 吹牛高手</a></li>
				<li><a href="/t/137365/20/#13348" title="第九十四章 老曹成为总监">第九十四章 老曹成为总监</a></li>
				<li><a href="/t/137365/20/#13349" title="第九十五章 工作和自由的矛盾">第九十五章 工作和自由的矛盾</a></li>
				<li><a href="/t/137365/20/#13350" title="第九十六章 问道紫霄宫">第九十六章 问道紫霄宫</a></li>
				<li><a href="/t/137365/20/#13351" title="第九十七章 为理想肆意妄为的岁月">第九十七章 为理想肆意妄为的岁月</a></li>
				<li><a href="/t/137365/21/#13352" title="第九十八章 国学讲师薛鹏">第九十八章 国学讲师薛鹏</a></li>
				<li><a href="/t/137365/21/#13353" title="第九十九章 老刘卷款失踪">第九十九章 老刘卷款失踪</a></li>
				<li><a href="/t/137365/21/#13354" title="第一百章 青旅客栈">第一百章 青旅客栈</a></li>
				<li><a href="/t/137365/21/#13355" title="第一百零一章 飞往深圳做电商">第一百零一章 飞往深圳做电商</a></li>
				<li><a href="/t/137365/21/#13356" title="第一百零二章 师公会">第一百零二章 师公会</a></li>
				<li><a href="/t/137365/22/#13357" title="第一百零三章 与陈旗相聚广州">第一百零三章 与陈旗相聚广州</a></li>
				<li><a href="/t/137365/22/#13358" title="第一百零四章 上上签">第一百零四章 上上签</a></li>
				<li><a href="/t/137365/22/#13359" title="第一百零五章 大师，我悟了！">第一百零五章 大师，我悟了！</a></li>
				<li><a href="/t/137365/22/#13360" title="第一百零六章 我被骗得身无分文">第一百零六章 我被骗得身无分文</a></li>
				<li><a href="/t/137365/22/#13361" title="第一百零七章 银龙网吧的名人">第一百零七章 银龙网吧的名人</a></li>
				<li><a href="/t/137365/23/#13362" title="第一百零八章 荣归故里">第一百零八章 荣归故里</a></li>
				
	   <li class="volume">长岛市的冒险</li>
		   <li><a href="/t/137365/23/#13364" title="第一百零九章 不良少年">第一百零九章 不良少年</a></li>
				<li><a href="/t/137365/23/#13365" title="第一百一十章 拜码头">第一百一十章 拜码头</a></li>
				<li><a href="/t/137365/23/#13366" title="第一百一十一章 刺探情报">第一百一十一章 刺探情报</a></li>
				<li><a href="/t/137365/23/#13367" title="第一百一十二章 少管所探访张源治">第一百一十二章 少管所探访张源治</a></li>
				<li><a href="/t/137365/24/#13368" title="第一百一十三章 FeelingClub酒吧齐聚">第一百一十三章 FeelingClub酒吧齐聚</a></li>
				<li><a href="/t/137365/24/#13369" title="第一百一十四章 大飞哥带头翘课">第一百一十四章 大飞哥带头翘课</a></li>
				<li><a href="/t/137365/24/#13370" title="第一百一十五章 战争导火线">第一百一十五章 战争导火线</a></li>
				<li><a href="/t/137365/24/#13371" title="第一百一十六章 仇恨的种子">第一百一十六章 仇恨的种子</a></li>
				<li><a href="/t/137365/24/#13372" title="第一百一十七章 前辈的希望">第一百一十七章 前辈的希望</a></li>
				<li><a href="/t/137365/25/#13373" title="第一百一十八章 造访雷霆学院">第一百一十八章 造访雷霆学院</a></li>
				<li><a href="/t/137365/25/#13374" title="第一百一十九章 战书">第一百一十九章 战书</a></li>
				<li><a href="/t/137365/25/#13375" title="第一百二十章 8卦">第一百二十章 8卦</a></li>
				<li><a href="/t/137365/25/#13376" title="第一百二十一章 设局">第一百二十一章 设局</a></li>
				<li><a href="/t/137365/25/#13377" title="第一百二十二章 裁员">第一百二十二章 裁员</a></li>
				<li><a href="/t/137365/26/#13378" title="第一百二十三章 谋略">第一百二十三章 谋略</a></li>
				<li><a href="/t/137365/26/#13379" title="第一百二十四章 战斗开始了">第一百二十四章 战斗开始了</a></li>
				<li><a href="/t/137365/26/#13380" title="第一百二十五章 铃木双雄">第一百二十五章 铃木双雄</a></li>
				<li><a href="/t/137365/26/#13381" title="第一百二十六章 新秩序">第一百二十六章 新秩序</a></li>
				<li><a href="/t/137365/26/#13382" title="第一百二十七章 联盟">第一百二十七章 联盟</a></li>
				<li><a href="/t/137365/27/#13383" title="第一百二十八章 招募人才">第一百二十八章 招募人才</a></li>
				<li><a href="/t/137365/27/#13384" title="第一百二十九章 新大陆网吧的电子竞技">第一百二十九章 新大陆网吧的电子竞技</a></li>
				<li><a href="/t/137365/27/#13385" title="第一百三十章 比武">第一百三十章 比武</a></li>
				<li><a href="/t/137365/27/#13386" title="第一百三十一章 避风港">第一百三十一章 避风港</a></li>
				<li><a href="/t/137365/27/#13387" title="第一百三十二章 雄心壮志">第一百三十二章 雄心壮志</a></li>
				<li><a href="/t/137365/28/#13388" title="第一百三十三章 黑关工业">第一百三十三章 黑关工业</a></li>
				<li><a href="/t/137365/28/#13389" title="第一百三十四章 龙腾">第一百三十四章 龙腾</a></li>
				<li><a href="/t/137365/28/#13390" title="第一百三十五章 白鹤与上官云当众赋诗">第一百三十五章 白鹤与上官云当众赋诗</a></li>
				<li><a href="/t/137365/28/#13391" title="第一百三十六章 游泳与对弈">第一百三十六章 游泳与对弈</a></li>
				<li><a href="/t/137365/28/#13392" title="第一百三十七章 铁血联盟干部会议">第一百三十七章 铁血联盟干部会议</a></li>
				<li><a href="/t/137365/29/#13393" title="第一百三十八章 突袭幽冥">第一百三十八章 突袭幽冥</a></li>
				<li><a href="/t/137365/29/#13394" title="第一百三十九章 神田真希VS成濑谦介">第一百三十九章 神田真希VS成濑谦介</a></li>
				<li><a href="/t/137365/29/#13395" title="第一百四十章 100箱啤酒的赌注">第一百四十章 100箱啤酒的赌注</a></li>
				<li><a href="/t/137365/29/#13396" title="第一百四十一章 总长9龙刺史上阵">第一百四十一章 总长9龙刺史上阵</a></li>
				<li><a href="/t/137365/29/#13397" title="第一百四十二章 幽灵旗被夺">第一百四十二章 幽灵旗被夺</a></li>
				<li><a href="/t/137365/30/#13398" title="第一百四十三章 援军抵达幽冥总部">第一百四十三章 援军抵达幽冥总部</a></li>
				<li><a href="/t/137365/30/#13399" title="第一百四十四章 合纵连横">第一百四十四章 合纵连横</a></li>
				<li><a href="/t/137365/30/#13400" title="第一百四十五章 背叛者神田真希">第一百四十五章 背叛者神田真希</a></li>
				<li><a href="/t/137365/30/#13401" title="第一百四十六章 极限网吧开业">第一百四十六章 极限网吧开业</a></li>
				<li><a href="/t/137365/30/#13402" title="第一百四十七章 魔兽争霸比赛">第一百四十七章 魔兽争霸比赛</a></li>
				<li><a href="/t/137365/31/#13403" title="第一百四十八章 星际争霸比赛">第一百四十八章 星际争霸比赛</a></li>
				<li><a href="/t/137365/31/#13404" title="第一百四十九章 决战，英雄联盟比赛">第一百四十九章 决战，英雄联盟比赛</a></li>
				<li><a href="/t/137365/31/#13405" title="第一百五十章 驱车逃离">第一百五十章 驱车逃离</a></li>
				<li><a href="/t/137365/31/#13406" title="第一百五十一章 众叛亲离">第一百五十一章 众叛亲离</a></li>
				<li><a href="/t/137365/31/#13407" title="第一百五十二章 复仇计划">第一百五十二章 复仇计划</a></li>
				<li><a href="/t/137365/32/#13408" title="第一百五十三章 花溪村">第一百五十三章 花溪村</a></li>
				<li><a href="/t/137365/32/#13409" title="第一百五十四章 黑客军团">第一百五十四章 黑客军团</a></li>
				<li><a href="/t/137365/32/#13410" title="第一百五十五章 网吧集体掉线">第一百五十五章 网吧集体掉线</a></li>
				<li><a href="/t/137365/32/#13411" title="第一百五十六章 双面间谍">第一百五十六章 双面间谍</a></li>
				<li><a href="/t/137365/32/#13412" title="第一百五十七章 金蝉脱壳">第一百五十七章 金蝉脱壳</a></li>
				<li><a href="/t/137365/33/#13413" title="第一百五十八章 9龙败退">第一百五十八章 9龙败退</a></li>
				<li><a href="/t/137365/33/#13414" title="第一百五十九章 山崎组被收购">第一百五十九章 山崎组被收购</a></li>
				<li><a href="/t/137365/33/#13415" title="第一百六十章 0岛会">第一百六十章 0岛会</a></li>
				<li><a href="/t/137365/33/#13416" title="第一百六十一章 神社决斗">第一百六十一章 神社决斗</a></li>
				<li><a href="/t/137365/33/#13417" title="第一百六十二章 秦泽赢了">第一百六十二章 秦泽赢了</a></li>
				<li><a href="/t/137365/34/#13418" title="第一百六十三章 4大天王">第一百六十三章 4大天王</a></li>
				<li><a href="/t/137365/34/#13419" title="第一百六十四章 转学铃木">第一百六十四章 转学铃木</a></li>
				<li><a href="/t/137365/34/#13420" title="第一百六十五章 我的未来">第一百六十五章 我的未来</a></li>
				
	   <li class="volume">闯荡北上广</li>
		   <li><a href="/t/137365/34/#13422" title="第一百六十六章 初到上海">第一百六十六章 初到上海</a></li>
				<li><a href="/t/137365/34/#13423" title="第一百六十七章 张江镇租房">第一百六十七章 张江镇租房</a></li>
				<li><a href="/t/137365/35/#13424" title="第一百六十八章 捷成科技">第一百六十八章 捷成科技</a></li>
				<li><a href="/t/137365/35/#13425" title="第一百六十九章 步入正轨">第一百六十九章 步入正轨</a></li>
				<li><a href="/t/137365/35/#13426" title="第一百七十章 人员不足">第一百七十章 人员不足</a></li>
				<li><a href="/t/137365/35/#13427" title="第一百七十一章 程序员加班猝死">第一百七十一章 程序员加班猝死</a></li>
				<li><a href="/t/137365/35/#13428" title="第一百七十二章 去铃铛家借宿">第一百七十二章 去铃铛家借宿</a></li>
				<li><a href="/t/137365/36/#13429" title="第一百七十三章 跳槽">第一百七十三章 跳槽</a></li>
				<li><a href="/t/137365/36/#13430" title="第一百七十四章 与上家公司的纠结">第一百七十四章 与上家公司的纠结</a></li>
				<li><a href="/t/137365/36/#13431" title="第一百七十五章 镜花水月的爱情">第一百七十五章 镜花水月的爱情</a></li>
				<li><a href="/t/137365/36/#13432" title="第一百七十六章 裁员">第一百七十六章 裁员</a></li>
				<li><a href="/t/137365/36/#13433" title="第一百七十七章 新官上任3把火">第一百七十七章 新官上任3把火</a></li>
				<li><a href="/t/137365/37/#13434" title="第一百七十八章 无休止的加班">第一百七十八章 无休止的加班</a></li>
				<li><a href="/t/137365/37/#13435" title="第一百七十九章 升职总监">第一百七十九章 升职总监</a></li>
				<li><a href="/t/137365/37/#13436" title="第一百八十章 房价问题">第一百八十章 房价问题</a></li>
				<li><a href="/t/137365/37/#13437" title="第一百八十一章 悼念青春">第一百八十一章 悼念青春</a></li>
				<li><a href="/t/137365/37/#13438" title="第一百八十二章 买房">第一百八十二章 买房</a></li>
				<li><a href="/t/137365/38/#13439" title="第一百八十三章 承接项目">第一百八十三章 承接项目</a></li>
				<li><a href="/t/137365/38/#13440" title="第一百八十四章 牵线搭桥">第一百八十四章 牵线搭桥</a></li>
				<li><a href="/t/137365/38/#13441" title="第一百八十五章 合作成功">第一百八十五章 合作成功</a></li>
				<li><a href="/t/137365/38/#13442" title="第一百八十六章 职场门道">第一百八十六章 职场门道</a></li>
				<li><a href="/t/137365/38/#13443" title="第一百八十七章 洞察先机">第一百八十七章 洞察先机</a></li>
				<li><a href="/t/137365/39/#13444" title="第一百八十八章 奖励宝马">第一百八十八章 奖励宝马</a></li>
				<li><a href="/t/137365/39/#13445" title="第一百八十九章 人生赢家">第一百八十九章 人生赢家</a></li>
				<li><a href="/t/137365/39/#13446" title="第一百九十章 停电事故">第一百九十章 停电事故</a></li>
				<li><a href="/t/137365/39/#13447" title="第一百九十一章 5环飙车">第一百九十一章 5环飙车</a></li>
				<li><a href="/t/137365/39/#13448" title="第一百九十二章 火车奇遇">第一百九十二章 火车奇遇</a></li>
				<li><a href="/t/137365/40/#13449" title="第一百九十三章 陈希廉调研溪贤市">第一百九十三章 陈希廉调研溪贤市</a></li>
				<li><a href="/t/137365/40/#13450" title="第一百九十四章 云游大理">第一百九十四章 云游大理</a></li>
				<li><a href="/t/137365/40/#13451" title="第一百九十五章 死生之地">第一百九十五章 死生之地</a></li>
				<li><a href="/t/137365/40/#13452" title="第一百九十六章 遥望北京">第一百九十六章 遥望北京</a></li>
				<li><a href="/t/137365/40/#13453" title="第一百九十七章 融资计划">第一百九十七章 融资计划</a></li>
				<li><a href="/t/137365/41/#13454" title="第一百九十八章 资本对赌协议">第一百九十八章 资本对赌协议</a></li>
				<li><a href="/t/137365/41/#13455" title="第一百九十九章 美股上市">第一百九十九章 美股上市</a></li>
				<li><a href="/t/137365/41/#13456" title="第一百二十章 人生巅峰">第一百二十章 人生巅峰</a></li>
				<li><a href="/t/137365/41/#13457" title="第二百零一章 陆应雄辞职">第二百零一章 陆应雄辞职</a></li>
				<li><a href="/t/137365/41/#13458" title="第二百零二章 周志伟被董事会罢免">第二百零二章 周志伟被董事会罢免</a></li>
				<li><a href="/t/137365/42/#13459" title="第二百零三章 花天酒地">第二百零三章 花天酒地</a></li>
				<li><a href="/t/137365/42/#13460" title="第二百零四章 比特币矿场">第二百零四章 比特币矿场</a></li>
				<li><a href="/t/137365/42/#13461" title="第二百零五章 夺回华途的控制权">第二百零五章 夺回华途的控制权</a></li>
				<li><a href="/t/137365/42/#13462" title="第二百零六章 玫瑰花的葬礼">第二百零六章 玫瑰花的葬礼</a></li>
				<li><a href="/t/137365/42/#13463" title="第二百零七章 大结局">第二百零七章 大结局</a></li>
				<li><a href="/t/137365/43/#13464" title="第二百零八章 回忆融资细节">第二百零八章 回忆融资细节</a></li>
				<li><a href="/t/137365/43/#13465" title="第二百零九章 垮掉的1代">第二百零九章 垮掉的1代</a></li>
				<li><a href="/t/137365/43/#13466" title="第二百一十章 重建Y寨">第二百一十章 重建Y寨</a></li>
				</ul>
			  </div>
	</div>
	<div class="post">
	
	<div class="postTitle"><a name="13250">求投票与</a><br/><span class="topic">垮掉的一代GoldenWind全文阅读</span><span class="author">作者：王波er</span><span class="opt" title="把本章节加入书架，方便下次阅读" onclick="addShujia('13250','求投票与'); return false;">加入书架</span></div>
	<div class="postContent">
	<p>　　《垮掉的一代GoldenWind》连载已经有一段时间了，感谢编辑大人和读者朋友们的厚爱！在这里，我真的要向你们三鞠躬了！我发誓我会非常认真地写作，为大家提供一个最佳的阅读体验，力争把该小说的故事讲得完美无瑕！为大家带来飞一般的视听体验。<p>　　可是，即便如此。俗话说：“酒香也怕巷子深！”我会很努力写好这部小说，同时，也希望大家能够多多投票、收藏，当然如果有打赏和订阅就更好了！在这里，我要为给我投票和订阅的朋友们诚挚地说声：“谢谢您的支持！能够让读者满意，我也很开心！”毕竟，生活中的很多事情，不就是为了让自己活得开心吗？能够得到很好的阅读体验也不错。<p>　　为了跟广大读者朋友们更好地交流，我已经开通了读者QQ群：4904903，欢迎大家进来闲聊，互相学习！谈论生活中的一切。最后，希望大家能够继续帮我投票哟、收藏、订阅三连，谢谢啦！您的支持会让我更加努力！我相信，我一定不会辜负大家的期望。如果您喜欢本作品，也可以帮忙推荐给其他朋友的，谢谢。<p> 
	</div>
	
		   <div class="dirAd" style="text-align:center;">
						<!-- 桌面全文横幅2 -->
							<ins class="adsbygoogle"
								 style="display:inline-block;width:970px;height:90px"
								 data-ad-client="ca-pub-1218476774185134"
								 data-ad-region="content"
								 data-ad-slot="2684719341"></ins>
						  
						<script>
							(adsbygoogle = window.adsbygoogle || []).push({});
						</script>
			</div>
	
	<div class="postTitle"><a name="13251">完结撒花</a><br/><span class="topic">垮掉的一代GoldenWind全文阅读</span><span class="author">作者：王波er</span><span class="opt" title="把本章节加入书架，方便下次阅读" onclick="addShujia('13251','完结撒花'); return false;">加入书架</span></div>
	<div class="postContent">
	<p>　　很荣幸，很自豪，《垮掉的一代GoldenWind》是我在起点的第一部完结的小说，这部小说从连载到完结花费了3个月的时间，在这3个月里，我认真地努力地去写作了，但还是有一些不足，相信在未来的写作之旅中，我会尽力克服这些不足，带给读者更多的经典作品，谢谢大家！<p>　　这部小说完结了，我的一个心愿也实现了，希望它能取得一个不错的成绩吧，也希望读者朋友们可以喜欢！最后，再一次感谢起点和编辑老师，谢谢你们慧眼识珠，还有提供了这个平台给我，给广大的网络作家，希望起点网越来越好！也希望读者朋友们生活愉快！<p>　　最后，庆祝一下自己的小说《垮掉的一代GoldenWind》正式完结，希望在不久的将来，我们还能够在其他的小说中，结下彼此的缘分，见证彼此的友谊！<p> 
	</div>
	
		   <div class="dirAd" style="text-align:center;">
						<!-- 桌面全文横幅2 -->
							<ins class="adsbygoogle"
								 style="display:inline-block;width:970px;height:90px"
								 data-ad-client="ca-pub-1218476774185134"
								 data-ad-region="content"
								 data-ad-slot="2684719341"></ins>
						  
						<script>
							(adsbygoogle = window.adsbygoogle || []).push({});
						</script>
			</div>
	
	<div class="postTitle"><a name="13253">第一章 在初中崭露头角</a><br/><span class="topic">垮掉的一代GoldenWind全文阅读</span><span class="author">作者：王波er</span><span class="opt" title="把本章节加入书架，方便下次阅读" onclick="addShujia('13253','第一章 在初中崭露头角'); return false;">加入书架</span></div>
	<div class="postContent">
	<p>　　张帆出生在遥远的西部城市张掖，这是个非常特别的城市，被称作小西安。原因很简单，四四方方的城市布局，再加上市中心那座鼓楼，俨然一副西安的模样，可是面积就小太多啦！张帆这个人，生性腼腆、胆小，甚至有点自闭，在小学和初中的时候，他都是班里的好学生，每次考试，总是名列前茅。有时候他在想，这每一天上学的日子，什么时候才能结束？他已经等不及要开始向世界证明自己，来施展自己的抱负了。为什么命运的齿轮不开始旋转呢？他很小的时候就有着这样的疑惑。<p>　　记得小学的时候，他努力学习，把老师交代给他的任务完成得特别出色，拿了不少金灿灿的奖状，都挂在家里的墙上，有三好学生、优秀班干部等等，这些奖状说实话值不了几个钱，但代表着一种荣誉。可他的小学生活与其他人没什么两样，他除了学习好点之外，就是象棋下得好，经常找那些老头对弈。去小学的路上，总是那样的风和日丽，即便是到了冬天，他也从来不慌不忙。可惜好景不长，到了初中的时候，这一切都变了。<p>　　那就是学习压力太紧张了，紧张得他透不过气来。初一的时候分班，他时常在想，自己身边要是坐一个长相俊美的女同学就好了，然而事实总是不遂心愿，他身旁坐了一个男同学，还是那种非常普通的。但这并不影响什么，反正老师隔三差五就要调座位嘛。可是，他等呀等呀，一连几周等没有等到这种机会，这让他非常难过。<p>　　他开始胡思乱想，自己小学的时候是什么？三好学生、优秀班干部，还担任了学校的护旗手，可是到了初中，他却什么都没有了！这让他愤愤不平。凭什么，这一切都凭什么？然而，年少的他很快找到了应对的办法，那就是努力学习，争取让老师们注意到他。初一期中考试，他考到了班级第一，年级第30，这可给他们这个初一16班争了光。这个班级，说实话太一般了，一般的名字，一般的学生，连教学楼都在最偏远的角落。<p>　　班主任非常看好他，亲自委任张帆为学习委员，这让他倍有面子！怎么说呢？张帆这孩子，在文学上有天赋，小小年纪就博览群书，在课堂上回答问题的时候，也颇有自己的一套见解，老师们都说他能考上清华大学，他也洋洋得意起来。而且，这个班级第一让他在班里有了很大的发言权，他以辅导功课为由，把班里的女生小霞调到了自己的旁边，这样，他总算是实现了人生的第一个目标，那就是和女生坐在一起。当然，这在若干年后，他攀登人生最高峰的时候，看起来简直是不值一提！可在当时，他就是觉得很有自豪感。<p>　　达成人生的第一个成就之后，张帆迷上了计算机课程，动不动就去微机室练习，还在课余的时候翻计算机书籍，试图成为一名编程高手。当然，程序都是英文写的，这让他备受打击，所以他坚持了一段时间就暂时搁浅了。<p>　　值得一提的是，语文老师仿佛看到了一颗文学界的新星正在冉冉升起，于是主动借书给张帆，让他学习，他借给张帆的第一本书是奥斯特洛夫斯基的《钢铁是怎样炼成的》，这书乍一看是炼钢的，张帆以为老师想让他当工人，所以这书他就不好意思的收下来，便放到书架子上，不去管他。倒是四大名著，他很喜欢，比如《三国演义》，以至于他有事没事就开始看三国，被称为三国迷。<p>　　但是，看三国也有好处，那就是他的写作能力得到了提升，以至于他轻而易举地在校刊上发表了一篇散文《我的电脑》，散文中讲解了他和电脑是怎么认识的，他怎么折腾电脑的，把电脑拆了又怎么装回去，总之那电脑挺惨的。这篇文章一发表，他就成了全校的知名人物了，可谓是无限风光！<p>　　然而，初一的他，却逐渐走上了命运安排的道路，当然，这路上有快乐也有悲伤。他自己折腾了好一阵子电脑，觉得没意思了，就开始和同学们三五成群的去网吧消遣！可这网吧，毕竟是鱼目混珠，有好人有坏人，也有传奇的人。张帆会遇见什么奇怪的事情呢，他的未来，大概与网吧有了很大的羁绊，每当他回忆起那个和网吧相遇的日子，他就感慨万千：<p>　　“那是一个阳光灿烂的午后，我在放学之后，从学校回家。当我走在街道上的时候，无聊地哼起了罗大佑的《童年》，可走着走着，我就看见了一个叫做启明星的网吧，这个招牌挂在楼上，很有吸引力，我不知不觉地就走了进去，却一下子见识到了互联网的魅力！”<p> 
	</div>
	
		   <div class="dirAd" style="text-align:center;">
						<!-- 桌面全文横幅2 -->
							<ins class="adsbygoogle"
								 style="display:inline-block;width:970px;height:90px"
								 data-ad-client="ca-pub-1218476774185134"
								 data-ad-region="content"
								 data-ad-slot="2684719341"></ins>
						  
						<script>
							(adsbygoogle = window.adsbygoogle || []).push({});
						</script>
			</div>
	
	<div class="postTitle"><a name="13254">第二章 网瘾少年</a><br/><span class="topic">垮掉的一代GoldenWind全文阅读</span><span class="author">作者：王波er</span><span class="opt" title="把本章节加入书架，方便下次阅读" onclick="addShujia('13254','第二章 网瘾少年'); return false;">加入书架</span></div>
	<div class="postContent">
	<p>　　自从张帆去了一次启明星网吧后，给他幼小的心灵带来了震撼！天呐，世界上还有这种地方。那个年代的网吧，非常朴素，可不像现在的网咖那么豪华。那个时代的网吧，分为平民区和包间区，根本没有电竞区，因为电子竞技什么的，根本不存在的！<p>　　张帆一下子打开了新世界的大门，一头扎在启明星网吧里度过了几十个小时，可以说是废寝忘食了。他第一个接触到的网络事务竟然不是QQ，还是网易聊天室。网易聊天室可有意思了，可以在163网站里选择聊天室，比如可以选择全国各地的某一个城市，其中就有甘肃。张帆在聊天室里聊啊聊，跟很多陌生人聊天，那可真算得上是网上冲浪了，不过聊来聊去也没什么意思，毕竟都是陌生人，也见不到人。接着，张帆就接触到了QQ，他注册了一个，网名叫做追风少年。<p>　　回到学校，张帆就可以吹嘘自己上网的事情，这不说不知道，一说吓一跳，原来班里已经有好几个同学开始泡网吧了。比如，王贵他就在玩大话西游，还有好几个人。张帆一听，惊讶了，他们说的是啥，我得补补。于是，他放学就回到了启明星，开始玩大话，然后他又发现了梦幻。于是乎，他就开始玩大话西游和梦幻西游这两款游戏。<p>　　日子就这样的确定下来，仿佛一切都步入了正规，又仿佛一切都那么虚无缥缈。记得有多少个阳光灿烂的午后，还有多少个大雨纷纷的傍晚，张帆就这样扎进启明星网吧，开始疯狂地寻找自己的人生。这一晃，就是一年过去了，他初二了。虽然他的学习成绩下降了一点，但仍然保持在班级第十。家人们开始留意他早出晚归的情况，最后总算是认定了他成为了一个网瘾少年，当然，这并没有什么，张帆并非控制不住自己。<p>　　有时候，张帆在想，人生的意义究竟是什么？这个天气晴朗的午后，我去网吧看电影，打CS，度过了一个美妙绝伦的日子，可是到了下机的时候，我揉揉疲惫惺忪的眼睛，又觉得有点虚度青春。他走在回家的路上，庸庸碌碌的不知道所以。在他幼小的心里，他时长在想外面的世界究竟是什么样子的？可以他的能力，别说是去北京了，就是连甘肃也出不去。<p>　　有一天，张帆下机后，走在欧式风情街，突然就下起了大雨。他庆幸自己拿了雨伞，撑开伞后，他漫步在街道上，看着四周的欧式建筑，还有那马可波罗的雕像，陷入了沉思：日子就这样一天一天度过了，我究竟在寻找什么？可能我永远也不知道答案了，或许我该去谈个恋爱，可是我同桌女生似乎对我不感兴趣！罢了，再说吧，反正时间多得是。<p>　　他漫步在欧式街上，又去鼓楼大街的新华书店里逛了一圈，决定看会书。他随便翻了翻连环画，又一抬头看见了这鳞次栉比的书架，这么多书籍，都是人类的宝贵财富，恐怕一辈子也看不完！接着，他从新华书店步行到东街，又回到了自己的家。<p>　　有一次，他去参加一个村里人的葬礼，顺便吃了点酒席，在回家的路上，春风时而呼啸，时而柔弱，就像一个淘气的孩子，拍打着张帆脸颊。张帆心想：他娘的这日子什么时候能过去？日子过得也太惬意了，真想一辈子都这样，可这风吹得人心里发慌，好像自己平静的生活之中总缺斤少两。算了，反正我也不知道答案，就这样吧，我希望有一天我能够明白生活的意义。都说我们是垮掉的一代，那我就当个网瘾少年吧，至少能在网吧里我能找到自己。<p>　　相比于大话和梦幻，张帆比较喜欢梦幻，大概是因为人物更Q吧，张帆觉得比较可爱。他选择了一个逍遥生，起了个飞雪寒江的名字，开始练级，这一练就到了40集。那时候，他很羡慕人家有瑞兽，觉得瑞兽放石头的技能太炫了，于是他去北俱芦洲抓了一只瑞兽，高兴极了，就开始组队炫耀。结果他们去大雁塔打怪的时候，张帆却遭到了对方的嘲笑，“你那瑞兽是野的，野的你也练？”一时间，张帆不知道回答什么，只好发了个表情敷衍过去。<p>　　后来，张帆才知道，练这些宠物需要从宝宝练起，成长好，后期战斗力也高，野的就不行了，各项技能都达不到。于是，张帆又寻思着怎么抓瑞兽宝宝，可他在北俱芦洲等了好久，也不见一个。着急的他，开始在洛阳城里乱逛，突然，他看到一个人在摆摊卖号，他就加了对方的QQ，对方只要一张60元的点卡，就把70级的剑侠客卖掉。张帆内心有点怀疑，真的可以买吗？算了，也就60，试试吧，骗子的话就当买个教训！于是，张帆给对方冲了60，没想到对方竟然特别讲义气，把那个号的信息发给了张帆。于是，张帆上号后一看，这个号简直是赚翻了，剑侠客70级，装备都有，还有瑞兽、凤凰、蛟龙、天兵宝宝，最牛逼的，这个天兵还带隐身技能！真是碉堡了，张帆高兴得睡不着觉，急急忙忙地跑同学跟前炫耀，结果玩游戏的同学都嗤之以鼻，“有什么了不起的，那些宝宝我也有！”从那时起，张帆认为自己是一个有财产的人了。<p>　　他开始没日没夜的练级，可以说是废寝忘食，学习成绩便一天天地下降了！而他却浑然不知，甚至不在乎，即便是掉到班级第20也无所谓。泡网吧的日子，张帆饿了就吃泡面，渴了就喝矿泉水，简直是乐在其中，不能自拔！可真正的危机，却不断地向他袭来！<p> 
	</div>
	
		   <div class="dirAd" style="text-align:center;">
						<!-- 桌面全文横幅2 -->
							<ins class="adsbygoogle"
								 style="display:inline-block;width:970px;height:90px"
								 data-ad-client="ca-pub-1218476774185134"
								 data-ad-region="content"
								 data-ad-slot="2684719341"></ins>
						  
						<script>
							(adsbygoogle = window.adsbygoogle || []).push({});
						</script>
			</div>
	
	<div class="postTitle"><a name="13255">第三章 车祸</a><br/><span class="topic">垮掉的一代GoldenWind全文阅读</span><span class="author">作者：王波er</span><span class="opt" title="把本章节加入书架，方便下次阅读" onclick="addShujia('13255','第三章 车祸'); return false;">加入书架</span></div>
	<div class="postContent">
	<p>　　张帆在网吧日复一日地练级，日子过得天昏地暗。他突然意识到，如果自己有一台电脑，平日里出去上学，就在家里登游戏挂机不就行了？这样，不但能练级，还能摆摊出售爆出来的装备。想到这里，张帆暗自庆幸，于是，他开始跟父母商量，谎称中学要举行计算机编程大赛，需要一台电脑，父母见他这样上进，就东拼西凑买了一台赛扬电脑。这可让张帆喜出望外了，他各种buff加身，一不小心，又拿了一次班级第一，大家都皆大欢喜。<p>　　这个周末，张帆摆摊挂机了，然后学会了抽烟。他点了一根，回忆起了从前。我出生在一个农民的家庭，境况最初很一般。因为是种田的，就算家里有很多亩地，在那个年代也赚不了什么钱，只能说在村里面还算可以的，但随着后来家人的一起努力，生活也逐渐变得富裕，但那也仅仅是农村里面的富裕。<p>　　关于童年的记忆，我只记得一些片段和特别难忘的场景。家里是典型的农家住宅，一个院子，几个房间，分别住着家人。那时家庭成员都尚在，就算家庭贫苦，一家也其乐融融，过着早耕晚归的生活。从出生那天起，我就好像一直在那个院子里玩，家人都出去务农了，没人照顾，就一个人在那里玩一些幼稚的游戏。<p>　　那时候天气一直很晴朗，似乎没经历过什么风雨吹打，或许是我只记得快乐的时光。一天，家人还是像往常一样出去忙碌了，我依然是一个人，就呆在家里玩。院子里种着一些树木很葱郁，下面是泥泞的路，我就在那里挖泥堆东西，还找来了很多道具放上去，表现出一个孩子心中的幻想。<p>　　后院里也很空旷，养着一条狗，几只猪，但我不太喜欢那些动物，我更喜欢天上的飞鸟。有时，我就一个人从梯子爬到房顶，仰望着湛蓝如洗的天空和自由飘浮的云朵，傻乎乎的只是望着。我并不知道未来会发生些什么，我的认知里只有家人亲戚，还有同村的伙伴，谁又能明白命中注定要发生的那些事情呢？几经变迁，心力交瘁，一心只为了最后的争取！<p>　　突然有一天，爷爷对我说：“你该上学了吧！”<p>　　我望着走在路上的那些学生说：“好吧，既然伙伴们都去上学了，那我也去！”爷爷高兴地领着我去报名，从此我开始读学前班。一个班里有几十个学生，都是我的伙伴，从不认识到认识，这也算是最初的缘分吧，其实关于对缘分的解读，这是个好例子。<p>　　我很喜欢学校生活，每天都从家高兴地去学校里上课，日子就有了规律。<p>　　并没有什么刻骨铭心的记忆，那时候是没有爱情烦恼的，就算在现在看来，那是很孤单的岁月，可又有什么关系呢？我们还是依然的非常快乐，因为这是属于我们的世界。<p>　　每天上课时的学习、讨论、做题，课下的玩耍、嬉闹、卫生，从未间断地持续，这就是我的童年我的生活。每当回忆起这些来，我总是非常感慨，为什么美丽的时光总要过去？而我的归属是哪里？苦并寻找着，痛并快乐着。<p>　　那后门的果园，多么的芬芳迷人，令人陶醉，洋溢着自然的魅力。爷爷在这里种植了很多果树，有杏树、桃树、梨树，每当它们开起花朵，那是何等的美丽，每当它们结出果实，那是何等的诱人。果园的脚下遍地是清香的泥土嫩绿的小草绽放的野花，经常会看见蝴蝶和蜜蜂的身影穿梭其中，饶有风趣啊，不但安静而且富有动态。<p>　　我在无忧无虑中度过了小学二年级，还过了一次不错的生日。那天，伙伴们都带着礼物来我家做客，奶奶买来了一个蛋糕，大家都吃得舒服，玩得开心，真是一次难忘的聚会啊！可是随之而来的就是无尽的分别，因为爷爷在村里享有名誉，而那时的乡镇企业又如雨后春笋，他去一家企业当了书记，所以家人也都跟着进城居住了。<p>　　“别了，我的朋友。别了，我的学校。别了，那院子里的一抹紫色葡萄架和盛夏时到处绽放的花朵。”你们是我人生中最初的美好际遇和风景，我总是这样的感慨。<p>　　我当时就住在这里生活条件非常艰辛但我总算是盼来了黎明的曙光，我的生活非常辛苦，因为家里穷，加上我本人比较弱小，我经常被一伙人欺负。为首的那个高个子，他经常嘲笑我，动不动就找各种理由来欺负我。有一次，我过生日，收到了很多礼物。在我们回家的路上，胖虎就和一个尖嘴猴腮的同伙拦住了我。<p>　　“你丫的干嘛去？东西都通通的交出来。”<p>　　“凭什么给你啊，这是我过生日朋友送的，为什么要交给你？”我很不服气。<p>　　“你不服气是吧，兄弟们，给我揍！”随着胖虎一声令下，我就被几个人围住了，他们一直在揍我！我的朋友都四散逃跑了，从那个时候起，我的心里面就有了浓重的阴影。我发誓，如果有我出头的一天，我一定也要欺负别人！<p>　　本来，我和胖虎的交锋持续了一段时间。我也因此被揍得不成人样，但后来，我们全家拔营去了城里。当我坐上小车走的那一刻，我看到了胖虎，他带领着兄弟们，站在桥那里，久久不肯离去。他望着我离去的身影，表现出了异常的愤怒！这不但是因为以后不能揍我了，而且是因为我终于离开了农村，去了城里，他又不服气了。我和胖虎的故事就告一段落了，我想着，他也不会带着他那一帮兄弟们来城里揍我！况且，我现在初中了，为了安全，投靠了班里的龙哥，我和龙哥关系非常好，因为我是学习委员，经常利用权力让龙哥不写作业，不交作业，龙哥也处处罩着我，真是好兄弟！好基友！胖虎就差远了，他从来就知道欺负人，一点也不体谅别人的内心。算了，不去想这些事情了，张帆的思绪回到了当前。<p>　　电脑屏幕上弹出了摆摊卖出东西的字眼，“哎呀，我这刚研制的金疮药总算卖出去了，这下又发了一笔，在游戏中自给自足就是好，如果生活中也能这样该多好呢？”张帆吸了一口说。算了，胖虎那帮人，就永远从我的世界中滚出去好了，他们算什么，连电脑都没见过，一帮乡巴佬！想到这里，张帆嬉皮咧嘴地笑了起来。结果，这一场景，正好被父母看见了，他们没有看到张帆在学习，反而用电脑玩游戏，还抽烟。结果，父母和张帆吵了起来，场面一度很尴尬！<p>　　“你……你给我滚出这个家！”父亲严厉地说道。他们已经争吵了半小时了，张帆实在吵不过，于是赌气说：“滚就滚，谁怕谁！”然后，他就一把拿上了外套，刚走出卧室门，心里一促，赶紧又回到电脑前，把游戏退了，还好还好。游戏可以不玩，但号一直要平安，这可是我的财产。<p>　　于是，张帆赌气在大街上乱逛，他骑了辆单车，速度70迈，心情是自由自在，刚才的怨气早就烟消云散了。整当他哼唱着速度70迈，心情是自由自在，随风奔跑自由是方向的时候，自行车一转弯，和摩托车碰上了。突然间，张帆的眼前一片黑暗，他被送到了医院，天知道他受伤有多严重！他昏迷了一阵子，睁开眼睛一看，亲朋好友都在，还有好兄弟龙哥，结果他刚起身，就看到了一个熟悉的身影，胖虎捧着花站在旁边，一下子簇拥过来了，于是他又昏倒了。脑海中还想起了胖虎高大狰狞的面孔，这哥们又胖了，活像一个猪八戒。<p>　　冥冥中，张帆不断地看见，《西游记》里的师徒四人，唐僧、孙悟空、猪八戒、沙和尚围着他的病床，不断地在对他说：“你醒了，手术很成功！”然后，张帆就又一次昏睡过去，这样的情况持续了无数遍，张帆快要疯了，可怜的孩子，他受到了惊吓！胖虎给他的心理阴影太强大了！<p>　　<p> 
	</div>
	
		   <div class="dirAd" style="text-align:center;">
						<!-- 桌面全文横幅2 -->
							<ins class="adsbygoogle"
								 style="display:inline-block;width:970px;height:90px"
								 data-ad-client="ca-pub-1218476774185134"
								 data-ad-region="content"
								 data-ad-slot="2684719341"></ins>
						  
						<script>
							(adsbygoogle = window.adsbygoogle || []).push({});
						</script>
			</div>
	<div class="page"  id="page"><span class="CurrentPage">1</span><a  href="/t/137365/2/">2</a><a  href="/t/137365/3/">3</a><a  href="/t/137365/4/">4</a><a  href="/t/137365/5/">5</a><a  href="/t/137365/6/">6</a><a  href="/t/137365/7/">7</a><a  href="/t/137365/8/">8</a><a  href="/t/137365/9/">9</a><a  href="/t/137365/10/">10</a><a  href="/t/137365/11/">11</a><a  href="/t/137365/12/">12</a>&nbsp;<a  href="/t/137365/2/">下一页</a>&nbsp;<a  href="/t/137365/43/">末页</a></div>
	  </div>
	  <div class="sao_mobile_box">
	  <div class="sao_box">
		<a class="sao_btn01" href="javascript:;" onclick="javascript:ShowQRCode(137365,0,1);">扫码</a>
		<div class="ma_box" style="display:none;" id="divQRCode">   
				<img style="border: 0" height="104" width="104" src="//img.uukanshu.com/static/www/images/index_qrcode.png">
		 <span>扫一扫，去手机上看</span>
		</div>
	  </div>
	</div>  
	<div  class="foot">
		<div class="xiaoshuo_shuoming" style="text-align:left">
	作者王波er所写的《垮掉的一代GoldenWind》为转载作品，垮掉的一代GoldenWind最新章节由网友发布，UU看书提供垮掉的一代GoldenWind全文阅读。<br />
	①如果您发现本小说垮掉的一代GoldenWind最新章节，而UU看书没有更新，请联系我们更新，您的热心是对网站最大的支持。<br />
	②书友如发现垮掉的一代GoldenWind内容有与法律抵触之处，请向本站举报，我们将马上处理。<br />
	③本小说垮掉的一代GoldenWind仅代表作者个人的观点，与UU看书的立场无关。<br />
	④如果您对垮掉的一代GoldenWind作品内容、版权等方面有质疑，或对本站有意见建议请发邮件给管理员，我们将第一时间作出相应处理。
	</div>
		<div class="copyright" style="text-align:center; margin-top:20px;">
		 <br/>Copyright &copy; <script type="text/javascript">myDate = new Date() ;myYear = myDate.getFullYear ();document.write(myYear);</script>  <a href="/">UU看书</a> All Rights Reserved.<br/>联系邮箱：uukanshu#gmail.com(#换成@) </div>   
	</div>
	
	<script>
		(function (i, s, o, g, r, a, m) {
			i['GoogleAnalyticsObject'] = r; i[r] = i[r] || function () {
				(i[r].q = i[r].q || []).push(arguments)
			}, i[r].l = 1 * new Date(); a = s.createElement(o),
	  m = s.getElementsByTagName(o)[0]; a.async = 1; a.src = g; m.parentNode.insertBefore(a, m)
		})(window, document, 'script', 'https://www.google-analytics.com/analytics.js', 'ga');
	
		ga('create', 'UA-59009863-7', 'auto');
		ga('send', 'pageview');
	
	</script>
	<!-- Go to www.addthis.com/dashboard to customize your tools --> <script type="text/javascript" src="//s7.addthis.com/js/300/addthis_widget.js#pubid=ra-593b35baefce2ecd"></script> 
	<script language="javascript">
		function addShujia(sid,stitle) {
			var tid = "137365";
			var sid = sid||"0";
			var st = stitle;
			$("#addsj").text("加入中...")
			$.get("/addshujia.ashx?rnd=" + Math.random(), { tid: tid, sid: sid, sTopic: st }, function (data) {
				switch (data) {
					case "-1":
						Common.showMsgBox("参数不合法，如有问题请和网站联系");
						break;
					case "0":
						$("#addsj").text("加入书架")
						User.showLogin();
						break;
					case "1":
						$("#addsj").text("已加书架")
						Common.showMsgBox("加入书架成功");
						break;
					case "2":
						$("#addsj").text("已加书架")
						Common.showMsgBox("请不要重复添加！");
						break;
					case "3":
						Common.showMsgBox("书架中最多只能存放200本小说<br/><br/>请删除书架中不再阅读的小说");
						break;
				}
			})
			return false;
		}
		$(function () {
			$("#closeMulu").click(function () {
				if ($("#mulu").is(":visible")) {
					$("#mulu").hide();
					$(this).text("显示目录");
				}
				else {
					$("#mulu").show();
					$(this).text("隐藏目录");
				}
				return false;
			});
		})
	
		function addNotice() {
			var tid = "137365";
			$("#addNotice").text("加入中...")
			$.get("/notice.ashx?rnd=" + Math.random(), { tid: tid }, function (data) {
				switch (data.status) {
					case -1:
						Common.showMsgBox("参数不合法，如有问题请和网站联系");
						break;
					case 0:
						$("#addNotice").text("更新提醒")
						User.showLogin();
						break;
					case 3:
						$("#addNotice").text("更新提醒")
						Common.showMsgBox("加入更新提醒成功。<br/>小说更新后会发送邮件到您的邮箱：" + data.email + "<br/><div style='font-size:12px; margin-top:10px'>邮箱不正确？<a href='/user/modifyuser.aspx' target='_blank'>点击修改</a></div>");
						break;
					case 2:
						$("#addNotice").text("更新提醒")
						Common.showMsgBox("已加入更新提醒。<br/>小说更新后会发送邮件到您的邮箱：" + data.email + "<br/><div style='font-size:12px; margin-top:10px'>邮箱不正确？<a href='/user/modifyuser.aspx' target='_blank'>点击修改</a></div>");
						break;
					case 1:
						$("#addNotice").text("更新提醒")
						Common.showMsgBox("最多只能添加30本书到更新提醒。<br/><div style='font-size:12px; margin-top:10px'>有不需要提醒的书？<a href='/user/notice.aspx' target='_blank'>点击删除</a></div>");
						break;
				}
			}, "json")
			return false;
		}
	
		function Report() {
			var tid = "137365";
			$.get("/report.ashx?rnd=" + Math.random(), { id: tid }, function (data) {
				switch (data) {
					case "1":
						Common.showMsgBox("提交成功，管理员将努力更新");
						break;
				}
			})
			return false;
		}
	</script>
	</body>
	</html> `), nil
}

```
