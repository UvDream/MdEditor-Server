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

)
