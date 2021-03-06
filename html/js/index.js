$(function () {
    function handlerLeftSecond() {
        var mobileNo = $("#phone").val();
        var leftSecond = getLeftSecond(mobileNo);
        if (leftSecond > 0) {
            $("#getValidCodeBtn").attr("disabled", true);
            $("#getValidCodeBtn").html(leftSecond + "s后重新发送");
        } else {
            $("#getValidCodeBtn").removeAttr("disabled");
            $("#getValidCodeBtn").html("获取验证码");
        }
    }

    //发送验证码
    $("#getValidCodeBtn").click(function () {
        var mobileNo = $("#phone").val();
        if (mobileNo == "") {
            showErrMsg("请填写手机号码");
            return
        }
        var leftSecond = getLeftSecond(mobileNo);
        if (leftSecond > 0) {
            setInterval(handlerLeftSecond, 1000);
        }
        var productId = $("#productId").val();
        $.get("/sendMsg", {"mobileNo": mobileNo, "productId": productId}, function (res) {
            if (res.isSucess) {
                showSucessMsg(res.message);
                setInterval(handlerLeftSecond, 1000);
            } else {
                if (res.message != undefined && res.message != "") {
                    showErrMsg(res.message);
                } else if (res.code != undefined && res.code != "") {
                    var alertpurchase = $("#alert-purchase");
                    if (res.code == "1") {
                        alertpurchase.removeClass("show")
                        var alertinfo2 = $("#alert-info2");//已经购买
                        alertinfo2.addClass("show");
                    }
                    else if (res.code == "2") {
                        alertpurchase.removeClass("show")
                        var alertinfo = $("#alert-info");//vip
                        alertinfo.addClass("show");
                    }
                }
            }
        });
    });

    function getLeftSecond(mobileNo) {
        var leftSecond = 0;
        var cookieValue = $.fn.cookie("code" + mobileNo);
        if (cookieValue != null && cookieValue != '') {
            cookieValue = cookieValue.substr(1);
            cookieValue = cookieValue.substr(0, cookieValue.length - 1);
            var cookieInt = Date.parse(cookieValue);
            leftSecond = (30 - parseInt((new Date().getTime() - cookieInt) / 1000));
            if (leftSecond < 0) {
                $.fn.cookie("code" + mobileNo, '');
            }
        }
        return leftSecond;
    }

    function get_cookie(Name) {
        var search = Name + "="//查询检索的值
        var returnvalue = "";//返回值
        if (document.cookie.length > 0) {
            sd = document.cookie.indexOf(search);
            if (sd != -1) {
                sd += search.length;
                end = document.cookie.indexOf(";", sd);
                if (end == -1)
                    end = document.cookie.length;
                //unescape() 函数可对通过 escape() 编码的字符串进行解码。
                returnvalue = unescape(document.cookie.substring(sd, end))
            }
        }
        return returnvalue;
    }

    $("#purchaseBtn").click(function () {
        if (!dataCheck()) return;
        var username = $("#username").val();
        var mobileNo = $("#phone").val();
        var code = $("#code").val();
        var productId = $("#productId").val();
        var channelcode = $("#channelcode").val();
        var alertpurchase = $("#alert-purchase")
        $.post("/addOrder", {
            "username": username,
            "mobileNo": mobileNo,
            "code": code,
            "productId": productId,
            "channelcode": channelcode
        }, function (res) {
            if (res.isSucess) {
                var intCode = parseInt(res.code);
                if (intCode < 0) {
                    showErrMsg(res.message);
                    return;
                }
                if (res.code == "1") {
                    alertpurchase.removeClass("show")
                    var alertinfo2 = $("#alert-info2");//已经购买
                    alertinfo2.addClass("show");
                }
                else if (res.code == "2") {
                    alertpurchase.removeClass("show")
                    var alertinfo = $("#alert-info");//vip
                    alertinfo.addClass("show");
                }
                else if (res.code == "3") {
                    //alert("临时单下单成功跳转支付");
                    //微信环境外拉起支付 微信环境则进行认证获取openid
                    window.location.href = res.redirect;
                }
                //var sucessAlert = $("#alert-success");
                //sucessAlert.addClass("show");
            } else {
                showErrMsg(res.message);
            }
        });
    })

    function dataCheck() {
        var username = $("#username").val();
        var mobileNo = $("#phone").val();
        var code = $("#code").val();
        if (username == "") {
            showErrMsg("姓名不得为空");
            return false;
        }
        if (mobileNo == "") {
            showErrMsg("手机号码不得为空");
            return false;
        }
        if (code == "") {
            showErrMsg("验证码不得为空");
            return false;
        }
        return true;
    }

    $("input").focus(function () {
        recoverMsg();
    })

    function recoverMsg() {
        var warnMsg = $("#warnMsg");
        warnMsg.hide();
    }

    function showSucessMsg(msg) {
        var warnMsg = $("#warnMsg");
        warnMsg.removeClass();
        warnMsg.addClass("tips-sucess");
        warnMsg.addClass("tips-info");
        warnMsg.html(msg);
        warnMsg.show();
    }

    function showErrMsg(msg) {
        var warnMsg = $("#warnMsg");
        warnMsg.removeClass();
        warnMsg.addClass("tips-err");
        warnMsg.addClass("tips-info");
        warnMsg.html(msg);
        warnMsg.show();
    }
})