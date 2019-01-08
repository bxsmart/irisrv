package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type CCoinAccount struct {
	BasicEntity

	ID        int64           `gorm: "column:id;type:bigint;primary_key;AUTO_INCREMENT"`            // auto-increment by-default by gorm
	Cid       int64           `gorm: "column:cid;type:bigint;;unique_index:unique_index:u_w_c_udx"` // 币种ID
	Wid       int64           `gorm: "column:wid;type:bigint;unique_index:u_w_c_udx"`               // 钱包ID
	Uid       string          `gorm: "column:uid;type:bigint;unique_index:u_w_c_udx"`               // 币种名称
	Coin      string          `gorm: "column:coin;type:varchar(20)"`                                // 币种名称
	Amount    decimal.Decimal `gorm: "column:amount;type:DECIMAL(28, 8)"`                           // 可用数量
	Balance   decimal.Decimal `gorm: "column:balance;type:DECIMAL(28, 8)"`                          // 总数量
	TxFrozen  decimal.Decimal `gorm: "column:tx_frozen;type:DECIMAL(28, 8)"`                        // 交易冻结
	PltFrozen decimal.Decimal `gorm: "column:plt_frozen;type:DECIMAL(28, 8)"`                       // 平台冻结
}

func (e *CCoinAccount) TableName() string {
	return "cc_position"
}

func (e *CCoinAccount) BeforeInsert() {
	e.CreatedAt = time.Now().UTC()
}

func (e *CCoinAccount) BeforeUpdate() {
	e.UpdatedAt = time.Now().UTC()
}

type FCoinAccount struct {
	BasicEntity

	ID        int64           `gorm: "column:id;type:bigint;primary_key;AUTO_INCREMENT"` // auto-increment by-default by gorm
	Uid       string          `gorm: "column:uid;type:bigint;unique_index:u_w_c_udx"`    // 币种名称
	Cid       int64           `gorm: "column:cid;type:bigint;unique_index:u_w_c_udx"`    // 币种ID
	Coin      string          `gorm: "column:coin;type:varchar(20)"`                     // 币种名称
	Amount    decimal.Decimal `gorm: "column:amount;type:DECIMAL(28, 8)"`                // 可用数量
	Balance   decimal.Decimal `gorm: "column:balance;type:DECIMAL(28, 8)"`               // 总数量
	TxFrozen  decimal.Decimal `gorm: "column:tx_frozen;type:DECIMAL(28, 8)"`             // 交易冻结
	PltFrozen decimal.Decimal `gorm: "column:plt_frozen;type:DECIMAL(28, 8)"`            // 平台冻结
}

func (e *FCoinAccount) TableName() string {
	return "fc_position"
}
