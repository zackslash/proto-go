// Code generated by protoc-gen-go.
// source: ftypes/finance.proto
// DO NOT EDIT!

/*
Package finance is a generated protocol buffer package.

It is generated from these files:
	ftypes/finance.proto

It has these top-level messages:
	Transaction
	BillingCycle
	PurchasedItem
*/
package finance

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import fortifi_product "github.com/fortifi/proto-go/ftypes/product"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionType int32

const (
	TransactionType_AUTH         TransactionType = 0
	TransactionType_CAPTURE_AUTH TransactionType = 1
	TransactionType_CAPTURE      TransactionType = 2
	TransactionType_CREDIT       TransactionType = 3
	TransactionType_REFUND       TransactionType = 4
	TransactionType_VOID         TransactionType = 5
	TransactionType_DISPUTE      TransactionType = 6
	TransactionType_VERIFY       TransactionType = 7
)

var TransactionType_name = map[int32]string{
	0: "AUTH",
	1: "CAPTURE_AUTH",
	2: "CAPTURE",
	3: "CREDIT",
	4: "REFUND",
	5: "VOID",
	6: "DISPUTE",
	7: "VERIFY",
}
var TransactionType_value = map[string]int32{
	"AUTH":         0,
	"CAPTURE_AUTH": 1,
	"CAPTURE":      2,
	"CREDIT":       3,
	"REFUND":       4,
	"VOID":         5,
	"DISPUTE":      6,
	"VERIFY":       7,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RefundType int32

const (
	RefundType_PARTIAL RefundType = 0
	RefundType_FULL    RefundType = 1
)

var RefundType_name = map[int32]string{
	0: "PARTIAL",
	1: "FULL",
}
var RefundType_value = map[string]int32{
	"PARTIAL": 0,
	"FULL":    1,
}

func (x RefundType) String() string {
	return proto.EnumName(RefundType_name, int32(x))
}
func (RefundType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RefundReason int32

const (
	RefundReason_REQUESTED  RefundReason = 0
	RefundReason_DUPLICATE  RefundReason = 1
	RefundReason_FRAUDULENT RefundReason = 2
)

var RefundReason_name = map[int32]string{
	0: "REQUESTED",
	1: "DUPLICATE",
	2: "FRAUDULENT",
}
var RefundReason_value = map[string]int32{
	"REQUESTED":  0,
	"DUPLICATE":  1,
	"FRAUDULENT": 2,
}

func (x RefundReason) String() string {
	return proto.EnumName(RefundReason_name, int32(x))
}
func (RefundReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DisputeStatus int32

const (
	DisputeStatus_OPEN DisputeStatus = 0
	DisputeStatus_WON  DisputeStatus = 1
	DisputeStatus_LOST DisputeStatus = 2
)

var DisputeStatus_name = map[int32]string{
	0: "OPEN",
	1: "WON",
	2: "LOST",
}
var DisputeStatus_value = map[string]int32{
	"OPEN": 0,
	"WON":  1,
	"LOST": 2,
}

func (x DisputeStatus) String() string {
	return proto.EnumName(DisputeStatus_name, int32(x))
}
func (DisputeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DisputeReason int32

const (
	DisputeReason_FRAUD        DisputeReason = 0
	DisputeReason_UNRECOGNIZED DisputeReason = 1
)

var DisputeReason_name = map[int32]string{
	0: "FRAUD",
	1: "UNRECOGNIZED",
}
var DisputeReason_value = map[string]int32{
	"FRAUD":        0,
	"UNRECOGNIZED": 1,
}

func (x DisputeReason) String() string {
	return proto.EnumName(DisputeReason_name, int32(x))
}
func (DisputeReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PaymentMethod int32

const (
	PaymentMethod_CASH            PaymentMethod = 0
	PaymentMethod_DEBIT_CARD      PaymentMethod = 1
	PaymentMethod_CREDIT_CARD     PaymentMethod = 2
	PaymentMethod_PREPAID_CARD    PaymentMethod = 3
	PaymentMethod_CHEQUE          PaymentMethod = 4
	PaymentMethod_BANK_TRANSFER   PaymentMethod = 5
	PaymentMethod_STANDING_ORDER  PaymentMethod = 6
	PaymentMethod_DIRECT_DEBIT    PaymentMethod = 7
	PaymentMethod_DIGITAL_WALLET  PaymentMethod = 8
	PaymentMethod_MOBILE_PAYMENT  PaymentMethod = 9
	PaymentMethod_BITCOIN         PaymentMethod = 10
	PaymentMethod_COUPON          PaymentMethod = 11
	PaymentMethod_CREDIT_NOTE     PaymentMethod = 12
	PaymentMethod_ACCOUNT_BALANCE PaymentMethod = 13
	PaymentMethod_DIGITAL_STORE   PaymentMethod = 14
)

var PaymentMethod_name = map[int32]string{
	0:  "CASH",
	1:  "DEBIT_CARD",
	2:  "CREDIT_CARD",
	3:  "PREPAID_CARD",
	4:  "CHEQUE",
	5:  "BANK_TRANSFER",
	6:  "STANDING_ORDER",
	7:  "DIRECT_DEBIT",
	8:  "DIGITAL_WALLET",
	9:  "MOBILE_PAYMENT",
	10: "BITCOIN",
	11: "COUPON",
	12: "CREDIT_NOTE",
	13: "ACCOUNT_BALANCE",
	14: "DIGITAL_STORE",
}
var PaymentMethod_value = map[string]int32{
	"CASH":            0,
	"DEBIT_CARD":      1,
	"CREDIT_CARD":     2,
	"PREPAID_CARD":    3,
	"CHEQUE":          4,
	"BANK_TRANSFER":   5,
	"STANDING_ORDER":  6,
	"DIRECT_DEBIT":    7,
	"DIGITAL_WALLET":  8,
	"MOBILE_PAYMENT":  9,
	"BITCOIN":         10,
	"COUPON":          11,
	"CREDIT_NOTE":     12,
	"ACCOUNT_BALANCE": 13,
	"DIGITAL_STORE":   14,
}

func (x PaymentMethod) String() string {
	return proto.EnumName(PaymentMethod_name, int32(x))
}
func (PaymentMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type BillingCycleDay int32

const (
	BillingCycleDay_SYNC_START   BillingCycleDay = 0
	BillingCycleDay_SYNC_END     BillingCycleDay = 1
	BillingCycleDay_RESYNC_END   BillingCycleDay = 2
	BillingCycleDay_SPECIFIC_DAY BillingCycleDay = 3
	BillingCycleDay_DAY_OF_WEEK  BillingCycleDay = 4
	BillingCycleDay_MONTH_END    BillingCycleDay = 5
)

var BillingCycleDay_name = map[int32]string{
	0: "SYNC_START",
	1: "SYNC_END",
	2: "RESYNC_END",
	3: "SPECIFIC_DAY",
	4: "DAY_OF_WEEK",
	5: "MONTH_END",
}
var BillingCycleDay_value = map[string]int32{
	"SYNC_START":   0,
	"SYNC_END":     1,
	"RESYNC_END":   2,
	"SPECIFIC_DAY": 3,
	"DAY_OF_WEEK":  4,
	"MONTH_END":    5,
}

func (x BillingCycleDay) String() string {
	return proto.EnumName(BillingCycleDay_name, int32(x))
}
func (BillingCycleDay) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type BillingCyclePeriod int32

const (
	BillingCyclePeriod_NONE  BillingCyclePeriod = 0
	BillingCyclePeriod_DAY   BillingCyclePeriod = 1
	BillingCyclePeriod_WEEK  BillingCyclePeriod = 2
	BillingCyclePeriod_MONTH BillingCyclePeriod = 3
	BillingCyclePeriod_YEAR  BillingCyclePeriod = 4
)

var BillingCyclePeriod_name = map[int32]string{
	0: "NONE",
	1: "DAY",
	2: "WEEK",
	3: "MONTH",
	4: "YEAR",
}
var BillingCyclePeriod_value = map[string]int32{
	"NONE":  0,
	"DAY":   1,
	"WEEK":  2,
	"MONTH": 3,
	"YEAR":  4,
}

func (x BillingCyclePeriod) String() string {
	return proto.EnumName(BillingCyclePeriod_name, int32(x))
}
func (BillingCyclePeriod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type PurchaseStatus int32

const (
	PurchaseStatus_UNKNOWN_PURCHASE_STATUS PurchaseStatus = 0
	PurchaseStatus_INTERESTED              PurchaseStatus = 1
	PurchaseStatus_CHECKOUT                PurchaseStatus = 2
	PurchaseStatus_SETUP                   PurchaseStatus = 3
	PurchaseStatus_ACTIVE                  PurchaseStatus = 4
	PurchaseStatus_INACTIVE                PurchaseStatus = 5
	PurchaseStatus_SUSPENDED               PurchaseStatus = 6
	PurchaseStatus_CANCELLED               PurchaseStatus = 7
	PurchaseStatus_ENDED                   PurchaseStatus = 8
	PurchaseStatus_TRIAL                   PurchaseStatus = 9
)

var PurchaseStatus_name = map[int32]string{
	0: "UNKNOWN_PURCHASE_STATUS",
	1: "INTERESTED",
	2: "CHECKOUT",
	3: "SETUP",
	4: "ACTIVE",
	5: "INACTIVE",
	6: "SUSPENDED",
	7: "CANCELLED",
	8: "ENDED",
	9: "TRIAL",
}
var PurchaseStatus_value = map[string]int32{
	"UNKNOWN_PURCHASE_STATUS": 0,
	"INTERESTED":              1,
	"CHECKOUT":                2,
	"SETUP":                   3,
	"ACTIVE":                  4,
	"INACTIVE":                5,
	"SUSPENDED":               6,
	"CANCELLED":               7,
	"ENDED":                   8,
	"TRIAL":                   9,
}

func (x PurchaseStatus) String() string {
	return proto.EnumName(PurchaseStatus_name, int32(x))
}
func (PurchaseStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AgreementAction int32

const (
	AgreementAction_CREATE    AgreementAction = 0
	AgreementAction_ACCEPTE   AgreementAction = 1
	AgreementAction_SUSPEND   AgreementAction = 2
	AgreementAction_UNSUSPEND AgreementAction = 3
	AgreementAction_CANCEL    AgreementAction = 4
	AgreementAction_RESTORE   AgreementAction = 5
	AgreementAction_TERMINATE AgreementAction = 6
)

var AgreementAction_name = map[int32]string{
	0: "CREATE",
	1: "ACCEPTE",
	2: "SUSPEND",
	3: "UNSUSPEND",
	4: "CANCEL",
	5: "RESTORE",
	6: "TERMINATE",
}
var AgreementAction_value = map[string]int32{
	"CREATE":    0,
	"ACCEPTE":   1,
	"SUSPEND":   2,
	"UNSUSPEND": 3,
	"CANCEL":    4,
	"RESTORE":   5,
	"TERMINATE": 6,
}

func (x AgreementAction) String() string {
	return proto.EnumName(AgreementAction_name, int32(x))
}
func (AgreementAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Transaction struct {
	Id                     string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                   TransactionType            `protobuf:"varint,2,opt,name=type,enum=fortifi.finance.TransactionType" json:"type,omitempty"`
	Currency               string                     `protobuf:"bytes,3,opt,name=currency" json:"currency,omitempty"`
	RawAmount              int64                      `protobuf:"varint,4,opt,name=raw_amount,json=rawAmount" json:"raw_amount,omitempty"`
	SetupAmount            int64                      `protobuf:"varint,5,opt,name=setup_amount,json=setupAmount" json:"setup_amount,omitempty"`
	ShippingAmount         int64                      `protobuf:"varint,6,opt,name=shipping_amount,json=shippingAmount" json:"shipping_amount,omitempty"`
	DiscountAmount         int64                      `protobuf:"varint,7,opt,name=discount_amount,json=discountAmount" json:"discount_amount,omitempty"`
	SetupDiscountAmount    int64                      `protobuf:"varint,8,opt,name=setup_discount_amount,json=setupDiscountAmount" json:"setup_discount_amount,omitempty"`
	ShippingDiscountAmount int64                      `protobuf:"varint,9,opt,name=shipping_discount_amount,json=shippingDiscountAmount" json:"shipping_discount_amount,omitempty"`
	TaxAmount              int64                      `protobuf:"varint,10,opt,name=tax_amount,json=taxAmount" json:"tax_amount,omitempty"`
	TotalAmount            int64                      `protobuf:"varint,11,opt,name=total_amount,json=totalAmount" json:"total_amount,omitempty"`
	TotalDiscountAmount    int64                      `protobuf:"varint,12,opt,name=total_discount_amount,json=totalDiscountAmount" json:"total_discount_amount,omitempty"`
	PaymentMethod          PaymentMethod              `protobuf:"varint,13,opt,name=payment_method,json=paymentMethod,enum=fortifi.finance.PaymentMethod" json:"payment_method,omitempty"`
	PaymentMethodDetail    string                     `protobuf:"bytes,14,opt,name=payment_method_detail,json=paymentMethodDetail" json:"payment_method_detail,omitempty"`
	GatewayId              string                     `protobuf:"bytes,15,opt,name=gateway_id,json=gatewayId" json:"gateway_id,omitempty"`
	Timestamp              *google_protobuf.Timestamp `protobuf:"bytes,16,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transaction) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type BillingCycle struct {
	Period    BillingCyclePeriod `protobuf:"varint,1,opt,name=period,enum=fortifi.finance.BillingCyclePeriod" json:"period,omitempty"`
	Frequency int32              `protobuf:"varint,2,opt,name=frequency" json:"frequency,omitempty"`
	Day       BillingCycleDay    `protobuf:"varint,3,opt,name=day,enum=fortifi.finance.BillingCycleDay" json:"day,omitempty"`
	DayValue  string             `protobuf:"bytes,4,opt,name=day_value,json=dayValue" json:"day_value,omitempty"`
}

func (m *BillingCycle) Reset()                    { *m = BillingCycle{} }
func (m *BillingCycle) String() string            { return proto.CompactTextString(m) }
func (*BillingCycle) ProtoMessage()               {}
func (*BillingCycle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PurchasedItem struct {
	Id               string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ProductCode      string                      `protobuf:"bytes,2,opt,name=product_code,json=productCode" json:"product_code,omitempty"`
	Sku              string                      `protobuf:"bytes,3,opt,name=sku" json:"sku,omitempty"`
	Name             string                      `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	UnitPrice        int64                       `protobuf:"varint,5,opt,name=unit_price,json=unitPrice" json:"unit_price,omitempty"`
	Quantity         int32                       `protobuf:"varint,6,opt,name=quantity" json:"quantity,omitempty"`
	ProductType      fortifi_product.ProductType `protobuf:"varint,7,opt,name=product_type,json=productType,enum=fortifi.product.ProductType" json:"product_type,omitempty"`
	Expiry           *google_protobuf.Timestamp  `protobuf:"bytes,9,opt,name=expiry" json:"expiry,omitempty"`
	StartDate        *google_protobuf.Timestamp  `protobuf:"bytes,10,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	Renews           bool                        `protobuf:"varint,11,opt,name=renews" json:"renews,omitempty"`
	ParentPurchaseId string                      `protobuf:"bytes,12,opt,name=parent_purchase_id,json=parentPurchaseId" json:"parent_purchase_id,omitempty"`
}

func (m *PurchasedItem) Reset()                    { *m = PurchasedItem{} }
func (m *PurchasedItem) String() string            { return proto.CompactTextString(m) }
func (*PurchasedItem) ProtoMessage()               {}
func (*PurchasedItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PurchasedItem) GetExpiry() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *PurchasedItem) GetStartDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "fortifi.finance.Transaction")
	proto.RegisterType((*BillingCycle)(nil), "fortifi.finance.BillingCycle")
	proto.RegisterType((*PurchasedItem)(nil), "fortifi.finance.PurchasedItem")
	proto.RegisterEnum("fortifi.finance.TransactionType", TransactionType_name, TransactionType_value)
	proto.RegisterEnum("fortifi.finance.RefundType", RefundType_name, RefundType_value)
	proto.RegisterEnum("fortifi.finance.RefundReason", RefundReason_name, RefundReason_value)
	proto.RegisterEnum("fortifi.finance.DisputeStatus", DisputeStatus_name, DisputeStatus_value)
	proto.RegisterEnum("fortifi.finance.DisputeReason", DisputeReason_name, DisputeReason_value)
	proto.RegisterEnum("fortifi.finance.PaymentMethod", PaymentMethod_name, PaymentMethod_value)
	proto.RegisterEnum("fortifi.finance.BillingCycleDay", BillingCycleDay_name, BillingCycleDay_value)
	proto.RegisterEnum("fortifi.finance.BillingCyclePeriod", BillingCyclePeriod_name, BillingCyclePeriod_value)
	proto.RegisterEnum("fortifi.finance.PurchaseStatus", PurchaseStatus_name, PurchaseStatus_value)
	proto.RegisterEnum("fortifi.finance.AgreementAction", AgreementAction_name, AgreementAction_value)
}

func init() { proto.RegisterFile("ftypes/finance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0xed, 0xd8, 0xb1, 0xd6, 0x7f, 0x42, 0x85, 0xe2, 0x49, 0x0b, 0x94, 0xf6, 0x02, 0xc6,
	0x53, 0x9c, 0x99, 0xc0, 0x05, 0x0c, 0xcc, 0x30, 0x8a, 0xb4, 0x6e, 0x34, 0x71, 0x56, 0x62, 0xb5,
	0x4a, 0x26, 0xbd, 0xd1, 0xa8, 0xb6, 0xec, 0x68, 0xb0, 0x25, 0x55, 0x96, 0x68, 0xfd, 0x0c, 0x3c,
	0x01, 0xf7, 0x3c, 0x05, 0xbc, 0x1c, 0x67, 0x77, 0xa5, 0x24, 0x76, 0x18, 0xb8, 0xb2, 0xce, 0xb7,
	0xdf, 0xf9, 0xd9, 0xef, 0x9c, 0xb3, 0x63, 0xf4, 0xf1, 0x22, 0xdf, 0xa6, 0xe1, 0xe6, 0x78, 0x11,
	0xc5, 0x41, 0x3c, 0x0b, 0xc7, 0x69, 0x96, 0xe4, 0x89, 0x36, 0x58, 0x24, 0x59, 0x1e, 0x2d, 0xa2,
	0x71, 0x09, 0x1f, 0x7d, 0xb1, 0x4c, 0x92, 0xe5, 0x2a, 0x3c, 0x16, 0xc7, 0x6f, 0x8b, 0xc5, 0x71,
	0x1e, 0xad, 0xc3, 0x4d, 0x1e, 0xac, 0x53, 0xe9, 0x71, 0x54, 0xc5, 0x01, 0x6b, 0x5e, 0xcc, 0x72,
	0x89, 0xbe, 0xf8, 0xbb, 0x89, 0x3a, 0x2c, 0x0b, 0xe2, 0x4d, 0x30, 0xcb, 0xa3, 0x24, 0xd6, 0xfa,
	0xa8, 0x1e, 0xcd, 0x87, 0xb5, 0xe7, 0xb5, 0xaf, 0x15, 0x0a, 0x5f, 0xda, 0x77, 0xe8, 0x80, 0xbb,
	0x0d, 0xeb, 0x80, 0xf4, 0x4f, 0x9e, 0x8f, 0xf7, 0xd2, 0x8e, 0xef, 0xf9, 0x32, 0xe0, 0x51, 0xc1,
	0xd6, 0x8e, 0x50, 0x7b, 0x56, 0x64, 0x59, 0x18, 0xcf, 0xb6, 0xc3, 0x86, 0x88, 0x75, 0x6b, 0x6b,
	0x9f, 0x21, 0x94, 0x05, 0xef, 0xfd, 0x60, 0x9d, 0x14, 0x71, 0x3e, 0x3c, 0x80, 0xd3, 0x06, 0x55,
	0x00, 0xd1, 0x05, 0xa0, 0x7d, 0x89, 0xba, 0x9b, 0x30, 0x2f, 0xd2, 0x8a, 0xd0, 0x14, 0x84, 0x8e,
	0xc0, 0x4a, 0xca, 0x57, 0x68, 0xb0, 0xb9, 0x89, 0xd2, 0x34, 0x8a, 0x97, 0x15, 0xab, 0x25, 0x58,
	0xfd, 0x0a, 0xbe, 0x23, 0xce, 0xa3, 0xcd, 0x8c, 0x7f, 0x57, 0xc4, 0x43, 0x49, 0xac, 0xe0, 0x92,
	0x78, 0x82, 0x3e, 0x91, 0x49, 0xf7, 0xe9, 0x6d, 0x41, 0x7f, 0x2c, 0x0e, 0xcd, 0x5d, 0x9f, 0xef,
	0xd1, 0xf0, 0xb6, 0x8a, 0x7d, 0x37, 0x45, 0xb8, 0x3d, 0xa9, 0xce, 0xf7, 0x3c, 0x41, 0x81, 0x3c,
	0xf8, 0x50, 0x71, 0x91, 0x54, 0x00, 0x90, 0x3b, 0x05, 0xf2, 0x24, 0x0f, 0x56, 0x15, 0xa1, 0x23,
	0x15, 0x10, 0xd8, 0x5d, 0xbd, 0x92, 0xb2, 0x9f, 0xb8, 0x2b, 0xeb, 0x15, 0x87, 0x7b, 0x59, 0x31,
	0xea, 0xa7, 0xc1, 0x76, 0x1d, 0x02, 0x79, 0x1d, 0xe6, 0x37, 0xc9, 0x7c, 0xd8, 0x13, 0x3d, 0xfd,
	0xfc, 0x41, 0x4f, 0x1d, 0x49, 0xbb, 0x10, 0x2c, 0xda, 0x4b, 0xef, 0x9b, 0x3c, 0xf5, 0x6e, 0x18,
	0x7f, 0x1e, 0xe6, 0x41, 0xb4, 0x1a, 0xf6, 0x45, 0x9f, 0x1f, 0xef, 0xb0, 0x4d, 0x71, 0xc4, 0x2f,
	0xbc, 0x0c, 0xf2, 0xf0, 0x7d, 0xb0, 0xf5, 0x61, 0xb8, 0x06, 0x82, 0xa8, 0x94, 0x88, 0x35, 0x07,
	0x25, 0x95, 0xdb, 0x61, 0x1d, 0xaa, 0x70, 0xda, 0x39, 0x39, 0x1a, 0xcb, 0x71, 0x1e, 0x57, 0xe3,
	0x3c, 0x66, 0x15, 0x83, 0xde, 0x91, 0x5f, 0xfc, 0x55, 0x43, 0xdd, 0xd3, 0x68, 0xb5, 0x02, 0x8d,
	0x8d, 0xed, 0x6c, 0x15, 0x6a, 0x3f, 0xa2, 0x56, 0x1a, 0x66, 0x51, 0x22, 0x47, 0xb8, 0x7f, 0xf2,
	0xf2, 0xc1, 0xe5, 0xee, 0xd3, 0x1d, 0x41, 0xa5, 0xa5, 0x8b, 0xf6, 0x0c, 0x29, 0x8b, 0x2c, 0x7c,
	0x57, 0x88, 0xb1, 0xe5, 0x03, 0xdf, 0xa4, 0x77, 0x00, 0x5c, 0xbc, 0x31, 0x0f, 0xe4, 0x38, 0xff,
	0xdb, 0x22, 0xdc, 0x8f, 0x6b, 0x06, 0x5b, 0xca, 0xc9, 0xda, 0x53, 0xa4, 0xc0, 0x8f, 0xff, 0x5b,
	0xb0, 0x2a, 0x42, 0x31, 0xea, 0xb0, 0x08, 0x00, 0x5c, 0x72, 0xfb, 0xc5, 0x1f, 0x0d, 0xd4, 0x73,
	0x8a, 0x6c, 0x76, 0x13, 0x6c, 0xc2, 0xb9, 0x95, 0x87, 0xeb, 0x07, 0xcb, 0x07, 0x93, 0x50, 0x6e,
	0xab, 0x3f, 0x4b, 0xe6, 0x72, 0x09, 0x15, 0xda, 0x29, 0x31, 0x03, 0x20, 0x4d, 0x45, 0x8d, 0xcd,
	0xaf, 0x45, 0xb9, 0x64, 0xfc, 0x53, 0xd3, 0xd0, 0x41, 0x1c, 0xac, 0xab, 0x74, 0xe2, 0x9b, 0x37,
	0xa0, 0x88, 0xa3, 0xdc, 0x4f, 0xb3, 0x68, 0x16, 0x96, 0x2b, 0xa5, 0x70, 0xc4, 0xe1, 0x00, 0x5f,
	0xd7, 0x77, 0x45, 0x10, 0xe7, 0x51, 0xbe, 0x15, 0x9b, 0xd4, 0xa4, 0xb7, 0xb6, 0xf6, 0xf3, 0x5d,
	0x0d, 0xe2, 0x21, 0x38, 0x14, 0xf7, 0x7f, 0x76, 0x7b, 0xff, 0xea, 0x39, 0x71, 0xe4, 0xaf, 0x78,
	0x04, 0xaa, 0x0a, 0xb9, 0x01, 0xba, 0xb5, 0xc2, 0x0f, 0x69, 0x94, 0x6d, 0xc5, 0x56, 0xfc, 0x77,
	0x6b, 0x4b, 0xa6, 0xf6, 0x03, 0x42, 0x00, 0x64, 0xb9, 0x3f, 0x87, 0x21, 0x11, 0x1b, 0xf2, 0x3f,
	0x23, 0x21, 0xd8, 0x26, 0x90, 0xb5, 0x27, 0xa8, 0x05, 0xef, 0x4c, 0xf8, 0x7e, 0x23, 0xf6, 0xa6,
	0x4d, 0x4b, 0x4b, 0x7b, 0x85, 0xb4, 0x34, 0xc8, 0xf8, 0xd8, 0xa6, 0xa5, 0xe6, 0x7c, 0x16, 0xbb,
	0x42, 0x24, 0x55, 0x9e, 0x54, 0xcd, 0xb0, 0xe6, 0xa3, 0x02, 0x0d, 0xf6, 0x5e, 0x36, 0xad, 0x8d,
	0x0e, 0x74, 0x8f, 0x9d, 0xa9, 0x8f, 0x40, 0xf3, 0xae, 0xa1, 0x3b, 0xcc, 0xa3, 0xd8, 0x17, 0x48,
	0x4d, 0xeb, 0xa0, 0xc3, 0x12, 0x51, 0xeb, 0x1a, 0x42, 0x2d, 0x83, 0x62, 0xd3, 0x62, 0x6a, 0x83,
	0x7f, 0x53, 0x3c, 0xf1, 0x88, 0xa9, 0x1e, 0xf0, 0x00, 0x97, 0xb6, 0x65, 0xaa, 0x4d, 0x4e, 0x37,
	0x2d, 0xd7, 0xf1, 0x18, 0x56, 0x5b, 0x9c, 0x72, 0x89, 0xa9, 0x35, 0xb9, 0x56, 0x0f, 0x47, 0x2f,
	0x11, 0xa2, 0xe1, 0xa2, 0x88, 0xe7, 0x22, 0x23, 0xd0, 0x1c, 0x9d, 0x32, 0x4b, 0x9f, 0x42, 0x52,
	0xf0, 0x9e, 0x78, 0xd3, 0xa9, 0x5a, 0x1b, 0xfd, 0x84, 0xba, 0x92, 0x44, 0xc3, 0x60, 0x03, 0x4f,
	0x76, 0x0f, 0x29, 0x14, 0xff, 0xe2, 0x61, 0x97, 0x61, 0x13, 0x88, 0x60, 0x9a, 0x9e, 0x33, 0xb5,
	0x0c, 0x1d, 0xc2, 0xd7, 0x60, 0xa6, 0xd0, 0x84, 0xea, 0x9e, 0xe9, 0x4d, 0x31, 0x61, 0x6a, 0x7d,
	0xf4, 0x0a, 0xf5, 0xe0, 0x61, 0x48, 0x8b, 0x3c, 0x74, 0xf3, 0x20, 0x2f, 0x36, 0x3c, 0xb0, 0xed,
	0x60, 0x02, 0x9e, 0x87, 0xa8, 0x71, 0x65, 0x13, 0xf0, 0x01, 0x68, 0x6a, 0xbb, 0xbb, 0xec, 0x32,
	0x99, 0x82, 0x9a, 0x22, 0x9c, 0x94, 0xc1, 0x23, 0x14, 0x1b, 0xf6, 0x6b, 0x62, 0xbd, 0x81, 0xd4,
	0xb5, 0xd1, 0xef, 0x75, 0x98, 0xe8, 0x9d, 0xd7, 0x02, 0x22, 0x19, 0xba, 0xcb, 0x45, 0x83, 0x3a,
	0x4c, 0x7c, 0x6a, 0x31, 0xdf, 0xd0, 0x29, 0x70, 0xb5, 0x01, 0xea, 0x48, 0x95, 0x24, 0x50, 0xe7,
	0xe1, 0x1c, 0x8a, 0x1d, 0xdd, 0x32, 0x25, 0x22, 0xc4, 0x33, 0xce, 0xf8, 0xcd, 0x40, 0xbc, 0x8f,
	0x50, 0xef, 0x54, 0x27, 0xe7, 0x3e, 0xa3, 0x3a, 0x71, 0x27, 0x98, 0x82, 0x8a, 0x1a, 0xea, 0xbb,
	0x4c, 0x27, 0xa6, 0x45, 0x5e, 0xfb, 0x36, 0x35, 0x01, 0x6b, 0xf1, 0x20, 0xa6, 0x05, 0x35, 0x31,
	0x5f, 0x24, 0x53, 0x0f, 0x39, 0xcb, 0xb4, 0x5e, 0x5b, 0x4c, 0x9f, 0xfa, 0x57, 0xfa, 0x74, 0x8a,
	0x99, 0xda, 0xe6, 0xd8, 0x85, 0x7d, 0x6a, 0x4d, 0xb1, 0xef, 0xe8, 0xd7, 0x17, 0x5c, 0x17, 0x85,
	0x8b, 0x0d, 0x0e, 0x86, 0x6d, 0x11, 0x15, 0x89, 0xcc, 0xb6, 0xe7, 0x80, 0x18, 0x9d, 0x7b, 0x85,
	0x12, 0x1b, 0x14, 0xed, 0x6a, 0x8f, 0xd1, 0x40, 0x37, 0xe0, 0x98, 0x30, 0xff, 0x54, 0x9f, 0xea,
	0xc4, 0xc0, 0x6a, 0x8f, 0xd7, 0x57, 0xa5, 0x71, 0x99, 0x0d, 0x73, 0xd0, 0x1f, 0xbd, 0x43, 0x83,
	0xbd, 0x47, 0x81, 0x8b, 0xe0, 0x5e, 0x13, 0x03, 0x28, 0xd0, 0x57, 0x10, 0xa5, 0x8b, 0xda, 0xc2,
	0xc6, 0xc4, 0x94, 0xad, 0xa2, 0xf8, 0xd6, 0x16, 0x8a, 0xb8, 0x0e, 0x36, 0xac, 0x89, 0x65, 0xf8,
	0xa6, 0x7e, 0x0d, 0x8a, 0x40, 0x2d, 0xf0, 0xe1, 0xdb, 0x13, 0xff, 0x0a, 0xe3, 0x73, 0x90, 0x05,
	0x9a, 0x7d, 0x61, 0x13, 0x76, 0x26, 0x3c, 0x9a, 0xa3, 0x09, 0xd2, 0x1e, 0xbe, 0x6f, 0xbc, 0x09,
	0xc4, 0x26, 0x58, 0x76, 0x98, 0x07, 0x12, 0x1d, 0x16, 0x11, 0xea, 0xbc, 0xa1, 0x22, 0x02, 0x44,
	0x07, 0xf0, 0x1a, 0xeb, 0x54, 0x3d, 0x18, 0xfd, 0x59, 0x43, 0xfd, 0x6a, 0x1b, 0xca, 0x31, 0x79,
	0x8a, 0x3e, 0xf5, 0xc8, 0x39, 0xb1, 0xaf, 0x88, 0xef, 0x78, 0xd4, 0x38, 0xd3, 0x5d, 0xcc, 0xaf,
	0xc1, 0x3c, 0x57, 0x36, 0xd7, 0x22, 0x0c, 0x53, 0x39, 0x83, 0x35, 0x7e, 0x2f, 0xe8, 0x9c, 0x71,
	0x6e, 0x7b, 0x4c, 0xa6, 0x70, 0x31, 0xf3, 0x1c, 0xd9, 0x52, 0xdd, 0x60, 0xd6, 0x25, 0x6f, 0x29,
	0x90, 0x2c, 0x52, 0x5a, 0x4d, 0x7e, 0x13, 0xd7, 0x83, 0xeb, 0x12, 0x13, 0x22, 0xb4, 0xb8, 0x69,
	0x70, 0x69, 0xa1, 0x63, 0x26, 0x74, 0x11, 0x42, 0xc8, 0x93, 0x36, 0xff, 0x64, 0x94, 0xef, 0x84,
	0x32, 0x8a, 0xa1, 0x13, 0xcb, 0x2c, 0x0c, 0xf9, 0xc0, 0xe9, 0xf2, 0xff, 0x8b, 0x5c, 0x3e, 0x3e,
	0xfa, 0x8f, 0x78, 0x4b, 0xa1, 0x51, 0xd8, 0x11, 0x7b, 0x00, 0x46, 0x19, 0x1f, 0x2a, 0x82, 0xe8,
	0x1e, 0xa9, 0x4c, 0x39, 0x68, 0x22, 0x19, 0x54, 0x05, 0x3c, 0x7e, 0x0d, 0xde, 0x42, 0x51, 0x14,
	0xdc, 0xea, 0x02, 0xca, 0xe4, 0xab, 0x7a, 0xfa, 0xea, 0xcd, 0x68, 0x19, 0xe5, 0x37, 0xc5, 0xdb,
	0xf1, 0x2c, 0x59, 0x1f, 0x97, 0x2f, 0xa0, 0xfc, 0xc7, 0xf5, 0xcd, 0x32, 0x39, 0xde, 0xfd, 0xa3,
	0xf6, 0xb6, 0x25, 0x0e, 0xbe, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x53, 0x30, 0xe2, 0xc1,
	0x09, 0x00, 0x00,
}
