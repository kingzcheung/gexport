grammar DDLGrammar;


ddl
    : createTable '(' columnsDefinition+ ')' table_options
    ;
createTable
    : K_CREATE K_TABLE (ifNotExists)* tablename
    ;

ifNotExists
    : K_IF K_NOT K_EXISTS
    ;

tablename
    : uid
    | ('`') uid ('`')
    ;


columnsDefinition
    : uid columnType defaultOption? isIndex? isNull? (K_AUTO_INCREMENT) ? (actionList)? (columnComment)? (',')*
    | uid columnType isNull? defaultOption? isIndex?  (K_AUTO_INCREMENT) ? (actionList)? (columnComment)? (',')*
    | uid columnType defaultOption? isNull? isIndex? (K_AUTO_INCREMENT) ? (actionList)? (columnComment)? (',')*
    | K_PRIMARY? K_KEY uid? '(' uid ')' (',')*
    ;

uid
    : BACKTICK? VAL BACKTICK?
    ;

columnType
    : K_INT '(' INT ')' K_UNSIGNED*
    | K_INT K_UNSIGNED*
    | K_INTEGER | K_INTEGER '(' INT ')'
    | K_BIGINT | K_BIGINT '(' INT ')'
    | K_TINYINT | K_TINYINT '(' INT ')'
    | K_SMALLINT | K_SMALLINT '(' INT ')'
    | K_MEDIUMINT | K_MEDIUMINT '(' INT ')'
    | K_FLOAT | K_FLOAT '(' INT ')'
    | K_DOUBLE | K_DOUBLE '(' INT ')'
    | K_DECIMAL | K_DECIMAL '(' INT ',' INT ')'
    | K_DATE
    | K_TIME
    | K_YEAR
    | K_DATETIME
    | K_TIMESTAMP
    | K_CHAR '(' INT ')'
    | K_VARCHAR '(' INT ')'
    | K_TINYBLOB
    | K_TINYTEXT
    | K_BLOB
    | K_TEXT
    | K_MEDIUMBLOB
    | K_MEDIUMTEXT
    | K_LONGBLOB
    | K_LONGTEXT
    ;

defaultOption
    : K_DEFAULT defaultValue
    ;

defaultValue
    : K_NULL
    | STRING_LITERAL
    | INT
    | FUNCTION_LITERAL
    ;

isIndex
    : K_PRIMARY K_KEY
    | K_UNIQUE (K_KEY)*
    ;

isNull
    : (K_NOT)* K_NULL
    ;

actionList
    : K_ON K_UPDATE
    | K_ON K_DELETE
    ;

columnComment
    : K_COMMENT STRING_LITERAL
    ;

table_options
    : (K_ENGINE '=' VAL)* (K_AUTO_INCREMENT '=' INT)* (K_CHARSET '=' VAL)* (K_COMMENT '\\' ANY '\\')* (';')*
    ;



WS : [ \t\n\r]+ -> channel(HIDDEN);
K_CREATE: C R E A T E;
K_TABLE : T A B L E;
K_IF: I F;
K_NOT : N O T;
K_EXISTS: E X I S T S;
K_INT: I N T;
K_INTEGER: I N T E G E R;
K_BIGINT: B I G I N T;
K_TINYINT: T I N Y I N T;
K_SMALLINT: S M A L L I N T;
K_MEDIUMINT:M E D I U M I N T;
K_FLOAT:F L O A T;
K_DOUBLE:D O U B L E;
K_DECIMAL:D E C I M A L;

K_DATE:D A T E;
K_TIME:T I M E;
K_YEAR:Y E A R;
K_DATETIME:D A T E T I M E;
K_TIMESTAMP:T I M E S T A M P;

K_CHAR:C H A R;
K_VARCHAR:V A R C H A R;
K_TINYBLOB:T I N Y B L O B;
K_TINYTEXT:T I N Y T E X T;
K_BLOB:B L O B;
K_TEXT:T E X T;
K_MEDIUMBLOB:M E D I U M B L O B;
K_MEDIUMTEXT:M E D I U M T E X T;
K_LONGBLOB:L O N G B L O B;
K_LONGTEXT:L O N G T E X T;
K_NULL: N U L L;
K_AUTO_INCREMENT: A U T O UNDER_LINE I N C R E M E N T;
K_PRIMARY: P R I M A R Y ;
K_KEY: K E Y ;
K_UNIQUE: U N I Q U E;
K_CURRENT: C U R R E N T;
K_ON: O N;
K_UPDATE: U P D A T E;
K_DELETE: D E L E T E;
K_COMMENT: C O M M E N T;
K_ENGINE:E N G I N E;
K_DEFAULT:D E F A U L T;
K_CHARSET:C H A R S E T;
K_UNSIGNED: U N S I G N E D;
INT: [0-9]+;
BACKTICK: '`';
//NAME: [a-zA-Z_0-9`]+;

STRING_LITERAL: '\'' ( ~'\'' | '\'\'' )* '\'';
VAL: [a-zA-Z_0-9]+;
FUNCTION_LITERAL: [a-zA-Z_0-9]+;
ANY: .;

fragment DIGIT : [0-9];
fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];

fragment UNDER_LINE: [_];


