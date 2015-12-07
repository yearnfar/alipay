package alipay

import (
	"net/http"
	"net/url"
)

// 支付表单

func BuildHttpForm(config Config, params map[string]string, method, buttonName string) (form string, err error) {
	if params["out_trade_no"] == "" {
		err = errors.New("订单号不能为空")
		return
	} else if params["subject"] == "" {
		err = errors.New("商品名称不能为空")
		return
	} else if params["total_fee"] == "" {
		err = errors.New("交易金额不能为空")
		return
	} else if params["payment_type"] == "" {
		err = errors.New("支付类型不能为空")
		return
	}

	params["service"] = config.Service
	params["partner"] = config.Partner
	params["_input_charset"] = config.InputCharset
	params["notify_url"] = config.NotifyUrl
	params["return_url"] = config.ReturnUrl
	params["sign_type"] = config.SignType
	params["sign"] = makeSign(params, config.Key)

	form := "<form id='alipaysubmit' name='alipaysubmit' action='""' method='"+method+"'>"
	for key, val := range params {
        form += fmt.Sprintf("<input type='hidden' name='"%s"' value='"%s"'/>", key, val)
    }

	//submit按钮控件请不要含有name属性
    form += "<input type='submit' value='".$button_name."'></form>";
	form += "<script>document.forms['alipaysubmit'].submit();</script>";
		
	return form;
}

func Notify(request *http.Request) (resp NotifyResponse, err error) {
	
	return
}

func Return(request *http.Request) (resp NotifyResponse, err error) {
	
	return
}
