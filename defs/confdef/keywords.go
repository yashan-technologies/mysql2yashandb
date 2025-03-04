package confdef

import "strings"

var _keywords = map[string]struct{}{
	"ABORT":                {},
	"ACCESS":               {},
	"ACCOUNT":              {},
	"ACTIONS":              {},
	"ADD":                  {},
	"ALL":                  {},
	"ALL_ROWS":             {},
	"ALTER":                {},
	"ANALYZE":              {},
	"AND":                  {},
	"ANY":                  {},
	"ARCHIVE":              {},
	"ARCHIVELOG":           {},
	"AS":                   {},
	"ASC":                  {},
	"AT":                   {},
	"AUDIT":                {},
	"BACKUP":               {},
	"BASE":                 {},
	"BEFORE":               {},
	"BEGIN":                {},
	"BETWEEN":              {},
	"BIGINT":               {},
	"BINARY":               {},
	"BINARY_BIGINT":        {},
	"BINARY_DOUBLE":        {},
	"BINARY_FLOAT":         {},
	"BINARY_INTEGER":       {},
	"BINARY_SMALLINT":      {},
	"BINARY_TINYINT":       {},
	"BIT":                  {},
	"BLOB":                 {},
	"BOOLEAN":              {},
	"BOUND":                {},
	"BUFFER_POOL":          {},
	"BUILD":                {},
	"BULK":                 {},
	"BULKLOAD":             {},
	"CACHE":                {},
	"CALL":                 {},
	"CANCEL":               {},
	"CASCADE":              {},
	"CASE":                 {},
	"CATEGORY":             {},
	"CELL_FLASH_CACHE":     {},
	"CHANGE":               {},
	"CHAR":                 {},
	"CHAR_CS":              {},
	"CHARACTER":            {},
	"CHECK":                {},
	"CHECKPOINT":           {},
	"CHOOSE":               {},
	"CHUNK":                {},
	"CLEAN":                {},
	"CLOB":                 {},
	"CLOSE":                {},
	"CLUSTER":              {},
	"COALESCE":             {},
	"COLUMN":               {},
	"COLUMNAR":             {},
	"COMMENT":              {},
	"COMMIT":               {},
	"COMPRESS":             {},
	"COMPRESSION":          {},
	"CONDITION":            {},
	"CONNECT":              {},
	"CONNECT_BY_ISCYCLE":   {},
	"CONNECT_BY_ISLEAF":    {},
	"CONNECT_BY_ROOT":      {},
	"CONSTRAINT":           {},
	"CONSTRAINTS":          {},
	"CONTINUE":             {},
	"CONTROLFILES":         {},
	"CONVERT":              {},
	"CREATE":               {},
	"CROSS":                {},
	"CUBE":                 {},
	"CURRENT":              {},
	"CURRENT_DATE":         {},
	"CURRENT_TIMESTAMP":    {},
	"CYCLE":                {},
	"DATABASE":             {},
	"DATAFILE":             {},
	"DATAOID_REUSE":        {},
	"DATASPACE":            {},
	"DATE":                 {},
	"DAY":                  {},
	"DECIMAL":              {},
	"DECLARE":              {},
	"DECRYPTION":           {},
	"DEDUPLICATE":          {},
	"DEFAULT":              {},
	"DEFINITION":           {},
	"DELETE":               {},
	"DESC":                 {},
	"DIRECTORY":            {},
	"DISABLE":              {},
	"DISTINCT":             {},
	"DISTRIBUTE":           {},
	"DOUBLE":               {},
	"DOUBLE_PRECISION":     {},
	"DOUBLE_WRITE":         {},
	"DROP":                 {},
	"DUMP":                 {},
	"DUPLICATE":            {},
	"DUPLICATED":           {},
	"EDITIONABLE":          {},
	"ELSE":                 {},
	"ELSIF":                {},
	"ENABLE":               {},
	"ENCODING":             {},
	"ENCRYPTION":           {},
	"END":                  {},
	"ENGINE_COL":           {},
	"ENGINE_ROW":           {},
	"ESCAPE":               {},
	"EXCEPT":               {},
	"EXCEPTION":            {},
	"EXEC":                 {},
	"EXECUTE":              {},
	"EXISTS":               {},
	"EXIT":                 {},
	"EXPLAIN":              {},
	"EXTEND":               {},
	"EXTERNAL":             {},
	"FAILOVER":             {},
	"FALSE":                {},
	"FETCH":                {},
	"FIRST":                {},
	"FIRST_ROWS":           {},
	"FLASH_CACHE":          {},
	"FLASHBACK":            {},
	"FLOAT":                {},
	"FLUSH":                {},
	"FOLLOWING":            {},
	"FOR":                  {},
	"FORALL":               {},
	"FORCE":                {},
	"FOREIGN":              {},
	"FORMAT":               {},
	"FREELIST":             {},
	"FREELISTS":            {},
	"FROM":                 {},
	"FULL":                 {},
	"FUNCTION":             {},
	"GLOBAL":               {},
	"GOTO":                 {},
	"GRANT":                {},
	"GROUP":                {},
	"GROUPING":             {},
	"HASH_AJ":              {},
	"HASH_SJ":              {},
	"HAVING":               {},
	"HEAP":                 {},
	"HORDER":               {},
	"HOUR":                 {},
	"IDENTIFIED":           {},
	"IF":                   {},
	"IGNORE":               {},
	"IMMEDIATE":            {},
	"IN":                   {},
	"INCLUDE":              {},
	"INCREMENT":            {},
	"INCREMENTAL":          {},
	"INDEPEND":             {},
	"INDEX":                {},
	"INDEX_ASC":            {},
	"INDEX_DESC":           {},
	"INDEX_FFS":            {},
	"INDEX_JOIN":           {},
	"INDEX_SS":             {},
	"INDEX_SS_ASC":         {},
	"INDEX_SS_DESC":        {},
	"INITIAL":              {},
	"INITRANS":             {},
	"INMEMORY":             {},
	"INNER":                {},
	"INSERT":               {},
	"INSTANCE":             {},
	"INSTANCES":            {},
	"INT":                  {},
	"INTEGER":              {},
	"INTERSECT":            {},
	"INTERVAL":             {},
	"INTO":                 {},
	"INVALIDATE":           {},
	"INVISIBLE":            {},
	"IS":                   {},
	"JOIN":                 {},
	"JSON":                 {},
	"KEEP_DUPLICATES":      {},
	"KILL":                 {},
	"LEADING":              {},
	"LEFT":                 {},
	"LEVEL":                {},
	"LIBRARY":              {},
	"LIKE":                 {},
	"LIMIT":                {},
	"LINK":                 {},
	"LOAD":                 {},
	"LOB":                  {},
	"LOCAL":                {},
	"LOCK":                 {},
	"LOG":                  {},
	"LOGFILE":              {},
	"LOGGING":              {},
	"LOGOFF":               {},
	"LOGON":                {},
	"LOOP":                 {},
	"LSC":                  {},
	"MATCHED":              {},
	"MATERIALIZED":         {},
	"MAX_WORKERS_PER_EXEC": {},
	"MAXDATABUCKETS":       {},
	"MAXDATAFILES":         {},
	"MAXEXTENTS":           {},
	"MAXINSTANCES":         {},
	"MAXLOGFILES":          {},
	"MAXLOGHISTORY":        {},
	"MAXSIZE":              {},
	"MAXTRANS":             {},
	"MAXVALUE":             {},
	"MCOL":                 {},
	"MERGE":                {},
	"MERGE_AJ":             {},
	"MERGE_SJ":             {},
	"MINEXTENTS":           {},
	"MINUS":                {},
	"MINUTE":               {},
	"MINVALUE":             {},
	"MODIFY":               {},
	"MONTH":                {},
	"MOUNT":                {},
	"NATIONAL":             {},
	"NATURAL":              {},
	"NCHAR":                {},
	"NCHAR_CS":             {},
	"NCLOB":                {},
	"NESTED":               {},
	"NEW":                  {},
	"NEXT":                 {},
	"NEXTVAL":              {},
	"NL_AJ":                {},
	"NL_SJ":                {},
	"NO":                   {},
	"NO_INDEX":             {},
	"NO_INDEX_FFS":         {},
	"NO_INDEX_SS":          {},
	"NO_USE_HASH":          {},
	"NO_USE_MERGE":         {},
	"NO_USE_NL":            {},
	"NOARCHIVELOG":         {},
	"NOAUDIT":              {},
	"NOCACHE":              {},
	"NOCOMPRESS":           {},
	"NOCYCLE":              {},
	"NOLOGGING":            {},
	"NOMAXVALUE":           {},
	"NOMINVALUE":           {},
	"NONEDITIONABLE":       {},
	"NOORDER":              {},
	"NOPARALLEL":           {},
	"NOREVERSE":            {},
	"NORMAL":               {},
	"NOT":                  {},
	"NOVALIDATE":           {},
	"NOWAIT":               {},
	"NULL":                 {},
	"NULLS":                {},
	"NUMBER":               {},
	"NUMERIC":              {},
	"NVARCHAR":             {},
	"NVARCHAR2":            {},
	"OBJNO_REUSE":          {},
	"OF":                   {},
	"OFFLINE":              {},
	"OFFSET":               {},
	"ON":                   {},
	"ONLINE":               {},
	"ONLY":                 {},
	"OPEN":                 {},
	"OR":                   {},
	"ORDER":                {},
	"ORDERED":              {},
	"ORGANIZATION":         {},
	"OUTER":                {},
	"OUTLINE":              {},
	"OVER":                 {},
	"PACKAGE":              {},
	"PARALLEL":             {},
	"PARALLELISM":          {},
	"PARTITION":            {},
	"PASSWORD":             {},
	"PCTFREE":              {},
	"PCTINCREASE":          {},
	"PCTUSED":              {},
	"PIVOT":                {},
	"PLS_INTEGER":          {},
	"PRAGMA":               {},
	"PRECEDING":            {},
	"PREPARE":              {},
	"PRESERVE":             {},
	"PRIMARY":              {},
	"PRIOR":                {},
	"PRIVATE":              {},
	"PRIVILEGE":            {},
	"PRIVILEGES":           {},
	"PROCEDURE":            {},
	"PROFILE":              {},
	"PUBLIC":               {},
	"PURGE":                {},
	"RAISE":                {},
	"RANGE":                {},
	"RAW":                  {},
	"READONLY":             {},
	"READWRITE":            {},
	"REAL":                 {},
	"REBUILD":              {},
	"RECLAIM":              {},
	"RECOVER":              {},
	"REFERENCES":           {},
	"REFRESH":              {},
	"REGISTER":             {},
	"RELEASE":              {},
	"RENAME":               {},
	"RESETLOGS":            {},
	"RESPECT":              {},
	"RESTART":              {},
	"RESTORE":              {},
	"RESTRICT":             {},
	"RETURN":               {},
	"RETURNING":            {},
	"REUSE":                {},
	"REVERSE":              {},
	"REVOKE":               {},
	"RIGHT":                {},
	"RLIKE":                {},
	"ROLE":                 {},
	"ROLES":                {},
	"ROLLBACK":             {},
	"ROLLUP":               {},
	"ROW":                  {},
	"ROWID":                {},
	"ROWNUM":               {},
	"ROWS":                 {},
	"RTREE":                {},
	"RULE":                 {},
	"SAMPLE":               {},
	"SAVEPOINT":            {},
	"SCHEMA":               {},
	"SCN":                  {},
	"SECOND":               {},
	"SECTION":              {},
	"SEGMENT":              {},
	"SELECT":               {},
	"SELECTIVITY":          {},
	"SEPARATOR":            {},
	"SEQUENCE":             {},
	"SESSION":              {},
	"SET":                  {},
	"SETS":                 {},
	"SHARDED":              {},
	"SHARED":               {},
	"SHRINK":               {},
	"SHUTDOWN":             {},
	"SIBLINGS":             {},
	"SKIP":                 {},
	"SLICE":                {},
	"SMALLINT":             {},
	"SOME":                 {},
	"SPLIT":                {},
	"SQL":                  {},
	"SQLMAP":               {},
	"STABLE":               {},
	"STANDBY":              {},
	"START":                {},
	"STORAGE":              {},
	"SUBPARTITION":         {},
	"SUCCESSFUL":           {},
	"SUPPLEMENTAL":         {},
	"SWAP":                 {},
	"SWITCH":               {},
	"SWITCHOVER":           {},
	"SYNONYM":              {},
	"SYS_REFCURSOR":        {},
	"SYSAUX":               {},
	"SYSDATE":              {},
	"SYSTEM":               {},
	"SYSTIMESTAMP":         {},
	"TABLE":                {},
	"TABLESPACE":           {},
	"TAC":                  {},
	"TAG":                  {},
	"TEMPFILE":             {},
	"TEMPORARY":            {},
	"THEN":                 {},
	"TIME":                 {},
	"TIMESTAMP":            {},
	"TINYINT":              {},
	"TO":                   {},
	"TRANSACTION":          {},
	"TRANSPORT":            {},
	"TRIGGER":              {},
	"TRUE":                 {},
	"TRUNCATE":             {},
	"TTL":                  {},
	"TYPE":                 {},
	"UNBOUNDED":            {},
	"UNDO":                 {},
	"UNDO_SEGMENTS":        {},
	"UNION":                {},
	"UNIQUE":               {},
	"UNUSABLE":             {},
	"UPDATE":               {},
	"UPGRADE":              {},
	"UROWID":               {},
	"USABLE":               {},
	"USE_HASH":             {},
	"USE_MERGE":            {},
	"USE_NL":               {},
	"USER":                 {},
	"USING":                {},
	"UTC_TIMESTAMP":        {},
	"VALIDATE":             {},
	"VALUES":               {},
	"VARCHAR":              {},
	"VARCHAR2":             {},
	"VIEW":                 {},
	"VISIBLE":              {},
	"WAIT":                 {},
	"WHEN":                 {},
	"WHENEVER":             {},
	"WHERE":                {},
	"WHILE":                {},
	"WITH":                 {},
	"WITHOUT":              {},
	"WRAPPED":              {},
	"XMLTYPE":              {},
	"YEAR":                 {},
	"ZORDER":               {},
}

func initKeywords(conf M2YConfig) {
	for _, s := range conf.Yashan.AddtionalKeywords {
		_keywords[strings.ToUpper(s)] = struct{}{}
	}
}

func IsKeyword(s string) bool {
	_, ok := _keywords[strings.ToUpper(s)]
	return ok
}
