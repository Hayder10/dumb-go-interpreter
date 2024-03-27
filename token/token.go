package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  //Variable or function identifiers
  IDENT = "IDENT"

  //Literals
  INT = "INT"

  //Operators
  ASSIGN = "="
  PLUS = "+"

  //Delimiters and parenthesis
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //KEYWORDS
  FUNCTION = "FUNCTION"
  LET = "LET"
)
