package desc

import "errors"

var srcFlag map[int64]string
var subType map[int64]string

func init() {
	srcFlag = make(map[int64]string)
	subType = make(map[int64]string)
	// 1. ITO 微信小程序，上传文件类型码
	srcFlag[WEB_MP_ITO_USER_AVATAR] = "文件源：微信小程序，用户头像"
	srcFlag[WEB_MP_ITO_SWIPER_HOME] = "文件源：微信小程序，轮播图，主页"
	srcFlag[WEB_MP_ITO_SWIPER_PRODUCT] = "文件源：微信小程序，轮播图，产品页"
	srcFlag[WEB_MP_ITO_CATEGORY_HOME] = "文件源：微信小程序，简单分类，主页"
	srcFlag[WEB_MP_ITO_CATEGORY_PRODUCT] = "文件源：微信小程序，简单分类，产品页"
	srcFlag[WEB_MP_ITO_HOT_GRID_HOME] = "文件源：微信小程序，热点，主副图，主页"
	srcFlag[WEB_MP_ITO_HOT_GRID_PRODUCT] = "文件源：微信小程序，热点，主副图，产品页"
	srcFlag[WEB_MP_ITO_HOT_GRID_CLUB] = "文件源：微信小程序，热点，主副图，社群页"
	srcFlag[WEB_MP_ITO_BAR_CAT_HOME] = "文件源：微信小程序，分类列表，主页"
	srcFlag[WEB_MP_ITO_BAR_CAT_COUPON] = "文件源：微信小程序，分类列表，卡券"
	srcFlag[WEB_MP_ITO_GUESS_HOME] = "文件源：微信小程序，猜你喜欢，产品信息组图，通用组件"
	srcFlag[WEB_MP_ITO_PRODUCT_MANAGER_SPU] = "文件源：微信小程序，产品管理，SPU"
	srcFlag[WEB_MP_ITO_PRODUCT_MANAGER_SKU] = "文件源：微信小程序，产品管理，SKU"

	// 2. ITO Web应用，上传文件类型码
	srcFlag[CUSTOMRIZED_PIC_ITO] = "文件源：ITO用户产品定制图片"

	//	3. ITO 文件上传 子类型
	subType[SPU_MAJOR_PIC] = "文件源：SPU 主图，主图需唯一，且Spu字段必须有值"
	subType[SPU_MINOR_PICS] = "文件源：SPU 副图，副图可以有多个，但必须先有SPU主图"
	subType[SPU_DETAIL_PICS] = "文件源：SPU 详情图，详情图可以有多个，但必须先有SPU主图"
	subType[SKU_MAJOR_PIC] = "文件源：SKU 主图，SKU主图唯一，且Spu，Sku两个字段必须有值"
	subType[SKU_THUMB_PICS] = "文件源：SKU 缩略图，SKU缩略图可以有多个，但必选先有SKU主图"
}

func MapSrcFlag(src int64) (string, error) {
	if flag, ok := srcFlag[src]; ok {
		return flag, nil
	} else {
		return "", errors.New("source flag 未定义")
	}
}

func MapSubType(sub int64) (string, error) {
	if flag, ok := subType[sub]; ok {
		return flag, nil
	} else {
		return "", errors.New("subType 未定义")
	}
}
