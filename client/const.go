package client

const defaultIPAddress = "0.0.0.0"

const (
	TOPIC_LOGGING_IRIS         = "topads_logger_event_sa"
	TOPIC_LOGGING_IRIS_DISPLAY = "topads_logger_event_su"
)

const (
	IMPRESSION_BUCKET                        = "impression_v2"
	CLICK_BUCKET                             = "click_v2"
	NON_UNIQ_IMPRESSION_BUCKET               = "non_unique_impression"
	NON_UNIQ_CLICK_BUCKET                    = "non_unique_click"
	CONVERSION_PRODUCT_BUCKET                = "conversion_product_v2"
	CONVERSION_SHOP_BUCKET                   = "conversion_shop_v2"
	DISPLAY_BUCKET                           = "display_request_v2"
	CRUD_DASHBOARD_BUCKET                    = "crud_dashboard"
	INSERT_DEPOSIT_BUCKET                    = "insert_deposit_v2"
	AUTO_CAMPAIGN_BUCKET                     = "auto_campaign"
	WISHLIST_BUCKET                          = "wishlist"
	CART_BUCKET                              = "cart"
	CHECKOUT_DEPOSIT_BUCKET                  = "checkout_deposit"
	AUTO_TOPUP_BUCKET                        = "auto_topup"
	SHOP_EVENT_SCHEDULE_BUCKET               = "shop_event_schedule_v2"
	AUTO_ADS_PERFORMANCE_BUCKET              = "auto_ads_performance"
	AUTO_ADS_QUALITY_BUCKET                  = "auto_ads_quality"
	ADS_FEEDBACK_BUCKET                      = "ads_feedback"
	TOPADS_DAILY_BUDGET_HIT                  = "daily_budget_hit"
	GROUP_RECOMMENDATION_BUCKET              = "group_recommendation"
	AUTO_ADS_HISTORY_BUCKET                  = "auto_ads_history"
	AFFILIATE_CPA                            = "order_attempt_tracker"
	PUSH_TO_IRIS_ATC_BUCKET                  = "addtocart"
	PUSH_TO_IRIS_CHECKOUT_BUCKET             = "order_checkout"
	EXTERNAL_CAMPAIGN_MAPPING_HISTORY_BUCKET = "external_campaign_mapping_history"
	TOPADS_USED_BUCKET                       = "used"
	TOPADS_USED_NEW_BUCKET                   = "used_new"
	INACTIVE_ALL_ADS                         = "inactivate_all_ads"
	ADS_ACTIVATED                            = "ads_activated"
	CREDIT_ELIGIBLE                          = "free_credit_eligible"
	CREDIT_CLAIMED                           = "free_credit_claimed"
	PURCHASE_START                           = "purchase_start"
	PURCHASE_FINISH                          = "purchase_finished"
	PURCHASE_WAITING                         = "purchase_waiting"
	SHOP_ADMIN_ACCESS                        = "shop_admin_access"
	CAMPAIGN                                 = "campaign"
	CAMPAIGN_ACTION                          = "campaign_action"
	HEADLINE_IMPRESSION_BUCKET               = "headline_impression"
	HOMEBANNER_AD_DETAIL_BUCKET              = "homebanner_ad_detail"
	// PUSH_TO_IRIS_CONVERTION_CPA const define bucket name for conversion cpa in iris
	PUSH_TO_IRIS_CONVERTION_CPA              = "cpa_freetrial_activation"
	AUTO_ADS_ECPC                            = "auto_ads_ecpc"
	EVENT_HISTORY                            = "event_history"
	USER_EVENT_HISTORY                       = "user_event_history"
	CHANNEL_EVENT_HISTORY                    = "channel_event_history"
	DORMANT_USER                             = "dormant_user"
	FIRST_CONVERSION                         = "first_conversion"
	FIRST_IMP                                = "first_imp"
	AD_SCORE_RECALCULATION_LOGGER            = "score_recalculation_logger"
	ADS_REVIEW_HISTORY                       = "ads_review_history"
	HEADLINE_USED                            = "headline_used"
	HEADLINE_DORMANT_USER                    = "headline_dormant_user"
	CPA_CAMPAIGN                             = "cpa_campaign"
	SHOP_GROUP_NEW_INSIGHT                   = "shop_group_new_insight"
	PRODUCT_RECOMMENDATION_INSIGHT           = "product_recommendation_insight"
	SHOP_WHITELIST_FEATURE                   = "shop_whitelist_feature"
	AUTO_ADS_KEYWORD_LENGTH_FILTERING_BUCKET = "auto_ads_keyword_length_filtering"
	TOPADS_LOW_CREDIT                        = "low_credit"
	TOPADS_NEW_INSIGHT                       = "new_insight"
	TOPADS_TOP_UP                            = "top_up"
	TOPADS_AUTOBID_ADS                       = "autobid_ads"
	AUTO_ADS_DAILY_BUDGET_HIT                = "auto_ads_daily_budget_hit"
	AUTO_ADS_ACTIVATION                      = "auto_ads_activation"
	DOWNLOAD_STATS_REPORT                    = "download_statistics_report"
	TOPADS_PRODUCT_ACQUISITION               = "product_acquisition"
	ADGROUP_BID_INSIGHT_CAMPAIGN             = "adgroup_bid_insight_campaign"
	DISPLAY_BROWSE_REQUEST                   = "display_browse_request"
	AUTO_BID_ACTIVATION                      = "auto_bid_activation"
	AUTO_ADS_DAILY_BUDGET_UPDATED            = "auto_ads_daily_budget_updated"
)

const (
	LogTypeImpression                     = 1
	LogTypeClick                          = 2
	LogTypeConversionProduct              = 3
	LogTypeConversionShop                 = 4
	LogTypeDisplayRequest                 = 5
	LogTypeDashboardEvent                 = 6
	LogTypeInsertDeposit                  = 7
	LogTypeWishlist                       = 8
	LogTypeAddToCart                      = 9
	LogTypeCheckoutDeposit                = 10
	LogTypeAutoTopup                      = 11
	LogTypeShopEventSchedule              = 12
	LogTypeAutoCampaign                   = 13
	LogTypeAutoAdsCalculatePerformance    = 14
	LogTypeAdsFeedback                    = 15
	LogTypeDailyBudgetHit                 = 16
	LogTypeGroupRecommendation            = 17
	LogTypeAutoAdsHistory                 = 18
	LogTypeExternalCampaignMappingHistory = 19
	LogTypeTopAdsUsed                     = 20
	LogTypeInactiveAllAds                 = 21
	LogAdsActivated                       = 22
	LogTypeCreditEligible                 = 23
	LogTypeCreditClaimed                  = 24
	LogTypePurchaseStart                  = 25
	LogTypePurchaseFinished               = 26
	LogTypePurchaseWaiting                = 27
	LogTypeCampaign                       = 28
	LogTypeShopAdminAccess                = 29
	LogTypeCampaignAction                 = 30
	LogTypeHeadlineImpression             = 31
	LogTypeHomebannerAdDetail             = 32
	LogTypeDormantUser                    = 33
	LogTypeAdsReviewHistory               = 36
	LogTypeFirstConversion                = 37
	LogTypeFirstImp                       = 38
	LogTypeAdRecalculationLogger          = 39
	LogTypeHeadlineUsed                   = 40
	LogTypeHeadlineDormantUser            = 41
	LogTypeCPACampaign                    = 42
	LogTypeShopGroupNewInsight            = 43
	LogTypeNonUniqImpression              = 44
	LogTypeNonUniqClick                   = 45
	LogTypeAutoAdsKeywordLengthFiltering  = 46
	LogTypeShopWhitelistFeature           = 47
	LogTypeTopadsLowCredit                = 48
	LogTypeTopadsNewInsight               = 49
	LogTypeAutoAdsQuality                 = 50
	LogTypeAutoAdsECPC                    = 51
	LogTypePushtoIrisATC                  = 52
	LogTypeAffiliateCPA                   = 53
	LogTypePushtoIrisCheckout             = 54
	LogTypePushToIrisConvertCPA           = 55
	LogTypePotentialProduct               = 56
	LogTypeUserEventHistory               = 57
	LogTypeEventHistory                   = 58
	LogTypeTopadsTopup                    = 59
	LogTypeAutoBidAds                     = 60
	LogTypeAutoAdsDailyBudgetHit          = 61
	LogTypeAutoAdsActivation              = 62
	LogTypeDownloadStatsReport            = 63
	LogTypeProductAcquisition             = 64
	LogTypeAdGroupBidInsightCampaign      = 65
	LogTypeDisplayBrowseRequest           = 66
	LogTypeAutoBidActivation              = 67
	LogTypeTopAdsUsedNew                  = 68
	LogTypeAutoAdsDailyBudgetUpdated      = 69
)

const (
	// Event Type (second level of event, divided log type into more specific event)
	EventTypeView                      = 0
	EventTypeClick                     = 1
	EventTypeConversionProduct         = 2
	EventTypeConversionProductIndirect = 3
	EventTypeConversionShop            = 4
	EventTypeProductDisplayRequest     = 5
	EventTypeShopDisplayRequest        = 6
	EventTypeCPMDisplayRequest         = 7
	EventTypeConversionShopIndirect    = 8
	EventTypeCreateAd                  = 9
	EventTypeCreateGroup               = 10
	EventTypeCreateKeyword             = 11
	EventTypeEditAd                    = 12
	EventTypeEditGroup                 = 13
	EventTypeEditKeyword               = 14
	EventTypeDeleteAd                  = 15
	EventTypeDeleteGroup               = 16
	EventTypeDeleteKeyword             = 17
	EventTypeSwitchOnAd                = 18
	EventTypeSwitchOnGroup             = 19
	EventTypeSwitchOnKeyword           = 20
	EventTypeSwitchOffAd               = 21
	EventTypeSwitchOffGroup            = 22
	EventTypeSwitchOffKeyword          = 23
	EventTypeMoveAd                    = 24
	EventTypeMoveKeyword               = 25
	EventTypeBidKeyword                = 26
	EventTypeTypeKeyword               = 27
	EventTypeWishlistDirect            = 28
	EventTypeCartDirect                = 29
	EventTypeCheckoutDeposit           = 30
	EventTypeInsertDeposit             = 31
	EventTypeInsertDepositIntools      = 32
	EventTypeSwitchOnAutoTopup         = 33
	EventTypeSwitchOffAutoTopup        = 34
	EventTypeApproveBanner             = 35
	EventTypeMerchantBroadcastUsage    = 36
	EventTypeInsightBidGroup           = 37
	EventTypeToggleOnAutoAds           = 38
	EventTypeToggleOffAutoAds          = 39
	EventTypeInsightBidAds             = 40
	EventTypeEditDailyBudgetGroup      = 41
	EventTypeEditDailyBudgetAutoAds    = 42
	EventTypeProductCPADisplayRequest  = 43
	EventTypeOverallConversion         = 44
	EventTypeBannerDisplayRequest      = 48
	EventTypeNonUniqView               = 49
	EventTypeNonUniqClick              = 50
	EventTypeCreateKeywordLevelAd      = 51
	EventTypeEditKeywordLevelAd        = 52
	EventTypeDeleteKeywordLevelAd      = 53
	EventTypePublishAd                 = 54
	EventTypeUnpublishAd               = 55
	EventTypeDisableAd                 = 56
	EventTypeEnableAd                  = 57
	EventTypeUnknown                   = -1
)

const (
	// Data Type (lowest level of event, some second level has same event but different data to be recorded
	DataTypeAd      = 1
	DataTypeGroup   = 2
	DataTypeKeyword = 3

	// auto topup
	DataTypeAutoTopup = 4

	// autoads
	DataTypeAutoAds = 5
)

const (
	// Action Type (action used for tracking AddToCart and Wishlist convertion)
	ActionTypeAdd     = "ADD"
	ActionTypeUpdate  = "UPDATE"
	ActionTypeDelete  = "DELETE"
	ActionTypeDefault = "XACTION"
)

const (
	//Error codes for Auto Topup error logging
	SALDO_LOW                = 1
	DAILY_TOPUP_LIMIT_EXCEED = 2
	PROCESS_AUTO_TOPUP_ERR   = 3
)

const (
	// Container
	CONTAINER_ADS       = "ads"
	CONTAINER_TOPADS    = "topads"
	CONTAINER_AFFILIATE = "content_affiliate"
	// CONTAINER_CONTENT used by content iris
	CONTAINER_CONTENT = "content"
)

const (
	// source refers to event from which app like customer, mitra, seller. For now, by default we are setting source=customer.
	SOURCE_CUSTOMER = "customer"
)

var CONTAINER_LIST = [...]string{
	CONTAINER_ADS,
	CONTAINER_TOPADS,
	CONTAINER_AFFILIATE,
	CONTAINER_CONTENT,
}

var EVENT_LIST = map[string]string{
	IMPRESSION_BUCKET:                        CONTAINER_ADS,
	CLICK_BUCKET:                             CONTAINER_ADS,
	NON_UNIQ_IMPRESSION_BUCKET:               CONTAINER_ADS,
	NON_UNIQ_CLICK_BUCKET:                    CONTAINER_ADS,
	CONVERSION_PRODUCT_BUCKET:                CONTAINER_ADS,
	CONVERSION_SHOP_BUCKET:                   CONTAINER_ADS,
	DISPLAY_BUCKET:                           CONTAINER_ADS,
	CRUD_DASHBOARD_BUCKET:                    CONTAINER_ADS,
	INSERT_DEPOSIT_BUCKET:                    CONTAINER_ADS,
	AUTO_CAMPAIGN_BUCKET:                     CONTAINER_ADS,
	WISHLIST_BUCKET:                          CONTAINER_ADS,
	CART_BUCKET:                              CONTAINER_ADS,
	CHECKOUT_DEPOSIT_BUCKET:                  CONTAINER_ADS,
	AUTO_TOPUP_BUCKET:                        CONTAINER_ADS,
	SHOP_EVENT_SCHEDULE_BUCKET:               CONTAINER_ADS,
	AUTO_ADS_PERFORMANCE_BUCKET:              CONTAINER_TOPADS,
	ADS_FEEDBACK_BUCKET:                      CONTAINER_TOPADS,
	TOPADS_DAILY_BUDGET_HIT:                  CONTAINER_TOPADS,
	GROUP_RECOMMENDATION_BUCKET:              CONTAINER_TOPADS,
	AUTO_ADS_HISTORY_BUCKET:                  CONTAINER_TOPADS,
	AFFILIATE_CPA:                            CONTAINER_AFFILIATE,
	PUSH_TO_IRIS_ATC_BUCKET:                  CONTAINER_AFFILIATE,
	PUSH_TO_IRIS_CHECKOUT_BUCKET:             CONTAINER_AFFILIATE,
	EXTERNAL_CAMPAIGN_MAPPING_HISTORY_BUCKET: CONTAINER_TOPADS,
	TOPADS_USED_BUCKET:                       CONTAINER_TOPADS,
	TOPADS_USED_NEW_BUCKET:                   CONTAINER_TOPADS,
	INACTIVE_ALL_ADS:                         CONTAINER_TOPADS,
	ADS_ACTIVATED:                            CONTAINER_TOPADS,
	CREDIT_ELIGIBLE:                          CONTAINER_TOPADS,
	CREDIT_CLAIMED:                           CONTAINER_TOPADS,
	PURCHASE_START:                           CONTAINER_TOPADS,
	PURCHASE_FINISH:                          CONTAINER_TOPADS,
	PURCHASE_WAITING:                         CONTAINER_TOPADS,
	SHOP_ADMIN_ACCESS:                        CONTAINER_TOPADS,
	CAMPAIGN:                                 CONTAINER_TOPADS,
	CAMPAIGN_ACTION:                          CONTAINER_TOPADS,
	PUSH_TO_IRIS_CONVERTION_CPA:              CONTAINER_CONTENT,
	AUTO_ADS_ECPC:                            CONTAINER_TOPADS,
	EVENT_HISTORY:                            CONTAINER_TOPADS,
	USER_EVENT_HISTORY:                       CONTAINER_TOPADS,
	CHANNEL_EVENT_HISTORY:                    CONTAINER_TOPADS,
	HEADLINE_IMPRESSION_BUCKET:               CONTAINER_TOPADS,
	HOMEBANNER_AD_DETAIL_BUCKET:              CONTAINER_TOPADS,
	AUTO_ADS_QUALITY_BUCKET:                  CONTAINER_TOPADS,
	DORMANT_USER:                             CONTAINER_TOPADS,
	ADS_REVIEW_HISTORY:                       CONTAINER_TOPADS,
	FIRST_CONVERSION:                         CONTAINER_TOPADS,
	FIRST_IMP:                                CONTAINER_TOPADS,
	AD_SCORE_RECALCULATION_LOGGER:            CONTAINER_TOPADS,
	HEADLINE_USED:                            CONTAINER_TOPADS,
	HEADLINE_DORMANT_USER:                    CONTAINER_TOPADS,
	CPA_CAMPAIGN:                             CONTAINER_TOPADS,
	SHOP_GROUP_NEW_INSIGHT:                   CONTAINER_TOPADS,
	SHOP_WHITELIST_FEATURE:                   CONTAINER_TOPADS,
	AUTO_ADS_KEYWORD_LENGTH_FILTERING_BUCKET: CONTAINER_TOPADS,
	TOPADS_LOW_CREDIT:                        CONTAINER_TOPADS,
	TOPADS_NEW_INSIGHT:                       CONTAINER_TOPADS,
	PRODUCT_RECOMMENDATION_INSIGHT:           CONTAINER_TOPADS,
	TOPADS_TOP_UP:                            CONTAINER_TOPADS,
	TOPADS_AUTOBID_ADS:                       CONTAINER_TOPADS,
	AUTO_ADS_DAILY_BUDGET_HIT:                CONTAINER_TOPADS,
	AUTO_ADS_ACTIVATION:                      CONTAINER_TOPADS,
	DOWNLOAD_STATS_REPORT:                    CONTAINER_TOPADS,
	TOPADS_PRODUCT_ACQUISITION:               CONTAINER_TOPADS,
	ADGROUP_BID_INSIGHT_CAMPAIGN:             CONTAINER_TOPADS,
	DISPLAY_BROWSE_REQUEST:                   CONTAINER_ADS,
	AUTO_BID_ACTIVATION:                      CONTAINER_TOPADS,
	AUTO_ADS_DAILY_BUDGET_UPDATED:            CONTAINER_TOPADS,
}
