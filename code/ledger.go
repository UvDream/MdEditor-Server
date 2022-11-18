package code

const (
	ErrLedgerExist                     = 11001 //	账本已存在
	ErrCreateLedger                    = 11002 //	创建账本失败
	ErrGetLedgerList                   = 11003 //	获取账本列表失败
	ErrCreateLedgerCategory            = 11004 //	创建账本分类失败
	ErrCreateLedgerCategoryRelation    = 11005 //	创建账本分类关联关系失败
	ErrorDeleteLedger                  = 11006 //	删除账本失败
	ErrorLedgerNotExist                = 11007 //	账本不存在
	ErrorGetLedgerCategoryList         = 11008 //	查询账本分类失败
	ErrorGetLedgerCategoryRelationList = 11009 //	查询账本分类关联关系失败
	ErrorDeleteLedgerCategoryRelation  = 11010 //	删除账本分类关联关系失败
	ErrorDeleteLedgerCategory          = 11011 //	删除账本分类失败
	ErrorUpdateLedger                  = 11012 //	更新账本失败
	ErrorGetLedger                     = 11013 //	查询账本失败
	ErrorCreateCategory                = 11014 //   创建分类失败
	ErrCategoryExist                   = 11015 //	分类已存在
	ErrorCategoryNotExist              = 11016 //	分类不存在
	ErrorUpdateCategory                = 11017 //	更新分类失败
	ErrorDeleteCategory                = 11018 //	删除分类失败
	ErrCreateBill                      = 11019 //	创建账单失败
	ErrorDeleteBill                    = 11020 //	删除账单失败
	ErrorUpdateBill                    = 11021 //	更新账单失败
	ErrorGetBill                       = 11022 //	查询账单失败
	ErrorMissingLedgerId               = 11023 //	缺少账本id
)
