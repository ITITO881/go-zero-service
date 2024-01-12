package desc

import "errors"

var srcFlag map[int64]string
var subType map[string]string

func init() {
	srcFlag = make(map[int64]string)
	subType = make(map[string]string)
	// 1. ITO 微信小程序，上传文件类型码
	srcFlag[WEB_MP_ITO_USER_AVATAR] = "文件源：微信小程序，用户头像"
	srcFlag[WEB_MP_ITO_SWIPER_HOME] = "文件源：微信小程序，轮播图，主页"
	srcFlag[WEB_MP_ITO_SWIPER_PRODUCT] = "文件源：微信小程序，轮播图，产品页"
	srcFlag[WEB_MP_ITO_CATEGORY_HOME] = "文件源：微信小程序，宫格列表，主页"
	srcFlag[WEB_MP_ITO_CATEGORY_PRODUCT] = "文件源：微信小程序，宫格列表，产品页"
	srcFlag[WEB_MP_ITO_CATEGORY_COUPON] = "文件源：微信小程序，宫格列表，卡券ICON"
	srcFlag[WEB_MP_ITO_HOT_GRID_HOME] = "文件源：微信小程序，热点，主副图，主页"
	srcFlag[WEB_MP_ITO_HOT_GRID_PRODUCT] = "文件源：微信小程序，热点，主副图，产品页"
	srcFlag[WEB_MP_ITO_HOT_GRID_CLUB] = "文件源：微信小程序，热点，主副图，社群页"
	srcFlag[WEB_MP_ITO_BAR_CAT_HOME] = "文件源：微信小程序，分类列表，主页"
	srcFlag[WEB_MP_ITO_BAR_CAT_Prod_CAT] = "文件源：微信小程序，分类列表，卡券"
	srcFlag[WEB_MP_ITO_BAR_CAT_HOT_PROPOSE] = "文件源：微信小程序，分类列表，热点推荐"
	srcFlag[WEB_MP_ITO_GUESS_HOME] = "文件源：微信小程序，猜你喜欢，产品信息组图，通用组件"
	srcFlag[WEB_MP_ITO_GUESS_HOT_PROPOSE] = "文件源：微信小程序，混入产品信息，热点推荐"

	srcFlag[WEB_MP_ITO_PRODUCT_MANAGER_SPU] = "文件源：微信小程序，产品管理，SPU"
	srcFlag[WEB_MP_ITO_PRODUCT_MANAGER_SKU] = "文件源：微信小程序，产品管理，SKU"

	srcFlag[WEB_MP_ITO_BAR_CAT_MIXIN_HOT_PROPOSE] = "文件源：微信小程序，分类混入列表，热点推荐"
	srcFlag[WEB_MP_ITO_BAR_CAT_MIXIN_PAGE] = "文件源：微信小程序，分类混入列表，页面素材"
	srcFlag[WEB_MP_ITO_BAR_CAT_MIXIN_SWIPER] = "文件源：微信小程序，分类混入列表，轮播素材"
	srcFlag[WEB_MP_ITO_DUAL_MIXIN_PROD] = "文件源：微信小程序，二级分类复合列表，产品页"
	srcFlag[WEB_MP_ITO_DUAL_BAR_PROD] = "文件源：微信小程序，二级分类，产品页"

	srcFlag[WEB_MP_ITO_COUPON_CARES] = "文件源: 售后卡素材"
	srcFlag[WEB_MP_ITO_COUPON_NUT] = "文件源: 售后卡素材"
	srcFlag[WEB_MP_ITO_COUPON_VIP] = "文件源: 售后卡素材"
	srcFlag[WEB_MP_ITO_COUPON_VOUCHER] = "文件源: 售后卡素材"

	// 2. ITO Web应用，上传文件类型码
	srcFlag[CUSTOMRIZED_PIC_ITO] = "文件源：ITO用户产品定制图片"

	//	3. ITO 文件上传 子类型
	subType[SPU_MAJOR_PIC] = "文件源：SPU 主图，主图需唯一，且Spu字段必须有值"
	subType[SPU_MINOR_PICS] = "文件源：SPU 副图，副图可以有多个，但必须先有SPU主图"
	subType[SPU_DETAIL_PICS] = "文件源：SPU 详情图，详情图可以有多个，但必须先有SPU主图"
	subType[SKU_MAJOR_PIC] = "文件源：SKU 主图，SKU主图唯一，且Spu，Sku两个字段必须有值"
	subType[SKU_THUMB_PICS] = "文件源：SKU 缩略图，SKU缩略图可以有多个，但必选先有SKU主图"
	subType[SUB_TAB_PANE] = "文件源：二级分类列表，子Tab组件的TAB_PANE"
	subType[SUB_TAB_LIST] = "文件源：二级分类列表，子List图形列表"
}

func MapSrcFlag(src int64) (string, error) {
	if flag, ok := srcFlag[src]; ok {
		return flag, nil
	} else {
		return "", errors.New("source flag 未定义")
	}
}

func MapSubType(sub string) (string, error) {
	if flag, ok := subType[sub]; ok {
		return flag, nil
	} else {
		return "", errors.New("subType 未定义")
	}
}
