package models

// AccountAddUnitRequest 添加子单元
type AccountAddUnitRequest struct {
	Request   AccountAddUnitData `json:"request"`
	Signature Signature          `json:"signature"`
}

type AccountAddUnitData struct {
	Body NtdmaaddxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtdmaaddxBody struct {
	Ntdmaaddx []*Ntdmaaddx `json:"ntdmaaddx,omitempty"`
}

type Ntdmaaddx struct {
	Accnbr string `json:"accnbr"`
	//Bcktyp string `json:"bcktyp"`
	//Clstyp string `json:"clstyp"`
	Dmanam string `json:"dmanam"`
	Dmanbr string `json:"dmanbr"`
	//Ovrctl string `json:"ovrctl"`
}

// AccountCloseUnitRequest 关闭子单元
type AccountCloseUnitRequest struct {
	Request   AccountCloseUnitData `json:"request"`
	Signature Signature            `json:"signature"`
}

type AccountCloseUnitData struct {
	Body NtbusmodyBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtbusmodyBody struct {
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdumdltx1 []*Ntdumdltx1 `json:"ntdumdltx1,omitempty"`
	Ntdumdltx2 []*Ntdumdltx2 `json:"ntdumdltx2,omitempty"`
}
type Ntdumdltx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Inbacc string `json:"inbacc,omitempty"`
}
type Ntdumdltx2 struct {
	Dyanbr string `json:"dyanbr,omitempty"`
	Yurref string `json:"yurref,omitempty"`
}
type Ntdmadltx1 struct {
	Accnbr string `json:"accnbr"`
}
type Ntdmadltx2 struct {
	Dmanbr string `json:"dmanbr"`
}
type Ntbusmody struct {
	Busmod string `json:"busmod,omitempty"`
}

// AccountUnitInfoRequest    获取金额
type AccountUnitInfoRequest struct {
	Request   AccountUnitBalanceData `json:"request"`
	Signature Signature              `json:"signature"`
}

type AccountUnitBalanceData struct {
	Body NtdmabalxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtdmabalxBody struct {
	Ntdmabalx []*Ntdmabalx `json:"ntdmabalx,omitempty"`
}
type Ntdmabalx struct {
	Accnbr string `json:"accnbr"`
	Dmanbr string `json:"dmanbr"`
}

//QueryUnitTransByBusNoRequest

// QueryUnitTransByBusNoRequest    按照交易流水获取
type QueryUnitTransByBusNoRequest struct {
	Request   QueryUnitTransByBusNoData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type QueryUnitTransByBusNoData struct {
	Body Ntdumredx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdumredx1Body struct {
	Ntdumredx1 []*Ntdumredx1 `json:"ntdumredx1,omitempty"`
}
type Ntdumredx1 struct {
	Yurref string `json:"yurref,omitempty"`
	Bgndat string `json:"bgndat,omitempty"`
	Enddat string `json:"enddat,omitempty"`
}

// QueryUnitAccountBalanceHistoryRequest    获取所有子单元账户历史余额
type QueryUnitAccountBalanceHistoryRequest struct {
	Request   QueryUnitAccountBalanceHistoryData `json:"request"`
	Signature Signature                          `json:"signature"`
}

type QueryUnitAccountBalanceHistoryData struct {
	Body Ntdmahbdz1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmahbdz1Body struct {
	Ntdmahbdx1 []*Ntdmahbdx1 `json:"ntdmahbdx1,omitempty"`
}
type Ntdmahbdx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Qrydat string `json:"qrydat,omitempty"`
	Dmanbr string `json:"dmanbr,omitempty"`
}

// QueryUnitAccountSingleBalanceHistoryRequest    获取单个交易历史余额
type QueryUnitAccountSingleBalanceHistoryRequest struct {
	Request   QueryUnitAccountSingleBalanceHistoryData `json:"request"`
	Signature Signature                                `json:"signature"`
}

type QueryUnitAccountSingleBalanceHistoryData struct {
	Body Ntdmahadx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}
type Ntdmahadx1Body struct {
	Ntdmahadx1 []*Ntdmahadx1 `json:"ntdmahadx1,omitempty"`
}

type Ntdmahadx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` // 分行号
	Accnbr string `json:"accnbr,omitempty"` //账户
	Dmanbr string `json:"dmanbr,omitempty"` // 记账单元
	Begdat string `json:"begdat,omitempty"` // 开始时间
	Enddat string `json:"enddat,omitempty"` // 结束时间
}
