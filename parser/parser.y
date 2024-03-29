%{
package parser

import "fmt"

%}

%union{
    val any
}

%token String
%token Number
%token Print
%token Label
%token Type
%token Get Put Post Delete Head
%token Header Body
%token Json
%token JsonStr
%token Comment
%token Bool
%token Write
%token Append
%token StringFunc
%token CodeBlock
%token If
%token For
%token Read
%token Map
%token Len
%token ForEach
%token Else

%right '='
%left '+' '-'
%left '*' '/'
%left Header Body

%start expressions

%%

expressions:
    expression
    | expressions expression

expression:
    assignment
    | statement
    | Comment

statement:
    print_st
    | label_st
    | type_st
    | network_st
    | write_st
    | read_st
    | append_st
    | if_st
    | for_st
    | foreach_st

label_st:
    Label { PrintYySymDebug($1) }

Value:
    String | Number | Bool

print_st:
    Print '(' Value ')' { PrintYySym($3) }
    | Print '(' Label ')' { PrintGlobalYySym($3) }
    | Print '(' ')' { println("") }
    | Print '(' field_st ')' { PrintYySym(yySymType{ val: WalkField($3.val.(string)) }) }
    | Print '(' network_st ')' { PrintYySym($3) }
    | Print '(' jsonstr_st ')' { PrintYySym($3) }
    | Print '(' arith_st ')' { PrintYySym($3) }
    | Print '(' string_st ')' { PrintYySym($3) }

write_st:
    Write '(' eval_expr ',' String ')' { WriteFile($3.val, $5.val.(string)) }

append_st:
    Append '(' eval_expr ',' String ')' { AppendFile($3.val, $5.val.(string)) }

read_st:
    Read '(' String ')' { $$ = yySymType{ val: ReadFile($3.val.(string)) } }

type_st:
    Type '(' Label ')' { PrintType($3) }

json_st:
    Json '(' String ')' { $$ = yySymType{ val: StrToJson($3.val) } }
    | Json '(' Label ')' { $$ = yySymType{ val: StrToJson(GlobalVarRead($3)) } }
    | Json '(' field_st ')' { $$ = yySymType{ val: StrToJson(WalkField($3.val.(string))) } }
    | Json '(' string_st ')' { $$ = yySymType{ val: StrToJson($3.val) } }

jsonstr_st:
    JsonStr '(' String ')' { $$ = yySymType{ val: ToJsonStr($3.val) } }
    | JsonStr '(' Label ')' { $$ = yySymType{ val: ToJsonStr(GlobalVarRead($3)) } }
    | JsonStr '(' field_st ')' { $$ = yySymType{ val: ToJsonStr(WalkField($3.val.(string))) } }

assignment:
    Label '=' eval_expr { GlobalVarWrite($1, $3.val) }
    | Label '=' network_st { GlobalVarWrite($1, $3.val) }
    | Label '=' json_st { GlobalVarWrite($1, $3.val) }
    | field_st '=' json_st { GlobalVarFieldWrite($1.val.(string), $3.val) }
    | field_st '=' eval_expr { GlobalVarFieldWrite($1.val.(string), $3.val) }
    | field_st '=' network_st { GlobalVarFieldWrite($1.val.(string), $3.val) }

arith_st:
    eval_expr '+' eval_expr { $$ = yySymType{ val: ValAdd($1.val, $3.val) } }
    | eval_expr '-' eval_expr { $$ = yySymType{ val: ValMinus($1.val, $3.val) } }
    | eval_expr '*' eval_expr { $$ = yySymType{ val: ValMul($1.val, $3.val) } }
    | eval_expr '/' eval_expr { $$ = yySymType{ val: ValDiv($1.val, $3.val) } }

eval_expr:
    Value { $$ = $1 }
    | Label { $$ = yySymType{ val: GlobalVarRead($1) } }
    | arith_st { $$ = $1 }
    | field_st { $$ = yySymType{ val: WalkField($1.val.(string)) } }
    | string_st { $$ = $1 }
    | read_st { $$ = $1 }
    | map_st { $$ = $1 }
    | len_st { $$ = $1 }
    | jsonstr_st { $$ = $1 }

header_sg:
    Header eval_expr { $$ = $2 }

header_st:
    /* empty */ { $$ = yySymType{ val : nil } }
    | header_st header_sg { $$ = joinHeaders($1, $2) }

body_st:
    Body String { $$ = yySymType{ val: $2.val } }
    | Body json_st { $$ = yySymType{ val: $2.val } }
    | /* empty */ { $$ = yySymType{ val: nil } }

network_st:
    Get eval_expr header_st { $$ = HttpSend("GET", $2.val.(string), $3, yySymType{}) }
    | Delete eval_expr header_st { $$ = HttpSend("DELETE", $2.val.(string), $3, yySymType{}) }
    | Head eval_expr header_st { $$ = HttpSend("HEAD", $2.val.(string), $3, yySymType{}) }
    | Put eval_expr header_st body_st { $$ = HttpSend("PUT", $2.val.(string), $3, $4) }
    | Post eval_expr header_st body_st { $$ = HttpSend("POST", $2.val.(string), $3, $4) }

/* this is lazy, it doesn't walk the fields recursively until meeting a terminal statement */
field_st:
    Label '.' Label {
        $$ = yySymType{ val: $1.val.(string) + "." + $3.val.(string) }
    }
    | field_st '.' Label {
        $$ = yySymType{ val: $1.val.(string) + "." + $3.val.(string) }
    }
    | Label '.' '[' Number ']' {
        $$ = yySymType{ val: fmt.Sprintf("%s.[%d]", $1.val, $4.val) }
    }
    | field_st '.' '[' Number ']' {
        $$ = yySymType{ val: fmt.Sprintf("%s.[%d]", $1.val, $4.val) }
    }

string_st:
    StringFunc '(' Label ')' { $$ = yySymType{ val: ToStr(GlobalVarRead($3)) } }
    | StringFunc '(' field_st ')' { $$ = yySymType{ val: ToStr(WalkField($3.val.(string))) } }

if_cond:
    Label { $$ = yySymType{ val: GlobalVarRead($1) } }
    | Bool { $$ = $1 }
    | field_st { $$ = yySymType{ val: WalkField($1.val.(string)) } }

else_st:
    Else CodeBlock { $$ = $2 }

if_st:
    If if_cond CodeBlock { $$ = yySymType{ val: RunIfCond($2.val, $3.val) } }
    | if_st else_st { $$ = yySymType{ val: RunIfCond(!($1.val.(bool)), $2.val ) } }

for_st:
    For Number CodeBlock { RepeatBlock($2.val, $3.val) }
    | For Label CodeBlock { RepeatBlock(GlobalVarRead($2), $3.val) }
    | For field_st CodeBlock { RepeatBlock(WalkField($2.val.(string)), $3.val) }
    | For len_st CodeBlock { RepeatBlock($2.val, $3.val) }

map_st:
    Map '(' ')' { $$ = yySymType{ val: map[string]any{} } }

len_st:
    Len '(' Label ')' { $$ = yySymType{ val: CalcLen(GlobalVarRead($3)) } }
    | Len '(' field_st ')' { $$ = yySymType{ val: CalcLen(WalkField($3.val.(string))) } }

foreach_st:
    ForEach Label CodeBlock { DoForEach(GlobalVarRead($2), $3.val) }
    | ForEach field_st CodeBlock { DoForEach(WalkField($2.val.(string)), $3.val) }
