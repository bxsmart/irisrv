package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type CCoinAccount struct {
	BasicEntity

	ID        int64           `xorm: "id pk not null autoincr"`   // auto-increment by-default by xorm
	Cid       int64           `xorm: "cid bigint"`                // 币种ID
	Wid       int64           `xorm: "wid bigint"`                // 钱包ID
	Coin      string          `xorm: "coin varchar(20)"`          // 币种名称
	Uid       string          `xorm: "uid bigint"`                // 币种名称
	Amount    decimal.Decimal `xorm: "amount DECIMAL(28, 8)"`     // 可用数量
	Balance   decimal.Decimal `xorm: "balance DECIMAL(28, 8)"`    // 总数量
	TxFrozen  decimal.Decimal `xorm: "tx_frozen DECIMAL(28, 8)"`  // 交易冻结
	PltFrozen decimal.Decimal `xorm: "plt_frozen DECIMAL(28, 8)"` // 平台冻结
	CreatedAt time.Time       `xorm: "create_at"`                 // 创建时间
	UpdatedAt time.Time       `xorm: "update_at"`                 // 更新时间
}

func (e *CCoinAccount) TableName() string {
	return "cc_position"
}

func (e *CCoinAccount) BeforeInsert() {
	e.CreatedAt = time.Now()
}

func (e *CCoinAccount) BeforeUpdate() {
	e.UpdatedAt = time.Now()
}

type FCoinAccount struct {
	BasicEntity

	ID        int64           `xorm: "id pk not null autoincr"`   // auto-increment by-default by xorm
	Uid       string          `xorm: "uid bigint"`                // 币种名称
	Cid       int64           `xorm: "cid bigint"`                // 币种ID
	Coin      string          `xorm: "coin varchar(20)"`          // 币种名称
	Amount    decimal.Decimal `xorm: "amount DECIMAL(28, 8)"`     // 可用数量
	Balance   decimal.Decimal `xorm: "balance DECIMAL(28, 8)"`    // 总数量
	TxFrozen  decimal.Decimal `xorm: "tx_frozen DECIMAL(28, 8)"`  // 交易冻结
	PltFrozen decimal.Decimal `xorm: "plt_frozen DECIMAL(28, 8)"` // 平台冻结
	CreatedAt time.Time       `xorm: "create_at"`                 // 创建时间
	UpdatedAt time.Time       `xorm: "update_at"`                 // 更新时间
}

func (e *FCoinAccount) TableName() string {
	return "fc_position"
}

func (e *FCoinAccount) BeforeInsert() {
	e.CreatedAt = time.Now()
}

func (e *FCoinAccount) BeforeUpdate() {
	e.UpdatedAt = time.Now()
}
