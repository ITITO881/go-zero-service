package desc

/**(�?位代表渠�?后三位代表子渠道)**/

const (
	WEB_MP_ITO_USER_AVATAR         int64 = 100110
	WEB_MP_ITO_SWIPER_HOME         int64 = 100120
	WEB_MP_ITO_SWIPER_PRODUCT      int64 = 100121
	WEB_MP_ITO_CATEGORY_HOME       int64 = 100130
	WEB_MP_ITO_CATEGORY_PRODUCT    int64 = 100131
	WEB_MP_ITO_CATEGORY_COUPON     int64 = 100132
	WEB_MP_ITO_HOT_GRID_HOME       int64 = 100140
	WEB_MP_ITO_HOT_GRID_PRODUCT    int64 = 100141
	WEB_MP_ITO_HOT_GRID_CLUB       int64 = 100142
	WEB_MP_ITO_BAR_CAT_HOME        int64 = 100150
	WEB_MP_ITO_BAR_CAT_COUPON      int64 = 100151
	WEB_MP_ITO_BAR_CAT_HOT_PROPOSE int64 = 100153
	WEB_MP_ITO_GUESS_HOME          int64 = 100160
	WEB_MP_ITO_GUESS_HOT_PROPOSE   int64 = 100161

	WEB_MP_ITO_PRODUCT_MANAGER_SPU int64 = 100170
	WEB_MP_ITO_PRODUCT_MANAGER_SKU int64 = 100171

	WEB_MP_ITO_DUAL_BAR_PROD             int64 = 100180
	WEB_MP_ITO_BAR_CAT_MIXIN_HOT_PROPOSE int64 = 100190
	WEB_MP_ITO_BAR_CAT_MIXIN_PAGE        int64 = 100191
	WEB_MP_ITO_DUAL_MIXIN_PROD           int64 = 100200
)

const (
	/** 1011 开头编码为 SPU图相关子�?**/
	SPU_MAJOR_PIC   string = "spuMajor"
	SPU_MINOR_PICS  string = "spuMinor"
	SPU_DETAIL_PICS string = "spuDetail"
	/**  1012 开头编码为 SKU图相关子�?**/
	SKU_MAJOR_PIC  string = "skuMajor"
	SKU_THUMB_PICS string = "skuThumb"
	/** 100180 二级分类 用于区分二级分类 **/
	SUB_TAB_PANE string = "subTabPane"
	SUT_TAB_LIST string = "subTabList"
)

const CUSTOMRIZED_PIC_ITO int64 = 200001
