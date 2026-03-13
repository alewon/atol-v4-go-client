package atol

type PostTokenRequest struct {
	Login  string `json:"login"`
	Pass   string `json:"pass"`
	Source string `json:"source,omitempty"`
}

type PostTokenResponse struct {
	Token     string                  `json:"token"`
	Error     *PostTokenResponseError `json:"error"`
	Timestamp string                  `json:"timestamp"`
}

type PostTokenResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type GetTokenRequest struct {
	Login  string
	Pass   string
	Source string
}

type GetTokenResponse struct {
	Token     string                 `json:"token"`
	Error     *GetTokenResponseError `json:"error"`
	Timestamp string                 `json:"timestamp"`
}

type GetTokenResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type SellRequest struct {
	Timestamp  string              `json:"timestamp"`
	SourceID   int                 `json:"source_id,omitempty"`
	ExternalID string              `json:"external_id"`
	Service    *SellRequestService `json:"service,omitempty"`
	Receipt    SellRequestReceipt  `json:"receipt"`
}

type SellRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type SellRequestReceipt struct {
	Client               SellRequestReceiptClient               `json:"client"`
	Company              SellRequestReceiptCompany              `json:"company"`
	AgentInfo            *SellRequestReceiptAgentInfo           `json:"agent_info,omitempty"`
	SupplierInfo         *SellRequestReceiptSupplierInfo        `json:"supplier_info,omitempty"`
	Items                []SellRequestReceiptItem               `json:"items"`
	Payments             []SellRequestReceiptPayment            `json:"payments"`
	Vats                 []SellRequestReceiptVat                `json:"vats,omitempty"`
	Total                float64                                `json:"total"`
	AdditionalCheckProps string                                 `json:"additional_check_props,omitempty"`
	Cashier              string                                 `json:"cashier,omitempty"`
	AdditionalUserProps  *SellRequestReceiptAdditionalUserProps `json:"additional_user_props,omitempty"`
	DeviceNumber         string                                 `json:"device_number,omitempty"`
	Internet             bool                                   `json:"internet,omitempty"`
	CashlessPayments     []SellRequestReceiptCashlessPayment    `json:"cashless_payments,omitempty"`
}

type SellRequestReceiptClient struct {
	Name  string `json:"name,omitempty"`
	INN   string `json:"inn,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type SellRequestReceiptCompany struct {
	Email          string `json:"email,omitempty"`
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type SellRequestReceiptAgentInfo struct {
	Type                    string                                              `json:"type,omitempty"`
	PayingAgent             *SellRequestReceiptAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *SellRequestReceiptAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *SellRequestReceiptAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type SellRequestReceiptAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type SellRequestReceiptAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type SellRequestReceiptAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type SellRequestReceiptSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type SellRequestReceiptItem struct {
	Name              string                              `json:"name"`
	Price             float64                             `json:"price"`
	Quantity          float64                             `json:"quantity"`
	Sum               *float64                            `json:"sum"`
	MeasurementUnit   string                              `json:"measurement_unit,omitempty"`
	PaymentMethod     string                              `json:"payment_method,omitempty"`
	PaymentObject     string                              `json:"payment_object,omitempty"`
	NomenclatureCode  string                              `json:"nomenclature_code,omitempty"`
	VAT               *SellRequestReceiptItemVAT          `json:"vat,omitempty"`
	AgentInfo         *SellRequestReceiptItemAgentInfo    `json:"agent_info,omitempty"`
	SupplierInfo      *SellRequestReceiptItemSupplierInfo `json:"supplier_info,omitempty"`
	UserData          string                              `json:"user_data,omitempty"`
	Excise            float64                             `json:"excise,omitempty"`
	CountryCode       string                              `json:"country_code,omitempty"`
	DeclarationNumber string                              `json:"declaration_number,omitempty"`
}

type SellRequestReceiptItemVAT struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRequestReceiptItemAgentInfo struct {
	Type                    string                                                  `json:"type,omitempty"`
	PayingAgent             *SellRequestReceiptItemAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *SellRequestReceiptItemAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *SellRequestReceiptItemAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type SellRequestReceiptItemAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type SellRequestReceiptItemAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type SellRequestReceiptItemAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type SellRequestReceiptItemSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type SellRequestReceiptPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRequestReceiptVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRequestReceiptAdditionalUserProps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SellRequestReceiptCashlessPayment struct {
	Sum            float64 `json:"sum"`
	Method         int     `json:"method"`
	ID             string  `json:"id"`
	AdditionalInfo string  `json:"additional_info,omitempty"`
}

type SellResponse struct {
	UUID        string             `json:"uuid"`
	Timestamp   string             `json:"timestamp"`
	Error       *SellResponseError `json:"error"`
	Status      string             `json:"status"`
	GroupCode   string             `json:"group_code"`
	DaemonCode  string             `json:"daemon_code"`
	DeviceCode  string             `json:"device_code"`
	CallbackURL string             `json:"callback_url"`
	Payload     any                `json:"payload"`
}

type SellResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type SellRefundRequest struct {
	Timestamp  string                    `json:"timestamp"`
	SourceID   int                       `json:"source_id,omitempty"`
	ExternalID string                    `json:"external_id"`
	Service    *SellRefundRequestService `json:"service,omitempty"`
	Receipt    SellRefundRequestReceipt  `json:"receipt"`
}

type SellRefundRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type SellRefundRequestReceipt struct {
	Client               SellRefundRequestReceiptClient               `json:"client"`
	Company              SellRefundRequestReceiptCompany              `json:"company"`
	AgentInfo            *SellRefundRequestReceiptAgentInfo           `json:"agent_info,omitempty"`
	SupplierInfo         *SellRefundRequestReceiptSupplierInfo        `json:"supplier_info,omitempty"`
	Items                []SellRefundRequestReceiptItem               `json:"items"`
	Payments             []SellRefundRequestReceiptPayment            `json:"payments"`
	Vats                 []SellRefundRequestReceiptVat                `json:"vats,omitempty"`
	Total                float64                                      `json:"total"`
	AdditionalCheckProps string                                       `json:"additional_check_props,omitempty"`
	Cashier              string                                       `json:"cashier,omitempty"`
	AdditionalUserProps  *SellRefundRequestReceiptAdditionalUserProps `json:"additional_user_props,omitempty"`
	DeviceNumber         string                                       `json:"device_number,omitempty"`
	Internet             bool                                         `json:"internet,omitempty"`
	CashlessPayments     []SellRefundRequestReceiptCashlessPayment    `json:"cashless_payments,omitempty"`
}

type SellRefundRequestReceiptClient struct {
	Name  string `json:"name,omitempty"`
	INN   string `json:"inn,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type SellRefundRequestReceiptCompany struct {
	Email          string `json:"email,omitempty"`
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type SellRefundRequestReceiptAgentInfo struct {
	Type                    string                                                    `json:"type,omitempty"`
	PayingAgent             *SellRefundRequestReceiptAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *SellRefundRequestReceiptAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *SellRefundRequestReceiptAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type SellRefundRequestReceiptAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type SellRefundRequestReceiptAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type SellRefundRequestReceiptAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type SellRefundRequestReceiptSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type SellRefundRequestReceiptItem struct {
	Name              string                                    `json:"name"`
	Price             float64                                   `json:"price"`
	Quantity          float64                                   `json:"quantity"`
	Sum               *float64                                  `json:"sum"`
	MeasurementUnit   string                                    `json:"measurement_unit,omitempty"`
	PaymentMethod     string                                    `json:"payment_method,omitempty"`
	PaymentObject     string                                    `json:"payment_object,omitempty"`
	NomenclatureCode  string                                    `json:"nomenclature_code,omitempty"`
	VAT               *SellRefundRequestReceiptItemVAT          `json:"vat,omitempty"`
	AgentInfo         *SellRefundRequestReceiptItemAgentInfo    `json:"agent_info,omitempty"`
	SupplierInfo      *SellRefundRequestReceiptItemSupplierInfo `json:"supplier_info,omitempty"`
	UserData          string                                    `json:"user_data,omitempty"`
	Excise            float64                                   `json:"excise,omitempty"`
	CountryCode       string                                    `json:"country_code,omitempty"`
	DeclarationNumber string                                    `json:"declaration_number,omitempty"`
}

type SellRefundRequestReceiptItemVAT struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRefundRequestReceiptItemAgentInfo struct {
	Type                    string                                                        `json:"type,omitempty"`
	PayingAgent             *SellRefundRequestReceiptItemAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *SellRefundRequestReceiptItemAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *SellRefundRequestReceiptItemAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type SellRefundRequestReceiptItemAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type SellRefundRequestReceiptItemAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type SellRefundRequestReceiptItemAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type SellRefundRequestReceiptItemSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type SellRefundRequestReceiptPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRefundRequestReceiptVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellRefundRequestReceiptAdditionalUserProps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SellRefundRequestReceiptCashlessPayment struct {
	Sum            float64 `json:"sum"`
	Method         int     `json:"method"`
	ID             string  `json:"id"`
	AdditionalInfo string  `json:"additional_info,omitempty"`
}

type SellRefundResponse struct {
	UUID        string                   `json:"uuid"`
	Timestamp   string                   `json:"timestamp"`
	Error       *SellRefundResponseError `json:"error"`
	Status      string                   `json:"status"`
	GroupCode   string                   `json:"group_code"`
	DaemonCode  string                   `json:"daemon_code"`
	DeviceCode  string                   `json:"device_code"`
	CallbackURL string                   `json:"callback_url"`
	Payload     any                      `json:"payload"`
}

type SellRefundResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type BuyRequest struct {
	Timestamp  string             `json:"timestamp"`
	SourceID   int                `json:"source_id,omitempty"`
	ExternalID string             `json:"external_id"`
	Service    *BuyRequestService `json:"service,omitempty"`
	Receipt    BuyRequestReceipt  `json:"receipt"`
}

type BuyRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type BuyRequestReceipt struct {
	Client               BuyRequestReceiptClient               `json:"client"`
	Company              BuyRequestReceiptCompany              `json:"company"`
	AgentInfo            *BuyRequestReceiptAgentInfo           `json:"agent_info,omitempty"`
	SupplierInfo         *BuyRequestReceiptSupplierInfo        `json:"supplier_info,omitempty"`
	Items                []BuyRequestReceiptItem               `json:"items"`
	Payments             []BuyRequestReceiptPayment            `json:"payments"`
	Vats                 []BuyRequestReceiptVat                `json:"vats,omitempty"`
	Total                float64                               `json:"total"`
	AdditionalCheckProps string                                `json:"additional_check_props,omitempty"`
	Cashier              string                                `json:"cashier,omitempty"`
	AdditionalUserProps  *BuyRequestReceiptAdditionalUserProps `json:"additional_user_props,omitempty"`
	DeviceNumber         string                                `json:"device_number,omitempty"`
	Internet             bool                                  `json:"internet,omitempty"`
	CashlessPayments     []BuyRequestReceiptCashlessPayment    `json:"cashless_payments,omitempty"`
}

type BuyRequestReceiptClient struct {
	Name  string `json:"name,omitempty"`
	INN   string `json:"inn,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type BuyRequestReceiptCompany struct {
	Email          string `json:"email,omitempty"`
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type BuyRequestReceiptAgentInfo struct {
	Type                    string                                             `json:"type,omitempty"`
	PayingAgent             *BuyRequestReceiptAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *BuyRequestReceiptAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *BuyRequestReceiptAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type BuyRequestReceiptAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type BuyRequestReceiptAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type BuyRequestReceiptAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type BuyRequestReceiptSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type BuyRequestReceiptItem struct {
	Name              string                             `json:"name"`
	Price             float64                            `json:"price"`
	Quantity          float64                            `json:"quantity"`
	Sum               *float64                           `json:"sum"`
	MeasurementUnit   string                             `json:"measurement_unit,omitempty"`
	PaymentMethod     string                             `json:"payment_method,omitempty"`
	PaymentObject     string                             `json:"payment_object,omitempty"`
	NomenclatureCode  string                             `json:"nomenclature_code,omitempty"`
	VAT               *BuyRequestReceiptItemVAT          `json:"vat,omitempty"`
	AgentInfo         *BuyRequestReceiptItemAgentInfo    `json:"agent_info,omitempty"`
	SupplierInfo      *BuyRequestReceiptItemSupplierInfo `json:"supplier_info,omitempty"`
	UserData          string                             `json:"user_data,omitempty"`
	Excise            float64                            `json:"excise,omitempty"`
	CountryCode       string                             `json:"country_code,omitempty"`
	DeclarationNumber string                             `json:"declaration_number,omitempty"`
}

type BuyRequestReceiptItemVAT struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRequestReceiptItemAgentInfo struct {
	Type                    string                                                 `json:"type,omitempty"`
	PayingAgent             *BuyRequestReceiptItemAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *BuyRequestReceiptItemAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *BuyRequestReceiptItemAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type BuyRequestReceiptItemAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type BuyRequestReceiptItemAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type BuyRequestReceiptItemAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type BuyRequestReceiptItemSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type BuyRequestReceiptPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRequestReceiptVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRequestReceiptAdditionalUserProps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BuyRequestReceiptCashlessPayment struct {
	Sum            float64 `json:"sum"`
	Method         int     `json:"method"`
	ID             string  `json:"id"`
	AdditionalInfo string  `json:"additional_info,omitempty"`
}

type BuyResponse struct {
	UUID        string            `json:"uuid"`
	Timestamp   string            `json:"timestamp"`
	Error       *BuyResponseError `json:"error"`
	Status      string            `json:"status"`
	GroupCode   string            `json:"group_code"`
	DaemonCode  string            `json:"daemon_code"`
	DeviceCode  string            `json:"device_code"`
	CallbackURL string            `json:"callback_url"`
	Payload     any               `json:"payload"`
}

type BuyResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type BuyRefundRequest struct {
	Timestamp  string                   `json:"timestamp"`
	SourceID   int                      `json:"source_id,omitempty"`
	ExternalID string                   `json:"external_id"`
	Service    *BuyRefundRequestService `json:"service,omitempty"`
	Receipt    BuyRefundRequestReceipt  `json:"receipt"`
}

type BuyRefundRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type BuyRefundRequestReceipt struct {
	Client               BuyRefundRequestReceiptClient               `json:"client"`
	Company              BuyRefundRequestReceiptCompany              `json:"company"`
	AgentInfo            *BuyRefundRequestReceiptAgentInfo           `json:"agent_info,omitempty"`
	SupplierInfo         *BuyRefundRequestReceiptSupplierInfo        `json:"supplier_info,omitempty"`
	Items                []BuyRefundRequestReceiptItem               `json:"items"`
	Payments             []BuyRefundRequestReceiptPayment            `json:"payments"`
	Vats                 []BuyRefundRequestReceiptVat                `json:"vats,omitempty"`
	Total                float64                                     `json:"total"`
	AdditionalCheckProps string                                      `json:"additional_check_props,omitempty"`
	Cashier              string                                      `json:"cashier,omitempty"`
	AdditionalUserProps  *BuyRefundRequestReceiptAdditionalUserProps `json:"additional_user_props,omitempty"`
	DeviceNumber         string                                      `json:"device_number,omitempty"`
	Internet             bool                                        `json:"internet,omitempty"`
	CashlessPayments     []BuyRefundRequestReceiptCashlessPayment    `json:"cashless_payments,omitempty"`
}

type BuyRefundRequestReceiptClient struct {
	Name  string `json:"name,omitempty"`
	INN   string `json:"inn,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type BuyRefundRequestReceiptCompany struct {
	Email          string `json:"email,omitempty"`
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type BuyRefundRequestReceiptAgentInfo struct {
	Type                    string                                                   `json:"type,omitempty"`
	PayingAgent             *BuyRefundRequestReceiptAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *BuyRefundRequestReceiptAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *BuyRefundRequestReceiptAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type BuyRefundRequestReceiptAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type BuyRefundRequestReceiptAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type BuyRefundRequestReceiptAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type BuyRefundRequestReceiptSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type BuyRefundRequestReceiptItem struct {
	Name              string                                   `json:"name"`
	Price             float64                                  `json:"price"`
	Quantity          float64                                  `json:"quantity"`
	Sum               *float64                                 `json:"sum"`
	MeasurementUnit   string                                   `json:"measurement_unit,omitempty"`
	PaymentMethod     string                                   `json:"payment_method,omitempty"`
	PaymentObject     string                                   `json:"payment_object,omitempty"`
	NomenclatureCode  string                                   `json:"nomenclature_code,omitempty"`
	VAT               *BuyRefundRequestReceiptItemVAT          `json:"vat,omitempty"`
	AgentInfo         *BuyRefundRequestReceiptItemAgentInfo    `json:"agent_info,omitempty"`
	SupplierInfo      *BuyRefundRequestReceiptItemSupplierInfo `json:"supplier_info,omitempty"`
	UserData          string                                   `json:"user_data,omitempty"`
	Excise            float64                                  `json:"excise,omitempty"`
	CountryCode       string                                   `json:"country_code,omitempty"`
	DeclarationNumber string                                   `json:"declaration_number,omitempty"`
}

type BuyRefundRequestReceiptItemVAT struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRefundRequestReceiptItemAgentInfo struct {
	Type                    string                                                       `json:"type,omitempty"`
	PayingAgent             *BuyRefundRequestReceiptItemAgentInfoPayingAgent             `json:"paying_agent,omitempty"`
	ReceivePaymentsOperator *BuyRefundRequestReceiptItemAgentInfoReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
	MoneyTransferOperator   *BuyRefundRequestReceiptItemAgentInfoMoneyTransferOperator   `json:"money_transfer_operator,omitempty"`
}

type BuyRefundRequestReceiptItemAgentInfoPayingAgent struct {
	Operation string   `json:"operation,omitempty"`
	Phones    []string `json:"phones,omitempty"`
}

type BuyRefundRequestReceiptItemAgentInfoReceivePaymentsOperator struct {
	Phones []string `json:"phones,omitempty"`
}

type BuyRefundRequestReceiptItemAgentInfoMoneyTransferOperator struct {
	Phones  []string `json:"phones,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address string   `json:"address,omitempty"`
	INN     string   `json:"inn,omitempty"`
}

type BuyRefundRequestReceiptItemSupplierInfo struct {
	Phones []string `json:"phones,omitempty"`
	Name   string   `json:"name,omitempty"`
	INN    string   `json:"inn,omitempty"`
}

type BuyRefundRequestReceiptPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRefundRequestReceiptVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyRefundRequestReceiptAdditionalUserProps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BuyRefundRequestReceiptCashlessPayment struct {
	Sum            float64 `json:"sum"`
	Method         int     `json:"method"`
	ID             string  `json:"id"`
	AdditionalInfo string  `json:"additional_info,omitempty"`
}

type BuyRefundResponse struct {
	UUID        string                  `json:"uuid"`
	Timestamp   string                  `json:"timestamp"`
	Error       *BuyRefundResponseError `json:"error"`
	Status      string                  `json:"status"`
	GroupCode   string                  `json:"group_code"`
	DaemonCode  string                  `json:"daemon_code"`
	DeviceCode  string                  `json:"device_code"`
	CallbackURL string                  `json:"callback_url"`
	Payload     any                     `json:"payload"`
}

type BuyRefundResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type SellCorrectionRequest struct {
	Timestamp  string                          `json:"timestamp"`
	SourceID   int                             `json:"source_id,omitempty"`
	ExternalID string                          `json:"external_id"`
	Service    *SellCorrectionRequestService   `json:"service,omitempty"`
	Correction SellCorrectionRequestCorrection `json:"correction"`
}

type SellCorrectionRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type SellCorrectionRequestCorrection struct {
	Company        SellCorrectionRequestCorrectionCompany        `json:"company"`
	Client         *SellCorrectionRequestCorrectionClient        `json:"client,omitempty"`
	CorrectionInfo SellCorrectionRequestCorrectionCorrectionInfo `json:"correction_info"`
	Payments       []SellCorrectionRequestCorrectionPayment      `json:"payments"`
	Vats           []SellCorrectionRequestCorrectionVat          `json:"vats"`
	Cashier        string                                        `json:"cashier,omitempty"`
	DeviceNumber   string                                        `json:"device_number,omitempty"`
	Internet       bool                                          `json:"internet,omitempty"`
}

type SellCorrectionRequestCorrectionCompany struct {
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type SellCorrectionRequestCorrectionClient struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type SellCorrectionRequestCorrectionCorrectionInfo struct {
	Type       string `json:"type"`
	BaseDate   string `json:"base_date"`
	BaseNumber string `json:"base_number,omitempty"`
}

type SellCorrectionRequestCorrectionPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellCorrectionRequestCorrectionVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type SellCorrectionResponse struct {
	UUID        string                       `json:"uuid"`
	Timestamp   string                       `json:"timestamp"`
	Error       *SellCorrectionResponseError `json:"error"`
	Status      string                       `json:"status"`
	GroupCode   string                       `json:"group_code"`
	DaemonCode  string                       `json:"daemon_code"`
	DeviceCode  string                       `json:"device_code"`
	CallbackURL string                       `json:"callback_url"`
	Payload     any                          `json:"payload"`
}

type SellCorrectionResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type BuyCorrectionRequest struct {
	Timestamp  string                         `json:"timestamp"`
	SourceID   int                            `json:"source_id,omitempty"`
	ExternalID string                         `json:"external_id"`
	Service    *BuyCorrectionRequestService   `json:"service,omitempty"`
	Correction BuyCorrectionRequestCorrection `json:"correction"`
}

type BuyCorrectionRequestService struct {
	CallbackURL string `json:"callback_url,omitempty"`
}

type BuyCorrectionRequestCorrection struct {
	Company        BuyCorrectionRequestCorrectionCompany        `json:"company"`
	Client         *BuyCorrectionRequestCorrectionClient        `json:"client,omitempty"`
	CorrectionInfo BuyCorrectionRequestCorrectionCorrectionInfo `json:"correction_info"`
	Payments       []BuyCorrectionRequestCorrectionPayment      `json:"payments"`
	Vats           []BuyCorrectionRequestCorrectionVat          `json:"vats"`
	Cashier        string                                       `json:"cashier,omitempty"`
	DeviceNumber   string                                       `json:"device_number,omitempty"`
	Internet       bool                                         `json:"internet,omitempty"`
}

type BuyCorrectionRequestCorrectionCompany struct {
	SNO            string `json:"sno,omitempty"`
	INN            string `json:"inn"`
	PaymentAddress string `json:"payment_address"`
	Location       string `json:"location,omitempty"`
}

type BuyCorrectionRequestCorrectionClient struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type BuyCorrectionRequestCorrectionCorrectionInfo struct {
	Type       string `json:"type"`
	BaseDate   string `json:"base_date"`
	BaseNumber string `json:"base_number,omitempty"`
}

type BuyCorrectionRequestCorrectionPayment struct {
	Type int      `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyCorrectionRequestCorrectionVat struct {
	Type string   `json:"type"`
	Sum  *float64 `json:"sum"`
}

type BuyCorrectionResponse struct {
	UUID        string                      `json:"uuid"`
	Timestamp   string                      `json:"timestamp"`
	Error       *BuyCorrectionResponseError `json:"error"`
	Status      string                      `json:"status"`
	GroupCode   string                      `json:"group_code"`
	DaemonCode  string                      `json:"daemon_code"`
	DeviceCode  string                      `json:"device_code"`
	CallbackURL string                      `json:"callback_url"`
	Payload     any                         `json:"payload"`
}

type BuyCorrectionResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type ReportResponse struct {
	UUID        string                  `json:"uuid"`
	Error       *ReportResponseError    `json:"error"`
	Status      string                  `json:"status"`
	Payload     *ReportResponsePayload  `json:"payload"`
	Timestamp   string                  `json:"timestamp"`
	GroupCode   string                  `json:"group_code"`
	DaemonCode  string                  `json:"daemon_code"`
	DeviceCode  string                  `json:"device_code"`
	ExternalID  string                  `json:"external_id"`
	CallbackURL string                  `json:"callback_url"`
	Warnings    *ReportResponseWarnings `json:"warnings"`
}

type ReportResponseError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type ReportResponsePayload struct {
	Total                   float64 `json:"total"`
	FNSSite                 string  `json:"fns_site"`
	FNNumber                string  `json:"fn_number"`
	ShiftNumber             int     `json:"shift_number"`
	ReceiptDateTime         string  `json:"receipt_datetime"`
	FiscalReceiptNumber     int     `json:"fiscal_receipt_number"`
	FiscalDocumentNumber    int     `json:"fiscal_document_number"`
	ECRRegistrationNumber   string  `json:"ecr_registration_number"`
	FiscalDocumentAttribute int     `json:"fiscal_document_attribute"`
	OFDINN                  string  `json:"ofd_inn"`
	OFDReceiptURL           string  `json:"ofd_receipt_url"`
}

type ReportResponseWarnings struct {
	CallbackURL string `json:"callback_url"`
}

type CallbackRequest struct {
	UUID        string                   `json:"uuid"`
	Error       *CallbackRequestError    `json:"error"`
	Status      string                   `json:"status"`
	Payload     *CallbackRequestPayload  `json:"payload"`
	Timestamp   string                   `json:"timestamp"`
	GroupCode   string                   `json:"group_code"`
	DaemonCode  string                   `json:"daemon_code"`
	DeviceCode  string                   `json:"device_code"`
	ExternalID  string                   `json:"external_id"`
	CallbackURL string                   `json:"callback_url"`
	Warnings    *CallbackRequestWarnings `json:"warnings"`
}

type CallbackRequestError struct {
	ErrorID string `json:"error_id"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Type    string `json:"type"`
}

type CallbackRequestPayload struct {
	Total                   float64 `json:"total"`
	FNSSite                 string  `json:"fns_site"`
	FNNumber                string  `json:"fn_number"`
	ShiftNumber             int     `json:"shift_number"`
	ReceiptDateTime         string  `json:"receipt_datetime"`
	FiscalReceiptNumber     int     `json:"fiscal_receipt_number"`
	FiscalDocumentNumber    int     `json:"fiscal_document_number"`
	ECRRegistrationNumber   string  `json:"ecr_registration_number"`
	FiscalDocumentAttribute int     `json:"fiscal_document_attribute"`
	OFDINN                  string  `json:"ofd_inn"`
	OFDReceiptURL           string  `json:"ofd_receipt_url"`
}

type CallbackRequestWarnings struct {
	CallbackURL string `json:"callback_url"`
}
