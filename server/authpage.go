package airdisk

import (
	"net/http"
	"fmt"
)

var indexHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<title>测试</title>
</head>
<body>
	<form action="/page" method="post">
		用户名：<br>
		<input name="username" type="text"><br>
		请输入文本：<br>
		<textarea name="usertext"></textarea><br>
		<input type="submit" value="提交">
	</form>
</body>
</html>`

var wechatPage = `
<!DOCTYPE HTML>
<html>
<head lang="zh-CN">
	<meta charset="UTF-8">
    <title>微信连Wi-Fi</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="format-detection" content="telephone=no">
	<script type="text/javascript">
		/**
		 * 微信连Wi-Fi协议3.1供运营商portal呼起微信浏览器使用
		 */
		var loadIframe = null;
		var noResponse = null;
		var callUpTimestamp = 0;

		function putNoResponse(ev){
			 clearTimeout(noResponse);
		}

		 function errorJump()
		 {
			 var now = new Date().getTime();
			 if((now - callUpTimestamp) > 4*1000){
				 return;
			 }
			 alert('该浏览器不支持自动跳转微信请手动打开微信\n如果已跳转请忽略此提示');
		 }

		 myHandler = function(error) {
			 errorJump();
		 };

		 function createIframe(){
			 var iframe = document.createElement("iframe");
		     iframe.style.cssText = "display:none;width:0px;height:0px;";
		     document.body.appendChild(iframe);
		     loadIframe = iframe;
		 }
		//注册回调函数
		function jsonpCallback(result){
			if(result && result.success){
			    alert('WeChat will call up : ' + result.success + '  data:' + result.data);
			    var ua=navigator.userAgent;
				if (ua.indexOf("iPhone") != -1 ||ua.indexOf("iPod")!=-1||ua.indexOf("iPad") != -1) {   //iPhone
					document.location = result.data;
				}else{

					if('false'=='true'){
						alert('[强制]该浏览器不支持自动跳转微信请手动打开微信\n如果已跳转请忽略此提示');
						return;
					}

				    createIframe();
				    callUpTimestamp = new Date().getTime();
				    loadIframe.src=result.data;
					noResponse = setTimeout(function(){
						errorJump();
			      	},3000);
				}
			}else if(result && !result.success){
				alert(result.data);
			}
		}

		function Wechat_GotoRedirect(appId, extend, timestamp, sign, shopId, authUrl, mac, ssid, bssid){

			//将回调函数名称带到服务器端
			var url = "https://wifi.weixin.qq.com/operator/callWechatBrowser.xhtml?appId=" + appId
																				+ "&extend=" + extend
																				+ "&timestamp=" + timestamp
																				+ "&sign=" + sign;

			//如果sign后面的参数有值，则是新3.1发起的流程
			if(authUrl && shopId){


				url = "https://wifi.weixin.qq.com/operator/callWechat.xhtml?appId=" + appId
																				+ "&extend=" + extend
																				+ "&timestamp=" + timestamp
																				+ "&sign=" + sign
																				+ "&shopId=" + shopId
																				+ "&authUrl=" + encodeURIComponent(authUrl)
																				+ "&mac=" + mac
																				+ "&ssid=" + ssid
																				+ "&bssid=" + bssid;

			}

			//通过dom操作创建script节点实现异步请求
			var script = document.createElement('script');
			script.setAttribute('src', url);
			document.getElementsByTagName('head')[0].appendChild(script);
		}
	</script>
    <link rel="stylesheet" href="static/css/style-simple-follow.css"/>
</head>
<body class="mod-simple-follow">
<div class="mod-simple-follow-page">
    <div class="mod-simple-follow-page__banner">
        <img class="mod-simple-follow-page__banner-bg" src="static/images/background.jpg" alt=""/>
        <div class="mod-simple-follow-page__img-shadow"></div>
        <div class="mod-simple-follow-page__logo">
            <img class="mod-simple-follow-page__logo-img" src="static/images/t.weixin.logo.png" alt=""/>
            <p class="mod-simple-follow-page__logo-name"></p>
            <p class="mod-simple-follow-page__logo-welcome">欢迎您</p>
        </div>
    </div>
    <div class="mod-simple-follow-page__attention">
        <p class="mod-simple-follow-page__attention-txt">欢迎使用微信连Wi-Fi</p>
        <a class="mod-simple-follow-page__attention-btn" onclick="callWechatBrowser()">一键打开微信连Wi-Fi</a>
    </div>
</div>
</body>
<script type="text/javascript">
	function callWechatBrowser(){
		//Wechat_GotoRedirect('wxd86056fe58061437', 'demoNew', '1486903282625', '7de58b9792874463538fac948a23116a', '7743860', 'http://wifi.weixin.qq.com/assistant/wifigw/auth.xhtml?httpCode=200', 'aa:aa:aa:aa:aa:aa', '2099', 'ff:ff:ff:ff:ff:ff');
		alert("This is Demo");
		window.location="http://192.168.1.1:3990/logon";
	}
</script>

<script type="text/javascript">
	document.addEventListener('visibilitychange', putNoResponse, false);
</script>
</html>`

func index(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintf(w, indexHTML)
	fmt.Fprintf(w, wechatPage)
}

