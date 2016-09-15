// Code generated by protoc-gen-go.
// source: finance.proto
// DO NOT EDIT!

/*
Package finance is a generated protocol buffer package.

It is generated from these files:
	finance.proto

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

func init() { proto.RegisterFile("finance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0x8d, 0x24, 0x4b, 0x16, 0x57, 0x7f, 0xfc, 0x18, 0x7c, 0xa9, 0xe0, 0xa4, 0x6d, 0x9a, 0x5c,
	0xb4, 0x10, 0x52, 0x19, 0x70, 0x7b, 0xd1, 0xa2, 0x05, 0x0a, 0x9a, 0x5c, 0xc5, 0x84, 0xe5, 0x25,
	0xbb, 0x5c, 0xda, 0x70, 0x6e, 0x08, 0x46, 0xa2, 0x64, 0xa2, 0x12, 0xc9, 0x50, 0x64, 0x1d, 0x3d,
	0x43, 0x9f, 0xa0, 0xf7, 0x7d, 0x8a, 0xf6, 0xe5, 0x3a, 0xbb, 0x4b, 0x5a, 0xb6, 0x5c, 0xb4, 0x57,
	0xdc, 0x39, 0x7b, 0x76, 0x66, 0xf6, 0xcc, 0xcc, 0x4a, 0xa8, 0xb7, 0x88, 0xe2, 0x20, 0x9e, 0x85,
	0xe3, 0x34, 0x4b, 0xf2, 0x44, 0x1b, 0x2c, 0x92, 0x2c, 0x8f, 0x16, 0xd1, 0xb8, 0x84, 0x8f, 0x3e,
	0x5f, 0x26, 0xc9, 0x72, 0x15, 0x1e, 0x8b, 0xed, 0xf7, 0xc5, 0xe2, 0x38, 0x8f, 0xd6, 0xe1, 0x26,
	0x0f, 0xd6, 0xa9, 0x3c, 0x71, 0xd4, 0x83, 0xcf, 0xbc, 0x98, 0xe5, 0xd2, 0x7c, 0xf5, 0x57, 0x13,
	0x75, 0x58, 0x16, 0xc4, 0x9b, 0x60, 0x96, 0x47, 0x49, 0xac, 0xf5, 0x51, 0x3d, 0x9a, 0x0f, 0x6b,
	0x2f, 0x6b, 0x5f, 0x29, 0x14, 0x56, 0xda, 0xb7, 0xe8, 0x20, 0xdf, 0xa6, 0xe1, 0xb0, 0x0e, 0x48,
	0xff, 0xe4, 0xe5, 0x78, 0x2f, 0xde, 0xf8, 0xde, 0x59, 0x06, 0x3c, 0x2a, 0xd8, 0xda, 0x11, 0x6a,
	0xcf, 0x8a, 0x2c, 0x0b, 0xe3, 0xd9, 0x76, 0xd8, 0x10, 0xbe, 0xee, 0x6c, 0xed, 0x53, 0x84, 0xb2,
	0xe0, 0xd6, 0x0f, 0xd6, 0x49, 0x11, 0xe7, 0xc3, 0x03, 0xd8, 0x6d, 0x50, 0x05, 0x10, 0x5d, 0x00,
	0xda, 0x17, 0xa8, 0xbb, 0x09, 0xf3, 0x22, 0xad, 0x08, 0x4d, 0x41, 0xe8, 0x08, 0xac, 0xa4, 0x7c,
	0x89, 0x06, 0x9b, 0x9b, 0x28, 0x4d, 0xa3, 0x78, 0x59, 0xb1, 0x5a, 0x82, 0xd5, 0xaf, 0xe0, 0x1d,
	0x71, 0x1e, 0x6d, 0x66, 0x7c, 0x5d, 0x11, 0x0f, 0x25, 0xb1, 0x82, 0x4b, 0xe2, 0x09, 0xfa, 0xbf,
	0x0c, 0xba, 0x4f, 0x6f, 0x0b, 0xfa, 0x53, 0xb1, 0x69, 0x3e, 0x3c, 0xf3, 0x1d, 0x1a, 0xde, 0x65,
	0xb1, 0x7f, 0x4c, 0x11, 0xc7, 0x9e, 0x55, 0xfb, 0x7b, 0x27, 0x41, 0x81, 0x3c, 0xf8, 0x58, 0x71,
	0x91, 0x54, 0x00, 0x90, 0x9d, 0x02, 0x79, 0x92, 0x07, 0xab, 0x8a, 0xd0, 0x91, 0x0a, 0x08, 0x6c,
	0x97, 0xaf, 0xa4, 0xec, 0x07, 0xee, 0xca, 0x7c, 0xc5, 0xe6, 0x5e, 0x54, 0x8c, 0xfa, 0x69, 0xb0,
	0x5d, 0x87, 0x40, 0x5e, 0x87, 0xf9, 0x4d, 0x32, 0x1f, 0xf6, 0x44, 0x4d, 0x3f, 0x7b, 0x54, 0x53,
	0x47, 0xd2, 0x2e, 0x04, 0x8b, 0xf6, 0xd2, 0xfb, 0x26, 0x0f, 0xfd, 0xd0, 0x8d, 0x3f, 0x0f, 0xf3,
	0x20, 0x5a, 0x0d, 0xfb, 0xa2, 0xce, 0x4f, 0x1f, 0xb0, 0x4d, 0xb1, 0xc5, 0x2f, 0xbc, 0x0c, 0xf2,
	0xf0, 0x36, 0xd8, 0xfa, 0xd0, 0x5c, 0x03, 0x41, 0x54, 0x4a, 0xc4, 0x9a, 0x83, 0x92, 0xca, 0x5d,
	0x97, 0x0e, 0x55, 0xd8, 0xed, 0x9c, 0x1c, 0x8d, 0x65, 0x1f, 0x8f, 0xab, 0x3e, 0x1e, 0xb3, 0x8a,
	0x41, 0x77, 0xe4, 0x57, 0x7f, 0xd6, 0x50, 0xf7, 0x34, 0x5a, 0xad, 0x40, 0x63, 0x63, 0x3b, 0x5b,
	0x85, 0xda, 0x0f, 0xa8, 0x95, 0x86, 0x59, 0x94, 0xc8, 0x16, 0xee, 0x9f, 0xbc, 0x7e, 0x74, 0xb9,
	0xfb, 0x74, 0x47, 0x50, 0x69, 0x79, 0x44, 0x7b, 0x81, 0x94, 0x45, 0x16, 0x7e, 0x28, 0x44, 0xdb,
	0xf2, 0x86, 0x6f, 0xd2, 0x1d, 0x00, 0x17, 0x6f, 0xcc, 0x03, 0xd9, 0xce, 0xff, 0x34, 0x08, 0xf7,
	0xfd, 0x9a, 0xc1, 0x96, 0x72, 0xb2, 0xf6, 0x1c, 0x29, 0xf0, 0xf1, 0x7f, 0x0d, 0x56, 0x45, 0x28,
	0x5a, 0x1d, 0x06, 0x01, 0x80, 0x4b, 0x6e, 0xbf, 0xfa, 0xbd, 0x81, 0x7a, 0x4e, 0x91, 0xcd, 0x6e,
	0x82, 0x4d, 0x38, 0xb7, 0xf2, 0x70, 0xfd, 0x68, 0xf8, 0xa0, 0x13, 0xca, 0x69, 0xf5, 0x67, 0xc9,
	0x5c, 0x0e, 0xa1, 0x42, 0x3b, 0x25, 0x66, 0x00, 0xa4, 0xa9, 0xa8, 0xb1, 0xf9, 0xa5, 0x28, 0x87,
	0x8c, 0x2f, 0x35, 0x0d, 0x1d, 0xc4, 0xc1, 0xba, 0x0a, 0x27, 0xd6, 0xbc, 0x00, 0x45, 0x1c, 0xe5,
	0x7e, 0x9a, 0x45, 0xb3, 0xb0, 0x1c, 0x29, 0x85, 0x23, 0x0e, 0x07, 0xf8, 0xb8, 0x7e, 0x28, 0x82,
	0x38, 0x8f, 0xf2, 0xad, 0x98, 0xa4, 0x26, 0xbd, 0xb3, 0xb5, 0x9f, 0x76, 0x39, 0x88, 0x87, 0xe0,
	0x50, 0xdc, 0xff, 0xc5, 0xdd, 0xfd, 0xab, 0xe7, 0xc4, 0x91, 0x5f, 0xf1, 0x08, 0x54, 0x19, 0x72,
	0x03, 0x74, 0x6b, 0x85, 0x1f, 0xd3, 0x28, 0xdb, 0x8a, 0xa9, 0xf8, 0xf7, 0xd2, 0x96, 0x4c, 0xed,
	0x7b, 0x84, 0x00, 0xc8, 0x72, 0x7f, 0x0e, 0x4d, 0x22, 0x26, 0xe4, 0x3f, 0x5a, 0x42, 0xb0, 0x4d,
	0x20, 0x6b, 0xcf, 0x50, 0x0b, 0xde, 0x99, 0xf0, 0x76, 0x23, 0xe6, 0xa6, 0x4d, 0x4b, 0x4b, 0x7b,
	0x83, 0xb4, 0x34, 0xc8, 0x78, 0xdb, 0xa6, 0xa5, 0xe6, 0xbc, 0x17, 0xbb, 0x42, 0x24, 0x55, 0xee,
	0x54, 0xc5, 0xb0, 0xe6, 0xa3, 0x02, 0x0d, 0xf6, 0x5e, 0x36, 0xad, 0x8d, 0x0e, 0x74, 0x8f, 0x9d,
	0xa9, 0x4f, 0x40, 0xf3, 0xae, 0xa1, 0x3b, 0xcc, 0xa3, 0xd8, 0x17, 0x48, 0x4d, 0xeb, 0xa0, 0xc3,
	0x12, 0x51, 0xeb, 0x1a, 0x42, 0x2d, 0x83, 0x62, 0xd3, 0x62, 0x6a, 0x83, 0xaf, 0x29, 0x9e, 0x78,
	0xc4, 0x54, 0x0f, 0xb8, 0x83, 0x4b, 0xdb, 0x32, 0xd5, 0x26, 0xa7, 0x9b, 0x96, 0xeb, 0x78, 0x0c,
	0xab, 0x2d, 0x4e, 0xb9, 0xc4, 0xd4, 0x9a, 0x5c, 0xab, 0x87, 0xa3, 0xd7, 0x08, 0xd1, 0x70, 0x51,
	0xc4, 0x73, 0x11, 0x11, 0x68, 0x8e, 0x4e, 0x99, 0xa5, 0x4f, 0x21, 0x28, 0x9c, 0x9e, 0x78, 0xd3,
	0xa9, 0x5a, 0x1b, 0xfd, 0x88, 0xba, 0x92, 0x44, 0xc3, 0x60, 0x03, 0x4f, 0x76, 0x0f, 0x29, 0x14,
	0xff, 0xec, 0x61, 0x97, 0x61, 0x13, 0x88, 0x60, 0x9a, 0x9e, 0x33, 0xb5, 0x0c, 0x1d, 0xdc, 0xd7,
	0xa0, 0xa7, 0xd0, 0x84, 0xea, 0x9e, 0xe9, 0x4d, 0x31, 0x61, 0x6a, 0x7d, 0xf4, 0x06, 0xf5, 0xe0,
	0x61, 0x48, 0x8b, 0x3c, 0x74, 0xf3, 0x20, 0x2f, 0x36, 0xdc, 0xb1, 0xed, 0x60, 0x02, 0x27, 0x0f,
	0x51, 0xe3, 0xca, 0x26, 0x70, 0x06, 0xa0, 0xa9, 0xed, 0x3e, 0x64, 0x97, 0xc1, 0x14, 0xd4, 0x14,
	0xee, 0xa4, 0x0c, 0x1e, 0xa1, 0xd8, 0xb0, 0xdf, 0x12, 0xeb, 0x1d, 0x84, 0xae, 0x8d, 0x7e, 0xab,
	0x43, 0x47, 0x3f, 0x78, 0x2d, 0xc0, 0x93, 0xa1, 0xbb, 0x5c, 0x34, 0xc8, 0xc3, 0xc4, 0xa7, 0x16,
	0xf3, 0x0d, 0x9d, 0x02, 0x57, 0x1b, 0xa0, 0x8e, 0x54, 0x49, 0x02, 0x75, 0xee, 0xce, 0xa1, 0xd8,
	0xd1, 0x2d, 0x53, 0x22, 0x42, 0x3c, 0xe3, 0x8c, 0xdf, 0x0c, 0xc4, 0xfb, 0x1f, 0xea, 0x9d, 0xea,
	0xe4, 0xdc, 0x67, 0x54, 0x27, 0xee, 0x04, 0x53, 0x50, 0x51, 0x43, 0x7d, 0x97, 0xe9, 0xc4, 0xb4,
	0xc8, 0x5b, 0xdf, 0xa6, 0x26, 0x60, 0x2d, 0xee, 0xc4, 0xb4, 0x20, 0x27, 0xe6, 0x8b, 0x60, 0xea,
	0x21, 0x67, 0x99, 0xd6, 0x5b, 0x8b, 0xe9, 0x53, 0xff, 0x4a, 0x9f, 0x4e, 0x31, 0x53, 0xdb, 0x1c,
	0xbb, 0xb0, 0x4f, 0xad, 0x29, 0xf6, 0x1d, 0xfd, 0xfa, 0x82, 0xeb, 0xa2, 0x70, 0xb1, 0xe1, 0x80,
	0x61, 0x5b, 0x44, 0x45, 0x22, 0xb2, 0xed, 0x39, 0x20, 0x46, 0xe7, 0x5e, 0xa2, 0xc4, 0x06, 0x45,
	0xbb, 0xda, 0x53, 0x34, 0xd0, 0x0d, 0xd8, 0x26, 0xcc, 0x3f, 0xd5, 0xa7, 0x3a, 0x31, 0xb0, 0xda,
	0xe3, 0xf9, 0x55, 0x61, 0x5c, 0x66, 0x43, 0x1f, 0xf4, 0x47, 0x1f, 0xd0, 0x60, 0xef, 0x51, 0xe0,
	0x22, 0xb8, 0xd7, 0xc4, 0x00, 0x0a, 0xd4, 0x15, 0x44, 0xe9, 0xa2, 0xb6, 0xb0, 0x31, 0x31, 0x65,
	0xa9, 0x28, 0xbe, 0xb3, 0x85, 0x22, 0xae, 0x83, 0x0d, 0x6b, 0x62, 0x19, 0xbe, 0xa9, 0x5f, 0x83,
	0x22, 0x90, 0x0b, 0x2c, 0x7c, 0x7b, 0xe2, 0x5f, 0x61, 0x7c, 0x0e, 0xb2, 0x40, 0xb1, 0x2f, 0x6c,
	0xc2, 0xce, 0xc4, 0x89, 0xe6, 0x68, 0x82, 0xb4, 0xc7, 0xef, 0x1b, 0x2f, 0x02, 0xb1, 0x09, 0x96,
	0x15, 0xe6, 0x8e, 0x44, 0x85, 0x85, 0x87, 0x3a, 0x2f, 0xa8, 0xf0, 0x00, 0xde, 0x01, 0xbc, 0xc6,
	0x3a, 0x55, 0x0f, 0x46, 0x7f, 0xd4, 0x50, 0xbf, 0x9a, 0x86, 0xb2, 0x4d, 0x9e, 0xa3, 0x4f, 0x3c,
	0x72, 0x4e, 0xec, 0x2b, 0xe2, 0x3b, 0x1e, 0x35, 0xce, 0x74, 0x17, 0xf3, 0x6b, 0x30, 0xcf, 0x95,
	0xc5, 0xb5, 0x08, 0xc3, 0x54, 0xf6, 0x60, 0x8d, 0xdf, 0x0b, 0x2a, 0x67, 0x9c, 0xdb, 0x1e, 0x93,
	0x21, 0x5c, 0xcc, 0x3c, 0x47, 0x96, 0x54, 0x37, 0x98, 0x75, 0xc9, 0x4b, 0x0a, 0x24, 0x8b, 0x94,
	0x56, 0x93, 0xdf, 0xc4, 0xf5, 0xe0, 0xba, 0xc4, 0x04, 0x0f, 0x2d, 0x6e, 0x1a, 0x5c, 0x5a, 0xa8,
	0x98, 0x09, 0x55, 0x04, 0x17, 0x72, 0xa7, 0xcd, 0x97, 0x8c, 0xf2, 0x99, 0x50, 0x46, 0x31, 0x54,
	0x62, 0x99, 0x85, 0x21, 0x6f, 0x38, 0x5d, 0xfe, 0x7f, 0x91, 0xc3, 0xc7, 0x5b, 0xff, 0x09, 0x2f,
	0x29, 0x14, 0x0a, 0x3b, 0x62, 0x0e, 0xc0, 0x28, 0xfd, 0x43, 0x46, 0xe0, 0xdd, 0x23, 0x95, 0x29,
	0x1b, 0x4d, 0x04, 0x83, 0xac, 0x80, 0xc7, 0xaf, 0xc1, 0x4b, 0x28, 0x92, 0x82, 0x5b, 0x5d, 0x40,
	0x9a, 0x7c, 0x54, 0x4f, 0xdf, 0xbc, 0x1b, 0x2d, 0xa3, 0xfc, 0xa6, 0x78, 0x3f, 0x9e, 0x25, 0xeb,
	0xe3, 0xf2, 0x05, 0x94, 0x7f, 0xb5, 0xbe, 0x5e, 0x26, 0xc7, 0x0b, 0xfe, 0x40, 0x6e, 0x8e, 0xcb,
	0x5f, 0x84, 0xf7, 0x2d, 0xb1, 0xf1, 0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x03, 0xd1,
	0xcb, 0xb3, 0x09, 0x00, 0x00,
}