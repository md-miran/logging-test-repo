package client

import (
	"net"

	"github.com/nsqio/go-nsq"
)

// WrapperEvent is a wrapper struct to collect all events happened in sync server then send it through NSQ to be processed in async server
type (
	WrapperEvent struct {
		Id              WrapperId   `json:"id,omitempty"`
		CreateTime      string      `json:"create_time,omitempty"`
		ClientTimestamp string      `json:"client_timestamp,omitempty"`
		LogType         int         `json:"log_type,omitempty"`
		Data            interface{} `json:"data,omitempty"`
		Nonce           string      `json:"nonce,omitempty"`

		NSQMsg *nsq.Message `json:"-"`
		Topic  string       `json:"-"`
	}

	WrapperId struct {
		UserId    int64 `json:"user_id,omitempty"`
		AdId      int64 `json:"ad_id,omitempty"`
		ShopId    int64 `json:"shop_id,omitempty"`
		ProductId int64 `json:"product_id,omitempty"`
	}

	AggregatorEvent struct {
		AdData      AdClass
		ProductData ProductEvent
		ShopData    ShopEvent
		UserData    UserEvent
	}

	AdClass struct {
		AdId       int64
		ShopId     int64
		Type       int
		ItemId     int64
		Status     int
		PriceBid   float64
		PriceDaily float64
		AdTitle    string
		AdReview   int
		AdStart    string
		AdEnd      string
		AdRefKey   string
		CreateBy   int64
		CreateTime string
		UpdateBy   int64
		UpdateTime string
		StickerId  int
		GroupId    int64

		// Private Field - don't use directly
		StickerImage string
	}
)

// These types are detail level
type (
	FormulaEvent struct {
		CTRSearch      float64 `json:"ctrSearch,omitempty"`
		CTRBrowse      float64 `json:"ctrBrowse,omitempty"`
		ScoreSearch    float64 `json:"scoreSearch,omitempty"`
		ScoreBrowse    float64 `json:"scoreBrowse,omitempty"`
		KeywordScore   float64 `json:"keywordScore,omitempty"`
		RelevanceScore float64 `json:"relevanceScore,omitempty"`
	}

	AdEvent struct {
		Id            string         `json:"id,omitempty"`
		GoalId        string         `json:"goal_id,omitempty"`
		Type          string         `json:"type,omitempty"`
		Score         float64        `json:"score,omitempty"`
		PriceBid      float64        `json:"priceBid,omitempty"`
		AdSequence    string         `json:"adSequence,omitempty"`
		HeadLine      HeadLineEvent  `json:"cpm,omitempty"`
		KeywordId     string         `json:"keywordId,omitempty"`
		KeywordTag    string         `json:"keywordTag,omitempty"`
		KeywordBid    float64        `json:"keywordBid,omitempty"`
		GroupId       string         `json:"groupId,omitempty"`
		StickerId     string         `json:"stickerId,omitempty"`
		AdLocation    string         `json:"adLocation,omitempty"`
		Formula       FormulaEvent   `json:"formula,omitempty"`
		Status        string         `json:"status,omitempty"`
		DailyBudget   float64        `json:"dailyBudget,omitempty"`
		ScheduleStart string         `json:"scheduleStart,omitempty"`
		ScheduleEnd   string         `json:"scheduleEnd,omitempty"`
		Statistic     StatisticEvent `json:"statistic,omitempty"`
	}

	StatisticEvent struct {
		TtlImpression  int64   `json:"ttlImpression,omitempty"`
		TtlClick       int64   `json:"ttlClick,omitempty"`
		TtlCTR         float64 `json:"ttlCTR,omitempty"`
		TtlSpent       float64 `json:"ttlSpent,omitempty"`
		PredictedOrder float64 `json:"predictedOrder,omitempty"`
		DateStart      string  `json:"dateStart,omitempty"`
		DateEnd        string  `json:"dateEnd,omitempty"`
	}

	HeadLineEvent struct {
		UseTemplateId string   `json:"useTemplateId,omitempty"`
		ReqTemplateId []string `json:"reqTemplateId,omitempty"`
		AviTemplateId []string `json:"aviTemplateId,omitempty"`
		ServiceType   string   `json:"serviceType,omitempty"`
	}

	ShopEvent struct {
		Id              string   `json:"id,omitempty"`
		Name            string   `json:"name,omitempty"`
		Location        string   `json:"location,omitempty"`
		ShippingId      []string `json:"shippingId,omitempty"`
		Category        []string `json:"category,omitempty"`
		OwnerId         string   `json:"ownerId,omitempty"`
		Rate            float64  `json:"rate,omitempty"`
		IsGm            string   `json:"isGm,omitempty"`
		IsOfficial      string   `json:"isOfficial,omitempty"`
		Reputation      int      `json:"reputation,omitempty"`
		TADeposit       float64  `json:"taDeposit,omitempty"`
		Status          string   `json:"status,omitempty"`
		MerchantVoucher []string `json:"merchant_voucher_type,omitempty"`
	}

	ShopEventV2 struct {
		Id              string   `json:"id,omitempty"`
		UserId          string   `json:"user_id,omitempty"`
		Name            string   `json:"name,omitempty"`
		IsGmFlag        string   `json:"is_gm_flag,omitempty"`
		IsOsFlag        string   `json:"is_os_flag,omitempty"`
		DistrictName    string   `json:"district_name,omitempty"`
		Category        []string `json:"category,omitempty"`
		Reputation      int      `json:"reputation_score,omitempty"`
		TADeposit       float64  `json:"topads_deposit_amount,omitempty"`
		Status          string   `json:"status,omitempty"`
		MerchantVoucher []string `json:"merchant_voucher_type,omitempty"`
	}

	UserEvent struct {
		Id        string  `json:"id,omitempty"`
		Name      string  `json:"name,omitempty"`
		Email     string  `json:"email,omitempty"`
		Gender    string  `json:"gender,omitempty"`
		Birthdate string  `json:"birthdate,omitempty"`
		Msisdn    string  `json:"msisdn,omitempty"`
		Saldo     float64 `json:"saldo,omitempty"`
	}

	PublisherEvent struct {
		UnitId   string `json:"unitId,omitempty"`
		DomainId string `json:"domainId,omitempty"`
		Id       string `json:"id,omitempty"`
		SizeId   string `json:"sizeId,omitempty"`
	}

	ProductEvent struct {
		Id         string   `json:"id,omitempty"`
		Name       string   `json:"name,omitempty"`
		DepId      []string `json:"depId,omitempty"`
		Price      int64    `json:"price,omitempty"`
		ParentId   string   `json:"parentId,omitempty"`
		Wholesale  string   `json:"wholesale,omitempty"`
		Cashback   string   `json:"cashback,omitempty"`
		Freereturn string   `json:"freereturn,omitempty"`
		Preorder   string   `json:"preorder,omitempty"`
		Rate       int      `json:"rate,omitempty"`
		Label      string   `json:"label,omitempty"`
		Score      float64  `json:"score,omitempty"`
		Condition  int      `json:"condition,omitempty"`
	}

	NetworkEvent struct {
		IpAddress  string        `json:"ipAddress,omitempty"`
		DeviceType string        `json:"deviceType,omitempty"`
		UserAgent  string        `json:"userAgent,omitempty"`
		Location   LocationEvent `json:"location,omitempty"`
		Source     string        `json:"source,omitempty"`
		SessionId  string        `json:"sessionId,omitempty"`
		Origin     string        `json:"origin,omitempty"`
		RefURL     string        `json:"refURL,omitempty"`
	}

	LocationEvent struct {
		City      string `json:"city,omitempty"`
		Province  string `json:"province,omitempty"`
		Country   string `json:"country,omitempty"`
		Latitude  string `json:"latitude,omitempty"`
		Longitude string `json:"longitude,omitempty"`
	}

	CommonEvent struct {
		Ad            AdEvent      `json:"ad,omitempty"`
		Shop          ShopEvent    `json:"shop,omitempty"`
		Network       NetworkEvent `json:"network,omitempty"`
		Nonce         string       `json:"nonce,omitempty"`
		EventType     string       `json:"eventType,omitempty"`
		CreateTime    string       `json:"createTime,omitempty"`
		User          UserEvent    `json:"user,omitempty"`
		Filter        FilterEvent  `json:"filter,omitempty"`
		Tag           TagEvent     `json:"tag,omitempty"`
		NCandidateAds int          `json:"nCandidateAds,omitempty"`
		NumberOfAds   int          `json:"numberOfAds,omitempty"`
		RefNonce      string       `json:"refNonce,omitempty"`
		Fraud         string       `json:"fraud,omitempty"`
	}

	FilterEvent struct {
		Keywords        string   `json:"keywords,omitempty"`
		NumberAdsReq    int      `json:"numberAdsReq,omitempty"`
		Page            string   `json:"page,omitempty"`
		Location        []string `json:"location,omitempty"`
		ShippingId      []string `json:"shippingId,omitempty"`
		Rating          []string `json:"rating,omitempty"`
		HotlistId       string   `json:"hotlistId,omitempty"`
		CatalogId       string   `json:"catalogId,omitempty"`
		DepartmentId    []string `json:"departmentId,omitempty"`
		PriceMin        float64  `json:"priceMin,omitempty"`
		PriceMax        float64  `json:"priceMax,omitempty"`
		Wholesale       string   `json:"wholesale,omitempty"`
		IsGm            string   `json:"isGm,omitempty"`
		IsOfficial      string   `json:"isOfficial,omitempty"`
		Preorder        string   `json:"preorder,omitempty"`
		Freereturn      string   `json:"freereturn,omitempty"`
		Cashback        string   `json:"cashback,omitempty"`
		BrandId         string   `json:"brandId,omitempty"`
		Condition       string   `json:"condition,omitempty"`
		Sorting         string   `json:"sorting,omitempty"`
		Variant         []string `json:"variant,omitempty"`
		KreasiLokal     string   `json:"kreasiLokal,omitempty"`
		MerchantVoucher []string `json:"merchant_voucher_type,omitempty"`
		AnnotationID    []string `json:"annotation_id,omitempty"`
	}

	TagEvent struct {
		AlgorithmTags          []string `json:"algorithmTags,omitempty"`
		IsSearchNotFound       string   `json:"isSearchNotFound,omitempty"`
		UserTarget             string   `json:"userTarget,omitempty"`
		PublisherRemarketing   string   `json:"publisherRemarketing,omitempty"`
		SearchQueryRemarketing string   `json:"searchQueryRemarketing,omitempty"`
		AbTest                 string   `json:"abTest,omitempty"`
		KeywordSuggest         string   `json:"keywordSuggest,omitempty"`
		NCandidateSuggest      int      `json:"nCandidateSuggest,omitempty"`
		OriginShopId           string   `json:"originShopId,omitempty"`
		OriginProductId        string   `json:"originProductId,omitempty"`
		OriginDepId            string   `json:"origindepid,omitempty"`
		PostAlgorithm          []string `json:"postAlgorithm,omitempty"`
		QueryTags              []string `json:"queryTags,omitempty"`
		TemporaryTags          []string `json:"temporaryTags,omitempty"`
		RollenceFlow           string   `json:"rollence_flow,omitempty"`
		RollenceExperimentName string   `json:"rollence_experiment_name,omitempty"`
		RollenceVariantName    string   `json:"rollence_variant_name,omitempty"`
		CampaignID             string   `json:"campaign_id,omitempty"`
		UtmSource              string   `json:"utm_source,omitempty"`
		UtmMedium              string   `json:"utm_medium,omitempty"`
	}

	OrderEvent struct {
		Id            string  `json:"id,omitempty"`
		Quantity      int     `json:"quantity,omitempty"`
		SubtotalPrice float64 `json:"subtotalPrice,omitempty"`
	}

	PaymentEvent struct {
		Id      string `json:"id,omitempty"`
		Gateway string `json:"gateway,omitempty"`
		Method  string `json:"method,omitempty"`
	}

	PromoEvent struct {
		Id             string  `json:"id,omitempty"`
		Code           string  `json:"code,omitempty"`
		DiscountAmount float64 `json:"discountAmount,omitempty"`
		ExtraAmount    float64 `json:"extraAmount,omitempty"`
		Source         string  `json:"source,omitempty"`
	}

	DepositEvent struct {
		Id       string  `json:"id,omitempty"`
		Amount   float64 `json:"amount,omitempty"`
		Quantity int     `json:"quantity,omitempty"`
		Type     string  `json:"type,omitempty"`
	}
)

// These types are high level of events which consist of some detail level events
//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=PushtoIrisATC -target=src/logger/type.go -destination=files/var/schema/addtocart.avsc
type PushtoIrisATC struct {
	RecID                     string `json:"rec_id"`
	DeviceName                string `json:"device_name"`
	CartID                    int64  `json:"cart_id"`
	ProductID                 int64  `json:"product_id"`
	ProductName               string `json:"product_name"`
	Quantity                  int    `json:"order_quantity"`
	Subtotal                  int    `json:"subtotal"`
	ShopID                    int64  `json:"shop_id"`
	ShopName                  string `json:"shopname"`
	LastClickAffiliateID      int64  `json:"last_click_affiliate_id"`
	LastClickAdID             int64  `json:"last_click_ad_id"`
	LastClickCuratedProductID int64  `json:"last_click_curated_product_id"`
	UserID                    int64  `json:"user_id"`
	ActionTrackerID           string `json:"action_tracker_id"`
	IsIndirect                bool   `json:"isindirect"`
	UniqueTrackerId           string `json:"unique_tracker_id"`
	ClientTimestamp           string `json:"client_timestamp"`
	SessionID                 string `json:"session_id"`
	IpAddress                 string `json:"ip_address"`
	CreateTime                string `json:"create_time"`
}

// unused, will be removed
type InsertDepositEvent struct {
	EventType      int     `json:"eventType,omitempty"`
	ShopId         string  `json:"shopId,omitempty"`
	UserId         string  `json:"userId,omitempty"`
	CreateTime     string  `json:"createTime,omitempty"`
	Nonce          string  `json:"nonce,omitempty"`
	PrevDeposit    float64 `json:"prevDeposit,omitempty"`
	DepositAmount  float64 `json:"depositAmount,omitempty"`
	DepositType    string  `json:"depositType,omitempty"`
	OrderId        string  `json:"orderId,omitempty"`
	PaymentMethod  string  `json:"paymentMethod,omitempty"`
	PaymentGateway string  `json:"paymentGateway,omitempty"`
	Device         string  `json:"device,omitempty"`
	DepositBonus   float64 `json:"depositBonus,omitempty"`
	DepositVoucher float64 `json:"depositVoucher,omitempty"`
	LastBuy        string  `json:"lastBuy,omitempty"`
	IsGm           string  `json:"isGm,omitempty"`
	IsOfficial     string  `json:"isOfficial,omitempty"`
	Source         string  `json:"source,omitempty"`
	PromoSource    string  `json:"promoSource,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AdsFeedbackEvent -target=src/logger/type.go -destination=files/var/schema/ads_feedback.avsc
type AdsFeedbackEvent struct {
	ShopId           string `json:"shop_id,omitempty"`
	ProductId        string `json:"product_id,omitempty"`
	ProductName      string `json:"product_name,omitempty"`
	AdId             string `json:"ad_id,omitempty"`
	UserId           string `json:"user_id,omitempty"`
	ReasonCode       string `json:"reason_code,omitempty"`
	AdKeywordTagName string `json:"ad_keyword_tag_name,omitempty"`
	SourceDesc       string `json:"source_desc,omitempty"`
	ProductClickUrl  string `json:"product_click_url,omitempty"`
	Level3Id         string `json:"level3_id,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	RefNonce         string `json:"refNonce,omitempty"`
	CreateTime       string `json:"create_time,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsHistory -target=src/logger/type.go -destination=files/var/schema/auto_ads_history.avsc
type AutoAdsHistory struct {
	ShopId                int64  `json:"shop_id,omitempty"`
	Status                int    `json:"status_code,omitempty"`
	IsSuccess             bool   `json:"is_success_flag,omitempty"`
	Notes                 string `json:"notes_text,omitempty"`
	CreateTime            string `json:"create_time,omitempty"`
	Nonce                 string `json:"nonce,omitempty"`
	AdExperimentMappingID int64  `json:"ad_experiment_mapping_id,omitempty"`
	SubStatusCode         int    `json:"sub_status_code,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsCalculatePerformanceEvent -target=src/logger/type.go -destination=files/var/schema/auto_ads_performance.avsc
type AutoAdsCalculatePerformanceEvent struct {
	Shop             ShopEventV2 `json:"shop,omitempty"`
	Nonce            string      `json:"nonce,omitempty"`
	Level3Id         int64       `json:"level3_id,omitempty"`
	Ads              []AdEvent   `json:"ads,omitempty"`
	CreateTime       string      `json:"createTime,omitempty"`
	SelectedLowAdsId []int64     `json:"selected_low_ads_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsQuality -target=src/logger/type.go -destination=files/var/schema/auto_ads_quality.avsc
type AutoAdsQuality struct {
	ShopID              int64   `json:"shop_id,omitempty"`
	AdID                int64   `json:"ad_id,omitempty"`
	AdQuality           string  `json:"ad_quality,omitempty"`
	ProductQuality      string  `json:"product_quality,omitempty"`
	AdQualityScore      float64 `json:"ad_quality_score,omitempty"`
	ProductQualityScore float64 `json:"product_quality_score,omitempty"`
	Source              string  `json:"source,omitempty"`
	CreateTime          string  `json:"create_time,omitempty"`
	ExperimentID        int64   `json:"experiment_id,omitempty"`
	ProductID           int64   `json:"product_id,omitempty"`
	ShopBudget          float64 `json:"shop_budget,omitempty"`
	ProductBudget       float64 `json:"product_budget,omitempty"`
	Priority            string  `json:"priority,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoCampaignEvent -target=src/logger/type.go -destination=files/var/schema/auto_campaign.avsc
type AutoCampaignEvent struct {
	CreateTime           string    `json:"createTime,omitempty"`
	Shop                 ShopEvent `json:"shop,omitempty"`
	User                 UserEvent `json:"user,omitempty"`
	FreeCreditCounter    int       `json:"fcCounter,omitempty"`
	FreeCreditReason     int       `json:"fcReason,omitempty"`
	FreeCreditBenefit    float64   `json:"fcBenefit,omitempty"`
	FreeCreditStartClaim string    `json:"fcStartClaim,omitempty"`
	FreeCreditEndClaim   string    `json:"fcEndClaim,omitempty"`
	FreeCreditExpireDate string    `json:"fcExpireDate,omitempty"`
	FreeCreditRemaining  float64   `json:"fcRemaining,omitempty"`
	Nonce                string    `json:"nonce,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoTopupEvent -target=src/logger/type.go -destination=files/var/schema/auto_topup.avsc
type AutoTopupEvent struct {
	User              UserEvent `json:"user,omitempty"`
	Shop              ShopEvent `json:"shop,omitempty"`
	Nonce             string    `json:"nonce,omitempty"`
	Attempt           int       `json:"attempt,omitempty"`
	NominalTopup      float64   `json:"nominalTopup,omitempty"`
	AvailSaldo        float64   `json:"availSaldo,omitempty"`
	AutoTopupCountDay int       `json:"autoTopupCountDay,omitempty"`
	FailureReasonCode int       `json:"failureReasonCode,omitempty"`
	CreateTime        string    `json:"createTime,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ConversionAddToCartEvent -target=src/logger/type.go -destination=files/var/schema/cart.avsc
type (
	ConversionAddToCartEvent struct {
		Common    CommonEvent    `json:"common"`
		Publisher PublisherEvent `json:"publisher"`
		Product   ProductEvent   `json:"product"`
		AddToCart AddToCartEvent `json:"addToCart"`
	}

	AddToCartEvent struct {
		Quantity   int    `json:"quantity"`
		ActionType string `json:"actionType"`
		Platform   string `json:"platform"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=CheckoutDepositEvent -target=src/logger/type.go -destination=files/var/schema/checkout_deposit.avsc
type CheckoutDepositEvent struct {
	Network        NetworkEvent `json:"network,omitempty"`
	User           UserEvent    `json:"user,omitempty"`
	Shop           ShopEvent    `json:"shop,omitempty"`
	Promo          PromoEvent   `json:"promo,omitempty"`
	Deposit        DepositEvent `json:"deposit,omitempty"`
	InstantPayment string       `json:"instantPayment,omitempty"`
	Nonce          string       `json:"nonce,omitempty"`
	EventType      string       `json:"eventType,omitempty"`
	CreateTime     string       `json:"createTime,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ClickEvent -target=src/logger/type.go -destination=files/var/schema/click_v2.avsc
//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ClickEvent -target=src/logger/type.go -destination=files/var/schema/non_unique_click.avsc
type (
	ClickEvent struct {
		Common                    ClickCommonEvent `json:"common,omitempty"`
		Publisher                 PublisherEvent   `json:"publisher,omitempty"`
		Product                   ProductEvent     `json:"product,omitempty"`
		AdCpc                     float64          `json:"adCpc,omitempty"`
		PubCpc                    float64          `json:"pubCpc,omitempty"`
		DiscountAmt               float64          `json:"charge_discount_amount,omitempty"`
		ClickType                 string           `json:"click_type_name,omitempty"`
		IsAutoAds                 bool             `json:"is_auto_ads_flag,omitempty"`
		TrackerId                 string           `json:"tracker_id,omitempty"`
		DeviceVersion             string           `json:"device_version,omitempty"`
		ClusterIdV2               int64            `json:"cluster_v2_id,omitempty"`
		CategorySuggestEvaluation []string         `json:"tag_category_suggest_evaluation,omitempty"`
		CategorySuggestAttribute  []string         `json:"tag_category_suggest_attribute,omitempty"`
		IsLoadTest                int              `json:"is_load_test,omitempty"` // https://tokopedia.atlassian.net/wiki/spaces/DE/pages/365592824/Data+Formatting+Guideline
		RequestType               string           `json:"request_type,omitempty"`
		IrisSessionId             string           `json:"iris_session_id,omitempty"`
	}

	ClickCommonEvent struct {
		Ad            ClickAdEvent       `json:"ad,omitempty"`
		Shop          ShopEvent          `json:"shop,omitempty"`
		Network       ClickNetworkEvent  `json:"network,omitempty"`
		Nonce         string             `json:"nonce,omitempty"`
		EventType     string             `json:"eventType,omitempty"`
		CreateTime    string             `json:"createTime,omitempty"`
		User          UserEvent          `json:"user,omitempty"`
		Filter        FilterEvent        `json:"filter,omitempty"`
		Tag           ClickTagEvent      `json:"tag,omitempty"`
		Carousel      ClickCarouselEvent `json:"carousel,omitempty"`
		NCandidateAds int                `json:"nCandidateAds,omitempty"`
		NumberOfAds   int                `json:"numberOfAds,omitempty"`
		RefNonce      string             `json:"refNonce,omitempty"`
		Fraud         string             `json:"fraud,omitempty"`
		ClickSource   string             `json:"click_from,omitempty"`
	}

	ClickTagEvent struct {
		AlgorithmTags               []string `json:"algorithmTags,omitempty"`
		IsSearchNotFound            string   `json:"isSearchNotFound,omitempty"`
		UserTarget                  string   `json:"userTarget,omitempty"`
		PublisherRemarketing        string   `json:"publisherRemarketing,omitempty"`
		SearchQueryRemarketing      string   `json:"searchQueryRemarketing,omitempty"`
		AbTest                      string   `json:"abTest,omitempty"`
		KeywordSuggest              string   `json:"keywordSuggest,omitempty"`
		NCandidateSuggest           int      `json:"nCandidateSuggest,omitempty"`
		OriginShopId                string   `json:"originShopId,omitempty"`
		OriginProductId             string   `json:"originProductId,omitempty"`
		OriginDepId                 string   `json:"origindepid,omitempty"`
		PostAlgorithm               []string `json:"postAlgorithm,omitempty"`
		QueryTags                   []string `json:"queryTags,omitempty"`
		TemporaryTags               []string `json:"temporaryTags,omitempty"`
		ProductTagName              []string `json:"product_tags_name,omitempty"`
		AlternateKeyword            string   `json:"alternate_keyword,omitempty"`
		IsIgnoredCategoryPrediction bool     `json:"is_ignored_category_prediction,omitempty"`
		RollenceFlow                string   `json:"rollence_flow,omitempty"`
		RollenceExperimentName      string   `json:"rollence_experiment_name,omitempty"`
		RollenceVariantName         string   `json:"rollence_variant_name,omitempty"`
		CampaignID                  string   `json:"campaign_id,omitempty"`
		UtmSource                   string   `json:"utm_source,omitempty"`
		UtmMedium                   string   `json:"utm_medium,omitempty"`
	}

	ClickNetworkEvent struct {
		IpAddress  string        `json:"ipAddress,omitempty"`
		DeviceType string        `json:"deviceType,omitempty"`
		UserAgent  string        `json:"userAgent,omitempty"`
		Location   LocationEvent `json:"location,omitempty"`
		Source     string        `json:"source,omitempty"`
		SessionId  string        `json:"sessionId,omitempty"`
		Origin     string        `json:"origin,omitempty"`
		RefURL     string        `json:"refURL,omitempty"`
	}

	// ClickCarouselEvent struct for carousel related information
	ClickCarouselEvent struct {
		Position int `json:"position,omitempty"`
	}

	ClickAdEvent struct {
		Id                  string         `json:"id,omitempty"`
		GoalId              string         `json:"goal_id,omitempty"`
		Type                string         `json:"type,omitempty"`
		Score               float64        `json:"score,omitempty"`
		PriceBid            float64        `json:"priceBid,omitempty"`
		AdSequence          string         `json:"adSequence,omitempty"`
		HeadLine            HeadLineEvent  `json:"cpm,omitempty"`
		KeywordId           string         `json:"keywordId,omitempty"`
		KeywordTag          string         `json:"keywordTag,omitempty"`
		KeywordBid          float64        `json:"keywordBid,omitempty"`
		KeywordCorrection   string         `json:"keywordCorrection,omitempty"`
		GroupId             string         `json:"groupId,omitempty"`
		StickerId           string         `json:"stickerId,omitempty"`
		AdLocation          string         `json:"adLocation,omitempty"`
		Formula             FormulaEvent   `json:"formula,omitempty"`
		Status              string         `json:"status,omitempty"`
		DailyBudget         float64        `json:"dailyBudget,omitempty"`
		ScheduleStart       string         `json:"scheduleStart,omitempty"`
		ScheduleEnd         string         `json:"scheduleEnd,omitempty"`
		Statistic           StatisticEvent `json:"statistic,omitempty"`
		SearchTermID        string         `json:"search_term_id,omitemtpy"`
		RTBBid              float64        `json:"rtb_bid,omitempty"`
		KeywordType         string         `json:"keyword_type,omitempty"`
		GroupStrategy       []int          `json:"group_strategy,omitempty"`
		RelevanceScore      float64        `json:"kt_relevance_score,omitempty"`
		K2KDistanceScore    float64        `json:"k2k_distance_score,omitempty"`
		FinalRelrankAdscore float64        `json:"final_relrank_adscore,omitempty"`
		QueryRelrankScore   float64        `json:"query_relrank_score,omitempty"`
		RankBestSeller      int64          `json:"rank_best_seller,omitempty"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ConversionProductEvent -target=src/logger/type.go -destination=files/var/schema/conversion_product_v2.avsc
type ConversionProductEvent struct {
	Common         CommonEvent    `json:"common,omitempty"`
	Publisher      PublisherEvent `json:"publisher,omitempty"`
	Product        ProductEvent   `json:"product,omitempty"`
	Order          OrderEvent     `json:"order,omitempty"`
	Payment        PaymentEvent   `json:"payment,omitempty"`
	ConversionType string         `json:"conversion_type"`
	TrackerId      string         `json:"tracker_id"`
	RequestType    string         `json:"request_type,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ConversionWishlistEvent -target=src/logger/type.go -destination=files/var/schema/wishlist.avsc
type (
	ConversionWishlistEvent struct {
		Common    CommonEvent    `json:"common"`
		Publisher PublisherEvent `json:"publisher"`
		Product   ProductEvent   `json:"product"`
		Wishlist  WishlistEvent  `json:"wishlist"`
	}

	WishlistEvent struct {
		WishlistId string `json:"wishlistId"`
		ActionType string `json:"actionType"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ConversionShopEvent -target=src/logger/type.go -destination=files/var/schema/conversion_shop_v2.avsc
type ConversionShopEvent struct {
	Common CommonEvent `json:"common,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=DashboardEvent -target=src/logger/type.go -destination=files/var/schema/crud_dashboard.avsc
type DashboardEvent struct {
	EventType                 int                     `json:"eventType,omitempty"`
	DataType                  int                     `json:"dataType,omitempty"`
	Nonce                     string                  `json:"nonce,omitempty"`
	Id                        string                  `json:"id,omitempty"`
	Type                      int                     `json:"type,omitempty"`
	ShopId                    string                  `json:"shopId,omitempty"`
	Status                    int                     `json:"status,omitempty"`
	PriceBid                  float64                 `json:"priceBid,omitempty"`
	CreateBy                  string                  `json:"createBy,omitempty"`
	CreateTime                string                  `json:"createTime,omitempty"`
	Device                    string                  `json:"device,omitempty"`
	Source                    string                  `json:"source,omitempty"`
	ReferrerNonce             string                  `json:"referrerNonce,omitempty"`
	Name                      string                  `json:"name,omitempty"`
	GroupId                   string                  `json:"groupId,omitempty"`
	ItemId                    string                  `json:"itemId,omitempty"`
	SuggestedBid              float64                 `json:"suggestedBid,omitempty"`
	SuggestionValue           int                     `json:"suggestionValue,omitempty"`
	SuggestionButton          int                     `json:"suggestionButton,omitempty"`
	PriceDaily                float64                 `json:"priceDaily,omitempty"`
	ScheduleStart             string                  `json:"scheduleStart,omitempty"`
	ScheduleEnd               string                  `json:"scheduleEnd,omitempty"`
	CustomBid                 int                     `json:"customBid,omitempty"`
	AdTotal                   int                     `json:"adTotal,omitempty"`
	KeywordTotal              int                     `json:"keywordTotal,omitempty"`
	PrevBid                   float64                 `json:"prevBid,omitempty"`
	Reason                    string                  `json:"reason,omitempty"`
	Channel                   string                  `json:"channel"`
	BidSettings               []BidSetting            `json:"bids,omitempty"`
	BidSettingsJSON           string                  `json:"bidSettings,omitempty"`
	SuggestionBidSettings     []SuggestionBidSettings `json:"suggestionBidSettingsData,omitempty"`
	SuggestionBidSettingsJSON string                  `json:"suggestionBidSettings,omitempty"`
}

type SuggestionBidSettings struct {
	BidType            int     `json:"bid_type"`
	SuggestionPriceBid float64 `json:"suggestion_price_bid"`
	SuggestionValue    int     `json:"suggestion_value"`
}

type BidSetting struct {
	BidType  int     `json:"bid_type"`
	PriceBid float64 `json:"price_bid"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AdGroupBidInsightCampaignEvent -target=src/logger/type.go -destination=files/var/schema/adgroup_bid_insight_campaign.avsc
type AdGroupBidInsightCampaignEvent struct {
	ShopID                   int64  `json:"shop_id,omitempty"`
	UserID                   int64  `json:"user_id,omitempty"`
	CountGroupWithRecom      string `json:"count_group_with_recom,omitempty"`
	PredictedTotalImpression int    `json:"predicted_total_impression,omitempty"`
	Timestamp                string `json:"timestamp,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=DailyBudgetHitEvent -target=src/logger/type.go -destination=files/var/schema/daily_budget_hit.avsc
type DailyBudgetHitEvent struct {
	ShopId        string  `json:"shop_id,omitempty"`
	AdIds         []int64 `json:"ad_ids,omitempty"`
	AdType        string  `json:"ad_type,omitempty"`
	GroupId       string  `json:"group_id,omitempty"`
	DailyBudget   float64 `json:"daily_budget,omitempty"`
	PriceBid      float64 `json:"price_bid,omitempty"`
	ScheduleStart string  `json:"schedule_start,omitempty"`
	ScheduleEnd   string  `json:"schedule_end,omitempty"`
	CreateTime    string  `json:"create_time,omitempty"`
	ProductID     int64   `json:"product_id,omitempty"`
	IsAutoAds     bool    `json:"is_auto_ads,omitempty"`
	BudgetScope   string  `json:"budget_scope,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=DisplayRequestEvent -target=src/logger/type.go -destination=files/var/schema/display_request_v2.avsc
type (
	DisplayRequestEvent struct {
		Network                        NetworkEvent              `json:"network,omitempty"`
		Filter                         DisplayRequestFilterEvent `json:"filter,omitempty"`
		Ads                            []AdEventDisplayRequest   `json:"ads,omitempty"`
		AdShown                        []string                  `json:"adShown,omitempty"`
		Tag                            DisplayTagEvent           `json:"tag,omitempty"`
		Nonce                          string                    `json:"nonce,omitempty"`
		EventType                      string                    `json:"eventType,omitempty"`
		CreateTime                     string                    `json:"createTime,omitempty"`
		NCandidateAds                  int                       `json:"nCandidateAds,omitempty"`
		NumberOfAds                    int                       `json:"numberOfAds,omitempty"`
		User                           UserEvent                 `json:"user,omitempty"`
		RefNonce                       string                    `json:"refNonce,omitempty"`
		AdsFromES                      []string                  `json:"list_ads,omitempty"`
		RemovedAds                     []AdsFilteredOut          `json:"list_ads_filter"`
		TrackerId                      string                    `json:"tracker_id,omitempty"`
		DeviceVersion                  string                    `json:"device_version,omitempty"`
		RequestType                    string                    `json:"request_type,omitempty"`
		AdsUnification                 string                    `json:"ads_unification,omitempty"`
		AuctionNParticipant            int                       `json:"auction_n_participant,omitempty"`
		AuctionLog                     string                    `json:"auction_log,omitempty"`
		GEResultCount                  int                       `json:"ge_result_count,omitempty"`
		GEThreshold                    int                       `json:"ge_threshold,omitempty"`
		GEAdCount                      int                       `json:"ge_ad_count,omitempty"`
		IrisSessionId                  string                    `json:"iris_session_id,omitempty"`
		PreRelevanceRankingKTAdsCount  int                       `json:"pre_relrank_kt_ads_count,omitempty"`
		PostRelevanceRankingKTAdsCount int                       `json:"post_relrank_kt_ads_count,omitempty"`
		CandidateAdsCount              string                    `json:"ncandidateads_by_algo,omitempty"`
	}

	DisplayTagEvent struct {
		AlgorithmTags               []string            `json:"algorithmTags,omitempty"`
		IsSearchNotFound            string              `json:"isSearchNotFound,omitempty"`
		UserTarget                  string              `json:"userTarget,omitempty"`
		PublisherRemarketing        string              `json:"publisherRemarketing,omitempty"`
		SearchQueryRemarketing      string              `json:"searchQueryRemarketing,omitempty"`
		AbTest                      string              `json:"abTest,omitempty"`
		KeywordCorrection           string              `json:"keywordCorrection,omitempty"`
		KeywordSuggest              string              `json:"keywordSuggest,omitempty"`
		NCandidateSuggest           int                 `json:"nCandidateSuggest,omitempty"`
		OriginShopId                string              `json:"originShopId,omitempty"`
		OriginProductId             string              `json:"originProductId,omitempty"`
		OriginDepId                 string              `json:"origindepid,omitempty"`
		PostAlgorithm               []string            `json:"postAlgorithm,omitempty"`
		QueryTags                   []string            `json:"queryTags,omitempty"`
		TemporaryTags               []string            `json:"temporaryTags,omitempty"`
		ProductTags                 []string            `json:"product_tags_name,omitempty"`
		SupplyTags                  []SupplyDetailEvent `json:"supply_tags,omitempty"`
		AlgoDetail                  []AlgoDetail        `json:"algo_detail,omitempty"`
		CategorySuggestEvaluation   []string            `json:"category_suggest_evaluation,omitempty"`
		CategorySuggestAttribute    []string            `json:"category_suggest_attribute,omitempty"`
		AlternateKeywords           []string            `json:"alternate_keywords,omitempty"`
		IsIgnoredCategoryPrediction bool                `json:"is_ignored_category_prediction,omitempty"`
		RollenceFlow                string              `json:"rollence_flow,omitempty"`
		RollenceExperimentName      string              `json:"rollence_experiment_name,omitempty"`
		RollenceVariantName         string              `json:"rollence_variant_name,omitempty"`
		CampaignID                  string              `json:"campaign_id,omitempty"`
		UtmSource                   string              `json:"utm_source,omitempty"`
		UtmMedium                   string              `json:"utm_medium,omitempty"`
	}

	SupplyDetailEvent struct {
		Algorithm       string `json:"algorithm,omitempty"`
		Keyword         string `json:"keyword,omitempty"`
		CategoryID      int    `json:"category_id,omitempty"`
		ClusterID       int64  `json:"cluster_id,omitempty"`
		SupplyCriteria  int64  `json:"supply_criteria,omitempty"`
		SupplyAlgorithm int64  `json:"supply_algorithm,omitempty"`
	}

	AlgoDetail struct {
		Algorithm string `json:"algorithm,omitempty"`
		AdType    int    `json:"ad_type,omitempty"`
		Total     int64  `json:"total,omitempty"`
	}

	DisplayRequestFilterEvent struct {
		Keywords        string   `json:"keywords,omitempty"`
		NumberAdsReq    int      `json:"numberAdsReq,omitempty"`
		Page            string   `json:"page,omitempty"`
		Location        []string `json:"location,omitempty"`
		ShippingId      []string `json:"shippingId,omitempty"`
		Rating          []string `json:"rating,omitempty"`
		HotlistId       string   `json:"hotlistId,omitempty"`
		CatalogId       string   `json:"catalogId,omitempty"`
		DepartmentId    []string `json:"departmentId,omitempty"`
		PriceMin        float64  `json:"priceMin,omitempty"`
		PriceMax        float64  `json:"priceMax,omitempty"`
		Wholesale       string   `json:"wholesale,omitempty"`
		IsGm            string   `json:"isGm,omitempty"`
		IsOfficial      string   `json:"isOfficial,omitempty"`
		Preorder        string   `json:"preorder,omitempty"`
		Freereturn      string   `json:"freereturn,omitempty"`
		Cashback        string   `json:"cashback,omitempty"`
		BrandId         string   `json:"brandId,omitempty"`
		Condition       string   `json:"condition,omitempty"`
		Sorting         string   `json:"sorting,omitempty"`
		Variant         []string `json:"variant,omitempty"`
		KreasiLokal     string   `json:"kreasiLokal,omitempty"`
		MerchantVoucher []string `json:"merchant_voucher_type,omitempty"`
		AnnotationID    []string `json:"annotationId,omitempty"`
	}

	AdEventDisplayRequest struct {
		Id               string   `json:"id,omitempty"`
		PostAlgorithm    []string `json:"post_algorithm,omitempty"`
		AdType           int      `json:"ad_type,omitempty"`
		AdRank           int      `json:"ad_rank,omitempty"`
		PriceBid         float64  `json:"price_bid,omitempty"`
		AlternateKeyword string   `json:"alternate_keyword,omitempty"`
		SearchTermID     string   `json:"search_term_id,omitempty"`
		K2KDistanceScore float64  `json:"k2k_distance_score,omitempty"`
		RankBestSeller   int64    `json:"rank_best_seller,omitempty"`
	}

	AdsFilteredOut struct {
		Id         string `json:"ad_id,omitempty"`
		FilteredBy string `json:"tag_filter,omitempty"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ExternalCampaignMappingHistory -target=src/logger/type.go -destination=files/var/schema/external_campaign_mapping_history.avsc
type ExternalCampaignMappingHistory struct {
	ShopID       int64  `json:"shop_id,omitempty"`
	GroupID      int64  `json:"group_id,omitempty"`
	CampaignID   int64  `json:"campaign_id,omitempty"`
	CampaignType int    `json:"campaign_type_code,omitempty"`
	Status       int    `json:"campaign_status_code,omitempty"`
	ActionCode   int    `json:"action_code,omitempty"`
	ReasonCode   int    `json:"reason_code,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
	Nonce        string `json:"nonce_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=HeadlineImpression -target=src/logger/type.go -destination=files/var/schema/headline_impression.avsc
type HeadlineImpression struct {
	ShopID          int64             `json:"shop_id,omitempty"`
	AdID            int64             `json:"ad_id,omitempty"`
	GroupID         int64             `json:"group_id,omitempty"`
	ViewerUserID    int64             `json:"user_id,omitempty"`
	ViewedProducts  []HeadlineProduct `json:"product_detail,omitempty"`
	ReqProductCount int               `json:"req_product_count,omitempty"`
	ReqTemplateID   []int             `json:"req_template_id,omitempty"`
	UsedTemplateID  int               `json:"used_template_id,omitempty"`
	ReqSearchTerm   string            `json:"req_search_term,omitempty"`
	KeywordID       int64             `json:"keyword_id,omitempty"`
	KeywordTag      string            `json:"keyword_tag,omitempty"`
	CreateTime      string            `json:"create_time,omitempty"`
	Nonce           string            `json:"nonce,omitempty"`
	RefNonce        string            `json:"ref_nonce,omitempty"`
}

type HeadlineProduct struct {
	ID       int64   `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	CatIDs   []int64 `json:"cat_ids,omitempty"`
	ParentID int64   `json:"parent_id,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=HomebannerAdDetail -target=src/logger/type.go -destination=files/var/schema/homebanner_ad_detail.avsc
type HomebannerAdDetail struct {
	AdID            int64   `json:"ad_id,omitempty"`
	Title           string  `json:"title,omitempty"`
	CategoryIDs     []int64 `json:"category_ids,omitempty"`
	PromoCode       string  `json:"promo_code,omitempty"`
	UserTarget      int     `json:"user_target,omitempty"`
	OwnerShopID     int64   `json:"owner_shop_id,omitempty"`
	CreateTime      string  `json:"create_time,omitempty"`
	CreateBy        int64   `json:"create_by,omitempty"`
	UpdateTime      string  `json:"update_time,omitempty"`
	UpdateBy        int64   `json:"update_by,omitempty"`
	EventType       int     `json:"event_type,omitempty"`
	EventCreateTime string  `json:"event_create_time,omitempty"`
	Nonce           string  `json:"nonce,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ImpressionEvent -target=src/logger/type.go -destination=files/var/schema/impression_v2.avsc
//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ImpressionEvent -target=src/logger/type.go -destination=files/var/schema/non_unique_impression.avsc
type (
	ImpressionEvent struct {
		Common                    ImpressionCommonEvent `json:"common,omitempty"`
		Publisher                 PublisherEvent        `json:"publisher,omitempty"`
		Product                   ProductEvent          `json:"product,omitempty"`
		AdCpm                     float64               `json:"adCpm,omitempty"`
		IsAutoAds                 bool                  `json:"is_auto_ads_flag,omitempty"`
		TrackerId                 string                `json:"tracker_id,omitempty"`
		DeviceVersion             string                `json:"device_version,omitempty"`
		ClusterIdV2               int64                 `json:"cluster_v2_id,omitempty"`
		CategorySuggestEvaluation []string              `json:"tag_category_suggest_evaluation,omitempty"`
		CategorySuggestAttribute  []string              `json:"tag_category_suggest_attribute,omitempty"`
		IsLoadTest                int                   `json:"is_load_test,omitempty"` // https://tokopedia.atlassian.net/wiki/spaces/DE/pages/365592824/Data+Formatting+Guideline
		RequestType               string                `json:"request_type,omitempty"`
		IrisSessionId             string                `json:"iris_session_id,omitempty"`
	}

	ImpressionCommonEvent struct {
		Ad            ImpAdEvent             `json:"ad,omitempty"`
		Shop          ShopEvent              `json:"shop,omitempty"`
		Network       ImpressionNetworkEvent `json:"network,omitempty"`
		Nonce         string                 `json:"nonce,omitempty"`
		EventType     string                 `json:"eventType,omitempty"`
		CreateTime    string                 `json:"createTime,omitempty"`
		User          UserEvent              `json:"user,omitempty"`
		Filter        FilterEvent            `json:"filter,omitempty"`
		Tag           ImpressionTagEvent     `json:"tag,omitempty"`
		Carousel      ImpCarouselEvent       `json:"carousel,omitempty"`
		NCandidateAds int                    `json:"nCandidateAds,omitempty"`
		NumberOfAds   int                    `json:"numberOfAds,omitempty"`
		RefNonce      string                 `json:"refNonce,omitempty"`
		Fraud         string                 `json:"fraud,omitempty"`
	}

	ImpressionTagEvent struct {
		AlgorithmTags               []string `json:"algorithmTags,omitempty"`
		IsSearchNotFound            string   `json:"isSearchNotFound,omitempty"`
		UserTarget                  string   `json:"userTarget,omitempty"`
		PublisherRemarketing        string   `json:"publisherRemarketing,omitempty"`
		SearchQueryRemarketing      string   `json:"searchQueryRemarketing,omitempty"`
		AbTest                      string   `json:"abTest,omitempty"`
		KeywordSuggest              string   `json:"keywordSuggest,omitempty"`
		NCandidateSuggest           int      `json:"nCandidateSuggest,omitempty"`
		OriginShopId                string   `json:"originShopId,omitempty"`
		OriginProductId             string   `json:"originProductId,omitempty"`
		OriginDepId                 string   `json:"origindepid,omitempty"`
		PostAlgorithm               []string `json:"postAlgorithm,omitempty"`
		QueryTags                   []string `json:"queryTags,omitempty"`
		TemporaryTags               []string `json:"temporaryTags,omitempty"`
		ProductTagName              []string `json:"product_tags_name,omitempty"`
		AlternateKeyword            string   `json:"alternate_keyword,omitempty"`
		IsIgnoredCategoryPrediction bool     `json:"is_ignored_category_prediction,omitempty"`
		RollenceFlow                string   `json:"rollence_flow,omitempty"`
		RollenceExperimentName      string   `json:"rollence_experiment_name,omitempty"`
		RollenceVariantName         string   `json:"rollence_variant_name,omitempty"`
		CampaignID                  string   `json:"campaign_id,omitempty"`
		UtmSource                   string   `json:"utm_source,omitempty"`
		UtmMedium                   string   `json:"utm_medium,omitempty"`
		TypoSuggestion              string   `json:"typosuggestion,omitempty"`
	}

	ImpressionNetworkEvent struct {
		IpAddress  string        `json:"ipAddress,omitempty"`
		DeviceType string        `json:"deviceType,omitempty"`
		UserAgent  string        `json:"userAgent,omitempty"`
		Location   LocationEvent `json:"location,omitempty"`
		Source     string        `json:"source,omitempty"`
		SessionId  string        `json:"sessionId,omitempty"`
		Origin     string        `json:"origin,omitempty"`
		RefURL     string        `json:"refURL,omitempty"`
	}

	ImpAdEvent struct {
		Id                      string         `json:"id,omitempty"`
		GoalId                  string         `json:"goal_id,omitempty"`
		Type                    string         `json:"type,omitempty"`
		Score                   float64        `json:"score,omitempty"`
		PriceBid                float64        `json:"priceBid,omitempty"`
		AdSequence              string         `json:"adSequence,omitempty"`
		HeadLine                HeadLineEvent  `json:"cpm,omitempty"`
		KeywordId               string         `json:"keywordId,omitempty"`
		KeywordCorrection       string         `json:"keywordCorrection,omitempty"`
		KeywordTag              string         `json:"keywordTag,omitempty"`
		KeywordBid              float64        `json:"keywordBid,omitempty"`
		GroupId                 string         `json:"groupId,omitempty"`
		StickerId               string         `json:"stickerId,omitempty"`
		AdLocation              string         `json:"adLocation,omitempty"`
		Formula                 FormulaEvent   `json:"formula,omitempty"`
		Status                  string         `json:"status,omitempty"`
		DailyBudget             float64        `json:"dailyBudget,omitempty"`
		ScheduleStart           string         `json:"scheduleStart,omitempty"`
		ScheduleEnd             string         `json:"scheduleEnd,omitempty"`
		Statistic               StatisticEvent `json:"statistic,omitempty"`
		SearchTermID            string         `json:"search_term_id,omitempty"`
		KeywordType             string         `json:"keyword_type,omitempty"`
		ScoreSearch             float64        `json:"score_search,omitempty"`
		ScoreAdsBooster         float64        `json:"score_ads_booster,omitempty"`
		ScoreFinalIndex         float64        `json:"score_final_index,omitempty"`
		ScoreFinalIndexCatBoost float64        `json:"score_es_query,omitempty"`
		ScoreFinalUnified       float64        `json:"score_final_unified,omitempty"`
		IsMatchSearch           bool           `json:"is_match_search,omitempty"`
		RTBBid                  float64        `json:"rtb_bid,omitempty"`
		ECVR                    float64        `json:"ecvr,omitempty"`
		PCVR                    float64        `json:"pcvr,omitempty"`
		GroupStrategy           []int          `json:"group_strategy,omitempty"`
		RelevanceScore          float64        `json:"kt_relevance_score,omitempty"`
		DeduplicateNKT          bool           `json:"deduplicate_nkt,omitempty"`
		K2KDistanceScore        float64        `json:"k2k_distance_score,omitempty"`
		FinalRelrankAdscore     float64        `json:"final_relrank_adscore,omitempty"`
		QueryRelrankScore       float64        `json:"query_relrank_score,omitempty"`
		PCTR                    float64        `json:"pctr,omitempty"`
		Position                int            `json:"position,omitempty"`
		PreRTBPosition          int            `json:"pre_rtb_position,omitempty"`
		PreRTBAdSequence        int            `json:"pre_rtb_adsequence,omitempty"`
		PreRTBPage              int            `json:"pre_rtb_page,omitempty"`
		RTBRank                 int            `json:"rtb_rank,omitempty"`
		RankBestSeller          int64          `json:"rank_best_seller,omitempty"`
	}

	// ImpCarouselEvent struct for carousel related information
	ImpCarouselEvent struct {
		Position int `json:"position,omitempty"`
	}

	// PushToIrisConvertCPA struct define iris data on seller
	// conver CPA, all json tag based on contract with DA team
	PushToIrisConvertCPA struct {
		TemplateName string  `json:"template_name"`
		ShopID       int64   `json:"shop_id"`
		UserID       int64   `json:"user_id"`
		IsSucceed    bool    `json:"is_succeed"`
		AdsConverted []int64 `json:"ads_converted"`
	}

	//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=EventHistory -target=src/logger/type.go -destination=files/var/schema/event_history.avsc
	EventHistory struct {
		EventID    int64  `json:"event_id"`
		Name       string `json:"name"`
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		Schedule   string `json:"schedule"`
		Type       int    `json:"type"`
		CreateBy   int64  `json:"create_by"`
		CreateTime string `json:"create_time"`
		UpdateBy   int64  `json:"update_by"`
		UpdateTime string `json:"update_time"`
		Status     int    `json:"status"`
		Action     string `json:"action"`
	}

	//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=UserEventHistory -target=src/logger/type.go -destination=files/var/schema/user_event_history.avsc
	UserEventHistory struct {
		EventID int64  `json:"event_id"`
		UserID  int64  `json:"user_id"`
		ShopID  int64  `json:"shop_id"`
		Action  string `json:"action"`
	}

	//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ChannelEventHistory -target=src/logger/type.go -destination=files/var/schema/channel_event_history.avsc
	ChannelEventHistory struct {
		EventID     int64  `json:"event_id"`
		ChannelType int    `json:"channel_type"`
		ChannelName string `json:"channel_name"`
		ChannelID   int64  `json:"channel_id"` //it can be ads is / ticket id etc..
		Action      string `json:"action"`
	}

	TopAdsUsed struct {
		Last7Days              float64 `json:"last_7_days,omitempty"`
		RemainingTopAdsBalance float64 `json:"remaining_topads_balance"`
		BudgetCategory         int     `json:"budget_category,omitempty"`
		UpdateTime             string  `json:"update_time,omitempty"`
		UserID                 int64   `json:"user_id,omitempty"`
		Source                 string  `json:"source,omitempty"`
		Click7Days             int64   `json:"click_7_days,omitempty"`
		Impression7Days        int64   `json:"impression_7_days,omitempty"`
		Conversion7Days        int     `json:"conversion_7_days,omitempty"`
		Revenue7Days           float64 `json:"revenue_7_days,omitempty"`
		Roas7Days              float64 `json:"roas_7_days,omitempty"`
		Conv1Day               int     `json:"conv_1_day,omitempty"`
		AverageBidCpc7Days     float64 `json:"average_bid_cpc_7_days,omitempty"`
		DayOfWeek              int     `json:"day_of_week,omitempty"`
		TriggerAutoTopUp       bool    `json:"trigger_auto_topup"`
		AutoAdsFlag            bool    `json:"auto_ads_flag"`
		Impression7DaysFmt     string  `json:"impression_7_days_fmt,omitempty"`
		Conversion7DaysFmt     string  `json:"conversion_7_days_fmt,omitempty"`
		Revenue7DaysFmt        string  `json:"revenue_7_days_fmt,omitempty"`
		Roas7DaysFmt           string  `json:"roas_7_days_fmt,omitempty"`
		Conv1DayFmt            string  `json:"conv_1_day_fmt,omitempty"`
		AverageBidCpc7DaysFmt  string  `json:"average_bid_cpc_7_days_fmt,omitempty"`
		Click1Day              int64   `json:"click_1_day,omitempty"`
		Conv30Days             int     `json:"conv_30_days,omitempty"`
		Impression1Day         int64   `json:"imp_1_day,omitempty"`
		Spend1Day              float64 `json:"spend_1_day,omitempty"`
		IsTopAdsUser           int     `json:"is_top_ads_user"`
		ShopID                 int64   `json:"shop_id,omitempty"`
		Imp30Days              int64   `json:"imp_30_days,omitempty"`
		Click30Days            int64   `json:"click_30_days,omitempty"`
		Spend30Day             float64 `json:"spend_30_day,omitempty"`
		RoasProduct7Days       float64 `json:"roas_product_7_days,omitempty"`
		RoasProduct7DaysFmt    string  `json:"roas_product_7_days_fmt,omitempty"`
		FreeCreditReason       int     `json:"free_credit_reason"`
		AutoBidPenetration     float64 `json:"autobid_penetration"`
		NonAutoBidGroupCount   int     `json:"non_autobid_group_count"`
	}

	TopAdsUsedNew struct {
		Last7Days                 float64 `json:"last_7_days,omitempty"`
		RemainingTopAdsBalance    float64 `json:"remaining_topads_balance"`
		RemainingTopAdsBalancefmt string  `json:"remaining_topads_balance_fmt,omitempty"`
		UsageLeft                 int     `json:"usage_left,omitempty"`
		Usage7Days                float64 `json:"usage_7_days,omitempty"`
		BudgetCategory            int     `json:"budget_category,omitempty"`
		UpdateTime                string  `json:"update_time,omitempty"`
		UserID                    int64   `json:"user_id,omitempty"`
		Source                    string  `json:"source,omitempty"`
		Click7Days                int64   `json:"click_7_days,omitempty"`
		Click7Daysfmt             string  `json:"click_7_days_fmt,omitempty"`
		Impression7Days           int64   `json:"impression_7_days,omitempty"`
		Conversion7Days           int     `json:"conversion_7_days,omitempty"`
		Revenue7Days              float64 `json:"revenue_7_days,omitempty"`
		Roas7Days                 float64 `json:"roas_7_days,omitempty"`
		Conv1Day                  int     `json:"conversion_1_day,omitempty"`
		AverageBidCpc7Days        float64 `json:"average_bid_cpc_7_days,omitempty"`
		DayOfWeek                 int     `json:"day_of_week,omitempty"`
		TriggerAutoTopUp          bool    `json:"trigger_auto_topup"`
		AutoAdsFlag               bool    `json:"auto_ads_flag"`
		Impression7DaysFmt        string  `json:"impression_7_days_fmt,omitempty"`
		Conversion7DaysFmt        string  `json:"conversion_7_days_fmt,omitempty"`
		Revenue7DaysFmt           string  `json:"revenue_7_days_fmt,omitempty"`
		Roas7DaysFmt              string  `json:"roas_7_days_fmt,omitempty"`
		Conv1DayFmt               string  `json:"conversion_1_day_fmt,omitempty"`
		AverageBidCpc7DaysFmt     string  `json:"average_bid_cpc_7_days_fmt,omitempty"`
		Click1Day                 int64   `json:"click_1_day,omitempty"`
		Click1Dayfmt              string  `json:"click_1_day_fmt,omitempty"`
		Conv30Days                int     `json:"conversion_30_days,omitempty"`
		Conv30Daysfmt             string  `json:"conversion_30_days_fmt,omitempty"`
		Impression1Day            int64   `json:"imp_1_day,omitempty"`
		Impression1Dayfmt         string  `json:"imp_1_day_fmt,omitempty"`
		Spend1Day                 float64 `json:"spend_1_day,omitempty"`
		IsTopAdsUser              int     `json:"is_top_ads_user"`
		ShopID                    int64   `json:"shop_id,omitempty"`
		Imp30Days                 int64   `json:"imp_30_days,omitempty"`
		Click30Days               int64   `json:"click_30_days,omitempty"`
		Click30Daysfmt            string  `json:"click_30_days_fmt,omitempty"`
		Spend30Day                float64 `json:"spend_30_day,omitempty"`
		RoasProduct7Days          float64 `json:"roas_product_7_days,omitempty"`
		RoasProduct7DaysFmt       string  `json:"roas_product_7_days_fmt,omitempty"`
		FreeCreditReason          int     `json:"free_credit_reason"`
		Roas1Days                 float64 `json:"roas_1_days,omitempty"`
		Roas1Daysfmt              string  `json:"roas_1_days_fmt,omitempty"`
		AdsPenetration            int64   `json:"ads_penetration"`
	}

	InactiveAllAds struct {
		UserID     int64  `json:"user_id,omitempty"`
		UpdateTime string `json:"update_time,omitempty"`
		Source     string `json:"source,omitempty"`
		ShopID     int64  `json:"shop_id,omitempty"`
	}

	AdsActivated struct {
		UserID     int64  `json:"user_id,omitempty"`
		ShopID     int64  `json:"shop_id,omitempty"`
		UpdateTime string `json:"update_time,omitempty"`
		Source     string `json:"source,omitempty"`
	}

	CreditEligible struct {
		FreeCreditReason int    `json:"free_credit_reason,omitempty"`
		StartClaim       string `json:"start_claim,omitempty"`
		EndClaim         string `json:"end_claim,omitempty"`
		BenefitFmt       string `json:"benefit_fmt,omitempty"`
		UpdateTime       string `json:"update_time,omitempty"`
		Source           string `json:"source,omitempty,omitempty"`
		UserID           int64  `json:"user_id,omitempty"`
		ShopID           int64  `json:"shop_id,omitempty"`
	}

	CreditClaimed struct {
		FreeCreditReason int    `json:"free_credit_reason,omitempty"`
		EndDate          string `json:"end_date,omitempty"`
		UpdateTime       string `json:"update_time,omitempty"`
		Source           string `json:"source,omitempty,omitempty"`
		UserID           int64  `json:"user_id,omitempty"`
		Nominal          int64  `json:"benefit,omitempty"`
		NominalFmt       string `json:"benefit_fmt,omitempty"`
		ShopID           int64  `json:"shop_id,omitempty"`
	}

	//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=TopAdsCampaign -target=src/logger/type.go -destination=files/var/schema/campaign.avsc
	TopAdsCampaign struct {
		UserID             int64   `json:"user_id,omitempty"`
		UpdateTime         string  `json:"update_time,omitempty"`
		Source             string  `json:"source,omitempty"`
		CampaignType       int     `json:"campaign_type,omitempty"`
		CampaignName       string  `json:"campaign_name,omitempty"`
		TotalData          int64   `json:"total_data,omitempty"`
		DayOfWeek          int     `json:"day_of_week,omitempty"`
		AverageBidCpc7Days float64 `json:"average_bid_cpc_7_days,omitempty"`
		ShopID             int64   `json:"shop_id,omitempty"`
	}

	PurchaseStartEvent struct {
		Uuid                   string  `json:"uuid,omitempty"`
		NominalFmt             string  `json:"nominal_fmt,omitempty"`
		Nominal                int64   `json:"nominal,omitempty"`
		UpdateTime             string  `json:"update_time,omitempty"`
		Source                 string  `json:"source,omitempty,omitempty"`
		UserID                 int64   `json:"user_id,omitempty"`
		ShopID                 int64   `json:"shop_id"`
		UsageLeft              int     `json:"usage_left,omitempty"`
		RemainingTopadsBalance float64 `json:"remaining_topads_balance,omitempty"`
	}

	PurchaseFinishedEvent struct {
		UserID                int64   `json:"user_id,omitempty"`
		Uuid                  string  `json:"uuid,omitempty"`
		PromoCode             string  `json:"promo_code,omitempty"`
		UpdateTime            string  `json:"update_time,omitempty"`
		Source                string  `json:"source,omitempty,omitempty"`
		TransactionID         int64   `json:"transaction_id,omitempty"`
		TopAdsPurchasePackage string  `json:"topads_purchase_package,omitempty"`
		TopAdsAmt             float64 `json:"topads_amt,omitempty"`
		PaymentMethod         int     `json:"payment_method,omitempty"`
		ShopID                int64   `json:"shop_id,omitempty"`
		PaymentAction         string  `json:"payment_action,omitempty"`
	}

	PurchaseWaitingEvent struct {
		UserID                 int64   `json:"user_id,omitempty"`
		Uuid                   string  `json:"uuid,omitempty"`
		UpdateTime             string  `json:"update_time,omitempty"`
		Source                 string  `json:"source,omitempty,omitempty"`
		ShopId                 int64   `json:"shop_id,omitempty"`
		PaymentAction          string  `json:"payment_action,omitempty"`
		TransactionId          int64   `json:"transaction_id,omitempty"`
		UsageLeft              int     `json:"usage_left,omitempty"`
		RemainingTopadsBalance float64 `json:"remaining_topads_balance,omitempty"`
	}

	ShopAdminAccess struct {
		AdminUserID int64  `json:"admin_user_id"`
		ShopID      int64  `json:"shop_id"`
		AdminType   int    `json:"admin_type"`
		Action      int    `json:"action"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		CreateTime  string `json:"create_time"`
		AdminName   string `json:"admin_name"`
	}

	TopAdsDormantUser struct {
		UserID           int64   `json:"user_id,omitempty"`
		UpdateTime       string  `json:"update_time,omitempty"`
		Roas             float64 `json:"ads_spend_roas"`
		Credit           float64 `json:"credit_limit"`
		Reputation       int     `json:"shop_reputation"`
		UsedAutoAds      bool    `json:"auto_ads_flag"`
		Source           string  `json:"source"`
		ShopID           int64   `json:"shop_id,omitempty"`
		DaysNoImpression int     `json:"days_no_impression"`
		RemainingCredit  float64 `json:"remaining_credit"`
		Imp7Days         int64   `json:"imp_7_days"`
		Revenue7Days     float64 `json:"revenue_7_days"`
		Click7Days       int64   `json:"click_7_days"`
		ItemSold7Days    int     `json:"item_sold_7_days"`
		AdsUsed          string  `json:"ads_used"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=TopAdsUsed -target=src/logger/type.go -destination=files/var/schema/used.avsc

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=InsertDepositEventV2 -target=src/logger/type.go -destination=files/var/schema/insert_deposit_v2.avsc
type (
	InsertDepositEventV2 struct {
		Network     InsertDepositNetworkEvent `json:"network,omitempty"`
		User        UserEvent                 `json:"user,omitempty"`
		Shop        ShopEvent                 `json:"shop,omitempty"`
		Deposit     DepositEvent              `json:"deposit,omitempty"`
		Order       OrderEvent                `json:"order,omitempty"`
		Payment     PaymentEvent              `json:"payment,omitempty"`
		Promo       PromoEvent                `json:"promo"`
		PrevDeposit float64                   `json:"previousDeposit,omitempty"`
		LastBought  string                    `json:"lastBought,omitempty"`
		Nonce       string                    `json:"nonce,omitempty"`
		EventType   string                    `json:"eventType,omitempty"`
		CreateTime  string                    `json:"createTime,omitempty"`
	}

	InsertDepositNetworkEvent struct {
		IpAddress  string        `json:"ipAddress,omitempty"`
		DeviceType string        `json:"deviceType,omitempty"`
		UserAgent  string        `json:"userAgent,omitempty"`
		Location   LocationEvent `json:"location,omitempty"`
		Source     string        `json:"source,omitempty"`
		SessionId  string        `json:"sessionId,omitempty"`
		Origin     string        `json:"origin,omitempty"`
		RefURL     string        `json:"refURL,omitempty"`
		Channel    string        `json:"channel,omitempty"`
	}
)

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=CPATrackEvent -target=src/logger/type.go -destination=files/var/schema/order_attempt_tracker.avsc
type CPATrackEvent struct {
	RecId           string `json:"rec_id"`
	SessionId       string `json:"session_id"`
	UserId          int64  `json:"user_id"`
	DeviceName      string `json:"device_name"`
	AdId            int64  `json:"ad_id"`
	AffiliateId     int64  `json:"affiliate_id"`
	UniqueTrackerId string `json:"unique_tracker_id"`
	ActionTrackerId string `json:"action_tracker_id"`
	IsSuccess       bool   `json:"is_success_flag"`
	IsError         bool   `json:"is_error_flag"`
	ReasonDesc      string `json:"reason_desc"`
	CreateTime      string `json:"create_time"`
	ProductName     string `json:"product_name"`
	ProductID       int64  `json:"product_id"`
	SourceType      string `json:"source_type"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=PushtoIrisCheckout -target=src/logger/type.go -destination=files/var/schema/order_checkout.avsc
type PushtoIrisCheckout struct {
	RecID                     string  `json:"rec_id"`
	DeviceName                string  `json:"device_name"`
	ProductName               string  `json:"affiliate_product_name"` // product_name is already used in iris as float so now not usable
	OrderID                   int64   `json:"order_id"`
	OrderDetailID             int64   `json:"orderdetailid"`
	PaymentID                 int64   `json:"payment_id"`
	TotalPaymentAffiliate     float64 `json:"total_payment_aff"`
	LastClickAffiliateID      int64   `json:"last_affiliate_id"`
	LastClickAdID             int64   `json:"last_ad_id"`
	LastClickCuratedProductID int64   `json:"last_curated_product_id"`
	UserID                    int64   `json:"user_id"`
	ActionTrackerID           string  `json:"action_tracker_id"`
	UniqueTrackerId           string  `json:"unique_tracker_id"`
	ClientTimestamp           string  `json:"client_timestamp"`
	SessionID                 string  `json:"session_id"`
	IpAddress                 string  `json:"ip_address"`
	CreateTime                string  `json:"create_time"`
	IsIndirect                bool    `json:"isindirect"`
	ProductID                 int64   `json:"product_id"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ShopEventSchedule -target=src/logger/type.go -destination=files/var/schema/shop_event_schedule_v2.avsc
type ShopEventSchedule struct {
	ShopId     int64  `json:"shop_id,omitempty"`
	Type       int    `json:"event_type_code,omitempty"`
	Status     int    `json:"event_status_code,omitempty"`
	StartTime  string `json:"start_time,omitempty"`
	EndTime    string `json:"end_time,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Nonce      string `json:"nonce,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=GroupRecommendation -target=src/logger/type.go -destination=files/var/schema/group_recommendation.avsc
type GroupRecommendation struct {
	GroupId    string   `json:"group_id,omitempty"`
	ShopId     string   `json:"shop_id,omitempty"`
	Type       int      `json:"event_type_code,omitempty"`
	Value      []string `json:"event_value,omitempty"`
	PriceBid   float64  `json:"price_bid_amount,omitempty"`
	CreateTime string   `json:"create_time,omitempty"`
	Nonce      string   `json:"nonce,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsECPC -target=src/logger/type.go -destination=files/var/schema/auto_ads_ecpc.avsc
type AutoAdsECPC struct {
	AdID               int64   `json:"ad_id,omitempty"`
	MaxBid             float64 `json:"max_bid,omitempty"`
	CompetitiveBid     float64 `json:"competitive_bid,omitempty"`
	RoasBasedCPC       float64 `json:"roas_based_cpc,omitempty"`
	CommissionBasedCPC float64 `json:"commission_based_cpc,omitempty"`
	ActualBid          float64 `json:"actual_bid,omitempty"`
	CreateTime         string  `json:"create_time,omitempty"`
	EffectiveCPC       float64 `json:"ecpc_product,omitempty"`
	CTRProductTA       float64 `json:"ctr_product_ta,omitempty"`
	CVRProductTA       float64 `json:"cvr_product_ta,omitempty"`
	Source             string  `json:"source,omitempty"`
	ExperimentID       int64   `json:"experiment_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=TopAdsCampaignAction -target=src/logger/type.go -destination=files/var/schema/campaign_action.avsc
type TopAdsCampaignAction struct {
	UserID       int64  `json:"user_id,omitempty"`
	Source       string `json:"source"`
	CampaignType int    `json:"campaign_type"`
	CampaignName string `json:"campaign_name"`
	Action       string `json:"action"`
	UpdateTime   string `json:"update_time"`
	ShopID       int64  `json:"shop_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=TopAdsFirstConversion -target=src/logger/type.go -destination=files/var/schema/first_conversion.avsc
type TopAdsFirstConversion struct {
	FreeCreditReason int   `json:"free_credit_reason"`
	UserID           int64 `json:"user_id,omitempty"`
	ShopID           int64 `json:"shop_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=TopAdsFirstImp -target=src/logger/type.go -destination=files/var/schema/first_imp.avsc
type TopAdsFirstImp struct {
	FreeCreditReason int   `json:"free_credit_reason"`
	UserID           int64 `json:"user_id,omitempty"`
	ShopID           int64 `json:"shop_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AdScoreRecalculationLogger -target=src/logger/type.go -destination=files/var/schema/score_recalculation_logger.avsc
type AdScoreRecalculationLogger struct {
	TimeTriggered            string  `json:"time_triggered,omitempty"`
	AdID                     int64   `json:"ad_id,omitempty"`
	KeywordID                int64   `json:"keyword_id,omitempty"`
	SearchTermID             int64   `json:"search_term_id,omitempty"`
	TriggeredType            int     `json:"triggered_type,omitempty"`
	TriggeredActivity        string  `json:"triggered_activity,omitempty"`
	SearchControlAdScore     float64 `json:"search_control_ad_score,omitempty"`
	BrowseControl            float64 `json:"browse_score,omitempty"`
	SearchExperimentAdScore1 float64 `json:"search_experiment_ad_score_1,omitempty"`
	SearchExperimentAdScore2 float64 `json:"search_experiment_ad_score_2,omitempty"`
	SearchExperimentAdScore3 float64 `json:"search_experiment_ad_score_3,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AdsReviewHistory -target=src/logger/type.go -destination=files/var/schema/ads_review_history.avsc
type AdsReviewHistory struct {
	AdID                int64  `json:"ad_id"`
	AdType              int    `json:"ad_type"`
	ShopID              int64  `json:"shop_id"`
	Status              int    `json:"status"`
	ReviewedTime        string `json:"reviewed_time,omitempty"`
	ReviewedBy          int64  `json:"reviewed_by,omitempty"`
	RejectionReasonText string `json:"rejection_reason_text,omitempty"`
	Source              string `json:"source"`
	ItemID              int64  `json:"item_id"`
	ProductName         string `json:"product_name"`
	ProductDesc         string `json:"product_description"`
	NormalPrice         int64  `json:"normal_price"`
	ChildCatID          int64  `json:"child_cat_id"`
	ProductImageURL     string `json:"product_image_url"`
	Action              string `json:"action"`
	CreateTime          string `json:"createtime"`
	FreeTextReason      string `json:"free_text_reason"`
	DSRejectionReason   string `json:"ds_rejection_reason"`
	RejectionIDs        string `json:"rejection_ids"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=HeadlineUsed -target=src/logger/type.go -destination=files/var/schema/headline_used.avsc
type HeadlineUsed struct {
	Roas7Days          float64 `json:"roas_7_days,omitempty"`
	Roas7DaysFmt       string  `json:"roas_7_days_fmt,omitempty"`
	Impression7Days    int64   `json:"impression_7_days,omitempty"`
	Impression7DaysFmt string  `json:"impression_7_days_fmt,omitempty"`
	Click7Days         int64   `json:"click_7_days,omitempty"`
	Conversion7Days    int     `json:"conversion_7_days,omitempty"`
	Conversion7DaysFmt string  `json:"conversion_7_days_fmt,omitempty"`
	Revenue7Days       float64 `json:"revenue_7_days,omitempty"`
	Revenue7DaysFmt    string  `json:"revenue_7_days_fmt,omitempty"`
	Impression1Day     int64   `json:"imp_1_day,omitempty"`
	Click1Day          int64   `json:"click_1_day,omitempty"`
	Conv1Day           int     `json:"conv_1_day,omitempty"`
	Conv1DayFmt        string  `json:"conv_1_day_fmt,omitempty"`
	Last7Days          float64 `json:"last_7_days,omitempty"`
	DayOfWeek          int     `json:"day_of_week,omitempty"`
	UserID             int64   `json:"user_id,omitempty"`
	ShopID             int64   `json:"shop_id,omitempty"`
	EverUsedHeadline   bool    `json:"ever_used_headline"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=HeadlineDormantUser -target=src/logger/type.go -destination=files/var/schema/headline_dormant_user.avsc
type HeadlineDormantUser struct {
	UserID      int64   `json:"user_id,omitempty"`
	Roas        float64 `json:"ads_spend_roas"`
	Credit      float64 `json:"credit_limit"`
	Reputation  int     `json:"shop_reputation"`
	UsedAutoAds bool    `json:"auto_ads_flag"`
	ShopID      int64   `json:"shop_id,omitempty"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=CPACampaign -target=src/logger/type.go -destination=files/var/schema/cpa_campaign.avsc
type CPACampaign struct {
	UserID     int64  `json:"user_id,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Source     string `json:"source,omitempty"`
	ShopID     int64  `json:"shop_id,omitempty"`
}

type ShopGroupNewInsightDataStruct struct {
	ShopID                    int64  `json:"shop_id,omitempty"`
	UserId                    int64  `json:"user_id"`
	AdType                    int    `json:"ad_type,omitempty"`
	CommonId                  int64  `json:"common_id,omitempty"`
	KeywordRecomCount         int    `json:"keyword_recom_count,omitempty"`
	BidRecomCount             int    `json:"bid_recom_count,omitempty"`
	NegativeKeywordRecomCount int    `json:"negative_keyword_recom_count,omitempty"`
	DateId                    string `json:"date_id,omitempty"`
	NewInsightCount           int    `json:"new_insight_count"`
	AvailableKeywordInsights  int    `json:"available_keyword_insights"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=ShopWhitelistFeature -target=src/logger/type.go -destination=files/var/schema/shop_whitelist_feature.avsc
type ShopWhitelistFeature struct {
	ShopID      int64  `json:"shop_id"`
	FeatureID   int64  `json:"feature_id"`
	FeatureName string `json:"feature_name"`
	Action      string `json:"action"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsKeywordLengthFiltering -target=src/logger/type.go -destination=files/var/schema/auto_ads_keyword_length_filtering.avsc
type AutoAdsKeywordLengthFiltering struct {
	AdID                   int64  `json:"ad_id,omitempty"`
	ProductID              int64  `json:"product_id,omitempty"`
	ShopID                 int64  `json:"shop_id,omitempty"`
	KeywordLengthThreshold []int  `json:"keyword_length_threshold,omitempty"`
	CreateTime             string `json:"create_time,omitempty"`
}

type TopadsLowCredit struct {
	UsageLeft        int     `json:"usage_left"`
	DaysNoImpression int     `json:"days_no_impression"`
	TopadsCredit     float64 `json:"remaining_credit"`
	Imp7Days         int64   `json:"imp_7_days"`
	Click7Days       int64   `json:"click_7_days"`
	ItemSold7Days    int     `json:"item_sold_7_days"`
	Revenue7Days     float64 `json:"revenue_7_days"`
	ShopId           int64   `json:"shop_id"`
	UserId           int64   `json:"user_id"`
}

type TopadsProductAcquisition struct {
	ProdAcqFCStart     string `json:"prod_acq_fc_start"`
	ProdAcqNonFCStart  string `json:"prod_acq_non_fc_start"`
	ProdAcqFCEnd       string `json:"prod_acq_fc_end"`
	ProdAcqNonFCSEnd   string `json:"prod_acq_non_fc_end"`
	TopadsStatus       string `json:"topads_status"`
	FreeCreditEligible bool   `json:"free_credit_eligible"`
	ShopId             int64  `json:"shop_id"`
	UserId             int64  `json:"user_id"`
}

type TopadsDownloadStatisticsReport struct {
	AdsType    string `json:"ads_type,omitempty"`
	AdminId    int64  `json:"admin_id,omitempty"`
	ShopID     int64  `json:"shop_id,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	StartDate  string `json:"start_date,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
}

type TopadsNewInsight struct {
	InsightType string `json:"insight_type"`
	Count       int    `json:"count"`
	Day         int    `json:"day"`
	ShopID      int64  `json:"shop_id"`
	UserID      int64  `json:"user_id"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=LogProductRecommendation -target=src/logger/type.go -destination=files/var/schema/product_recommendation_insight.avsc
type LogProductRecommendation struct {
	ShopID         int64   `json:"shop_id,omitempty"`
	ProductID      int64   `json:"product_id,omitempty"`
	SearchCount    int64   `json:"search_count,omitempty"`
	RecommendedBid float64 `json:"recommended_bid,omitempty"`
}

type TopadsTopUp struct {
	IsAutoTopUp bool   `json:"is_auto_top_up"`
	TopUpCount  int    `json:"top_up_count"`
	UserID      int64  `json:"user_id"`
	UpdateTime  string `json:"update_time"`
	ShopID      int64  `json:"shop_id"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoBidAds -target=src/logger/type.go -destination=files/var/schema/autobid_ads.avsc
type AutoBidAds struct {
	ShopID        int64   `json:"shop_id"`
	Source        string  `json:"source"`
	GroupID       int64   `json:"group_id"`
	GroupBudget   float64 `json:"group_budget"`
	AdID          int64   `json:"ad_id"`
	ProductID     int64   `json:"product_id"`
	ProductBudget float64 `json:"product_budget"`
	CreateTime    string  `json:"create_time"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsDailyBudgetHit -target=src/logger/type.go -destination=files/var/schema/auto_ads_daily_budget_hit.avsc
type AutoAdsDailyBudgetHit struct {
	ShopID      int64   `json:"shop_id"`
	UserID      int64   `json:"user_id"`
	EventID     int64   `json:"event_id"`
	DailyBudget float64 `json:"daily_budget"`
	CreateTime  string  `json:"create_time"`
	Roas1Day    float64 `json:"roas_1_day"`
	Roas7Day    float64 `json:"roas_7_day"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsActivation -target=src/logger/type.go -destination=files/var/schema/auto_ads_activation.avsc
type AutoAdsActivation struct {
	ShopID        int64   `json:"shop_id"`
	UserID        int64   `json:"user_id"`
	Credit        float64 `json:"credit"`
	EventID       int64   `json:"event_id"`
	TriggerAction string  `json:"trigger_action"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoBidActivation -target=src/logger/type.go -destination=files/var/schema/auto_bid_activation.avsc
type AutoBidActivation struct {
	ShopID             int64   `json:"shop_id"`
	UserID             int64   `json:"user_id"`
	GroupID            int64   `json:"group_id"`
	AutoBidPenetration float64 `json:"autobid_penetration"`
}

//go:generate $GOPATH/src/github.com/tokopedia/topads-logging/tools/avrogenerator/avrogenerator -struct-name=AutoAdsDailyBudgetUpdated -target=src/logger/type.go -destination=files/var/schema/auto_ads_daily_budget_updated.avsc
type AutoAdsDailyBudgetUpdated struct {
	DailyBudgetOld float64 `json:"daily_budget_old"`
	DailyBudgetNew float64 `json:"daily_budget_new"`
	UserID         int64   `json:"user_id"`
	ShopID         int64   `json:"shop_id"`
}

type DisplayBrowseRequestEvent struct {
	Network          NetworkEvent    `json:"network,omitempty"`
	EventType        string          `json:"eventType,omitempty"`
	CreateTime       string          `json:"createTime,omitempty"`
	NumberOfAds      int             `json:"numberOfAds,omitempty"`
	User             UserEvent       `json:"user,omitempty"`
	AdsFromES        []string        `json:"list_ads,omitempty"`
	RequestType      string          `json:"request_type,omitempty"`
	IrisSessionId    string          `json:"iris_session_id,omitempty"`
	ListOfProductIds []string        `json:"list_product_ids,omitempty"`
	IsBrowseAds      bool            `json:"is_browse_ads,omitempty"`
	ClickUrl         string          `json:"click_url,omitempty"`
	ImpressionUrl    string          `json:"impression_url,omitempty"`
	Tag              DisplayTagEvent `json:"tag,omitempty"`
	DeviceType       string          `json:"deviceType,omitempty"`
}

// IsValidIP validate textual representation of an IP address
func (n NetworkEvent) IsValidIP() bool {
	if net.ParseIP(n.IpAddress) != nil {
		return true
	}
	return false
}

// IsValidIP validate textual representation of an IP address
func (i ImpressionNetworkEvent) IsValidIP() bool {
	if net.ParseIP(i.IpAddress) != nil {
		return true
	}
	return false
}

// IsValidIP validate textual representation of an IP address
func (c ClickNetworkEvent) IsValidIP() bool {
	if net.ParseIP(c.IpAddress) != nil {
		return true
	}
	return false
}

// IsValidIP validate textual representation of an IP address
func (i InsertDepositNetworkEvent) IsValidIP() bool {
	if net.ParseIP(i.IpAddress) != nil {
		return true
	}
	return false
}
